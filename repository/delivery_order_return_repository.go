package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderReturnRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderReturn *model.DeliveryOrderReturn) error
	InsertMany(ctx context.Context, deliveryOrderReturns []model.DeliveryOrderReturn, options ...data_type.RepositoryOption) error

	// read
	FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderReturn, error)

	// update
	Update(ctx context.Context, deliveryOrderReturn *model.DeliveryOrderReturn) error

	// delete
	Truncate(ctx context.Context) error
}

type deliveryOrderReturnRepository struct {
	db infrastructure.DBTX
}

func NewDeliveryOrderReturnRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderReturnRepository {
	return &deliveryOrderReturnRepository{
		db: db,
	}
}

func (r *deliveryOrderReturnRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderReturn, error) {
	deliveryOrderReturns := []model.DeliveryOrderReturn{}
	if err := fetch(r.db, ctx, &deliveryOrderReturns, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderReturns, nil
}

func (r *deliveryOrderReturnRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderReturn, error) {
	deliveryOrderReturn := model.DeliveryOrderReturn{}
	if err := get(r.db, ctx, &deliveryOrderReturn, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderReturn, nil
}

func (r *deliveryOrderReturnRepository) Insert(ctx context.Context, deliveryOrderReturn *model.DeliveryOrderReturn) error {
	return defaultInsert(r.db, ctx, deliveryOrderReturn, "*")
}

func (r *deliveryOrderReturnRepository) InsertMany(ctx context.Context, deliveryOrderReturns []model.DeliveryOrderReturn, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderReturns {
		arr = append(arr, &deliveryOrderReturns[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderReturnRepository) FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderReturn, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderReturnTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderIds})

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderReturnRepository) Update(ctx context.Context, deliveryOrderReturn *model.DeliveryOrderReturn) error {
	return defaultUpdate(r.db, ctx, deliveryOrderReturn, "*", nil)
}

func (r *deliveryOrderReturnRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderReturnTableName)
}
