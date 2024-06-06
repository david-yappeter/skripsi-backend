package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReturnRepository interface {
	// create
	Insert(ctx context.Context, productReturn *model.ProductReturn) error
	InsertMany(ctx context.Context, productReturns []model.ProductReturn, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductReturnQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductReturnQueryOption) ([]model.ProductReturn, error)
	Get(ctx context.Context, id string) (*model.ProductReturn, error)
	IsExistBySupplierId(ctx context.Context, supplierId string) (bool, error)

	// update
	Update(ctx context.Context, productReturn *model.ProductReturn) error

	// delete
	Delete(ctx context.Context, productReturn *model.ProductReturn) error
	Truncate(ctx context.Context) error
}

type productReturnRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductReturnRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReturnRepository {
	return &productReturnRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productReturnRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReturn, error) {
	productReturns := []model.ProductReturn{}
	if err := fetch(r.db, ctx, &productReturns, stmt); err != nil {
		return nil, err
	}

	return productReturns, nil
}

func (r *productReturnRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReturn, error) {
	productReturn := model.ProductReturn{}
	if err := get(r.db, ctx, &productReturn, stmt); err != nil {
		return nil, err
	}

	return &productReturn, nil
}

func (r *productReturnRepository) prepareQuery(option model.ProductReturnQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s pr", model.ProductReturnTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"pr.invoice_number": phrase},
		})
	}

	if option.Status != nil {
		stmt = stmt.Where(squirrel.Eq{"pr.status": option.Status})
	}

	if option.SupplierId != nil {
		stmt = stmt.Where(squirrel.Eq{"pr.supplier_id": option.SupplierId})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productReturnRepository) Insert(ctx context.Context, productReturn *model.ProductReturn) error {
	return defaultInsert(r.db, ctx, productReturn, "*")
}

func (r *productReturnRepository) InsertMany(ctx context.Context, productReturns []model.ProductReturn, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReturns {
		arr = append(arr, &productReturns[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReturnRepository) Count(ctx context.Context, options ...model.ProductReturnQueryOption) (int, error) {
	option := model.ProductReturnQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productReturnRepository) Fetch(ctx context.Context, options ...model.ProductReturnQueryOption) ([]model.ProductReturn, error) {
	option := model.ProductReturnQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productReturnRepository) Get(ctx context.Context, id string) (*model.ProductReturn, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReturnTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productReturnRepository) IsExistBySupplierId(ctx context.Context, supplierId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductReturnTableName).
			Where(squirrel.Eq{"supplier_id": supplierId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productReturnRepository) Update(ctx context.Context, productReturn *model.ProductReturn) error {
	return defaultUpdate(r.db, ctx, productReturn, "*", nil)
}

func (r *productReturnRepository) Delete(ctx context.Context, productReturn *model.ProductReturn) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productReturn, nil)
}

func (r *productReturnRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReturnTableName)
}
