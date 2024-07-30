package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type CustomerTypeDiscountRepository interface {
	// create
	Insert(ctx context.Context, customerTypeDiscount *model.CustomerTypeDiscount) error
	InsertMany(ctx context.Context, customerTypeDiscounts []model.CustomerTypeDiscount, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.CustomerTypeDiscountQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CustomerTypeDiscountQueryOption) ([]model.CustomerTypeDiscount, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.CustomerTypeDiscount, error)
	FetchByCustomerTypeIds(ctx context.Context, customerTypeIds []string) ([]model.CustomerTypeDiscount, error)
	Get(ctx context.Context, id string) (*model.CustomerTypeDiscount, error)
	GetByCustomerTypeIdAndProductId(ctx context.Context, customerTypeId string, productId string) (*model.CustomerTypeDiscount, error)
	GetByIdAndCustomerTypeId(ctx context.Context, id string, customerTypeId string) (*model.CustomerTypeDiscount, error)

	// update
	Update(ctx context.Context, customerTypeDiscount *model.CustomerTypeDiscount) error

	// delete
	Delete(ctx context.Context, customerTypeDiscount *model.CustomerTypeDiscount) error
	Truncate(ctx context.Context) error
}

type customerTypeDiscountRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewCustomerTypeDiscountRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) CustomerTypeDiscountRepository {
	return &customerTypeDiscountRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *customerTypeDiscountRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.CustomerTypeDiscount, error) {
	customerTypeDiscounts := []model.CustomerTypeDiscount{}
	if err := fetch(r.db, ctx, &customerTypeDiscounts, stmt); err != nil {
		return nil, err
	}

	return customerTypeDiscounts, nil
}

func (r *customerTypeDiscountRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.CustomerTypeDiscount, error) {
	customerTypeDiscount := model.CustomerTypeDiscount{}
	if err := get(r.db, ctx, &customerTypeDiscount, stmt); err != nil {
		return nil, err
	}

	return &customerTypeDiscount, nil
}

func (r *customerTypeDiscountRepository) prepareQuery(option model.CustomerTypeDiscountQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s ctd", model.CustomerTypeDiscountTableName)).
		InnerJoin(fmt.Sprintf("%s p ON ctc.product_id = p.id", model.ProductTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"p.name": phrase},
		})
	}

	if option.CustomerTypeId != nil {
		stmt = stmt.Where(squirrel.Eq{"ctd.customer_type_id": option.CustomerTypeId})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *customerTypeDiscountRepository) Insert(ctx context.Context, customerTypeDiscount *model.CustomerTypeDiscount) error {
	return defaultInsert(r.db, ctx, customerTypeDiscount, "*")
}

func (r *customerTypeDiscountRepository) InsertMany(ctx context.Context, customerTypeDiscounts []model.CustomerTypeDiscount, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range customerTypeDiscounts {
		arr = append(arr, &customerTypeDiscounts[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *customerTypeDiscountRepository) Count(ctx context.Context, options ...model.CustomerTypeDiscountQueryOption) (int, error) {
	option := model.CustomerTypeDiscountQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *customerTypeDiscountRepository) Fetch(ctx context.Context, options ...model.CustomerTypeDiscountQueryOption) ([]model.CustomerTypeDiscount, error) {
	option := model.CustomerTypeDiscountQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *customerTypeDiscountRepository) FetchByIds(ctx context.Context, ids []string) ([]model.CustomerTypeDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTypeDiscountTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *customerTypeDiscountRepository) FetchByCustomerTypeIds(ctx context.Context, customerTypeIds []string) ([]model.CustomerTypeDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTypeDiscountTableName).
		Where(squirrel.Eq{"customer_type_id": customerTypeIds})

	return r.fetch(ctx, stmt)
}

func (r *customerTypeDiscountRepository) Get(ctx context.Context, id string) (*model.CustomerTypeDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTypeDiscountTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *customerTypeDiscountRepository) GetByCustomerTypeIdAndProductId(ctx context.Context, customerTypeId string, productId string) (*model.CustomerTypeDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTypeDiscountTableName).
		Where(squirrel.Eq{"customer_type_id": customerTypeId}).
		Where(squirrel.Eq{"product_id": productId})

	return r.get(ctx, stmt)
}

func (r *customerTypeDiscountRepository) GetByIdAndCustomerTypeId(ctx context.Context, id string, customerTypeId string) (*model.CustomerTypeDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CustomerTypeDiscountTableName).
		Where(squirrel.Eq{"id": id}).
		Where(squirrel.Eq{"customer_type_id": customerTypeId})

	return r.get(ctx, stmt)
}

func (r *customerTypeDiscountRepository) Update(ctx context.Context, customerTypeDiscount *model.CustomerTypeDiscount) error {
	return defaultUpdate(r.db, ctx, customerTypeDiscount, "*", nil)
}

func (r *customerTypeDiscountRepository) Delete(ctx context.Context, customerTypeDiscount *model.CustomerTypeDiscount) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, customerTypeDiscount, nil)
}

func (r *customerTypeDiscountRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.CustomerTypeDiscountTableName)
}
