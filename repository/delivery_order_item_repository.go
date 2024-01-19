package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderItemRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderItem *model.DeliveryOrderItem) error
	InsertMany(ctx context.Context, deliveryOrderItems []model.DeliveryOrderItem, options ...data_type.RepositoryOption) error

	// read
	FetchByProductReceiveIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderItem, error)
	Get(ctx context.Context, id string) (*model.DeliveryOrderItem, error)
	GetByDeliveryOrderIdAndProductUnitId(ctx context.Context, deliveryOrderId string, productUnitId string) (*model.DeliveryOrderItem, error)

	// update
	Update(ctx context.Context, deliveryOrderItem *model.DeliveryOrderItem) error

	// delete
	Delete(ctx context.Context, deliveryOrderItem *model.DeliveryOrderItem) error
	DeleteManyByDeliveryOrderId(ctx context.Context, deliveryOrderId string) error
	Truncate(ctx context.Context) error
}

type deliveryOrderItemRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDeliveryOrderItemRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderItemRepository {
	return &deliveryOrderItemRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *deliveryOrderItemRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderItem, error) {
	deliveryOrderItems := []model.DeliveryOrderItem{}
	if err := fetch(r.db, ctx, &deliveryOrderItems, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderItems, nil
}

func (r *deliveryOrderItemRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderItem, error) {
	deliveryOrderItem := model.DeliveryOrderItem{}
	if err := get(r.db, ctx, &deliveryOrderItem, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderItem, nil
}

func (r *deliveryOrderItemRepository) Insert(ctx context.Context, deliveryOrderItem *model.DeliveryOrderItem) error {
	return defaultInsert(r.db, ctx, deliveryOrderItem, "*")
}

func (r *deliveryOrderItemRepository) InsertMany(ctx context.Context, deliveryOrderItems []model.DeliveryOrderItem, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderItems {
		arr = append(arr, &deliveryOrderItems[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderItemRepository) FetchByProductReceiveIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderItemTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderIds})

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderItemRepository) Get(ctx context.Context, id string) (*model.DeliveryOrderItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderItemTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderItemRepository) GetByDeliveryOrderIdAndProductUnitId(ctx context.Context, productReceiveId string, productUnitId string) (*model.DeliveryOrderItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderItemTableName).
		Where(squirrel.Eq{"delivery_order_id": productReceiveId}).
		Where(squirrel.Eq{"product_unit_id": productUnitId})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderItemRepository) Update(ctx context.Context, deliveryOrderItem *model.DeliveryOrderItem) error {
	return defaultUpdate(r.db, ctx, deliveryOrderItem, "*", nil)
}

func (r *deliveryOrderItemRepository) Delete(ctx context.Context, deliveryOrderItem *model.DeliveryOrderItem) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, deliveryOrderItem, nil)
}

func (r *deliveryOrderItemRepository) DeleteManyByDeliveryOrderId(ctx context.Context, deliveryOrderId string) error {
	whereStmt := squirrel.Eq{
		"delivery_order_id": deliveryOrderId,
	}
	return destroy(r.db, ctx, model.DeliveryOrderItemTableName, whereStmt)
}

func (r *deliveryOrderItemRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderItemTableName)
}
