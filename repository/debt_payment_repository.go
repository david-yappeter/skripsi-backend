package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DebtPaymentRepository interface {
	// create
	Insert(ctx context.Context, debtPayment *model.DebtPayment) error
	InsertMany(ctx context.Context, debtPayments []model.DebtPayment, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.DebtPaymentQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.DebtPaymentQueryOption) ([]model.DebtPayment, error)
	FetchByDebtIds(ctx context.Context, debtIds []string) ([]model.DebtPayment, error)
	Get(ctx context.Context, id string) (*model.DebtPayment, error)

	// update
	Update(ctx context.Context, debtPayment *model.DebtPayment) error

	// delete
	Delete(ctx context.Context, debtPayment *model.DebtPayment) error
	Truncate(ctx context.Context) error
}

type debtPaymentRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDebtPaymentRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DebtPaymentRepository {
	return &debtPaymentRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *debtPaymentRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DebtPayment, error) {
	debtPayments := []model.DebtPayment{}
	if err := fetch(r.db, ctx, &debtPayments, stmt); err != nil {
		return nil, err
	}

	return debtPayments, nil
}

func (r *debtPaymentRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DebtPayment, error) {
	debtPayment := model.DebtPayment{}
	if err := get(r.db, ctx, &debtPayment, stmt); err != nil {
		return nil, err
	}

	return &debtPayment, nil
}

func (r *debtPaymentRepository) prepareQuery(option model.DebtPaymentQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s cp", model.DebtPaymentTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"cp.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *debtPaymentRepository) Insert(ctx context.Context, debtPayment *model.DebtPayment) error {
	return defaultInsert(r.db, ctx, debtPayment, "*")
}

func (r *debtPaymentRepository) InsertMany(ctx context.Context, debtPayments []model.DebtPayment, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range debtPayments {
		arr = append(arr, &debtPayments[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *debtPaymentRepository) Count(ctx context.Context, options ...model.DebtPaymentQueryOption) (int, error) {
	option := model.DebtPaymentQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *debtPaymentRepository) Fetch(ctx context.Context, options ...model.DebtPaymentQueryOption) ([]model.DebtPayment, error) {
	option := model.DebtPaymentQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *debtPaymentRepository) FetchByDebtIds(ctx context.Context, debtIds []string) ([]model.DebtPayment, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DebtPaymentTableName).
		Where(squirrel.Eq{"id": debtIds})

	return r.fetch(ctx, stmt)
}

func (r *debtPaymentRepository) Get(ctx context.Context, id string) (*model.DebtPayment, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DebtPaymentTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *debtPaymentRepository) Update(ctx context.Context, debtPayment *model.DebtPayment) error {
	return defaultUpdate(r.db, ctx, debtPayment, "*", nil)
}

func (r *debtPaymentRepository) Delete(ctx context.Context, debtPayment *model.DebtPayment) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, debtPayment, nil)
}

func (r *debtPaymentRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DebtPaymentTableName)
}
