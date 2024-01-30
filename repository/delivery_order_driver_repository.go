package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderDriverRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderDriver *model.DeliveryOrderDriver) error
	InsertMany(ctx context.Context, deliveryOrderDrivers []model.DeliveryOrderDriver, options ...data_type.RepositoryOption) error

	// read
	FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderDriver, error)
	Get(ctx context.Context, id string) (*model.DeliveryOrderDriver, error)
	GetByDeliveryOrderIdAndDriverUserId(ctx context.Context, deliveryOrderId string, driverUserId string) (*model.DeliveryOrderDriver, error)
	IsExistByDeliveryOrderIdAndDriverUserId(ctx context.Context, deliveryOrderId string, driverUserId string) (bool, error)

	// delete
	Delete(ctx context.Context, deliveryOrderDriver *model.DeliveryOrderDriver) error
	DeleteManyByDeliveryOrderId(ctx context.Context, deliveryOrderId string) error
	Truncate(ctx context.Context) error
}

type deliveryOrderDriverRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDeliveryOrderDriverRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderDriverRepository {
	return &deliveryOrderDriverRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *deliveryOrderDriverRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderDriver, error) {
	deliveryOrderDrivers := []model.DeliveryOrderDriver{}
	if err := fetch(r.db, ctx, &deliveryOrderDrivers, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderDrivers, nil
}

func (r *deliveryOrderDriverRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderDriver, error) {
	deliveryOrderDriver := model.DeliveryOrderDriver{}
	if err := get(r.db, ctx, &deliveryOrderDriver, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderDriver, nil
}

func (r *deliveryOrderDriverRepository) Insert(ctx context.Context, deliveryOrderDriver *model.DeliveryOrderDriver) error {
	return defaultInsert(r.db, ctx, deliveryOrderDriver, "*")
}

func (r *deliveryOrderDriverRepository) InsertMany(ctx context.Context, deliveryOrderDrivers []model.DeliveryOrderDriver, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderDrivers {
		arr = append(arr, &deliveryOrderDrivers[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderDriverRepository) FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderDriver, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderDriverTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderIds})

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderDriverRepository) Get(ctx context.Context, id string) (*model.DeliveryOrderDriver, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderDriverTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderDriverRepository) GetByDeliveryOrderIdAndDriverUserId(ctx context.Context, deliveryOrderId string, driverUserId string) (*model.DeliveryOrderDriver, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderDriverTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderId}).
		Where(squirrel.Eq{"user_id": driverUserId})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderDriverRepository) IsExistByDeliveryOrderIdAndDriverUserId(ctx context.Context, deliveryOrderId string, driverUserId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.DeliveryOrderDriverTableName).
			Where(squirrel.Eq{"delivery_order_id": deliveryOrderId}).
			Where(squirrel.Eq{"driver_user_id": driverUserId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *deliveryOrderDriverRepository) Update(ctx context.Context, deliveryOrderDriver *model.DeliveryOrderDriver) error {
	return defaultUpdate(r.db, ctx, deliveryOrderDriver, "*", nil)
}

func (r *deliveryOrderDriverRepository) Delete(ctx context.Context, deliveryOrderDriver *model.DeliveryOrderDriver) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, deliveryOrderDriver, nil)
}

func (r *deliveryOrderDriverRepository) DeleteManyByDeliveryOrderId(ctx context.Context, deliveryOrderId string) error {
	whereStmt := squirrel.Eq{
		"delivery_order_id": deliveryOrderId,
	}
	return destroy(r.db, ctx, model.DeliveryOrderDriverTableName, whereStmt)
}

func (r *deliveryOrderDriverRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderDriverTableName)
}
