package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type TransactionPaymentRepository interface {
	// create
	Insert(ctx context.Context, transactionPayment *model.TransactionPayment) error
	InsertMany(ctx context.Context, transactionPayments []model.TransactionPayment, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context) (int, error)
	FetchByTransactionIds(ctx context.Context, transactionIds []string) ([]model.TransactionPayment, error)
	Get(ctx context.Context, id string) (*model.TransactionPayment, error)
	GetByTransactionId(ctx context.Context, transactionId string) (*model.TransactionPayment, error)

	// update
	Update(ctx context.Context, transactionPayment *model.TransactionPayment) error

	// delete
	Delete(ctx context.Context, transactionPayment *model.TransactionPayment) error
	Truncate(ctx context.Context) error
}

type transactionPaymentRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewTransactionPaymentRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) TransactionPaymentRepository {
	return &transactionPaymentRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *transactionPaymentRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.TransactionPayment, error) {
	transactionPayments := []model.TransactionPayment{}
	if err := fetch(r.db, ctx, &transactionPayments, stmt); err != nil {
		return nil, err
	}

	return transactionPayments, nil
}

func (r *transactionPaymentRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.TransactionPayment, error) {
	transactionPayment := model.TransactionPayment{}
	if err := get(r.db, ctx, &transactionPayment, stmt); err != nil {
		return nil, err
	}

	return &transactionPayment, nil
}

func (r *transactionPaymentRepository) Insert(ctx context.Context, transactionPayment *model.TransactionPayment) error {
	return defaultInsert(r.db, ctx, transactionPayment, "*")
}

func (r *transactionPaymentRepository) InsertMany(ctx context.Context, transactionPayments []model.TransactionPayment, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range transactionPayments {
		arr = append(arr, &transactionPayments[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *transactionPaymentRepository) Count(ctx context.Context) (int, error) {
	stmt := stmtBuilder.Select("COUNT(*) as count").
		From(model.TransactionPaymentTableName)

	return count(r.db, ctx, stmt)
}

func (r *transactionPaymentRepository) FetchByTransactionIds(ctx context.Context, transactionIds []string) ([]model.TransactionPayment, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TransactionPaymentTableName).
		Where(squirrel.Eq{"transaction_id": transactionIds})

	return r.fetch(ctx, stmt)
}

func (r *transactionPaymentRepository) Get(ctx context.Context, id string) (*model.TransactionPayment, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TransactionPaymentTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *transactionPaymentRepository) GetByTransactionId(ctx context.Context, transactionId string) (*model.TransactionPayment, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TransactionPaymentTableName).
		Where(squirrel.Eq{"transaction_id": transactionId})

	return r.get(ctx, stmt)
}

func (r *transactionPaymentRepository) Update(ctx context.Context, transactionPayment *model.TransactionPayment) error {
	return defaultUpdate(r.db, ctx, transactionPayment, "*", nil)
}

func (r *transactionPaymentRepository) Delete(ctx context.Context, transactionPayment *model.TransactionPayment) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, transactionPayment, nil)
}

func (r *transactionPaymentRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.TransactionPaymentTableName)
}
