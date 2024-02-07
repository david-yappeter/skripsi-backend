package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CustomerDebtRepository interface {
	// create
	Insert(ctx context.Context, customerDebt *model.CustomerDebt) error
	InsertMany(ctx context.Context, customerDebts []model.CustomerDebt, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.CustomerDebtQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CustomerDebtQueryOption) ([]model.CustomerDebt, error)
	Get(ctx context.Context, id string) (*model.CustomerDebt, error)
	GetByDebtSourceAndDebtSourceId(ctx context.Context, debtSource data_type.CustomerDebtDebtSource, debtSourceId string) (*model.CustomerDebt, error)

	// update
	Update(ctx context.Context, customerDebt *model.CustomerDebt) error

	// delete
	Delete(ctx context.Context, customerDebt *model.CustomerDebt) error
	Truncate(ctx context.Context) error
}

type customerDebtRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCustomerDebtRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CustomerDebtRepository {
	return &customerDebtRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *customerDebtRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.CustomerDebt, error) {
	customerDebts := []model.CustomerDebt{}
	if err := fetch(r.db, ctx, &customerDebts, stmt); err != nil {
		return nil, err
	}

	return customerDebts, nil
}

func (r *customerDebtRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.CustomerDebt, error) {
	customerDebt := model.CustomerDebt{}
	if err := get(r.db, ctx, &customerDebt, stmt); err != nil {
		return nil, err
	}

	return &customerDebt, nil
}

func (r *customerDebtRepository) prepareQuery(option model.CustomerDebtQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s cd", model.CustomerDebtTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"cd.name": phrase},
		})
	}

	if option.CustomerId != nil {
		stmt = stmt.Where(squirrel.Eq{
			"cd.customer_id": option.CustomerId,
		})
	}

	if option.Status != nil {
		stmt = stmt.Where(squirrel.Eq{
			"cd.status": option.Status,
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *customerDebtRepository) Insert(ctx context.Context, customerDebt *model.CustomerDebt) error {
	return defaultInsert(r.db, ctx, customerDebt, "*")
}

func (r *customerDebtRepository) InsertMany(ctx context.Context, customerDebts []model.CustomerDebt, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range customerDebts {
		arr = append(arr, &customerDebts[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *customerDebtRepository) Count(ctx context.Context, options ...model.CustomerDebtQueryOption) (int, error) {
	option := model.CustomerDebtQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *customerDebtRepository) Fetch(ctx context.Context, options ...model.CustomerDebtQueryOption) ([]model.CustomerDebt, error) {
	option := model.CustomerDebtQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *customerDebtRepository) Get(ctx context.Context, id string) (*model.CustomerDebt, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerDebtTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *customerDebtRepository) GetByDebtSourceAndDebtSourceId(ctx context.Context, debtSource data_type.CustomerDebtDebtSource, debtSourceId string) (*model.CustomerDebt, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerDebtTableName).
		Where(squirrel.Eq{"debt_source": debtSource}).
		Where(squirrel.Eq{"debt_source_id": debtSourceId})

	return r.get(ctx, stmt)
}

func (r *customerDebtRepository) Update(ctx context.Context, customerDebt *model.CustomerDebt) error {
	return defaultUpdate(r.db, ctx, customerDebt, "*", nil)
}

func (r *customerDebtRepository) Delete(ctx context.Context, customerDebt *model.CustomerDebt) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, customerDebt, nil)
}

func (r *customerDebtRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CustomerDebtTableName)
}
