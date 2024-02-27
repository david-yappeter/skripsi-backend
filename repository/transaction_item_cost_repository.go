package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type TransactionItemCostRepository interface {
	// create
	Insert(ctx context.Context, transactionItemCost *model.TransactionItemCost) error
	InsertMany(ctx context.Context, transactionItemCosts []model.TransactionItemCost, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.TransactionItemCostQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.TransactionItemCostQueryOption) ([]model.TransactionItemCost, error)
	Get(ctx context.Context, id string) (*model.TransactionItemCost, error)

	// update
	Update(ctx context.Context, transactionItemCost *model.TransactionItemCost) error

	// delete
	Delete(ctx context.Context, transactionItemCost *model.TransactionItemCost) error
	Truncate(ctx context.Context) error
}

type transactionItemCostRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewTransactionItemCostRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) TransactionItemCostRepository {
	return &transactionItemCostRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *transactionItemCostRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.TransactionItemCost, error) {
	transactionItemCosts := []model.TransactionItemCost{}
	if err := fetch(r.db, ctx, &transactionItemCosts, stmt); err != nil {
		return nil, err
	}

	return transactionItemCosts, nil
}

func (r *transactionItemCostRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.TransactionItemCost, error) {
	transactionItemCost := model.TransactionItemCost{}
	if err := get(r.db, ctx, &transactionItemCost, stmt); err != nil {
		return nil, err
	}

	return &transactionItemCost, nil
}

func (r *transactionItemCostRepository) prepareQuery(option model.TransactionItemCostQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.TransactionItemCostTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *transactionItemCostRepository) Insert(ctx context.Context, transactionItemCost *model.TransactionItemCost) error {
	return defaultInsert(r.db, ctx, transactionItemCost, "*")
}

func (r *transactionItemCostRepository) InsertMany(ctx context.Context, transactionItemCosts []model.TransactionItemCost, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range transactionItemCosts {
		arr = append(arr, &transactionItemCosts[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *transactionItemCostRepository) Count(ctx context.Context, options ...model.TransactionItemCostQueryOption) (int, error) {
	option := model.TransactionItemCostQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *transactionItemCostRepository) Fetch(ctx context.Context, options ...model.TransactionItemCostQueryOption) ([]model.TransactionItemCost, error) {
	option := model.TransactionItemCostQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *transactionItemCostRepository) Get(ctx context.Context, id string) (*model.TransactionItemCost, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TransactionItemCostTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *transactionItemCostRepository) Update(ctx context.Context, transactionItemCost *model.TransactionItemCost) error {
	return defaultUpdate(r.db, ctx, transactionItemCost, "*", nil)
}

func (r *transactionItemCostRepository) Delete(ctx context.Context, transactionItemCost *model.TransactionItemCost) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, transactionItemCost, nil)
}

func (r *transactionItemCostRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.TransactionItemCostTableName)
}
