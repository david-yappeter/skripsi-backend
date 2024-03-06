package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type UnitRepository interface {
	// create
	Insert(ctx context.Context, unit *model.Unit) error
	InsertMany(ctx context.Context, units []model.Unit, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.UnitQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.UnitQueryOption) ([]model.Unit, error)
	Get(ctx context.Context, id string) (*model.Unit, error)
	IsExistByName(ctx context.Context, name string) (bool, error)

	// update
	Update(ctx context.Context, unit *model.Unit) error

	// delete
	Delete(ctx context.Context, unit *model.Unit) error
	Truncate(ctx context.Context) error
}

type unitRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewUnitRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) UnitRepository {
	return &unitRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *unitRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Unit, error) {
	units := []model.Unit{}
	if err := fetch(r.db, ctx, &units, stmt); err != nil {
		return nil, err
	}

	return units, nil
}

func (r *unitRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Unit, error) {
	unit := model.Unit{}
	if err := get(r.db, ctx, &unit, stmt); err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *unitRepository) prepareQuery(option model.UnitQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.UnitTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	if option.ProductIdNotExist != nil {
		stmt = stmt.Where(
			stmtBuilder.Select("1").
				From(fmt.Sprintf("%s pu", model.ProductUnitTableName)).
				Where(squirrel.Eq{"pu.product_id": option.ProductIdNotExist}).
				Where(squirrel.Expr("u.id = pu.unit_id")).
				Prefix("NOT EXISTS (").Suffix(")"),
		)
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *unitRepository) Insert(ctx context.Context, unit *model.Unit) error {
	return defaultInsert(r.db, ctx, unit, "*")
}

func (r *unitRepository) InsertMany(ctx context.Context, units []model.Unit, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range units {
		arr = append(arr, &units[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *unitRepository) Count(ctx context.Context, options ...model.UnitQueryOption) (int, error) {
	option := model.UnitQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *unitRepository) Fetch(ctx context.Context, options ...model.UnitQueryOption) ([]model.Unit, error) {
	option := model.UnitQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *unitRepository) Get(ctx context.Context, id string) (*model.Unit, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UnitTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *unitRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.UnitTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *unitRepository) Update(ctx context.Context, unit *model.Unit) error {
	return defaultUpdate(r.db, ctx, unit, "*", nil)
}

func (r *unitRepository) Delete(ctx context.Context, unit *model.Unit) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, unit, nil)
}

func (r *unitRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.UnitTableName)
}
