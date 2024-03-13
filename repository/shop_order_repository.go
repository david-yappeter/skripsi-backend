package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ShopOrderRepository interface {
	// create
	Insert(ctx context.Context, shopOrder *model.ShopOrder) error
	InsertMany(ctx context.Context, shopOrders []model.ShopOrder, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ShopOrderQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ShopOrderQueryOption) ([]model.ShopOrder, error)
	Get(ctx context.Context, id string) (*model.ShopOrder, error)
	GetByPlatformTypeAndPlatformIdentifierId(ctx context.Context, platformType data_type.ShopOrderPlatformType, platformIdentifier string) (*model.ShopOrder, error)

	// update
	Update(ctx context.Context, shopOrder *model.ShopOrder) error

	// delete
	Delete(ctx context.Context, shopOrder *model.ShopOrder) error
	Truncate(ctx context.Context) error
}

type shopOrderRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewShopOrderRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ShopOrderRepository {
	return &shopOrderRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *shopOrderRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ShopOrder, error) {
	shopOrders := []model.ShopOrder{}
	if err := fetch(r.db, ctx, &shopOrders, stmt); err != nil {
		return nil, err
	}

	return shopOrders, nil
}

func (r *shopOrderRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ShopOrder, error) {
	shopOrder := model.ShopOrder{}
	if err := get(r.db, ctx, &shopOrder, stmt); err != nil {
		return nil, err
	}

	return &shopOrder, nil
}

func (r *shopOrderRepository) prepareQuery(option model.ShopOrderQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s so", model.ShopOrderTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"so.tracking_number": phrase},
			squirrel.ILike{"so.recipient_name": phrase},
			squirrel.ILike{"so.recipient_phone_number": phrase},
		})
	}

	if option.PlatformType != nil {
		stmt = stmt.Where(squirrel.Eq{"so.platform_type": option.PlatformType})
	}

	if option.TrackingStatus != nil {
		stmt = stmt.Where(squirrel.Eq{"so.tracking_status": option.TrackingStatus})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *shopOrderRepository) Insert(ctx context.Context, shopOrder *model.ShopOrder) error {
	return defaultInsert(r.db, ctx, shopOrder, "*")
}

func (r *shopOrderRepository) InsertMany(ctx context.Context, shopOrders []model.ShopOrder, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range shopOrders {
		arr = append(arr, &shopOrders[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *shopOrderRepository) Count(ctx context.Context, options ...model.ShopOrderQueryOption) (int, error) {
	option := model.ShopOrderQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *shopOrderRepository) Fetch(ctx context.Context, options ...model.ShopOrderQueryOption) ([]model.ShopOrder, error) {
	option := model.ShopOrderQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *shopOrderRepository) Get(ctx context.Context, id string) (*model.ShopOrder, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ShopOrderTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *shopOrderRepository) GetByPlatformTypeAndPlatformIdentifierId(ctx context.Context, platformType data_type.ShopOrderPlatformType, platformIdentifier string) (*model.ShopOrder, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ShopOrderTableName).
		Where(squirrel.Eq{"platform_type": platformType}).
		Where(squirrel.Eq{"platform_identifier": platformIdentifier})

	return r.get(ctx, stmt)
}

func (r *shopOrderRepository) Update(ctx context.Context, shopOrder *model.ShopOrder) error {
	return defaultUpdate(r.db, ctx, shopOrder, "*", nil)
}

func (r *shopOrderRepository) Delete(ctx context.Context, shopOrder *model.ShopOrder) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, shopOrder, nil)
}

func (r *shopOrderRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ShopOrderTableName)
}
