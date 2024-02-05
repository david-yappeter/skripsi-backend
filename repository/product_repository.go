package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductRepository interface {
	// create
	Insert(ctx context.Context, product *model.Product) error
	InsertMany(ctx context.Context, products []model.Product, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductQueryOption) ([]model.Product, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.Product, error)
	Get(ctx context.Context, id string) (*model.Product, error)
	IsExistByName(ctx context.Context, name string) (bool, error)

	// update
	Update(ctx context.Context, product *model.Product) error

	// delete
	Delete(ctx context.Context, product *model.Product) error
	Truncate(ctx context.Context) error
}

type productRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductRepository {
	return &productRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Product, error) {
	products := []model.Product{}
	if err := fetch(r.db, ctx, &products, stmt); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Product, error) {
	product := model.Product{}
	if err := get(r.db, ctx, &product, stmt); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) prepareQuery(option model.ProductQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(model.ProductTableName)

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.And{
			squirrel.ILike{"name": phrase},
		})
	}

	if option.IsActive != nil {
		stmt = stmt.Where(squirrel.And{
			squirrel.Eq{"is_active": option.IsActive},
		})
	}

	if len(option.ExcludeIds) > 0 {
		stmt = stmt.Where(squirrel.NotEq{
			"id": option.ExcludeIds,
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productRepository) Insert(ctx context.Context, product *model.Product) error {
	return defaultInsert(r.db, ctx, product, "*")
}

func (r *productRepository) InsertMany(ctx context.Context, products []model.Product, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range products {
		arr = append(arr, &products[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productRepository) Count(ctx context.Context, options ...model.ProductQueryOption) (int, error) {
	option := model.ProductQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productRepository) Fetch(ctx context.Context, options ...model.ProductQueryOption) ([]model.Product, error) {
	option := model.ProductQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productRepository) FetchByIds(ctx context.Context, ids []string) ([]model.Product, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *productRepository) Get(ctx context.Context, id string) (*model.Product, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productRepository) Update(ctx context.Context, product *model.Product) error {
	return defaultUpdate(r.db, ctx, product, "*", nil)
}

func (r *productRepository) Delete(ctx context.Context, product *model.Product) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, product, nil)
}

func (r *productRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductTableName)
}
