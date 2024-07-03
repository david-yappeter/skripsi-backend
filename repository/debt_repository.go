package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DebtRepository interface {
	// create
	Insert(ctx context.Context, Debt *model.Debt) error
	InsertMany(ctx context.Context, Debts []model.Debt, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.DebtQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.DebtQueryOption) ([]model.Debt, error)
	Get(ctx context.Context, id string) (*model.Debt, error)
	GetByDebtSourceAndDebtSourceId(ctx context.Context, debtSource data_type.DebtSource, debtSourceId string) (*model.Debt, error)

	// update
	Update(ctx context.Context, Debt *model.Debt) error

	// delete
	Delete(ctx context.Context, Debt *model.Debt) error
	Truncate(ctx context.Context) error
}

type debtRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewDebtRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DebtRepository {
	return &debtRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *debtRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Debt, error) {
	Debts := []model.Debt{}
	if err := fetch(r.db, ctx, &Debts, stmt); err != nil {
		return nil, err
	}

	return Debts, nil
}

func (r *debtRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Debt, error) {
	Debt := model.Debt{}
	if err := get(r.db, ctx, &Debt, stmt); err != nil {
		return nil, err
	}

	return &Debt, nil
}

func (r *debtRepository) prepareQuery(option model.DebtQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s d", model.DebtTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"d.name": phrase},
		})
	}

	if option.Id != nil {
		stmt = stmt.Where(squirrel.Eq{
			"d._id": option.Id,
		})
	}

	if option.Status != nil {
		stmt = stmt.Where(squirrel.Eq{
			"d.status": option.Status,
		})
	}

	if option.StartDate.DateP() != nil {
		stmt = stmt.Where(squirrel.LtOrEq{"d.created_at": option.StartDate})
	}

	if option.EndDate.DateP() != nil {
		stmt = stmt.Where(squirrel.GtOrEq{"d.created_at": option.EndDate})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *debtRepository) Insert(ctx context.Context, Debt *model.Debt) error {
	return defaultInsert(r.db, ctx, Debt, "*")
}

func (r *debtRepository) InsertMany(ctx context.Context, Debts []model.Debt, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range Debts {
		arr = append(arr, &Debts[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *debtRepository) Count(ctx context.Context, options ...model.DebtQueryOption) (int, error) {
	option := model.DebtQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *debtRepository) Fetch(ctx context.Context, options ...model.DebtQueryOption) ([]model.Debt, error) {
	option := model.DebtQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *debtRepository) Get(ctx context.Context, id string) (*model.Debt, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DebtTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *debtRepository) GetByDebtSourceAndDebtSourceId(ctx context.Context, debtSource data_type.DebtSource, debtSourceId string) (*model.Debt, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DebtTableName).
		Where(squirrel.Eq{"debt_source": debtSource}).
		Where(squirrel.Eq{"debt_source_identifier": debtSourceId})

	return r.get(ctx, stmt)
}

func (r *debtRepository) Update(ctx context.Context, Debt *model.Debt) error {
	return defaultUpdate(r.db, ctx, Debt, "*", nil)
}

func (r *debtRepository) Delete(ctx context.Context, Debt *model.Debt) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, Debt, nil)
}

func (r *debtRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DebtTableName)
}
