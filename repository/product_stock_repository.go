package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductStockRepository interface {
	// create
	Insert(ctx context.Context, productStock *model.ProductStock) error
	InsertMany(ctx context.Context, productStocks []model.ProductStock, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductStockQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductStockQueryOption) ([]model.ProductStock, error)
	FetchByProductIds(ctx context.Context, productIds []string) ([]model.ProductStock, error)
	Get(ctx context.Context, id string) (*model.ProductStock, error)
	GetByProductId(ctx context.Context, productId string) (*model.ProductStock, error)
	IsExistByName(ctx context.Context, name string) (bool, error)
	IsExistByProductId(ctx context.Context, productId string) (bool, error)

	// update
	Update(ctx context.Context, productStock *model.ProductStock) error

	// delete
	Delete(ctx context.Context, productStock *model.ProductStock) error
	Truncate(ctx context.Context) error
}

type productStockRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductStockRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductStockRepository {
	return &productStockRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productStockRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductStock, error) {
	productStocks := []model.ProductStock{}
	if err := fetch(r.db, ctx, &productStocks, stmt); err != nil {
		return nil, err
	}

	return productStocks, nil
}

func (r *productStockRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductStock, error) {
	productStock := model.ProductStock{}
	if err := get(r.db, ctx, &productStock, stmt); err != nil {
		return nil, err
	}

	return &productStock, nil
}

func (r *productStockRepository) prepareQuery(option model.ProductStockQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.ProductStockTableName))

	andStatements := squirrel.And{}

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		andStatements = append(
			andStatements,
			squirrel.Or{
				squirrel.ILike{"u.name": phrase},
			},
		)
	}

	stmt = stmt.Where(andStatements)

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productStockRepository) Insert(ctx context.Context, productStock *model.ProductStock) error {
	return defaultInsert(r.db, ctx, productStock, "*")
}

func (r *productStockRepository) InsertMany(ctx context.Context, productStocks []model.ProductStock, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productStocks {
		arr = append(arr, &productStocks[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productStockRepository) Count(ctx context.Context, options ...model.ProductStockQueryOption) (int, error) {
	option := model.ProductStockQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productStockRepository) Fetch(ctx context.Context, options ...model.ProductStockQueryOption) ([]model.ProductStock, error) {
	option := model.ProductStockQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productStockRepository) FetchByProductIds(ctx context.Context, productIds []string) ([]model.ProductStock, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductStockTableName).
		Where(squirrel.Eq{"product_id": productIds})

	return r.fetch(ctx, stmt)
}

func (r *productStockRepository) Get(ctx context.Context, id string) (*model.ProductStock, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductStockTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productStockRepository) GetByProductId(ctx context.Context, productId string) (*model.ProductStock, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductStockTableName).
		Where(squirrel.Eq{"product_id": productId})

	return r.get(ctx, stmt)
}

func (r *productStockRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductStockTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productStockRepository) IsExistByProductId(ctx context.Context, productId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductStockTableName).
			Where(squirrel.Eq{"product_id": productId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productStockRepository) Update(ctx context.Context, productStock *model.ProductStock) error {
	return defaultUpdate(r.db, ctx, productStock, "*", nil)
}

func (r *productStockRepository) Delete(ctx context.Context, productStock *model.ProductStock) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productStock, nil)
}

func (r *productStockRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductStockTableName)
}
