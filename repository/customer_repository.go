package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CustomerRepository interface {
	// create
	Insert(ctx context.Context, customer *model.Customer) error
	InsertMany(ctx context.Context, customers []model.Customer, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.CustomerQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CustomerQueryOption) ([]model.Customer, error)
	FetchByCustomerTypeId(ctx context.Context, customerTypeId *string) ([]model.Customer, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.Customer, error)
	Get(ctx context.Context, id string) (*model.Customer, error)
	IsExistByCustomerTypeId(ctx context.Context, customerTypeId string) (bool, error)

	// update
	Update(ctx context.Context, customer *model.Customer) error

	// delete
	Delete(ctx context.Context, customer *model.Customer) error
	Truncate(ctx context.Context) error
}

type customerRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCustomerRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CustomerRepository {
	return &customerRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *customerRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Customer, error) {
	customers := []model.Customer{}
	if err := fetch(r.db, ctx, &customers, stmt); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Customer, error) {
	customer := model.Customer{}
	if err := get(r.db, ctx, &customer, stmt); err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) prepareQuery(option model.CustomerQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s c", model.CustomerTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(
			squirrel.Or{
				squirrel.ILike{"c.name": phrase},
				squirrel.ILike{"c.email": phrase},
				squirrel.ILike{"c.phone": phrase},
			},
		)
	}

	if option.IsActive != nil {
		stmt = stmt.Where(squirrel.And{
			squirrel.Eq{"c.is_active": option.IsActive},
		})
	}

	if option.CustomerTypeId != nil {
		stmt = stmt.Where(squirrel.Eq{
			"c.customer_type_id": option.CustomerTypeId,
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *customerRepository) Insert(ctx context.Context, customer *model.Customer) error {
	return defaultInsert(r.db, ctx, customer, "*")
}

func (r *customerRepository) InsertMany(ctx context.Context, customers []model.Customer, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range customers {
		arr = append(arr, &customers[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *customerRepository) Count(ctx context.Context, options ...model.CustomerQueryOption) (int, error) {
	option := model.CustomerQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *customerRepository) Fetch(ctx context.Context, options ...model.CustomerQueryOption) ([]model.Customer, error) {
	option := model.CustomerQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *customerRepository) FetchByCustomerTypeId(ctx context.Context, customerTypeId *string) ([]model.Customer, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTableName).
		Where(squirrel.Eq{"customer_type_id": customerTypeId})

	return r.fetch(ctx, stmt)
}

func (r *customerRepository) FetchByIds(ctx context.Context, ids []string) ([]model.Customer, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *customerRepository) Get(ctx context.Context, id string) (*model.Customer, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *customerRepository) IsExistByCustomerTypeId(ctx context.Context, customerTypeId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("*").
			From(model.CustomerTableName).
			Where(squirrel.Eq{"customer_type_id": customerTypeId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *customerRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.CustomerTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *customerRepository) Update(ctx context.Context, customer *model.Customer) error {
	return defaultUpdate(r.db, ctx, customer, "*", nil)
}

func (r *customerRepository) Delete(ctx context.Context, customer *model.Customer) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, customer, nil)
}

func (r *customerRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CustomerTableName)
}
