package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductDiscountRepository interface {
	// create
	Insert(ctx context.Context, productDiscount *model.ProductDiscount) error
	InsertMany(ctx context.Context, productDiscounts []model.ProductDiscount, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductDiscountQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductDiscountQueryOption) ([]model.ProductDiscount, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.ProductDiscount, error)
	FetchByProductIds(ctx context.Context, productIds []string) ([]model.ProductDiscount, error)
	Get(ctx context.Context, id string) (*model.ProductDiscount, error)
	GetByProductId(ctx context.Context, productId string) (*model.ProductDiscount, error)
	IsExistByProductId(ctx context.Context, productId string) (bool, error)

	// update
	Update(ctx context.Context, productDiscount *model.ProductDiscount) error

	// delete
	Delete(ctx context.Context, productDiscount *model.ProductDiscount) error
	Truncate(ctx context.Context) error
}

type productDiscountRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductDiscountRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductDiscountRepository {
	return &productDiscountRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productDiscountRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductDiscount, error) {
	productDiscounts := []model.ProductDiscount{}
	if err := fetch(r.db, ctx, &productDiscounts, stmt); err != nil {
		return nil, err
	}

	return productDiscounts, nil
}

func (r *productDiscountRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductDiscount, error) {
	productDiscount := model.ProductDiscount{}
	if err := get(r.db, ctx, &productDiscount, stmt); err != nil {
		return nil, err
	}

	return &productDiscount, nil
}

func (r *productDiscountRepository) prepareQuery(option model.ProductDiscountQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s pd", model.ProductDiscountTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"pd.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productDiscountRepository) Insert(ctx context.Context, productDiscount *model.ProductDiscount) error {
	return defaultInsert(r.db, ctx, productDiscount, "*")
}

func (r *productDiscountRepository) InsertMany(ctx context.Context, productDiscounts []model.ProductDiscount, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productDiscounts {
		arr = append(arr, &productDiscounts[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productDiscountRepository) Count(ctx context.Context, options ...model.ProductDiscountQueryOption) (int, error) {
	option := model.ProductDiscountQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productDiscountRepository) Fetch(ctx context.Context, options ...model.ProductDiscountQueryOption) ([]model.ProductDiscount, error) {
	option := model.ProductDiscountQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productDiscountRepository) FetchByIds(ctx context.Context, ids []string) ([]model.ProductDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductDiscountTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *productDiscountRepository) FetchByProductIds(ctx context.Context, productIds []string) ([]model.ProductDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductDiscountTableName).
		Where(squirrel.Eq{"product_id": productIds})

	return r.fetch(ctx, stmt)
}

func (r *productDiscountRepository) Get(ctx context.Context, id string) (*model.ProductDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductDiscountTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productDiscountRepository) GetByProductId(ctx context.Context, productId string) (*model.ProductDiscount, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductDiscountTableName).
		Where(squirrel.Eq{"product_id": productId})

	return r.get(ctx, stmt)
}

func (r *productDiscountRepository) IsExistByProductId(ctx context.Context, productId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductDiscountTableName).
			Where(squirrel.Eq{"product_id": productId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productDiscountRepository) Update(ctx context.Context, productDiscount *model.ProductDiscount) error {
	return defaultUpdate(r.db, ctx, productDiscount, "*", nil)
}

func (r *productDiscountRepository) Delete(ctx context.Context, productDiscount *model.ProductDiscount) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productDiscount, nil)
}

func (r *productDiscountRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductDiscountTableName)
}
