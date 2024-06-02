package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type PurchaseOrderRepository interface {
	// create
	Insert(ctx context.Context, purchaseOrder *model.PurchaseOrder) error
	InsertMany(ctx context.Context, purchaseOrders []model.PurchaseOrder, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.PurchaseOrderQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.PurchaseOrderQueryOption) ([]model.PurchaseOrder, error)
	Get(ctx context.Context, id string) (*model.PurchaseOrder, error)
	IsExistBySupplierId(ctx context.Context, supplierId string) (bool, error)

	// update
	Update(ctx context.Context, purchaseOrder *model.PurchaseOrder) error

	// delete
	Delete(ctx context.Context, purchaseOrder *model.PurchaseOrder) error
	Truncate(ctx context.Context) error
}

type purchaseOrderRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewPurchaseOrderRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) PurchaseOrderRepository {
	return &purchaseOrderRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *purchaseOrderRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.PurchaseOrder, error) {
	purchaseOrders := []model.PurchaseOrder{}
	if err := fetch(r.db, ctx, &purchaseOrders, stmt); err != nil {
		return nil, err
	}

	return purchaseOrders, nil
}

func (r *purchaseOrderRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.PurchaseOrder, error) {
	purchaseOrder := model.PurchaseOrder{}
	if err := get(r.db, ctx, &purchaseOrder, stmt); err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}

func (r *purchaseOrderRepository) prepareQuery(option model.PurchaseOrderQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s po", model.PurchaseOrderTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"po.invoice_number": phrase},
		})
	}

	if option.Status != nil {
		stmt = stmt.Where(squirrel.Eq{"po.status": option.Status})
	}

	if option.SupplierId != nil {
		stmt = stmt.Where(squirrel.Eq{"po.supplier_id": option.SupplierId})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *purchaseOrderRepository) Insert(ctx context.Context, purchaseOrder *model.PurchaseOrder) error {
	return defaultInsert(r.db, ctx, purchaseOrder, "*")
}

func (r *purchaseOrderRepository) InsertMany(ctx context.Context, purchaseOrders []model.PurchaseOrder, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range purchaseOrders {
		arr = append(arr, &purchaseOrders[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *purchaseOrderRepository) Count(ctx context.Context, options ...model.PurchaseOrderQueryOption) (int, error) {
	option := model.PurchaseOrderQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *purchaseOrderRepository) Fetch(ctx context.Context, options ...model.PurchaseOrderQueryOption) ([]model.PurchaseOrder, error) {
	option := model.PurchaseOrderQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *purchaseOrderRepository) Get(ctx context.Context, id string) (*model.PurchaseOrder, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PurchaseOrderTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *purchaseOrderRepository) IsExistBySupplierId(ctx context.Context, supplierId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.PurchaseOrderTableName).
			Where(squirrel.Eq{"supplier_id": supplierId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *purchaseOrderRepository) Update(ctx context.Context, purchaseOrder *model.PurchaseOrder) error {
	return defaultUpdate(r.db, ctx, purchaseOrder, "*", nil)
}

func (r *purchaseOrderRepository) Delete(ctx context.Context, purchaseOrder *model.PurchaseOrder) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, purchaseOrder, nil)
}

func (r *purchaseOrderRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.PurchaseOrderTableName)
}
