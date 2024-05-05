package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReceiveReturnImageRepository interface {
	// create
	Insert(ctx context.Context, productReceiveReturnImage *model.ProductReceiveReturnImage) error
	InsertMany(ctx context.Context, productReceiveReturnImages []model.ProductReceiveReturnImage, options ...data_type.RepositoryOption) error

	// read
	FetchByProductReceiveReturnImageIds(ctx context.Context, productReceiveReturnImageIds []string) ([]model.ProductReceiveReturnImage, error)

	// update
	Update(ctx context.Context, productReceiveReturnImage *model.ProductReceiveReturnImage) error

	// delete
	Truncate(ctx context.Context) error
}

type productReceiveReturnImageRepository struct {
	db infrastructure.DBTX
}

func NewProductReceiveReturnImageRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReceiveReturnImageRepository {
	return &productReceiveReturnImageRepository{
		db: db,
	}
}

func (r *productReceiveReturnImageRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReceiveReturnImage, error) {
	productReceiveReturnImages := []model.ProductReceiveReturnImage{}
	if err := fetch(r.db, ctx, &productReceiveReturnImages, stmt); err != nil {
		return nil, err
	}

	return productReceiveReturnImages, nil
}

func (r *productReceiveReturnImageRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReceiveReturnImage, error) {
	productReceiveReturnImage := model.ProductReceiveReturnImage{}
	if err := get(r.db, ctx, &productReceiveReturnImage, stmt); err != nil {
		return nil, err
	}

	return &productReceiveReturnImage, nil
}

func (r *productReceiveReturnImageRepository) Insert(ctx context.Context, productReceiveReturnImage *model.ProductReceiveReturnImage) error {
	return defaultInsert(r.db, ctx, productReceiveReturnImage, "*")
}

func (r *productReceiveReturnImageRepository) InsertMany(ctx context.Context, productReceiveReturnImages []model.ProductReceiveReturnImage, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReceiveReturnImages {
		arr = append(arr, &productReceiveReturnImages[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReceiveReturnImageRepository) FetchByProductReceiveReturnImageIds(ctx context.Context, productReceiveReturnImageIds []string) ([]model.ProductReceiveReturnImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveReturnImageTableName).
		Where(squirrel.Eq{"delivery_order_return_id": productReceiveReturnImageIds})

	return r.fetch(ctx, stmt)
}

func (r *productReceiveReturnImageRepository) Update(ctx context.Context, productReceiveReturnImage *model.ProductReceiveReturnImage) error {
	return defaultUpdate(r.db, ctx, productReceiveReturnImage, "*", nil)
}

func (r *productReceiveReturnImageRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReceiveReturnImageTableName)
}
