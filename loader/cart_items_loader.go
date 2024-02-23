package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type CartItemsLoader struct {
	loaderByCartId dataloader.Loader
}

func (l *CartItemsLoader) loadByCartId(id string) ([]model.CartItem, error) {
	thunk := l.loaderByCartId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.CartItem), nil
}

func (l *CartItemsLoader) CartFn(cart *model.Cart) func() error {
	return func() error {
		if cart != nil {
			cartItems, err := l.loadByCartId(cart.Id)
			if err != nil {
				return err
			}

			cart.CartItems = cartItems
		}

		return nil
	}
}

func NewCartItemsLoader(cartItemRepository repository.CartItemRepository) *CartItemsLoader {
	batchByCartIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		cartItems, err := cartItemRepository.FetchByCartIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		cartItemByCartId := map[string][]model.CartItem{}
		for _, cartItem := range cartItems {
			cartItemByCartId[cartItem.CartId] = append(cartItemByCartId[cartItem.CartId], cartItem)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var userRoles []model.CartItem
			if v, ok := cartItemByCartId[k.String()]; ok {
				userRoles = v
			}

			result := &dataloader.Result{Data: userRoles, Error: nil}
			if userRoles == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &CartItemsLoader{
		loaderByCartId: NewDataloader(batchByCartIdFn),
	}
}
