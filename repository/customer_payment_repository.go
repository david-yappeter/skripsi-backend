package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CustomerPaymentRepository interface {
	// create
	Insert(ctx context.Context, customerPayment *model.CustomerPayment) error
	InsertMany(ctx context.Context, customerPayments []model.CustomerPayment, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.CustomerPaymentQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CustomerPaymentQueryOption) ([]model.CustomerPayment, error)
	FetchByCustomerDebtIds(ctx context.Context, customerDebtIds []string) ([]model.CustomerPayment, error)
	Get(ctx context.Context, id string) (*model.CustomerPayment, error)

	// update
	Update(ctx context.Context, customerPayment *model.CustomerPayment) error

	// delete
	Delete(ctx context.Context, customerPayment *model.CustomerPayment) error
	Truncate(ctx context.Context) error
}

type customerPaymentRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCustomerPaymentRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CustomerPaymentRepository {
	return &customerPaymentRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *customerPaymentRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.CustomerPayment, error) {
	customerPayments := []model.CustomerPayment{}
	if err := fetch(r.db, ctx, &customerPayments, stmt); err != nil {
		return nil, err
	}

	return customerPayments, nil
}

func (r *customerPaymentRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.CustomerPayment, error) {
	customerPayment := model.CustomerPayment{}
	if err := get(r.db, ctx, &customerPayment, stmt); err != nil {
		return nil, err
	}

	return &customerPayment, nil
}

func (r *customerPaymentRepository) prepareQuery(option model.CustomerPaymentQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s cp", model.CustomerPaymentTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"cp.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *customerPaymentRepository) Insert(ctx context.Context, customerPayment *model.CustomerPayment) error {
	return defaultInsert(r.db, ctx, customerPayment, "*")
}

func (r *customerPaymentRepository) InsertMany(ctx context.Context, customerPayments []model.CustomerPayment, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range customerPayments {
		arr = append(arr, &customerPayments[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *customerPaymentRepository) Count(ctx context.Context, options ...model.CustomerPaymentQueryOption) (int, error) {
	option := model.CustomerPaymentQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *customerPaymentRepository) Fetch(ctx context.Context, options ...model.CustomerPaymentQueryOption) ([]model.CustomerPayment, error) {
	option := model.CustomerPaymentQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *customerPaymentRepository) FetchByCustomerDebtIds(ctx context.Context, customerDebtIds []string) ([]model.CustomerPayment, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerPaymentTableName).
		Where(squirrel.Eq{"customer_debt_id": customerDebtIds})

	return r.fetch(ctx, stmt)
}

func (r *customerPaymentRepository) Get(ctx context.Context, id string) (*model.CustomerPayment, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerPaymentTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *customerPaymentRepository) Update(ctx context.Context, customerPayment *model.CustomerPayment) error {
	return defaultUpdate(r.db, ctx, customerPayment, "*", nil)
}

func (r *customerPaymentRepository) Delete(ctx context.Context, customerPayment *model.CustomerPayment) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, customerPayment, nil)
}

func (r *customerPaymentRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CustomerPaymentTableName)
}
