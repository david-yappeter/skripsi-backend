package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ShopOrderItemRepository interface {
	// create
	Insert(ctx context.Context, shopOrderItem *model.ShopOrderItem) error
	InsertMany(ctx context.Context, shopOrderItems []model.ShopOrderItem, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ShopOrderItemQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ShopOrderItemQueryOption) ([]model.ShopOrderItem, error)
	Get(ctx context.Context, id string) (*model.ShopOrderItem, error)

	// update
	Update(ctx context.Context, shopOrderItem *model.ShopOrderItem) error

	// delete
	Delete(ctx context.Context, shopOrderItem *model.ShopOrderItem) error
	Truncate(ctx context.Context) error
}

type shopOrderItemRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewShopOrderItemRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ShopOrderItemRepository {
	return &shopOrderItemRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *shopOrderItemRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ShopOrderItem, error) {
	shopOrderItems := []model.ShopOrderItem{}
	if err := fetch(r.db, ctx, &shopOrderItems, stmt); err != nil {
		return nil, err
	}

	return shopOrderItems, nil
}

func (r *shopOrderItemRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ShopOrderItem, error) {
	shopOrderItem := model.ShopOrderItem{}
	if err := get(r.db, ctx, &shopOrderItem, stmt); err != nil {
		return nil, err
	}

	return &shopOrderItem, nil
}

func (r *shopOrderItemRepository) prepareQuery(option model.ShopOrderItemQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s so", model.ShopOrderItemTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"so.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *shopOrderItemRepository) Insert(ctx context.Context, shopOrderItem *model.ShopOrderItem) error {
	return defaultInsert(r.db, ctx, shopOrderItem, "*")
}

func (r *shopOrderItemRepository) InsertMany(ctx context.Context, shopOrderItems []model.ShopOrderItem, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range shopOrderItems {
		arr = append(arr, &shopOrderItems[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *shopOrderItemRepository) Count(ctx context.Context, options ...model.ShopOrderItemQueryOption) (int, error) {
	option := model.ShopOrderItemQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *shopOrderItemRepository) Fetch(ctx context.Context, options ...model.ShopOrderItemQueryOption) ([]model.ShopOrderItem, error) {
	option := model.ShopOrderItemQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *shopOrderItemRepository) Get(ctx context.Context, id string) (*model.ShopOrderItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ShopOrderItemTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *shopOrderItemRepository) Update(ctx context.Context, shopOrderItem *model.ShopOrderItem) error {
	return defaultUpdate(r.db, ctx, shopOrderItem, "*", nil)
}

func (r *shopOrderItemRepository) Delete(ctx context.Context, shopOrderItem *model.ShopOrderItem) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, shopOrderItem, nil)
}

func (r *shopOrderItemRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ShopOrderItemTableName)
}
