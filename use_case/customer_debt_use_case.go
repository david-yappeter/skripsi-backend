package use_case

import (
	"context"
	"fmt"
	"io"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"path"

	"golang.org/x/sync/errgroup"
)

type customerDebtLoaderParams struct {
	customer         bool
	customerPayments bool
	deliveryOrder    bool
}

type CustomerDebtUseCase interface {
	// create
	UploadImage(ctx context.Context, request dto_request.CustomerDebtUploadImageRequest) string
	DownloadReport(ctx context.Context, request dto_request.CustomerDebtDownloadReportRequest) (io.ReadCloser, int64, string, string)

	//  read
	Fetch(ctx context.Context, request dto_request.CustomerDebtFetchRequest) ([]model.CustomerDebt, int)
	Get(ctx context.Context, request dto_request.CustomerDebtGetRequest) model.CustomerDebt

	//  update
	Payment(ctx context.Context, request dto_request.CustomerDebtPaymentRequest) model.CustomerDebt
}

type customerDebtUseCase struct {
	repositoryManager repository.RepositoryManager

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client

	baseFileUseCase
}

func NewCustomerDebtUseCase(
	repositoryManager repository.RepositoryManager,
	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) CustomerDebtUseCase {
	return &customerDebtUseCase{
		repositoryManager: repositoryManager,
		mainFilesystem:    mainFilesystem,
		tmpFilesystem:     tmpFilesystem,
		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
	}
}

func (u *customerDebtUseCase) mustLoadCustomerDebtsData(ctx context.Context, customerDebts []*model.CustomerDebt, option customerDebtLoaderParams) {
	customerLoader := loader.NewCustomerLoader(u.repositoryManager.CustomerRepository())
	customerPaymentsLoader := loader.NewCustomerPaymentsLoader(u.repositoryManager.CustomerPaymentRepository())
	deliveryOrderLoader := loader.NewDeliveryOrderLoader(u.repositoryManager.DeliveryOrderRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range customerDebts {
			if option.customerPayments {
				group.Go(customerPaymentsLoader.CustomerDebtFn(customerDebts[i]))
			}

			if option.customer {
				group.Go(customerLoader.CustomerDebtFn(customerDebts[i]))
			}

			if option.deliveryOrder {
				group.Go(deliveryOrderLoader.CustomerDebtFn(customerDebts[i]))
			}
		}
	}))

	customerTypeLoader := loader.NewCustomerTypeLoader(u.repositoryManager.CustomerTypeRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		if option.customerPayments {
			for i := range customerDebts {
				for j := range customerDebts[i].CustomerPayments {
					group.Go(fileLoader.CustomerPaymentFn(&customerDebts[i].CustomerPayments[j]))
				}

				if option.customer {
					group.Go(customerTypeLoader.CustomerFn(customerDebts[i].Customer))
				}
			}
		}
	}))

	for i := range customerDebts {
		for j := range customerDebts[i].CustomerPayments {
			file := customerDebts[i].CustomerPayments[j].ImageFile
			if file != nil {
				file.SetLink(u.mainFilesystem)
			}
		}
	}
}

func (u *customerDebtUseCase) UploadImage(ctx context.Context, request dto_request.CustomerDebtUploadImageRequest) string {
	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.CustomerPaymentImagePath,
		request.File.Filename,
		request.File,
		fileUploadTemporaryParams{
			supportedExtensions: listSupportedExtension([]string{
				extensionTypeImage,
			}),
			maxFileSizeInBytes: util.Pointer[int64](2 << 20),
		},
	)
}

func (u *customerDebtUseCase) DownloadReport(ctx context.Context, request dto_request.CustomerDebtDownloadReportRequest) (io.ReadCloser, int64, string, string) {
	var customer *model.Customer
	var err error
	if request.CustomerId != nil {
		customer, err = u.repositoryManager.CustomerRepository().Get(ctx, *request.CustomerId)
		panicIfErr(err)
	}

	queryOption := model.CustomerDebtQueryOption{
		QueryOption: model.QueryOption{
			Sorts: model.Sorts{
				{Field: "created_at", Direction: "DESC"},
			},
		},
		StartDate:  request.StartDate.NullDate(),
		EndDate:    request.EndDate.NullDate(),
		CustomerId: request.CustomerId,
	}

	customerDebts, err := u.repositoryManager.CustomerDebtRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadCustomerDebtsData(ctx, util.SliceValueToSlicePointer(customerDebts), customerDebtLoaderParams{
		customerPayments: false,
		customer:         true,
		deliveryOrder:    true,
	})

	reportExcel, err := NewReportCustomerDebtExcel(
		util.CurrentDateTime(),
		request.StartDate,
		request.EndDate,
		customer,
	)
	panicIfErr(err)

	for _, customerDebt := range customerDebts {
		reportExcel.AddSheet1Data(ReportCustomerDebtExcelSheet1Data{
			Id:                           customerDebt.Id,
			CustomerDebtSource:           customerDebt.DebtSource.String(),
			CustomerDebtSourceIdentifier: customerDebt.DebtSourceId,
			Status:                       customerDebt.Status.String(),
			Amount:                       customerDebt.Amount,
			RemainingAmount:              customerDebt.RemainingAmount,
			CustomerId:                   customerDebt.Customer.Id,
			CustomerName:                 customerDebt.Customer.Name,
			CreatedAt:                    customerDebt.CreatedAt.Time(),
		})
	}

	readCloser, contentLength, err := reportExcel.ToReadSeekCloserWithContentLength()
	panicIfErr(err)

	return readCloser, contentLength, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "customer_debt.xlsx"
}

func (u *customerDebtUseCase) Fetch(ctx context.Context, request dto_request.CustomerDebtFetchRequest) ([]model.CustomerDebt, int) {
	queryOption := model.CustomerDebtQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Status:     request.Status,
		CustomerId: request.CustomerId,
		Phrase:     request.Phrase,
	}

	customerDebts, err := u.repositoryManager.CustomerDebtRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CustomerDebtRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadCustomerDebtsData(ctx, util.SliceValueToSlicePointer(customerDebts), customerDebtLoaderParams{
		customerPayments: false,
		customer:         true,
		deliveryOrder:    true,
	})

	return customerDebts, total
}

func (u *customerDebtUseCase) Get(ctx context.Context, request dto_request.CustomerDebtGetRequest) model.CustomerDebt {
	customerDebt := mustGetCustomerDebt(ctx, u.repositoryManager, request.CustomerDebtId, true)

	u.mustLoadCustomerDebtsData(ctx, []*model.CustomerDebt{&customerDebt}, customerDebtLoaderParams{
		customerPayments: true,
		customer:         true,
		deliveryOrder:    true,
	})

	return customerDebt
}

func (u *customerDebtUseCase) Payment(ctx context.Context, request dto_request.CustomerDebtPaymentRequest) model.CustomerDebt {
	var (
		currentTime  = util.CurrentDateTime()
		authUser     = model.MustGetUserCtx(ctx)
		customerDebt = mustGetCustomerDebt(ctx, u.repositoryManager, request.CustomerDebtId, true)
	)

	if customerDebt.Status != data_type.CustomerDebtStatusUnpaid {
		panic(dto_response.NewBadRequestErrorResponse("CUSTOMER_DEBT.STATUS_MUST_BE_UNPAID"))
	}

	if customerDebt.RemainingAmount < request.Amount {
		panic(dto_response.NewBadRequestErrorResponse("CUSTOMER_DEBT.PAYMENT_INVALID_AMOUNT"))
	}

	// change status to paid and change remaining amount
	customerDebt.RemainingAmount -= request.Amount
	if customerDebt.RemainingAmount == 0 {
		customerDebt.Status = data_type.CustomerDebtStatusPaid
	} else {
		customerDebt.Status = data_type.CustomerDebtStatusHalfPaid
	}

	// initialize customer payment
	customerPayment := model.CustomerPayment{
		Id:             util.NewUuid(),
		UserId:         authUser.Id,
		ImageFileId:    "",
		CustomerDebtId: customerDebt.Id,
		Amount:         request.Amount,
		Description:    request.Description,
		PaidAt:         currentTime,
	}

	// upload file
	imageFile := model.File{
		Id:        util.NewUuid(),
		Name:      "",
		Type:      data_type.FileTypeCustomerPaymentImage,
		Path:      "",
		Timestamp: model.Timestamp{},
		Link:      new(string),
	}

	imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
		ctx,
		constant.CustomerPaymentImagePath,
		customerPayment.Id,
		fmt.Sprintf("%s%s", imageFile.Id, path.Ext(request.ImageFilePath)),
		request.ImageFilePath,
		fileUploadTemporaryToMainParams{
			deleteTmpOnSuccess: true,
		},
	)

	customerPayment.ImageFileId = imageFile.Id

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			customerDebtRepository := u.repositoryManager.CustomerDebtRepository()
			customerPaymentRepository := u.repositoryManager.CustomerPaymentRepository()
			fileRepository := u.repositoryManager.FileRepository()

			if err := fileRepository.Insert(ctx, &imageFile); err != nil {
				return err
			}

			if err := customerPaymentRepository.Insert(ctx, &customerPayment); err != nil {
				return err
			}

			if err := customerDebtRepository.Update(ctx, &customerDebt); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadCustomerDebtsData(ctx, []*model.CustomerDebt{&customerDebt}, customerDebtLoaderParams{
		customerPayments: true,
		customer:         true,
		deliveryOrder:    true,
	})

	return customerDebt
}
