package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderItemCostRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderItemCost *model.DeliveryOrderItemCost) error
	InsertMany(ctx context.Context, deliveryOrderItemCosts []model.DeliveryOrderItemCost, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.DeliveryOrderItemCostQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.DeliveryOrderItemCostQueryOption) ([]model.DeliveryOrderItemCost, error)
	Get(ctx context.Context, id string) (*model.DeliveryOrderItemCost, error)

	// update
	Update(ctx context.Context, deliveryOrderItemCost *model.DeliveryOrderItemCost) error

	// delete
	Delete(ctx context.Context, deliveryOrderItemCost *model.DeliveryOrderItemCost) error
	Truncate(ctx context.Context) error
}

type deliveryOrderItemCostRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDeliveryOrderItemCostRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderItemCostRepository {
	return &deliveryOrderItemCostRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *deliveryOrderItemCostRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderItemCost, error) {
	deliveryOrderItemCosts := []model.DeliveryOrderItemCost{}
	if err := fetch(r.db, ctx, &deliveryOrderItemCosts, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderItemCosts, nil
}

func (r *deliveryOrderItemCostRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderItemCost, error) {
	deliveryOrderItemCost := model.DeliveryOrderItemCost{}
	if err := get(r.db, ctx, &deliveryOrderItemCost, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderItemCost, nil
}

func (r *deliveryOrderItemCostRepository) prepareQuery(option model.DeliveryOrderItemCostQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s doic", model.DeliveryOrderItemCostTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *deliveryOrderItemCostRepository) Insert(ctx context.Context, deliveryOrderItemCost *model.DeliveryOrderItemCost) error {
	return defaultInsert(r.db, ctx, deliveryOrderItemCost, "*")
}

func (r *deliveryOrderItemCostRepository) InsertMany(ctx context.Context, deliveryOrderItemCosts []model.DeliveryOrderItemCost, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderItemCosts {
		arr = append(arr, &deliveryOrderItemCosts[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderItemCostRepository) Count(ctx context.Context, options ...model.DeliveryOrderItemCostQueryOption) (int, error) {
	option := model.DeliveryOrderItemCostQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *deliveryOrderItemCostRepository) Fetch(ctx context.Context, options ...model.DeliveryOrderItemCostQueryOption) ([]model.DeliveryOrderItemCost, error) {
	option := model.DeliveryOrderItemCostQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderItemCostRepository) Get(ctx context.Context, id string) (*model.DeliveryOrderItemCost, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderItemCostTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderItemCostRepository) Update(ctx context.Context, deliveryOrderItemCost *model.DeliveryOrderItemCost) error {
	return defaultUpdate(r.db, ctx, deliveryOrderItemCost, "*", nil)
}

func (r *deliveryOrderItemCostRepository) Delete(ctx context.Context, deliveryOrderItemCost *model.DeliveryOrderItemCost) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, deliveryOrderItemCost, nil)
}

func (r *deliveryOrderItemCostRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderItemCostTableName)
}
