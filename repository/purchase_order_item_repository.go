package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type PurchaseOrderItemRepository interface {
	// create
	Insert(ctx context.Context, purchaseOrderItem *model.PurchaseOrderItem) error
	InsertMany(ctx context.Context, purchaseOrderItems []model.PurchaseOrderItem, options ...data_type.RepositoryOption) error

	// read
	FetchByPurchaseOrderIds(ctx context.Context, purchaseOrderIds []string) ([]model.PurchaseOrderItem, error)
	Get(ctx context.Context, id string) (*model.PurchaseOrderItem, error)
	GetByPurchaseOrderIdAndProductUnitId(ctx context.Context, purchaseOrderId string, productUnitId string) (*model.PurchaseOrderItem, error)
	IsExistByProductIdAndHavePurchaseOrder(ctx context.Context, productId string) (bool, error)

	// update
	Update(ctx context.Context, purchaseOrderItem *model.PurchaseOrderItem) error

	// delete
	Delete(ctx context.Context, purchaseOrderItem *model.PurchaseOrderItem) error
	DeleteManyByPurchaseOrderId(ctx context.Context, purchaseOrderId string) error
	Truncate(ctx context.Context) error
}

type purchaseOrderItemRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewPurchaseOrderItemRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) PurchaseOrderItemRepository {
	return &purchaseOrderItemRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *purchaseOrderItemRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.PurchaseOrderItem, error) {
	purchaseOrderItems := []model.PurchaseOrderItem{}
	if err := fetch(r.db, ctx, &purchaseOrderItems, stmt); err != nil {
		return nil, err
	}

	return purchaseOrderItems, nil
}

func (r *purchaseOrderItemRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.PurchaseOrderItem, error) {
	purchaseOrderItem := model.PurchaseOrderItem{}
	if err := get(r.db, ctx, &purchaseOrderItem, stmt); err != nil {
		return nil, err
	}

	return &purchaseOrderItem, nil
}

func (r *purchaseOrderItemRepository) Insert(ctx context.Context, purchaseOrderItem *model.PurchaseOrderItem) error {
	return defaultInsert(r.db, ctx, purchaseOrderItem, "*")
}

func (r *purchaseOrderItemRepository) InsertMany(ctx context.Context, purchaseOrderItems []model.PurchaseOrderItem, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range purchaseOrderItems {
		arr = append(arr, &purchaseOrderItems[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *purchaseOrderItemRepository) FetchByPurchaseOrderIds(ctx context.Context, purchaseOrderIds []string) ([]model.PurchaseOrderItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PurchaseOrderItemTableName).
		Where(squirrel.Eq{"purchase_order_id": purchaseOrderIds})

	return r.fetch(ctx, stmt)
}

func (r *purchaseOrderItemRepository) Get(ctx context.Context, id string) (*model.PurchaseOrderItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PurchaseOrderItemTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *purchaseOrderItemRepository) GetByPurchaseOrderIdAndProductUnitId(ctx context.Context, purchaseOrderId string, productUnitId string) (*model.PurchaseOrderItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PurchaseOrderItemTableName).
		Where(squirrel.Eq{"purchase_order_id": purchaseOrderId}).
		Where(squirrel.Eq{"product_unit_id": productUnitId})

	return r.get(ctx, stmt)
}

func (r *purchaseOrderItemRepository) IsExistByProductIdAndHavePurchaseOrder(ctx context.Context, productId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(fmt.Sprintf("%s pri", model.PurchaseOrderItemTableName)).
			InnerJoin(fmt.Sprintf("%s pr ON pr.id = pri.purchase_order_id", model.PurchaseOrderTableName)).
			InnerJoin(fmt.Sprintf("%s pu ON pu.id = pri.product_unit_id", model.ProductUnitTableName)).
			Where(squirrel.Eq{"pu.product_id": productId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *purchaseOrderItemRepository) Update(ctx context.Context, purchaseOrderItem *model.PurchaseOrderItem) error {
	return defaultUpdate(r.db, ctx, purchaseOrderItem, "*", nil)
}

func (r *purchaseOrderItemRepository) Delete(ctx context.Context, purchaseOrderItem *model.PurchaseOrderItem) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, purchaseOrderItem, nil)
}

func (r *purchaseOrderItemRepository) DeleteManyByPurchaseOrderId(ctx context.Context, purchaseOrderId string) error {
	whereStmt := squirrel.Eq{
		"purchase_order_id": purchaseOrderId,
	}
	return destroy(r.db, ctx, model.PurchaseOrderItemTableName, whereStmt)
}

func (r *purchaseOrderItemRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.PurchaseOrderItemTableName)
}
