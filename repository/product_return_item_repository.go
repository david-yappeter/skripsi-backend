package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReturnItemRepository interface {
	// create
	Insert(ctx context.Context, productReturnItem *model.ProductReturnItem) error
	InsertMany(ctx context.Context, productReturnItems []model.ProductReturnItem, options ...data_type.RepositoryOption) error

	// read
	FetchByProductReturnIds(ctx context.Context, productReturnIds []string) ([]model.ProductReturnItem, error)
	Get(ctx context.Context, id string) (*model.ProductReturnItem, error)
	GetByProductReturnIdAndProductUnitId(ctx context.Context, productReturnId string, productUnitId string) (*model.ProductReturnItem, error)

	// update
	Update(ctx context.Context, productReturnItem *model.ProductReturnItem) error

	// delete
	Delete(ctx context.Context, productReturnItem *model.ProductReturnItem) error
	DeleteManyByProductReturnId(ctx context.Context, productReturnId string) error
	Truncate(ctx context.Context) error
}

type productReturnItemRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductReturnItemRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReturnItemRepository {
	return &productReturnItemRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productReturnItemRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReturnItem, error) {
	productReturnItems := []model.ProductReturnItem{}
	if err := fetch(r.db, ctx, &productReturnItems, stmt); err != nil {
		return nil, err
	}

	return productReturnItems, nil
}

func (r *productReturnItemRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReturnItem, error) {
	productReturnItem := model.ProductReturnItem{}
	if err := get(r.db, ctx, &productReturnItem, stmt); err != nil {
		return nil, err
	}

	return &productReturnItem, nil
}

func (r *productReturnItemRepository) Insert(ctx context.Context, productReturnItem *model.ProductReturnItem) error {
	return defaultInsert(r.db, ctx, productReturnItem, "*")
}

func (r *productReturnItemRepository) InsertMany(ctx context.Context, productReturnItems []model.ProductReturnItem, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReturnItems {
		arr = append(arr, &productReturnItems[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReturnItemRepository) FetchByProductReturnIds(ctx context.Context, productReturnIds []string) ([]model.ProductReturnItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReturnItemTableName).
		Where(squirrel.Eq{"product_return_id": productReturnIds})

	return r.fetch(ctx, stmt)
}

func (r *productReturnItemRepository) Get(ctx context.Context, id string) (*model.ProductReturnItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReturnItemTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productReturnItemRepository) GetByProductReturnIdAndProductUnitId(ctx context.Context, productReturnId string, productUnitId string) (*model.ProductReturnItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReturnItemTableName).
		Where(squirrel.Eq{"product_return_id": productReturnId}).
		Where(squirrel.Eq{"product_unit_id": productUnitId})

	return r.get(ctx, stmt)
}

func (r *productReturnItemRepository) Update(ctx context.Context, productReturnItem *model.ProductReturnItem) error {
	return defaultUpdate(r.db, ctx, productReturnItem, "*", nil)
}

func (r *productReturnItemRepository) Delete(ctx context.Context, productReturnItem *model.ProductReturnItem) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productReturnItem, nil)
}

func (r *productReturnItemRepository) DeleteManyByProductReturnId(ctx context.Context, productReturnId string) error {
	whereStmt := squirrel.Eq{
		"product_return_id": productReturnId,
	}
	return destroy(r.db, ctx, model.ProductReturnItemTableName, whereStmt)
}

func (r *productReturnItemRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReturnItemTableName)
}
