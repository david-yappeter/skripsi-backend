package repository

import (
	"context"
	"fmt"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CartItemRepository interface {
	// create
	Insert(ctx context.Context, cartItem *model.CartItem) error

	// read
	Count(ctx context.Context, options ...model.CartItemQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CartItemQueryOption) ([]model.CartItem, error)
	Get(ctx context.Context, id string) (*model.CartItem, error)
	GetByCartIdAndProductUnitId(ctx context.Context, cartId string, productUnitId string) (*model.CartItem, error)

	// update
	Update(ctx context.Context, cartItem *model.CartItem) error

	// delete
	Delete(ctx context.Context, cartItem *model.CartItem) error
	Truncate(ctx context.Context) error
}

type cartItemRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCartItemRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CartItemRepository {
	return &cartItemRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *cartItemRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.CartItem, error) {
	cartItems := []model.CartItem{}
	if err := fetch(r.db, ctx, &cartItems, stmt); err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (r *cartItemRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.CartItem, error) {
	cartItem := model.CartItem{}
	if err := get(r.db, ctx, &cartItem, stmt); err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (r *cartItemRepository) prepareQuery(option model.CartItemQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.CartItemTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *cartItemRepository) Insert(ctx context.Context, cartItem *model.CartItem) error {
	return defaultInsert(r.db, ctx, cartItem, "*")
}

func (r *cartItemRepository) Count(ctx context.Context, options ...model.CartItemQueryOption) (int, error) {
	option := model.CartItemQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *cartItemRepository) Fetch(ctx context.Context, options ...model.CartItemQueryOption) ([]model.CartItem, error) {
	option := model.CartItemQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *cartItemRepository) Get(ctx context.Context, id string) (*model.CartItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CartItemTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *cartItemRepository) GetByCartIdAndProductUnitId(ctx context.Context, cartId string, productUnitId string) (*model.CartItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CartItemTableName).
		Where(squirrel.Eq{"cart_id": cartId}).
		Where(squirrel.Eq{"product_unit_id": productUnitId})

	return r.get(ctx, stmt)
}

func (r *cartItemRepository) Update(ctx context.Context, cartItem *model.CartItem) error {
	return defaultUpdate(r.db, ctx, cartItem, "*", nil)
}

func (r *cartItemRepository) Delete(ctx context.Context, cartItem *model.CartItem) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, cartItem, nil)
}

func (r *cartItemRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CartItemTableName)
}
