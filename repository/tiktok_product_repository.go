package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type TiktokProductRepository interface {
	// create
	Insert(ctx context.Context, tiktokProduct *model.TiktokProduct) error
	InsertMany(ctx context.Context, tiktokProducts []model.TiktokProduct, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.TiktokProductQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.TiktokProductQueryOption) ([]model.TiktokProduct, error)
	Get(ctx context.Context, tiktokProductId string) (*model.TiktokProduct, error)
	GetByProductId(ctx context.Context, productId string) (*model.TiktokProduct, error)
	IsExistByProductId(ctx context.Context, productId string) (bool, error)

	// update
	Update(ctx context.Context, tiktokProduct *model.TiktokProduct) error

	// delete
	Delete(ctx context.Context, tiktokProduct *model.TiktokProduct) error
	Truncate(ctx context.Context) error
}

type tiktokProductRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewTiktokProductRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) TiktokProductRepository {
	return &tiktokProductRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *tiktokProductRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.TiktokProduct, error) {
	tiktokProducts := []model.TiktokProduct{}
	if err := fetch(r.db, ctx, &tiktokProducts, stmt); err != nil {
		return nil, err
	}

	return tiktokProducts, nil
}

func (r *tiktokProductRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.TiktokProduct, error) {
	tiktokProduct := model.TiktokProduct{}
	if err := get(r.db, ctx, &tiktokProduct, stmt); err != nil {
		return nil, err
	}

	return &tiktokProduct, nil
}

func (r *tiktokProductRepository) prepareQuery(option model.TiktokProductQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.TiktokProductTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *tiktokProductRepository) Insert(ctx context.Context, tiktokProduct *model.TiktokProduct) error {
	return defaultInsert(r.db, ctx, tiktokProduct, "*")
}

func (r *tiktokProductRepository) InsertMany(ctx context.Context, tiktokProducts []model.TiktokProduct, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range tiktokProducts {
		arr = append(arr, &tiktokProducts[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *tiktokProductRepository) Count(ctx context.Context, options ...model.TiktokProductQueryOption) (int, error) {
	option := model.TiktokProductQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *tiktokProductRepository) Fetch(ctx context.Context, options ...model.TiktokProductQueryOption) ([]model.TiktokProduct, error) {
	option := model.TiktokProductQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *tiktokProductRepository) Get(ctx context.Context, tiktokProductId string) (*model.TiktokProduct, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TiktokProductTableName).
		Where(squirrel.Eq{"tiktok_product_id": tiktokProductId})

	return r.get(ctx, stmt)
}

func (r *tiktokProductRepository) GetByProductId(ctx context.Context, productId string) (*model.TiktokProduct, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TiktokProductTableName).
		Where(squirrel.Eq{"product_id": productId})

	return r.get(ctx, stmt)
}

func (r *tiktokProductRepository) IsExistByProductId(ctx context.Context, productId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.TiktokProductTableName).
			Where(squirrel.Eq{"product_id": productId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *tiktokProductRepository) Update(ctx context.Context, tiktokProduct *model.TiktokProduct) error {
	return defaultUpdate(r.db, ctx, tiktokProduct, "*", nil)
}

func (r *tiktokProductRepository) Delete(ctx context.Context, tiktokProduct *model.TiktokProduct) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, tiktokProduct, nil)
}

func (r *tiktokProductRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.TiktokProductTableName)
}
