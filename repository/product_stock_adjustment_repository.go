package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductStockAdjustmentRepository interface {
	// create
	Insert(ctx context.Context, productStockAdjustment *model.ProductStockAdjustment) error
	InsertMany(ctx context.Context, productStockAdjustments []model.ProductStockAdjustment, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductStockAdjustmentQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductStockAdjustmentQueryOption) ([]model.ProductStockAdjustment, error)

	// delete
	Truncate(ctx context.Context) error
}

type productStockAdjustmentRepository struct {
	db infrastructure.DBTX
}

func NewProductStockAdjustmentRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductStockAdjustmentRepository {
	return &productStockAdjustmentRepository{
		db: db,
	}
}

func (r *productStockAdjustmentRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductStockAdjustment, error) {
	productStockAdjustments := []model.ProductStockAdjustment{}
	if err := fetch(r.db, ctx, &productStockAdjustments, stmt); err != nil {
		return nil, err
	}

	return productStockAdjustments, nil
}

func (r *productStockAdjustmentRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductStockAdjustment, error) {
	productStockAdjustment := model.ProductStockAdjustment{}
	if err := get(r.db, ctx, &productStockAdjustment, stmt); err != nil {
		return nil, err
	}

	return &productStockAdjustment, nil
}

func (r *productStockAdjustmentRepository) prepareQuery(option model.ProductStockAdjustmentQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s psa", model.ProductStockAdjustmentTableName))

	if option.ProductStockId != nil {
		stmt = stmt.Where(squirrel.Eq{"psa.product_stock_id": option.ProductStockId})
	}

	if option.UserId != nil {
		stmt = stmt.Where(squirrel.Eq{"psa.user_id": option.UserId})
	}

	if option.ProductId != nil {
		stmt = stmt.Where(
			stmtBuilder.Select("1").
				From(fmt.Sprintf("%s ps", model.ProductStockTableName)).
				Where(squirrel.Expr("ps.id = psa.product_stock_id")).
				Where(squirrel.Eq{"ps.product_id": option.ProductId}).
				Prefix("EXISTS (").Suffix(")"),
		)
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productStockAdjustmentRepository) Insert(ctx context.Context, productStockAdjustment *model.ProductStockAdjustment) error {
	return defaultInsert(r.db, ctx, productStockAdjustment, "*")
}

func (r *productStockAdjustmentRepository) InsertMany(ctx context.Context, productStockAdjustments []model.ProductStockAdjustment, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productStockAdjustments {
		arr = append(arr, &productStockAdjustments[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productStockAdjustmentRepository) Count(ctx context.Context, options ...model.ProductStockAdjustmentQueryOption) (int, error) {
	option := model.ProductStockAdjustmentQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productStockAdjustmentRepository) Fetch(ctx context.Context, options ...model.ProductStockAdjustmentQueryOption) ([]model.ProductStockAdjustment, error) {
	option := model.ProductStockAdjustmentQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productStockAdjustmentRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductStockAdjustmentTableName)
}
