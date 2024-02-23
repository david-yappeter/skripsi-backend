package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type TransactionRepository interface {
	// create
	Insert(ctx context.Context, transaction *model.Transaction) error
	InsertMany(ctx context.Context, transactions []model.Transaction, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.TransactionQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.TransactionQueryOption) ([]model.Transaction, error)
	Get(ctx context.Context, id string) (*model.Transaction, error)

	// update
	Update(ctx context.Context, transaction *model.Transaction) error

	// delete
	Delete(ctx context.Context, transaction *model.Transaction) error
	Truncate(ctx context.Context) error
}

type transactionRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewTransactionRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) TransactionRepository {
	return &transactionRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *transactionRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	if err := fetch(r.db, ctx, &transactions, stmt); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Transaction, error) {
	transaction := model.Transaction{}
	if err := get(r.db, ctx, &transaction, stmt); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) prepareQuery(option model.TransactionQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s t", model.TransactionTableName))

	if option.CashierSessionId != nil {
		stmt = stmt.Where(squirrel.Eq{"t.cashier_session_id": option.CashierSessionId})
	}

	if option.Status != nil {
		stmt = stmt.Where(squirrel.Eq{"t.status": option.Status})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *transactionRepository) Insert(ctx context.Context, transaction *model.Transaction) error {
	return defaultInsert(r.db, ctx, transaction, "*")
}

func (r *transactionRepository) InsertMany(ctx context.Context, transactions []model.Transaction, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range transactions {
		arr = append(arr, &transactions[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *transactionRepository) Count(ctx context.Context, options ...model.TransactionQueryOption) (int, error) {
	option := model.TransactionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *transactionRepository) Fetch(ctx context.Context, options ...model.TransactionQueryOption) ([]model.Transaction, error) {
	option := model.TransactionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *transactionRepository) Get(ctx context.Context, id string) (*model.Transaction, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TransactionTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *transactionRepository) Update(ctx context.Context, transaction *model.Transaction) error {
	return defaultUpdate(r.db, ctx, transaction, "*", nil)
}

func (r *transactionRepository) Delete(ctx context.Context, transaction *model.Transaction) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, transaction, nil)
}

func (r *transactionRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.TransactionTableName)
}
