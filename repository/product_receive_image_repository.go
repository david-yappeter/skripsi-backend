package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ProductReceiveImageRepository interface {
	// create
	Insert(ctx context.Context, productReceiveImage *model.ProductReceiveImage) error
	InsertMany(ctx context.Context, productReceiveImages []model.ProductReceiveImage, options ...data_type.RepositoryOption) error

	// read
	FetchByProductReceiveIds(ctx context.Context, productReceiveIds []string) ([]model.ProductReceiveImage, error)
	Get(ctx context.Context, id string) (*model.ProductReceiveImage, error)
	GetByProductReceiveIdAndFileId(ctx context.Context, productReceiveId string, fileId string) (*model.ProductReceiveImage, error)
	IsExistByName(ctx context.Context, name string) (bool, error)
	IsExistByProductReceiveIds(ctx context.Context, productReceiveIds []string) (bool, error)

	// update
	Update(ctx context.Context, productReceiveImage *model.ProductReceiveImage) error

	// delete
	Delete(ctx context.Context, productReceiveImage *model.ProductReceiveImage) error
	DeleteManyByProductReceiveId(ctx context.Context, productReceiveId string) error
	Truncate(ctx context.Context) error
}

type productReceiveImageRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewProductReceiveImageRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ProductReceiveImageRepository {
	return &productReceiveImageRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *productReceiveImageRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.ProductReceiveImage, error) {
	productReceiveImages := []model.ProductReceiveImage{}
	if err := fetch(r.db, ctx, &productReceiveImages, stmt); err != nil {
		return nil, err
	}

	return productReceiveImages, nil
}

func (r *productReceiveImageRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ProductReceiveImage, error) {
	productReceiveImage := model.ProductReceiveImage{}
	if err := get(r.db, ctx, &productReceiveImage, stmt); err != nil {
		return nil, err
	}

	return &productReceiveImage, nil
}

func (r *productReceiveImageRepository) Insert(ctx context.Context, productReceiveImage *model.ProductReceiveImage) error {
	return defaultInsert(r.db, ctx, productReceiveImage, "*")
}

func (r *productReceiveImageRepository) InsertMany(ctx context.Context, productReceiveImages []model.ProductReceiveImage, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range productReceiveImages {
		arr = append(arr, &productReceiveImages[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *productReceiveImageRepository) FetchByProductReceiveIds(ctx context.Context, productReceiveIds []string) ([]model.ProductReceiveImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveImageTableName).
		Where(squirrel.Eq{"product_receive_id": productReceiveIds})

	return r.fetch(ctx, stmt)
}
func (r *productReceiveImageRepository) Get(ctx context.Context, id string) (*model.ProductReceiveImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveImageTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *productReceiveImageRepository) GetByProductReceiveIdAndFileId(ctx context.Context, productReceiveId string, fileId string) (*model.ProductReceiveImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ProductReceiveImageTableName).
		Where(squirrel.Eq{"product_receive_id": productReceiveId}).
		Where(squirrel.Eq{"file_id": fileId})

	return r.get(ctx, stmt)
}

func (r *productReceiveImageRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductReceiveImageTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productReceiveImageRepository) IsExistByProductReceiveIds(ctx context.Context, productReceiveIds []string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.ProductReceiveImageTableName).
			Where(squirrel.Eq{"product_receive_id": productReceiveIds}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *productReceiveImageRepository) Update(ctx context.Context, productReceiveImage *model.ProductReceiveImage) error {
	return defaultUpdate(r.db, ctx, productReceiveImage, "*", nil)
}

func (r *productReceiveImageRepository) Delete(ctx context.Context, productReceiveImage *model.ProductReceiveImage) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, productReceiveImage, nil)
}

func (r *productReceiveImageRepository) DeleteManyByProductReceiveId(ctx context.Context, productReceiveId string) error {
	whereStmt := squirrel.Eq{
		"product_receive_id": productReceiveId,
	}
	return destroy(r.db, ctx, model.ProductReceiveImageTableName, whereStmt)
}

func (r *productReceiveImageRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.ProductReceiveImageTableName)
}
