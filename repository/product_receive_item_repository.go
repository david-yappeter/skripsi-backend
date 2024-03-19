package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReceiveItemRepository interface {
	// create
	Insert(ctx context.Context, productReceiveItem *model.ProductReceiveItem) error
	InsertMany(ctx context.Context, productReceiveItems []model.ProductReceiveItem, options ...data_type.RepositoryOption) error

	// read
	FetchByProductReceiveIds(ctx context.Context, productReceiveIds []string) ([]model.ProductReceiveItem, error)
	Get(ctx context.Context, id string) (*model.ProductReceiveItem, error)
	GetByProductReceiveIdAndProductUnitId(ctx context.Context, productReceiveId string, productUnitId string) (*model.ProductReceiveItem, error)
	IsExistByProductId(ctx context.Context, productId string) (bool, error)

	// update
	Update(ctx context.Context, productReceiveItem *model.ProductReceiveItem) error

	// delete
	Delete(ctx context.Context, productReceiveItem *model.ProductReceiveItem) error
	DeleteManyByProductReceiveId(ctx context.Context, productReceiveId string) error
	Truncate(ctx context.Context) error
}

type productReceiveItemRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductReceiveItemRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReceiveItemRepository {
	return &productReceiveItemRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productReceiveItemRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReceiveItem, error) {
	productReceiveItems := []model.ProductReceiveItem{}
	if err := fetch(r.db, ctx, &productReceiveItems, stmt); err != nil {
		return nil, err
	}

	return productReceiveItems, nil
}

func (r *productReceiveItemRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReceiveItem, error) {
	productReceiveItem := model.ProductReceiveItem{}
	if err := get(r.db, ctx, &productReceiveItem, stmt); err != nil {
		return nil, err
	}

	return &productReceiveItem, nil
}

func (r *productReceiveItemRepository) Insert(ctx context.Context, productReceiveItem *model.ProductReceiveItem) error {
	return defaultInsert(r.db, ctx, productReceiveItem, "*")
}

func (r *productReceiveItemRepository) InsertMany(ctx context.Context, productReceiveItems []model.ProductReceiveItem, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReceiveItems {
		arr = append(arr, &productReceiveItems[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReceiveItemRepository) FetchByProductReceiveIds(ctx context.Context, productReceiveIds []string) ([]model.ProductReceiveItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveItemTableName).
		Where(squirrel.Eq{"product_receive_id": productReceiveIds})

	return r.fetch(ctx, stmt)
}

func (r *productReceiveItemRepository) Get(ctx context.Context, id string) (*model.ProductReceiveItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveItemTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productReceiveItemRepository) GetByProductReceiveIdAndProductUnitId(ctx context.Context, productReceiveId string, productUnitId string) (*model.ProductReceiveItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveItemTableName).
		Where(squirrel.Eq{"product_receive_id": productReceiveId}).
		Where(squirrel.Eq{"product_unit_id": productUnitId})

	return r.get(ctx, stmt)
}

func (r *productReceiveItemRepository) IsExistByProductId(ctx context.Context, productId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(fmt.Sprintf("%s pri", model.ProductReceiveItemTableName)).
			InnerJoin(fmt.Sprintf("%s pu ON pu.id = pri.product_unit_id", model.ProductUnitTableName)).
			Where(squirrel.Eq{"pu.product_id": productId}),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productReceiveItemRepository) Update(ctx context.Context, productReceiveItem *model.ProductReceiveItem) error {
	return defaultUpdate(r.db, ctx, productReceiveItem, "*", nil)
}

func (r *productReceiveItemRepository) Delete(ctx context.Context, productReceiveItem *model.ProductReceiveItem) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productReceiveItem, nil)
}

func (r *productReceiveItemRepository) DeleteManyByProductReceiveId(ctx context.Context, productReceiveId string) error {
	whereStmt := squirrel.Eq{
		"product_receive_id": productReceiveId,
	}
	return destroy(r.db, ctx, model.ProductReceiveItemTableName, whereStmt)
}

func (r *productReceiveItemRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReceiveItemTableName)
}
