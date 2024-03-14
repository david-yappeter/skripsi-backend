package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CustomerTypeRepository interface {
	// create
	Insert(ctx context.Context, customerType *model.CustomerType) error
	InsertMany(ctx context.Context, customerTypes []model.CustomerType, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.CustomerTypeQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CustomerTypeQueryOption) ([]model.CustomerType, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.CustomerType, error)
	Get(ctx context.Context, id string) (*model.CustomerType, error)
	IsExistByName(ctx context.Context, name string) (bool, error)

	// update
	Update(ctx context.Context, customerType *model.CustomerType) error

	// delete
	Delete(ctx context.Context, customerType *model.CustomerType) error
	Truncate(ctx context.Context) error
}

type customerTypeRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCustomerTypeRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CustomerTypeRepository {
	return &customerTypeRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *customerTypeRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.CustomerType, error) {
	customerTypes := []model.CustomerType{}
	if err := fetch(r.db, ctx, &customerTypes, stmt); err != nil {
		return nil, err
	}

	return customerTypes, nil
}

func (r *customerTypeRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.CustomerType, error) {
	customerType := model.CustomerType{}
	if err := get(r.db, ctx, &customerType, stmt); err != nil {
		return nil, err
	}

	return &customerType, nil
}

func (r *customerTypeRepository) prepareQuery(option model.CustomerTypeQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s ct", model.CustomerTypeTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"ct.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *customerTypeRepository) Insert(ctx context.Context, customerType *model.CustomerType) error {
	return defaultInsert(r.db, ctx, customerType, "*")
}

func (r *customerTypeRepository) InsertMany(ctx context.Context, customerTypes []model.CustomerType, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range customerTypes {
		arr = append(arr, &customerTypes[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *customerTypeRepository) Count(ctx context.Context, options ...model.CustomerTypeQueryOption) (int, error) {
	option := model.CustomerTypeQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *customerTypeRepository) Fetch(ctx context.Context, options ...model.CustomerTypeQueryOption) ([]model.CustomerType, error) {
	option := model.CustomerTypeQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *customerTypeRepository) FetchByIds(ctx context.Context, ids []string) ([]model.CustomerType, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTypeTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *customerTypeRepository) Get(ctx context.Context, id string) (*model.CustomerType, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTypeTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *customerTypeRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.CustomerTypeTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *customerTypeRepository) Update(ctx context.Context, customerType *model.CustomerType) error {
	return defaultUpdate(r.db, ctx, customerType, "*", nil)
}

func (r *customerTypeRepository) Delete(ctx context.Context, customerType *model.CustomerType) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, customerType, nil)
}

func (r *customerTypeRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CustomerTypeTableName)
}
