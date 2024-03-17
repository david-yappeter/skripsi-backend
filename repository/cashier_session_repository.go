package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CashierSessionRepository interface {
	// create
	Insert(ctx context.Context, cashierSession *model.CashierSession) error

	// read
	Count(ctx context.Context, options ...model.CashierSessionQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CashierSessionQueryOption) ([]model.CashierSession, error)
	Get(ctx context.Context, id string) (*model.CashierSession, error)
	GetByUserIdAndStatusActive(ctx context.Context, userId string) (*model.CashierSession, error)
	IsExistByUserIdAndStatusActive(ctx context.Context, userId string) (bool, error)

	// update
	Update(ctx context.Context, cashierSession *model.CashierSession) error

	// delete
	Delete(ctx context.Context, cashierSession *model.CashierSession) error
	Truncate(ctx context.Context) error
}

type cashierSessionRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCashierSessionRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CashierSessionRepository {
	return &cashierSessionRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *cashierSessionRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.CashierSession, error) {
	cashierSessions := []model.CashierSession{}
	if err := fetch(r.db, ctx, &cashierSessions, stmt); err != nil {
		return nil, err
	}

	return cashierSessions, nil
}

func (r *cashierSessionRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.CashierSession, error) {
	cashierSession := model.CashierSession{}
	if err := get(r.db, ctx, &cashierSession, stmt); err != nil {
		return nil, err
	}

	return &cashierSession, nil
}

func (r *cashierSessionRepository) prepareQuery(option model.CashierSessionQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s cs", model.CashierSessionTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"cs.name": phrase},
		})
	}

	if option.StartedAtLte.DateTimeP() != nil {
		stmt = stmt.Where(squirrel.LtOrEq{
			"cs.started_at": option.StartedAtLte,
		})
	}

	if option.EndedAtGte.DateTimeP() != nil {
		stmt = stmt.Where(squirrel.GtOrEq{
			"cs.ended_at": option.EndedAtGte,
		})
	}

	if option.UserId != nil {
		stmt = stmt.Where(squirrel.Eq{
			"cs.user_id": option.UserId,
		})
	}

	if option.Status != nil {
		stmt = stmt.Where(squirrel.Eq{
			"cs.status": option.Status,
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *cashierSessionRepository) Insert(ctx context.Context, cashierSession *model.CashierSession) error {
	return defaultInsert(r.db, ctx, cashierSession, "*")
}

func (r *cashierSessionRepository) Count(ctx context.Context, options ...model.CashierSessionQueryOption) (int, error) {
	option := model.CashierSessionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *cashierSessionRepository) Fetch(ctx context.Context, options ...model.CashierSessionQueryOption) ([]model.CashierSession, error) {
	option := model.CashierSessionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *cashierSessionRepository) Get(ctx context.Context, id string) (*model.CashierSession, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CashierSessionTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *cashierSessionRepository) GetByUserIdAndStatusActive(ctx context.Context, userId string) (*model.CashierSession, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CashierSessionTableName).
		Where(squirrel.Eq{"user_id": userId}).
		Where(squirrel.Eq{"status": data_type.CashierSessionStatusActive})

	return r.get(ctx, stmt)
}

func (r *cashierSessionRepository) IsExistByUserIdAndStatusActive(ctx context.Context, userId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.CashierSessionTableName).
			Where(squirrel.Eq{"user_id": userId}).
			Where(squirrel.Eq{"status": data_type.CashierSessionStatusActive}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *cashierSessionRepository) Update(ctx context.Context, cashierSession *model.CashierSession) error {
	return defaultUpdate(r.db, ctx, cashierSession, "*", nil)
}

func (r *cashierSessionRepository) Delete(ctx context.Context, cashierSession *model.CashierSession) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, cashierSession, nil)
}

func (r *cashierSessionRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CashierSessionTableName)
}
