package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReceiveReturnRepository interface {
	// create
	Insert(ctx context.Context, productReceiveReturn *model.ProductReceiveReturn) error
	InsertMany(ctx context.Context, productReceiveReturns []model.ProductReceiveReturn, options ...data_type.RepositoryOption) error

	// read
	FetchByProductReceiveIds(ctx context.Context, productReceiveIds []string) ([]model.ProductReceiveReturn, error)

	// update
	Update(ctx context.Context, productReceiveReturn *model.ProductReceiveReturn) error

	// delete
	Truncate(ctx context.Context) error
}

type productReceiveReturnRepository struct {
	db infrastructure.DBTX
}

func NewProductReceiveReturnRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReceiveReturnRepository {
	return &productReceiveReturnRepository{
		db: db,
	}
}

func (r *productReceiveReturnRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReceiveReturn, error) {
	productReceiveReturns := []model.ProductReceiveReturn{}
	if err := fetch(r.db, ctx, &productReceiveReturns, stmt); err != nil {
		return nil, err
	}

	return productReceiveReturns, nil
}

func (r *productReceiveReturnRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReceiveReturn, error) {
	productReceiveReturn := model.ProductReceiveReturn{}
	if err := get(r.db, ctx, &productReceiveReturn, stmt); err != nil {
		return nil, err
	}

	return &productReceiveReturn, nil
}

func (r *productReceiveReturnRepository) Insert(ctx context.Context, productReceiveReturn *model.ProductReceiveReturn) error {
	return defaultInsert(r.db, ctx, productReceiveReturn, "*")
}

func (r *productReceiveReturnRepository) InsertMany(ctx context.Context, productReceiveReturns []model.ProductReceiveReturn, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReceiveReturns {
		arr = append(arr, &productReceiveReturns[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReceiveReturnRepository) FetchByProductReceiveIds(ctx context.Context, productReceiveIds []string) ([]model.ProductReceiveReturn, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveReturnTableName).
		Where(squirrel.Eq{"product_receive_id": productReceiveIds})

	return r.fetch(ctx, stmt)
}

func (r *productReceiveReturnRepository) Update(ctx context.Context, productReceiveReturn *model.ProductReceiveReturn) error {
	return defaultUpdate(r.db, ctx, productReceiveReturn, "*", nil)
}

func (r *productReceiveReturnRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReceiveReturnTableName)
}
