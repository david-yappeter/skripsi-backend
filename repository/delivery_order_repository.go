package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrder *model.DeliveryOrder) error
	InsertMany(ctx context.Context, deliveryOrders []model.DeliveryOrder, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.DeliveryOrderQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.DeliveryOrderQueryOption) ([]model.DeliveryOrder, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.DeliveryOrder, error)
	Get(ctx context.Context, id string) (*model.DeliveryOrder, error)
	GetByDriverUserIdAndStatusDelivering(ctx context.Context, driverUserId string) (*model.DeliveryOrder, error)
	IsExistByDriverUserIdAndStatusDelivering(ctx context.Context, driverUserId string) (bool, error)
	IsExistByCustomerId(ctx context.Context, customerId string) (bool, error)

	// update
	Update(ctx context.Context, deliveryOrder *model.DeliveryOrder) error

	// delete
	Delete(ctx context.Context, deliveryOrder *model.DeliveryOrder) error
	Truncate(ctx context.Context) error
}

type deliveryOrderRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDeliveryOrderRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderRepository {
	return &deliveryOrderRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *deliveryOrderRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrder, error) {
	deliveryOrders := []model.DeliveryOrder{}
	if err := fetch(r.db, ctx, &deliveryOrders, stmt); err != nil {
		return nil, err
	}

	return deliveryOrders, nil
}

func (r *deliveryOrderRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrder, error) {
	deliveryOrder := model.DeliveryOrder{}
	if err := get(r.db, ctx, &deliveryOrder, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrder, nil
}

func (r *deliveryOrderRepository) prepareQuery(option model.DeliveryOrderQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s dorder", model.DeliveryOrderTableName)).
		InnerJoin(fmt.Sprintf("%s c ON c.id = dorder.customer_id", model.CustomerTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"

		stmt = stmt.Where(
			squirrel.Or{
				squirrel.ILike{"c.name": phrase},
				squirrel.ILike{"dorder.invoice_number": phrase},
			},
		)
	}

	if !option.IsCount && option.SortStatusImportance {
		stmt = stmt.OrderBy(`CASE dorder.status
WHEN 'DELIVERING' THEN 1
WHEN 'ONGOING' THEN 2
WHEN 'COMPLETED' THEN 3
WHEN 'CANCELED' THEN 4
ELSE 999
END`,
		)
	}

	if option.DriverUserId != nil {
		stmt = stmt.Where(
			stmtBuilder.Select("1").
				From(fmt.Sprintf("%s dod", model.DeliveryOrderDriverTableName)).
				Where(squirrel.Expr("dorder.id = dod.delivery_order_id")).
				Where(squirrel.Eq{"dod.driver_user_id": option.DriverUserId}).
				Prefix("EXISTS (").Suffix(")"),
		)
	}

	if option.Status != nil {
		stmt = stmt.Where(squirrel.Eq{"dorder.status": option.Status})
	}

	if option.CustomerId != nil {
		stmt = stmt.Where(squirrel.Eq{
			"dorder.customer_id": option.CustomerId,
		})
	}

	if len(option.ExcludeStatuses) > 0 {
		stmt = stmt.Where(squirrel.NotEq{"dorder.status": option.ExcludeStatuses})
	}

	if option.StartDateTime.DateTimeP() != nil {
		stmt = stmt.Where(squirrel.GtOrEq{"dorder.date": option.StartDateTime})
	}

	if option.EndDateTime.DateTimeP() != nil {
		stmt = stmt.Where(squirrel.LtOrEq{"dorder.date": option.EndDateTime})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *deliveryOrderRepository) Insert(ctx context.Context, deliveryOrder *model.DeliveryOrder) error {
	return defaultInsert(r.db, ctx, deliveryOrder, "*")
}

func (r *deliveryOrderRepository) InsertMany(ctx context.Context, deliveryOrders []model.DeliveryOrder, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrders {
		arr = append(arr, &deliveryOrders[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderRepository) Count(ctx context.Context, options ...model.DeliveryOrderQueryOption) (int, error) {
	option := model.DeliveryOrderQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *deliveryOrderRepository) Fetch(ctx context.Context, options ...model.DeliveryOrderQueryOption) ([]model.DeliveryOrder, error) {
	option := model.DeliveryOrderQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderRepository) FetchByIds(ctx context.Context, ids []string) ([]model.DeliveryOrder, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderRepository) Get(ctx context.Context, id string) (*model.DeliveryOrder, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderRepository) GetByDriverUserIdAndStatusDelivering(ctx context.Context, driverUserId string) (*model.DeliveryOrder, error) {
	stmt := stmtBuilder.Select("dorder.*").
		From(fmt.Sprintf("%s dorder", model.DeliveryOrderTableName)).
		InnerJoin(fmt.Sprintf("%s dod ON dorder.id = dod.delivery_order_id", model.DeliveryOrderDriverTableName)).
		Where(squirrel.Eq{"dorder.status": data_type.DeliveryOrderStatusDelivering}).
		Where(squirrel.Eq{"dod.driver_user_id": driverUserId})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderRepository) IsExistByDriverUserIdAndStatusDelivering(ctx context.Context, driverUserId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(fmt.Sprintf("%s dorder", model.DeliveryOrderTableName)).
			InnerJoin(fmt.Sprintf("%s dod ON dorder.id = dod.delivery_order_id", model.DeliveryOrderDriverTableName)).
			Where(squirrel.Eq{"dorder.status": data_type.DeliveryOrderStatusDelivering}).
			Where(squirrel.Eq{"dod.driver_user_id": driverUserId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *deliveryOrderRepository) IsExistByCustomerId(ctx context.Context, customerId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(fmt.Sprintf("%s dorder", model.DeliveryOrderTableName)).
			Where(squirrel.Eq{"dorder.customer_id": customerId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *deliveryOrderRepository) Update(ctx context.Context, deliveryOrder *model.DeliveryOrder) error {
	return defaultUpdate(r.db, ctx, deliveryOrder, "*", nil)
}

func (r *deliveryOrderRepository) Delete(ctx context.Context, deliveryOrder *model.DeliveryOrder) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, deliveryOrder, nil)
}

func (r *deliveryOrderRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderTableName)
}
