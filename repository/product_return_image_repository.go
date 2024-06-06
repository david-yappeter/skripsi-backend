package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReturnImageRepository interface {
	// create
	Insert(ctx context.Context, productReturnImage *model.ProductReturnImage) error
	InsertMany(ctx context.Context, productReturnImages []model.ProductReturnImage, options ...data_type.RepositoryOption) error

	// read
	FetchByProductReturnIds(ctx context.Context, productReturnIds []string) ([]model.ProductReturnImage, error)
	Get(ctx context.Context, id string) (*model.ProductReturnImage, error)
	GetByProductReturnIdAndFileId(ctx context.Context, productReturnId string, fileId string) (*model.ProductReturnImage, error)
	IsExistByName(ctx context.Context, name string) (bool, error)
	IsExistByProductReturnIds(ctx context.Context, productReturnIds []string) (bool, error)

	// update
	Update(ctx context.Context, productReturnImage *model.ProductReturnImage) error

	// delete
	Delete(ctx context.Context, productReturnImage *model.ProductReturnImage) error
	DeleteManyByProductReturnId(ctx context.Context, productReturnId string) error
	Truncate(ctx context.Context) error
}

type productReturnImageRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductReturnImageRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReturnImageRepository {
	return &productReturnImageRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productReturnImageRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReturnImage, error) {
	productReturnImages := []model.ProductReturnImage{}
	if err := fetch(r.db, ctx, &productReturnImages, stmt); err != nil {
		return nil, err
	}

	return productReturnImages, nil
}

func (r *productReturnImageRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReturnImage, error) {
	productReturnImage := model.ProductReturnImage{}
	if err := get(r.db, ctx, &productReturnImage, stmt); err != nil {
		return nil, err
	}

	return &productReturnImage, nil
}

func (r *productReturnImageRepository) Insert(ctx context.Context, productReturnImage *model.ProductReturnImage) error {
	return defaultInsert(r.db, ctx, productReturnImage, "*")
}

func (r *productReturnImageRepository) InsertMany(ctx context.Context, productReturnImages []model.ProductReturnImage, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReturnImages {
		arr = append(arr, &productReturnImages[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReturnImageRepository) FetchByProductReturnIds(ctx context.Context, productReturnIds []string) ([]model.ProductReturnImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReturnImageTableName).
		Where(squirrel.Eq{"product_return_id": productReturnIds})

	return r.fetch(ctx, stmt)
}
func (r *productReturnImageRepository) Get(ctx context.Context, id string) (*model.ProductReturnImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReturnImageTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productReturnImageRepository) GetByProductReturnIdAndFileId(ctx context.Context, productReturnId string, fileId string) (*model.ProductReturnImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReturnImageTableName).
		Where(squirrel.Eq{"product_return_id": productReturnId}).
		Where(squirrel.Eq{"file_id": fileId})

	return r.get(ctx, stmt)
}

func (r *productReturnImageRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductReturnImageTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productReturnImageRepository) IsExistByProductReturnIds(ctx context.Context, productReturnIds []string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductReturnImageTableName).
			Where(squirrel.Eq{"product_return_id": productReturnIds}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productReturnImageRepository) Update(ctx context.Context, productReturnImage *model.ProductReturnImage) error {
	return defaultUpdate(r.db, ctx, productReturnImage, "*", nil)
}

func (r *productReturnImageRepository) Delete(ctx context.Context, productReturnImage *model.ProductReturnImage) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productReturnImage, nil)
}

func (r *productReturnImageRepository) DeleteManyByProductReturnId(ctx context.Context, productReturnId string) error {
	whereStmt := squirrel.Eq{
		"product_return_id": productReturnId,
	}
	return destroy(r.db, ctx, model.ProductReturnImageTableName, whereStmt)
}

func (r *productReturnImageRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReturnImageTableName)
}
