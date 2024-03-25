package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReceiveRepository interface {
	// create
	Insert(ctx context.Context, productReceive *model.ProductReceive) error
	InsertMany(ctx context.Context, productReceives []model.ProductReceive, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ProductReceiveQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductReceiveQueryOption) ([]model.ProductReceive, error)
	Get(ctx context.Context, id string) (*model.ProductReceive, error)
	IsExistBySupplierId(ctx context.Context, supplierId string) (bool, error)

	// update
	Update(ctx context.Context, productReceive *model.ProductReceive) error

	// delete
	Delete(ctx context.Context, productReceive *model.ProductReceive) error
	Truncate(ctx context.Context) error
}

type productReceiveRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductReceiveRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReceiveRepository {
	return &productReceiveRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productReceiveRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReceive, error) {
	productReceives := []model.ProductReceive{}
	if err := fetch(r.db, ctx, &productReceives, stmt); err != nil {
		return nil, err
	}

	return productReceives, nil
}

func (r *productReceiveRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReceive, error) {
	productReceive := model.ProductReceive{}
	if err := get(r.db, ctx, &productReceive, stmt); err != nil {
		return nil, err
	}

	return &productReceive, nil
}

func (r *productReceiveRepository) prepareQuery(option model.ProductReceiveQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s pr", model.ProductReceiveTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"pr.invoice_number": phrase},
		})
	}

	if option.SupplierId != nil {
		stmt = stmt.Where(squirrel.Eq{"pr.supplier_id": option.SupplierId})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *productReceiveRepository) Insert(ctx context.Context, productReceive *model.ProductReceive) error {
	return defaultInsert(r.db, ctx, productReceive, "*")
}

func (r *productReceiveRepository) InsertMany(ctx context.Context, productReceives []model.ProductReceive, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReceives {
		arr = append(arr, &productReceives[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReceiveRepository) Count(ctx context.Context, options ...model.ProductReceiveQueryOption) (int, error) {
	option := model.ProductReceiveQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *productReceiveRepository) Fetch(ctx context.Context, options ...model.ProductReceiveQueryOption) ([]model.ProductReceive, error) {
	option := model.ProductReceiveQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *productReceiveRepository) Get(ctx context.Context, id string) (*model.ProductReceive, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productReceiveRepository) IsExistBySupplierId(ctx context.Context, supplierId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductReceiveTableName).
			Where(squirrel.Eq{"supplier_id": supplierId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productReceiveRepository) Update(ctx context.Context, productReceive *model.ProductReceive) error {
	return defaultUpdate(r.db, ctx, productReceive, "*", nil)
}

func (r *productReceiveRepository) Delete(ctx context.Context, productReceive *model.ProductReceive) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productReceive, nil)
}

func (r *productReceiveRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReceiveTableName)
}
