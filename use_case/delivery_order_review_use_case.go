package use_case

import (
	"context"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type deliveryOrderReviewLoaderParams struct {
}

type DeliveryOrderReviewUseCase interface {
	// create
	CreateGuest(ctx context.Context, request dto_request.DeliveryOrderReviewCreateGuestRequest) model.DeliveryOrderReview

	//  read
	Fetch(ctx context.Context, request dto_request.DeliveryOrderReviewFetchRequest) ([]model.DeliveryOrderReview, int)
	Get(ctx context.Context, request dto_request.DeliveryOrderReviewGetRequest) model.DeliveryOrderReview
	IsExistByDeliveryOrder(ctx context.Context, request dto_request.DeliveryOrderReviewIsExistByDeliveryOrderRequest) bool
}

type deliveryOrderReviewUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewDeliveryOrderReviewUseCase(
	repositoryManager repository.RepositoryManager,
) DeliveryOrderReviewUseCase {
	return &deliveryOrderReviewUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *deliveryOrderReviewUseCase) mustLoadDeliveryOrderReviewsData(ctx context.Context, deliveryOrderReviews []*model.DeliveryOrderReview, option deliveryOrderReviewLoaderParams) {
	deliveryOrderLoader := loader.NewDeliveryOrderLoader(u.repositoryManager.DeliveryOrderRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range deliveryOrderReviews {
			group.Go(deliveryOrderLoader.DeliveryOrderReviewFn(deliveryOrderReviews[i]))
		}
	}))

	customerLoader := loader.NewCustomerLoader(u.repositoryManager.CustomerRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range deliveryOrderReviews {
			group.Go(customerLoader.DeliveryOrderFn(deliveryOrderReviews[i].DeliveryOrder))
		}
	}))
}

func (u *deliveryOrderReviewUseCase) CreateGuest(ctx context.Context, request dto_request.DeliveryOrderReviewCreateGuestRequest) model.DeliveryOrderReview {
	// currentDateTime := util.CurrentDateTime()
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusCompleted {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER_REVIEW.DELIVERY_ORDER_MUST_BE_COMPLETED"))
	}

	// if deliveryOrder.UpdatedAt.Add(time.Hour * 24 * 2).IsLessThan(currentDateTime) {
	// 	panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER_REVIEW.PASS_2_DAYS_CANNOT_REVIEW_ANYMORE"))
	// }

	isExist, err := u.repositoryManager.DeliveryOrderReviewRepository().IsExistByDeliveryOrderId(ctx, request.DeliveryOrderId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER_REVIEW.ALREADY_EXIST"))
	}

	deliveryOrderReview := model.DeliveryOrderReview{
		Id:              util.NewUuid(),
		DeliveryOrderId: request.DeliveryOrderId,
		StarRating:      request.StarRating,
		Description:     request.Description,
	}

	panicIfErr(
		u.repositoryManager.DeliveryOrderReviewRepository().Insert(ctx, &deliveryOrderReview),
	)

	return deliveryOrderReview
}

func (u *deliveryOrderReviewUseCase) Fetch(ctx context.Context, request dto_request.DeliveryOrderReviewFetchRequest) ([]model.DeliveryOrderReview, int) {
	queryOption := model.DeliveryOrderReviewQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		StarRating: request.StarRating,
	}

	deliveryOrderReviews, err := u.repositoryManager.DeliveryOrderReviewRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.DeliveryOrderReviewRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return deliveryOrderReviews, total
}

func (u *deliveryOrderReviewUseCase) Get(ctx context.Context, request dto_request.DeliveryOrderReviewGetRequest) model.DeliveryOrderReview {
	deliveryOrderReview := mustGetDeliveryOrderReview(ctx, u.repositoryManager, request.DeliveryOrderReviewId, true)

	u.mustLoadDeliveryOrderReviewsData(ctx, []*model.DeliveryOrderReview{&deliveryOrderReview}, deliveryOrderReviewLoaderParams{})

	return deliveryOrderReview
}

func (u *deliveryOrderReviewUseCase) IsExistByDeliveryOrder(ctx context.Context, request dto_request.DeliveryOrderReviewIsExistByDeliveryOrderRequest) bool {
	isExist, err := u.repositoryManager.DeliveryOrderReviewRepository().IsExistByDeliveryOrderId(ctx, request.DeliveryOrderId)
	panicIfErr(err)

	return isExist
}
