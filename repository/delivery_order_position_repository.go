package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderPositionRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderPosition *model.DeliveryOrderPosition) error
	InsertMany(ctx context.Context, deliveryOrderPositions []model.DeliveryOrderPosition, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.DeliveryOrderPositionQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.DeliveryOrderPositionQueryOption) ([]model.DeliveryOrderPosition, error)
	Get(ctx context.Context, id string) (*model.DeliveryOrderPosition, error)
	GetByDeliveryOrderId(ctx context.Context, deliveryOrderId string) (*model.DeliveryOrderPosition, error)
	IsExistByDeliveryOrderId(ctx context.Context, deliveryOrderId string) (bool, error)

	// update
	Update(ctx context.Context, deliveryOrderPosition *model.DeliveryOrderPosition) error

	// delete
	Delete(ctx context.Context, deliveryOrderPosition *model.DeliveryOrderPosition) error
	Truncate(ctx context.Context) error
}

type deliveryOrderPositionRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDeliveryOrderPositionRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderPositionRepository {
	return &deliveryOrderPositionRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *deliveryOrderPositionRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderPosition, error) {
	deliveryOrderPositions := []model.DeliveryOrderPosition{}
	if err := fetch(r.db, ctx, &deliveryOrderPositions, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderPositions, nil
}

func (r *deliveryOrderPositionRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderPosition, error) {
	deliveryOrderPosition := model.DeliveryOrderPosition{}
	if err := get(r.db, ctx, &deliveryOrderPosition, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderPosition, nil
}

func (r *deliveryOrderPositionRepository) prepareQuery(option model.DeliveryOrderPositionQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.DeliveryOrderPositionTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *deliveryOrderPositionRepository) Insert(ctx context.Context, deliveryOrderPosition *model.DeliveryOrderPosition) error {
	return defaultInsert(r.db, ctx, deliveryOrderPosition, "*")
}

func (r *deliveryOrderPositionRepository) InsertMany(ctx context.Context, deliveryOrderPositions []model.DeliveryOrderPosition, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderPositions {
		arr = append(arr, &deliveryOrderPositions[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderPositionRepository) Count(ctx context.Context, options ...model.DeliveryOrderPositionQueryOption) (int, error) {
	option := model.DeliveryOrderPositionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *deliveryOrderPositionRepository) Fetch(ctx context.Context, options ...model.DeliveryOrderPositionQueryOption) ([]model.DeliveryOrderPosition, error) {
	option := model.DeliveryOrderPositionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderPositionRepository) Get(ctx context.Context, id string) (*model.DeliveryOrderPosition, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderPositionTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderPositionRepository) GetByDeliveryOrderId(ctx context.Context, deliveryOrderId string) (*model.DeliveryOrderPosition, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderPositionTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderId})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderPositionRepository) IsExistByDeliveryOrderId(ctx context.Context, deliveryOrderId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.DeliveryOrderPositionTableName).
			Where(squirrel.Eq{"delivery_order_id": deliveryOrderId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *deliveryOrderPositionRepository) Update(ctx context.Context, deliveryOrderPosition *model.DeliveryOrderPosition) error {
	return defaultUpdate(r.db, ctx, deliveryOrderPosition, "*", nil)
}

func (r *deliveryOrderPositionRepository) Delete(ctx context.Context, deliveryOrderPosition *model.DeliveryOrderPosition) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, deliveryOrderPosition, nil)
}

func (r *deliveryOrderPositionRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderPositionTableName)
}
