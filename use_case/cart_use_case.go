package use_case

import (
	"context"
	"myapp/constant"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type cartLoaderParams struct {
	items bool
}

type CartUseCase interface {
	//  create
	AddItem(ctx context.Context, request dto_request.CartAddItemRequest) model.Cart

	//  read
	Fetch(ctx context.Context, request dto_request.CartFetchRequest) ([]model.Cart, int)
	FetchInActive(ctx context.Context, request dto_request.CartFetchInActiveRequest) []model.Cart
	Get(ctx context.Context, request dto_request.CartGetRequest) model.Cart
	GetCurrent(ctx context.Context) *model.Cart

	//  update
	UpdateItem(ctx context.Context, request dto_request.CartUpdateItemRequest) model.Cart
	SetActive(ctx context.Context, request dto_request.CartSetActiveRequest) model.Cart
	SetInActive(ctx context.Context) model.Cart

	//  delete
	Delete(ctx context.Context, request dto_request.CartDeleteRequest)
	DeleteItem(ctx context.Context, request dto_request.CartDeleteItemRequest) model.Cart
}

type cartUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewCartUseCase(
	repositoryManager repository.RepositoryManager,
) CartUseCase {
	return &cartUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *cartUseCase) mustGetCurrentUserActiveCashierSession(ctx context.Context) model.CashierSession {
	authUser := model.MustGetUserCtx(ctx)

	cashierSession, err := u.repositoryManager.CashierSessionRepository().GetByUserIdAndStatusActive(ctx, authUser.Id)
	panicIfErr(err, constant.ErrNoData)

	if cashierSession == nil {
		panic(dto_response.NewBadRequestErrorResponse("CART.USER_MUST_HAVE_ACTIVE_CASHIER_SESSION"))
	}

	return *cashierSession
}

func (u *cartUseCase) mustGetActiveCartByCashierSessionId(ctx context.Context, cashierSessionId string) model.Cart {
	cart := u.shouldGetActiveCartByCashierSessionId(ctx, cashierSessionId)

	if cart == nil {
		panic(dto_response.NewBadRequestErrorResponse("CART.ACTIVE_NOT_FOUND"))
	}

	return *cart
}

func (u *cartUseCase) mustGetCartItemByCartIdAndProductUnitId(ctx context.Context, cartId string, productUnitId string) model.CartItem {
	cartItem := u.shouldGetCartItemByCartIdAndProductUnitId(ctx, cartId, productUnitId)

	if cartItem == nil {
		panic(dto_response.NewBadRequestErrorResponse("CART.CART_ITEM_NOT_FOUND"))
	}

	return *cartItem
}

func (u *cartUseCase) mustLoadCartDatas(ctx context.Context, carts []*model.Cart, option cartLoaderParams) {
	cartItemsLoader := loader.NewCartItemsLoader(u.repositoryManager.CartItemRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range carts {
				if option.items {
					group.Go(cartItemsLoader.CartFn(carts[i]))
				}
			}
		}),
	)

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range carts {
				if option.items {
					for j := range carts[i].CartItems {
						group.Go(productUnitLoader.CartItemFn(&carts[i].CartItems[j]))
					}
				}
			}
		}),
	)
}

func (u *cartUseCase) shouldGetActiveCartByCashierSessionId(ctx context.Context, cashierSessionId string) *model.Cart {
	cart, err := u.repositoryManager.CartRepository().GetByCashierSessionIdAndIsActive(ctx, cashierSessionId, true)
	panicIfErr(err, constant.ErrNoData)

	return cart
}

func (u *cartUseCase) shouldGetCartItemByCartIdAndProductUnitId(ctx context.Context, cartId string, productUnitId string) *model.CartItem {
	cartItem, err := u.repositoryManager.CartItemRepository().GetByCartIdAndProductUnitId(ctx, cartId, productUnitId)
	panicIfErr(err, constant.ErrNoData)

	return cartItem
}

func (u *cartUseCase) AddItem(ctx context.Context, request dto_request.CartAddItemRequest) model.Cart {
	var (
		cartItem *model.CartItem
	)

	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	cart := u.shouldGetActiveCartByCashierSessionId(ctx, cashierSession.Id)
	isNewCart := cart == nil
	isNewCartItem := false

	if isNewCart {
		cart = &model.Cart{
			Id:               util.NewUuid(),
			CashierSessionId: cashierSession.Id,
			Name:             nil,
			IsActive:         true,
		}

		cartItem = &model.CartItem{
			Id:            util.NewUuid(),
			CartId:        cart.Id,
			ProductUnitId: request.ProductUnitId,
			Qty:           request.Qty,
		}
	} else {
		cartItem = u.shouldGetCartItemByCartIdAndProductUnitId(ctx, cart.Id, request.ProductUnitId)

		if cartItem == nil {
			cartItem = &model.CartItem{
				Id:            util.NewUuid(),
				CartId:        cart.Id,
				ProductUnitId: request.ProductUnitId,
				Qty:           request.Qty,
			}
		} else {
			cartItem.Qty += request.Qty
		}
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			cartRepository := u.repositoryManager.CartRepository()
			cartItemRepository := u.repositoryManager.CartItemRepository()

			if isNewCart {
				if err := cartRepository.Insert(ctx, cart); err != nil {
					return err
				}
			}

			if isNewCartItem {
				if err := cartItemRepository.Insert(ctx, cartItem); err != nil {
					return err
				}
			} else {
				if err := cartItemRepository.Update(ctx, cartItem); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	u.mustLoadCartDatas(ctx, []*model.Cart{cart}, cartLoaderParams{
		items: true,
	})

	return *cart
}

func (u *cartUseCase) Fetch(ctx context.Context, request dto_request.CartFetchRequest) ([]model.Cart, int) {
	queryOption := model.CartQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	carts, err := u.repositoryManager.CartRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CartRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return carts, total
}

func (u *cartUseCase) FetchInActive(ctx context.Context, request dto_request.CartFetchInActiveRequest) []model.Cart {
	queryOption := model.CartQueryOption{
		IsActive: util.BoolP(false),
		Phrase:   request.Phrase,
	}

	carts, err := u.repositoryManager.CartRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	return carts
}

func (u *cartUseCase) Get(ctx context.Context, request dto_request.CartGetRequest) model.Cart {
	cart := mustGetCart(ctx, u.repositoryManager, request.CartId, true)

	u.mustLoadCartDatas(ctx, []*model.Cart{&cart}, cartLoaderParams{
		items: true,
	})

	return cart
}

func (u *cartUseCase) GetCurrent(ctx context.Context) *model.Cart {
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	cart := u.shouldGetActiveCartByCashierSessionId(ctx, cashierSession.Id)

	u.mustLoadCartDatas(ctx, []*model.Cart{cart}, cartLoaderParams{
		items: true,
	})

	return cart
}

func (u *cartUseCase) UpdateItem(ctx context.Context, request dto_request.CartUpdateItemRequest) model.Cart {
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	cart := u.mustGetActiveCartByCashierSessionId(ctx, cashierSession.Id)
	cartItem := u.mustGetCartItemByCartIdAndProductUnitId(ctx, cart.Id, request.ProductUnitId)

	cartItem.Qty = request.Qty

	panicIfErr(
		u.repositoryManager.CartRepository().Update(ctx, &cart),
	)

	u.mustLoadCartDatas(ctx, []*model.Cart{&cart}, cartLoaderParams{
		items: true,
	})

	return cart
}

func (u *cartUseCase) SetActive(ctx context.Context, request dto_request.CartSetActiveRequest) model.Cart {
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	activeCart := u.shouldGetActiveCartByCashierSessionId(ctx, cashierSession.Id)

	if activeCart != nil {
		panic(dto_response.NewBadRequestErrorResponse("CART.THERE_IS_AN_ACTIVE_CART"))
	}

	cart := mustGetCart(ctx, u.repositoryManager, request.CartId, false)

	if cart.CashierSessionId != cashierSession.Id {
		panic(dto_response.NewBadRequestErrorResponse("CART.NOT_FOUND"))
	}

	cart.IsActive = true

	panicIfErr(
		u.repositoryManager.CartRepository().Update(ctx, &cart),
	)

	u.mustLoadCartDatas(ctx, []*model.Cart{&cart}, cartLoaderParams{
		items: true,
	})

	return cart
}

func (u *cartUseCase) SetInActive(ctx context.Context) model.Cart {
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	cart := u.mustGetActiveCartByCashierSessionId(ctx, cashierSession.Id)

	cart.IsActive = false

	panicIfErr(
		u.repositoryManager.CartRepository().Update(ctx, &cart),
	)

	u.mustLoadCartDatas(ctx, []*model.Cart{&cart}, cartLoaderParams{
		items: true,
	})

	return cart
}

func (u *cartUseCase) Delete(ctx context.Context, request dto_request.CartDeleteRequest) {
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	cart := u.mustGetActiveCartByCashierSessionId(ctx, cashierSession.Id)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			cartRepository := u.repositoryManager.CartRepository()
			cartItemRepository := u.repositoryManager.CartItemRepository()

			if err := cartItemRepository.DeleteManyByCartId(ctx, cart.Id); err != nil {
				return err
			}

			if err := cartRepository.Delete(ctx, &cart); err != nil {
				return err
			}

			return nil
		}),
	)
}

func (u *cartUseCase) DeleteItem(ctx context.Context, request dto_request.CartDeleteItemRequest) model.Cart {
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	cart := u.mustGetActiveCartByCashierSessionId(ctx, cashierSession.Id)
	cartItem := u.mustGetCartItemByCartIdAndProductUnitId(ctx, cart.Id, request.ProductUnitId)

	cartItemCount, err := u.repositoryManager.CartItemRepository().Count(ctx, model.CartItemQueryOption{
		CartId: &cart.Id,
	})
	panicIfErr(err)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			cartRepository := u.repositoryManager.CartRepository()
			cartItemRepository := u.repositoryManager.CartItemRepository()

			if err := cartItemRepository.Delete(ctx, &cartItem); err != nil {
				return err
			}

			if cartItemCount == 1 {
				if err := cartRepository.Delete(ctx, &cart); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	u.mustLoadCartDatas(ctx, []*model.Cart{&cart}, cartLoaderParams{
		items: true,
	})

	return cart
}
