package use_case

import (
	"context"
	"fmt"
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

type debtLoaderParams struct {
	payments bool
}

type DebtUseCase interface {
	// create
	UploadImage(ctx context.Context, request dto_request.DebtUploadImageRequest) string

	//  read
	Fetch(ctx context.Context, request dto_request.DebtFetchRequest) ([]model.Debt, int)
	Get(ctx context.Context, request dto_request.DebtGetRequest) model.Debt

	//  update
	Payment(ctx context.Context, request dto_request.DebtPaymentRequest) model.Debt
}

type debtUseCase struct {
	repositoryManager repository.RepositoryManager

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client

	baseFileUseCase
}

func NewDebtUseCase(
	repositoryManager repository.RepositoryManager,
	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) DebtUseCase {
	return &debtUseCase{
		repositoryManager: repositoryManager,
		mainFilesystem:    mainFilesystem,
		tmpFilesystem:     tmpFilesystem,
		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
	}
}

func (u *debtUseCase) mustLoadDebtsData(ctx context.Context, debts []*model.Debt, option debtLoaderParams) {
	debtPaymentsLoader := loader.NewDebtPaymentsLoader(u.repositoryManager.DebtPaymentRepository())
	productReceiveLoader := loader.NewProductReceiveLoader(u.repositoryManager.ProductReceiveRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range debts {
			if option.payments {
				group.Go(debtPaymentsLoader.DebtFn(debts[i]))
			}

			switch debts[i].DebtSource {
			case data_type.DebtSourceProductReceive:
				group.Go(productReceiveLoader.DebtFn(debts[i]))
			}
		}
	}))

	supplierLoader := loader.NewSupplierLoader(u.repositoryManager.SupplierRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range debts {
			if option.payments {
				for j := range debts[i].DebtPayments {
					group.Go(fileLoader.DebtPaymentFn(&debts[i].DebtPayments[j]))
				}
			}

			if debts[i].ProductReceive != nil {
				group.Go(supplierLoader.ProductReceiveFn(debts[i].ProductReceive))
			}
		}
	}))

	for i := range debts {
		for j := range debts[i].DebtPayments {
			file := debts[i].DebtPayments[j].ImageFile
			if file != nil {
				file.SetLink(u.mainFilesystem)
			}
		}
	}
}

func (u *debtUseCase) UploadImage(ctx context.Context, request dto_request.DebtUploadImageRequest) string {
	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.DebtPaymentImagePath,
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

func (u *debtUseCase) Fetch(ctx context.Context, request dto_request.DebtFetchRequest) ([]model.Debt, int) {
	queryOption := model.DebtQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Status: request.Status,
		Phrase: request.Phrase,
	}

	debts, err := u.repositoryManager.DebtRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.DebtRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadDebtsData(ctx, util.SliceValueToSlicePointer(debts), debtLoaderParams{})

	return debts, total
}

func (u *debtUseCase) Get(ctx context.Context, request dto_request.DebtGetRequest) model.Debt {
	debt := mustGetDebt(ctx, u.repositoryManager, request.DebtId, true)

	u.mustLoadDebtsData(ctx, []*model.Debt{&debt}, debtLoaderParams{
		payments: true,
	})

	return debt
}

func (u *debtUseCase) Payment(ctx context.Context, request dto_request.DebtPaymentRequest) model.Debt {
	var (
		currentTime = util.CurrentDateTime()
		authUser    = model.MustGetUserCtx(ctx)
		debt        = mustGetDebt(ctx, u.repositoryManager, request.DebtId, true)
	)

	if debt.Status != data_type.DebtStatusUnpaid {
		panic(dto_response.NewBadRequestErrorResponse("DEBT.STATUS_MUST_BE_UNPAID"))
	}

	if debt.RemainingAmount < request.Amount {
		panic(dto_response.NewBadRequestErrorResponse("DEBT.PAYMENT_INVALID_AMOUNT"))
	}

	// change status to paid and change remaining amount
	debt.RemainingAmount -= request.Amount
	if debt.RemainingAmount == 0 {
		debt.Status = data_type.DebtStatusPaid
	} else {
		debt.Status = data_type.DebtStatusHalfPaid
	}

	// initialize debt payment
	debtPayment := model.DebtPayment{
		Id:          util.NewUuid(),
		UserId:      authUser.Id,
		ImageFileId: "",
		DebtId:      debt.Id,
		Amount:      request.Amount,
		Description: request.Description,
		PaidAt:      currentTime,
	}

	// upload file
	imageFile := model.File{
		Id:   util.NewUuid(),
		Name: "",
		Type: data_type.FileTypeDebtPaymentImage,
		Path: "",
		Link: new(string),
	}

	imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
		ctx,
		constant.DebtPaymentImagePath,
		debtPayment.Id,
		fmt.Sprintf("%s%s", imageFile.Id, path.Ext(request.ImageFilePath)),
		request.ImageFilePath,
		fileUploadTemporaryToMainParams{
			deleteTmpOnSuccess: true,
		},
	)

	debtPayment.ImageFileId = imageFile.Id

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			debtRepository := u.repositoryManager.DebtRepository()
			debtPaymentRepository := u.repositoryManager.DebtPaymentRepository()
			fileRepository := u.repositoryManager.FileRepository()

			if err := fileRepository.Insert(ctx, &imageFile); err != nil {
				return err
			}

			if err := debtPaymentRepository.Insert(ctx, &debtPayment); err != nil {
				return err
			}

			if err := debtRepository.Update(ctx, &debt); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadDebtsData(ctx, []*model.Debt{&debt}, debtLoaderParams{
		payments: true,
	})

	return debt
}
