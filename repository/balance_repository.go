package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type BalanceRepository interface {
	// create
	Insert(ctx context.Context, balance *model.Balance) error
	InsertMany(ctx context.Context, balances []model.Balance, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.BalanceQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.BalanceQueryOption) ([]model.Balance, error)
	Get(ctx context.Context, id string) (*model.Balance, error)

	// update
	Update(ctx context.Context, balance *model.Balance) error

	// delete
	Delete(ctx context.Context, balance *model.Balance) error
	Truncate(ctx context.Context) error
}

type balanceRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewBalanceRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) BalanceRepository {
	return &balanceRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *balanceRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Balance, error) {
	balances := []model.Balance{}
	if err := fetch(r.db, ctx, &balances, stmt); err != nil {
		return nil, err
	}

	return balances, nil
}

func (r *balanceRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Balance, error) {
	balance := model.Balance{}
	if err := get(r.db, ctx, &balance, stmt); err != nil {
		return nil, err
	}

	return &balance, nil
}

func (r *balanceRepository) prepareQuery(option model.BalanceQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s b", model.BalanceTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"b.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *balanceRepository) Insert(ctx context.Context, balance *model.Balance) error {
	return defaultInsert(r.db, ctx, balance, "*")
}

func (r *balanceRepository) InsertMany(ctx context.Context, balances []model.Balance, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range balances {
		arr = append(arr, &balances[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *balanceRepository) Count(ctx context.Context, options ...model.BalanceQueryOption) (int, error) {
	option := model.BalanceQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *balanceRepository) Fetch(ctx context.Context, options ...model.BalanceQueryOption) ([]model.Balance, error) {
	option := model.BalanceQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *balanceRepository) Get(ctx context.Context, id string) (*model.Balance, error) {
	stmt := stmtBuilder.Select("*").
		From(model.BalanceTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *balanceRepository) Update(ctx context.Context, balance *model.Balance) error {
	return defaultUpdate(r.db, ctx, balance, "*", nil)
}

func (r *balanceRepository) Delete(ctx context.Context, balance *model.Balance) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, balance, nil)
}

func (r *balanceRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.BalanceTableName)
}
