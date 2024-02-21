package use_case

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"path"
)

type CustomerDebtUseCase interface {
	// create
	UploadImage(ctx context.Context, request dto_request.CustomerDebtUploadImageRequest) string

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

func (u *customerDebtUseCase) Fetch(ctx context.Context, request dto_request.CustomerDebtFetchRequest) ([]model.CustomerDebt, int) {
	queryOption := model.CustomerDebtQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	customerDebts, err := u.repositoryManager.CustomerDebtRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CustomerDebtRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return customerDebts, total
}

func (u *customerDebtUseCase) Get(ctx context.Context, request dto_request.CustomerDebtGetRequest) model.CustomerDebt {
	customerDebt := mustGetCustomerDebt(ctx, u.repositoryManager, request.CustomerDebtId, true)

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
		Type:      0,
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

	return customerDebt
}