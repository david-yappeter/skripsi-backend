package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductUnitRepository interface {
	// create
	Insert(ctx context.Context, productProductUnit *model.ProductUnit) error
	InsertMany(ctx context.Context, productProductUnits []model.ProductUnit, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductUnitQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductUnitQueryOption) ([]model.ProductUnit, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.ProductUnit, error)
	FetchByProductIds(ctx context.Context, productIds []string) ([]model.ProductUnit, error)
	FetchBaseByProductIds(ctx context.Context, productIds []string) ([]model.ProductUnit, error)
	Get(ctx context.Context, id string) (*model.ProductUnit, error)
	GetByProductIdAndUnitId(ctx context.Context, productId string, unitId string) (*model.ProductUnit, error)
	GetByProductIdAndToUnitId(ctx context.Context, productId string, toUnitId string) (*model.ProductUnit, error)
	GetBaseProductUnitByProductId(ctx context.Context, productId string) (*model.ProductUnit, error)
	IsExistByProductId(ctx context.Context, productId string) (bool, error)
	IsExistByProductIdAndUnitId(ctx context.Context, productId string, unitId string) (bool, error)
	IsExistByProductIdAndToUnitId(ctx context.Context, productId string, toUnitId string) (bool, error)

	// update
	Update(ctx context.Context, productProductUnit *model.ProductUnit) error

	// delete
	Delete(ctx context.Context, productProductUnit *model.ProductUnit) error
	Truncate(ctx context.Context) error
}

type productProductUnitRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductUnitRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductUnitRepository {
	return &productProductUnitRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productProductUnitRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductUnit, error) {
	productProductUnits := []model.ProductUnit{}
	if err := fetch(r.db, ctx, &productProductUnits, stmt); err != nil {
		return nil, err
	}

	return productProductUnits, nil
}

func (r *productProductUnitRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductUnit, error) {
	productProductUnit := model.ProductUnit{}
	if err := get(r.db, ctx, &productProductUnit, stmt); err != nil {
		return nil, err
	}

	return &productProductUnit, nil
}

func (r *productProductUnitRepository) prepareQuery(option model.ProductUnitQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s pu", model.ProductUnitTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"pu.name": phrase},
		})
	}

	if len(option.ExcludeIds) > 0 {
		stmt = stmt.Where(squirrel.NotEq{
			"pu.id": option.ExcludeIds,
		})
	}

	if option.ProductId != nil {
		stmt = stmt.Where(squirrel.Eq{
			"pu.product_id": option.ProductId,
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productProductUnitRepository) Insert(ctx context.Context, productProductUnit *model.ProductUnit) error {
	return defaultInsert(r.db, ctx, productProductUnit, "*")
}

func (r *productProductUnitRepository) InsertMany(ctx context.Context, productProductUnits []model.ProductUnit, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productProductUnits {
		arr = append(arr, &productProductUnits[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productProductUnitRepository) Count(ctx context.Context, options ...model.ProductUnitQueryOption) (int, error) {
	option := model.ProductUnitQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productProductUnitRepository) Fetch(ctx context.Context, options ...model.ProductUnitQueryOption) ([]model.ProductUnit, error) {
	option := model.ProductUnitQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productProductUnitRepository) FetchByIds(ctx context.Context, ids []string) ([]model.ProductUnit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductUnitTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *productProductUnitRepository) FetchByProductIds(ctx context.Context, productIds []string) ([]model.ProductUnit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductUnitTableName).
		Where(squirrel.Eq{"product_id": productIds})

	return r.fetch(ctx, stmt)
}

func (r *productProductUnitRepository) FetchBaseByProductIds(ctx context.Context, productIds []string) ([]model.ProductUnit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductUnitTableName).
		Where(squirrel.Eq{"product_id": productIds}).
		Where(squirrel.Expr("to_unit_id IS NULL"))

	return r.fetch(ctx, stmt)
}

func (r *productProductUnitRepository) Get(ctx context.Context, id string) (*model.ProductUnit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductUnitTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productProductUnitRepository) GetByProductIdAndUnitId(ctx context.Context, productId string, unitId string) (*model.ProductUnit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductUnitTableName).
		Where(squirrel.Eq{"product_id": productId}).
		Where(squirrel.Eq{"unit_id": unitId})

	return r.get(ctx, stmt)
}

func (r *productProductUnitRepository) GetByProductIdAndToUnitId(ctx context.Context, productId string, toUnitId string) (*model.ProductUnit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductUnitTableName).
		Where(squirrel.Eq{"product_id": productId}).
		Where(squirrel.Eq{"to_unit_id": toUnitId})

	return r.get(ctx, stmt)
}

func (r *productProductUnitRepository) GetBaseProductUnitByProductId(ctx context.Context, productId string) (*model.ProductUnit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductUnitTableName).
		Where(squirrel.Eq{"product_id": productId}).
		Where(squirrel.Eq{"to_unit_id": nil})

	return r.get(ctx, stmt)
}

func (r *productProductUnitRepository) IsExistByProductId(ctx context.Context, productId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductUnitTableName).
			Where(squirrel.Eq{"product_id": productId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productProductUnitRepository) IsExistByProductIdAndUnitId(ctx context.Context, productId string, unitId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductUnitTableName).
			Where(squirrel.Eq{"product_id": productId}).
			Where(squirrel.Eq{"unit_id": unitId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productProductUnitRepository) IsExistByProductIdAndToUnitId(ctx context.Context, productId string, toUnitId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductUnitTableName).
			Where(squirrel.Eq{"product_id": productId}).
			Where(squirrel.Eq{"to_unit_id": toUnitId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productProductUnitRepository) Update(ctx context.Context, productProductUnit *model.ProductUnit) error {
	return defaultUpdate(r.db, ctx, productProductUnit, "*", nil)
}

func (r *productProductUnitRepository) Delete(ctx context.Context, productProductUnit *model.ProductUnit) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productProductUnit, nil)
}

func (r *productProductUnitRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductUnitTableName)
}
