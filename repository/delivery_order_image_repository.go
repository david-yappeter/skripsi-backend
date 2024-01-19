package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderImageRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderImage *model.DeliveryOrderImage) error
	InsertMany(ctx context.Context, deliveryOrderImages []model.DeliveryOrderImage, options ...data_type.RepositoryOption) error

	// read
	FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderImage, error)
	Get(ctx context.Context, id string) (*model.DeliveryOrderImage, error)
	GetByDeliveryOrderIdAndFileId(ctx context.Context, deliveryOrderId string, fileId string) (*model.DeliveryOrderImage, error)
	IsExistByName(ctx context.Context, name string) (bool, error)
	IsExistByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) (bool, error)

	// update
	Update(ctx context.Context, deliveryOrderImage *model.DeliveryOrderImage) error

	// delete
	Delete(ctx context.Context, deliveryOrderImage *model.DeliveryOrderImage) error
	DeleteManyByDeliveryOrderId(ctx context.Context, deliveryOrderId string) error
	Truncate(ctx context.Context) error
}

type deliveryOrderImageRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDeliveryOrderImageRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderImageRepository {
	return &deliveryOrderImageRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *deliveryOrderImageRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderImage, error) {
	deliveryOrderImages := []model.DeliveryOrderImage{}
	if err := fetch(r.db, ctx, &deliveryOrderImages, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderImages, nil
}

func (r *deliveryOrderImageRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderImage, error) {
	deliveryOrderImage := model.DeliveryOrderImage{}
	if err := get(r.db, ctx, &deliveryOrderImage, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderImage, nil
}

func (r *deliveryOrderImageRepository) Insert(ctx context.Context, deliveryOrderImage *model.DeliveryOrderImage) error {
	return defaultInsert(r.db, ctx, deliveryOrderImage, "*")
}

func (r *deliveryOrderImageRepository) InsertMany(ctx context.Context, deliveryOrderImages []model.DeliveryOrderImage, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderImages {
		arr = append(arr, &deliveryOrderImages[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderImageRepository) FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderImageTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderIds})

	return r.fetch(ctx, stmt)
}
func (r *deliveryOrderImageRepository) Get(ctx context.Context, id string) (*model.DeliveryOrderImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderImageTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderImageRepository) GetByDeliveryOrderIdAndFileId(ctx context.Context, deliveryOrderId string, fileId string) (*model.DeliveryOrderImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderImageTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderId}).
		Where(squirrel.Eq{"file_id": fileId})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderImageRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.DeliveryOrderImageTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *deliveryOrderImageRepository) IsExistByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.DeliveryOrderImageTableName).
			Where(squirrel.Eq{"delivery_order_id": deliveryOrderIds}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *deliveryOrderImageRepository) Update(ctx context.Context, deliveryOrderImage *model.DeliveryOrderImage) error {
	return defaultUpdate(r.db, ctx, deliveryOrderImage, "*", nil)
}

func (r *deliveryOrderImageRepository) Delete(ctx context.Context, deliveryOrderImage *model.DeliveryOrderImage) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, deliveryOrderImage, nil)
}

func (r *deliveryOrderImageRepository) DeleteManyByDeliveryOrderId(ctx context.Context, deliveryOrderId string) error {
	whereStmt := squirrel.Eq{
		"delivery_order_id": deliveryOrderId,
	}
	return destroy(r.db, ctx, model.DeliveryOrderImageTableName, whereStmt)
}

func (r *deliveryOrderImageRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderImageTableName)
}
