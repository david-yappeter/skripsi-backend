package repository

import (
	"context"
	"fmt"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CartRepository interface {
	// create
	Insert(ctx context.Context, cart *model.Cart) error

	// read
	Count(ctx context.Context, options ...model.CartQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CartQueryOption) ([]model.Cart, error)
	Get(ctx context.Context, id string) (*model.Cart, error)
	GetByCashierSessionIdAndIsActive(ctx context.Context, cashierSessionId string, isActive bool) (*model.Cart, error)
	IsExistByCashierSessionId(ctx context.Context, cashierSessionId string) (bool, error)
	IsExistByCashierSessionIdAndIsActive(ctx context.Context, cashierSessionId string, isActive bool) (bool, error)

	// update
	Update(ctx context.Context, cart *model.Cart) error

	// delete
	Delete(ctx context.Context, cart *model.Cart) error
	Truncate(ctx context.Context) error
}

type cartRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCartRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CartRepository {
	return &cartRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *cartRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Cart, error) {
	carts := []model.Cart{}
	if err := fetch(r.db, ctx, &carts, stmt); err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *cartRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Cart, error) {
	cart := model.Cart{}
	if err := get(r.db, ctx, &cart, stmt); err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepository) prepareQuery(option model.CartQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.CartTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *cartRepository) Insert(ctx context.Context, cart *model.Cart) error {
	return defaultInsert(r.db, ctx, cart, "*")
}

func (r *cartRepository) Count(ctx context.Context, options ...model.CartQueryOption) (int, error) {
	option := model.CartQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *cartRepository) Fetch(ctx context.Context, options ...model.CartQueryOption) ([]model.Cart, error) {
	option := model.CartQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *cartRepository) Get(ctx context.Context, id string) (*model.Cart, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CartTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *cartRepository) GetByCashierSessionIdAndIsActive(ctx context.Context, cashierSessionId string, isActive bool) (*model.Cart, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CartTableName).
		Where(squirrel.Eq{"cashier_session_id": cashierSessionId}).
		Where(squirrel.Eq{"is_active": isActive})

	return r.get(ctx, stmt)
}

func (r *cartRepository) IsExistByCashierSessionId(ctx context.Context, cashierSessionId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.CartTableName).
			Where(squirrel.Eq{"cashier_session_id": cashierSessionId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *cartRepository) IsExistByCashierSessionIdAndIsActive(ctx context.Context, cashierSessionId string, isActive bool) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.CartTableName).
			Where(squirrel.Eq{"cashier_session_id": cashierSessionId}).
			Where(squirrel.Eq{"is_active": isActive}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *cartRepository) Update(ctx context.Context, cart *model.Cart) error {
	return defaultUpdate(r.db, ctx, cart, "*", nil)
}

func (r *cartRepository) Delete(ctx context.Context, cart *model.Cart) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, cart, nil)
}

func (r *cartRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CartTableName)
}
