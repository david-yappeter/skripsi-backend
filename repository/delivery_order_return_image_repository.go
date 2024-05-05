package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderReturnImageRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderReturnImage *model.DeliveryOrderReturnImage) error
	InsertMany(ctx context.Context, deliveryOrderReturnImages []model.DeliveryOrderReturnImage, options ...data_type.RepositoryOption) error

	// read
	FetchByDeliveryOrderReturnIds(ctx context.Context, deliveryOrderReturnIds []string) ([]model.DeliveryOrderReturnImage, error)

	// update
	Update(ctx context.Context, deliveryOrderReturnImage *model.DeliveryOrderReturnImage) error

	// delete
	Truncate(ctx context.Context) error
}

type deliveryOrderReturnImageRepository struct {
	db infrastructure.DBTX
}

func NewDeliveryOrderReturnImageRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderReturnImageRepository {
	return &deliveryOrderReturnImageRepository{
		db: db,
	}
}

func (r *deliveryOrderReturnImageRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderReturnImage, error) {
	deliveryOrderReturnImages := []model.DeliveryOrderReturnImage{}
	if err := fetch(r.db, ctx, &deliveryOrderReturnImages, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderReturnImages, nil
}

func (r *deliveryOrderReturnImageRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderReturnImage, error) {
	deliveryOrderReturnImage := model.DeliveryOrderReturnImage{}
	if err := get(r.db, ctx, &deliveryOrderReturnImage, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderReturnImage, nil
}

func (r *deliveryOrderReturnImageRepository) Insert(ctx context.Context, deliveryOrderReturnImage *model.DeliveryOrderReturnImage) error {
	return defaultInsert(r.db, ctx, deliveryOrderReturnImage, "*")
}

func (r *deliveryOrderReturnImageRepository) InsertMany(ctx context.Context, deliveryOrderReturnImages []model.DeliveryOrderReturnImage, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderReturnImages {
		arr = append(arr, &deliveryOrderReturnImages[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderReturnImageRepository) FetchByDeliveryOrderReturnIds(ctx context.Context, deliveryOrderReturnIds []string) ([]model.DeliveryOrderReturnImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderReturnImageTableName).
		Where(squirrel.Eq{"delivery_order_return_id": deliveryOrderReturnIds})

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderReturnImageRepository) Update(ctx context.Context, deliveryOrderReturnImage *model.DeliveryOrderReturnImage) error {
	return defaultUpdate(r.db, ctx, deliveryOrderReturnImage, "*", nil)
}

func (r *deliveryOrderReturnImageRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderReturnImageTableName)
}
