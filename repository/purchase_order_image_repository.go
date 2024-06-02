package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type PurchaseOrderImageRepository interface {
	// create
	Insert(ctx context.Context, purchaseOrderImage *model.PurchaseOrderImage) error
	InsertMany(ctx context.Context, purchaseOrderImages []model.PurchaseOrderImage, options ...data_type.RepositoryOption) error

	// read
	FetchByPurchaseOrderIds(ctx context.Context, purchaseOrderIds []string) ([]model.PurchaseOrderImage, error)
	Get(ctx context.Context, id string) (*model.PurchaseOrderImage, error)
	GetByPurchaseOrderIdAndFileId(ctx context.Context, purchaseOrderId string, fileId string) (*model.PurchaseOrderImage, error)
	IsExistByName(ctx context.Context, name string) (bool, error)
	IsExistByPurchaseOrderIds(ctx context.Context, purchaseOrderIds []string) (bool, error)

	// update
	Update(ctx context.Context, purchaseOrderImage *model.PurchaseOrderImage) error

	// delete
	Delete(ctx context.Context, purchaseOrderImage *model.PurchaseOrderImage) error
	DeleteManyByPurchaseOrderId(ctx context.Context, purchaseOrderId string) error
	Truncate(ctx context.Context) error
}

type purchaseOrderImageRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewPurchaseOrderImageRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) PurchaseOrderImageRepository {
	return &purchaseOrderImageRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *purchaseOrderImageRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.PurchaseOrderImage, error) {
	purchaseOrderImages := []model.PurchaseOrderImage{}
	if err := fetch(r.db, ctx, &purchaseOrderImages, stmt); err != nil {
		return nil, err
	}

	return purchaseOrderImages, nil
}

func (r *purchaseOrderImageRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.PurchaseOrderImage, error) {
	purchaseOrderImage := model.PurchaseOrderImage{}
	if err := get(r.db, ctx, &purchaseOrderImage, stmt); err != nil {
		return nil, err
	}

	return &purchaseOrderImage, nil
}

func (r *purchaseOrderImageRepository) Insert(ctx context.Context, purchaseOrderImage *model.PurchaseOrderImage) error {
	return defaultInsert(r.db, ctx, purchaseOrderImage, "*")
}

func (r *purchaseOrderImageRepository) InsertMany(ctx context.Context, purchaseOrderImages []model.PurchaseOrderImage, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range purchaseOrderImages {
		arr = append(arr, &purchaseOrderImages[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *purchaseOrderImageRepository) FetchByPurchaseOrderIds(ctx context.Context, purchaseOrderIds []string) ([]model.PurchaseOrderImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PurchaseOrderImageTableName).
		Where(squirrel.Eq{"purchase_order_id": purchaseOrderIds})

	return r.fetch(ctx, stmt)
}
func (r *purchaseOrderImageRepository) Get(ctx context.Context, id string) (*model.PurchaseOrderImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PurchaseOrderImageTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *purchaseOrderImageRepository) GetByPurchaseOrderIdAndFileId(ctx context.Context, purchaseOrderId string, fileId string) (*model.PurchaseOrderImage, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PurchaseOrderImageTableName).
		Where(squirrel.Eq{"purchase_order_id": purchaseOrderId}).
		Where(squirrel.Eq{"file_id": fileId})

	return r.get(ctx, stmt)
}

func (r *purchaseOrderImageRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.PurchaseOrderImageTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *purchaseOrderImageRepository) IsExistByPurchaseOrderIds(ctx context.Context, purchaseOrderIds []string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.PurchaseOrderImageTableName).
			Where(squirrel.Eq{"purchase_order_id": purchaseOrderIds}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *purchaseOrderImageRepository) Update(ctx context.Context, purchaseOrderImage *model.PurchaseOrderImage) error {
	return defaultUpdate(r.db, ctx, purchaseOrderImage, "*", nil)
}

func (r *purchaseOrderImageRepository) Delete(ctx context.Context, purchaseOrderImage *model.PurchaseOrderImage) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, purchaseOrderImage, nil)
}

func (r *purchaseOrderImageRepository) DeleteManyByPurchaseOrderId(ctx context.Context, purchaseOrderId string) error {
	whereStmt := squirrel.Eq{
		"purchase_order_id": purchaseOrderId,
	}
	return destroy(r.db, ctx, model.PurchaseOrderImageTableName, whereStmt)
}

func (r *purchaseOrderImageRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.PurchaseOrderImageTableName)
}
