package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductStockMutationRepository interface {
	// create
	Insert(ctx context.Context, productStockMutation *model.ProductStockMutation) error
	InsertMany(ctx context.Context, productStockMutations []model.ProductStockMutation, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductStockMutationQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductStockMutationQueryOption) ([]model.ProductStockMutation, error)
	FetchHaveQtyLeft(ctx context.Context) ([]model.ProductStockMutation, error)
	FetchByTypeAndIdentifierIds(ctx context.Context, _type data_type.ProductStockMutationType, identifierIds []string) ([]model.ProductStockMutation, error)
	Get(ctx context.Context, id string) (*model.ProductStockMutation, error)
	GetByTypeAndIdentifierId(ctx context.Context, _type data_type.ProductStockMutationType, identifierId string) (*model.ProductStockMutation, error)
	GetFIFOByProductIdAndBaseQtyLeftNotZero(ctx context.Context, productId string) (*model.ProductStockMutation, error)

	// update
	Update(ctx context.Context, productStockMutation *model.ProductStockMutation) error

	// delete
	Delete(ctx context.Context, productStockMutation *model.ProductStockMutation) error
	DeleteManyByIds(ctx context.Context, ids []string) error
	Truncate(ctx context.Context) error
}

type productStockMutationRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductStockMutationRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductStockMutationRepository {
	return &productStockMutationRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productStockMutationRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductStockMutation, error) {
	productStockMutations := []model.ProductStockMutation{}
	if err := fetch(r.db, ctx, &productStockMutations, stmt); err != nil {
		return nil, err
	}

	return productStockMutations, nil
}

func (r *productStockMutationRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductStockMutation, error) {
	productStockMutation := model.ProductStockMutation{}
	if err := get(r.db, ctx, &productStockMutation, stmt); err != nil {
		return nil, err
	}

	return &productStockMutation, nil
}

func (r *productStockMutationRepository) prepareQuery(option model.ProductStockMutationQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s psm", model.ProductStockMutationTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"psm.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productStockMutationRepository) Insert(ctx context.Context, productStockMutation *model.ProductStockMutation) error {
	return defaultInsert(r.db, ctx, productStockMutation, "*")
}

func (r *productStockMutationRepository) InsertMany(ctx context.Context, productStockMutations []model.ProductStockMutation, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productStockMutations {
		arr = append(arr, &productStockMutations[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productStockMutationRepository) Count(ctx context.Context, options ...model.ProductStockMutationQueryOption) (int, error) {
	option := model.ProductStockMutationQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productStockMutationRepository) Fetch(ctx context.Context, options ...model.ProductStockMutationQueryOption) ([]model.ProductStockMutation, error) {
	option := model.ProductStockMutationQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productStockMutationRepository) FetchHaveQtyLeft(ctx context.Context) ([]model.ProductStockMutation, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductStockMutationTableName).
		Where(squirrel.Gt{"base_qty_left": 0})

	return r.fetch(ctx, stmt)
}

func (r *productStockMutationRepository) FetchByTypeAndIdentifierIds(ctx context.Context, _type data_type.ProductStockMutationType, identifierIds []string) ([]model.ProductStockMutation, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductStockMutationTableName).
		Where(squirrel.Eq{"type": _type}).
		Where(squirrel.Eq{"identifier_id": identifierIds})

	return r.fetch(ctx, stmt)
}

func (r *productStockMutationRepository) Get(ctx context.Context, id string) (*model.ProductStockMutation, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductStockMutationTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productStockMutationRepository) GetByTypeAndIdentifierId(ctx context.Context, _type data_type.ProductStockMutationType, identifierId string) (*model.ProductStockMutation, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductStockMutationTableName).
		Where(squirrel.Eq{"type": _type}).
		Where(squirrel.Eq{"identifier_id": identifierId})

	return r.get(ctx, stmt)
}

func (r *productStockMutationRepository) GetFIFOByProductIdAndBaseQtyLeftNotZero(ctx context.Context, productId string) (*model.ProductStockMutation, error) {
	stmt := stmtBuilder.Select("psm.*").
		From(fmt.Sprintf("%s psm", model.ProductStockMutationTableName)).
		InnerJoin(fmt.Sprintf("%s pu ON pu.id = psm.product_unit_id", model.ProductUnitTableName)).
		Where(squirrel.Eq{"pu.product_id": productId}).
		Where(squirrel.Gt{"psm.base_qty_left": 0}).
		OrderBy("psm.mutated_at ASC")

	return r.get(ctx, stmt)
}

func (r *productStockMutationRepository) Update(ctx context.Context, productStockMutation *model.ProductStockMutation) error {
	return defaultUpdate(r.db, ctx, productStockMutation, "*", nil)
}

func (r *productStockMutationRepository) Delete(ctx context.Context, productStockMutation *model.ProductStockMutation) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productStockMutation, nil)
}

func (r *productStockMutationRepository) DeleteManyByIds(ctx context.Context, ids []string) error {
	whereStmt := squirrel.Eq{
		"id": ids,
	}
	return destroy(r.db, ctx, model.ProductStockMutationTableName, whereStmt)
}

func (r *productStockMutationRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductStockMutationTableName)
}
