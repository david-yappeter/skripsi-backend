package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type SupplierRepository interface {
	// create
	Insert(ctx context.Context, supplier *model.Supplier) error
	InsertMany(ctx context.Context, suppliers []model.Supplier, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.SupplierQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.SupplierQueryOption) ([]model.Supplier, error)
	Get(ctx context.Context, id string) (*model.Supplier, error)
	IsExistByCode(ctx context.Context, code string) (bool, error)
	IsExistBySupplierTypeId(ctx context.Context, supplierTypeId string) (bool, error)

	// update
	Update(ctx context.Context, supplier *model.Supplier) error

	// delete
	Delete(ctx context.Context, supplier *model.Supplier) error
	Truncate(ctx context.Context) error
}

type supplierRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewSupplierRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) SupplierRepository {
	return &supplierRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *supplierRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	if err := fetch(r.db, ctx, &suppliers, stmt); err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *supplierRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Supplier, error) {
	supplier := model.Supplier{}
	if err := get(r.db, ctx, &supplier, stmt); err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (r *supplierRepository) prepareQuery(option model.SupplierQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s s", model.SupplierTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"s.name": phrase},
		})
	}

	if option.IsActive != nil {
		stmt = stmt.Where(
			squirrel.Eq{"s.is_active": option.IsActive},
		)
	}

	if len(option.SupplierTypeIds) > 0 {
		stmt = stmt.Where(
			squirrel.Eq{"s.supplier_type_id": option.SupplierTypeIds},
		)
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *supplierRepository) Insert(ctx context.Context, supplier *model.Supplier) error {
	return defaultInsert(r.db, ctx, supplier, "*")
}

func (r *supplierRepository) InsertMany(ctx context.Context, suppliers []model.Supplier, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range suppliers {
		arr = append(arr, &suppliers[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *supplierRepository) Count(ctx context.Context, options ...model.SupplierQueryOption) (int, error) {
	option := model.SupplierQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *supplierRepository) Fetch(ctx context.Context, options ...model.SupplierQueryOption) ([]model.Supplier, error) {
	option := model.SupplierQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *supplierRepository) Get(ctx context.Context, id string) (*model.Supplier, error) {
	stmt := stmtBuilder.Select("*").
		From(model.SupplierTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *supplierRepository) IsExistByCode(ctx context.Context, code string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("*").
			From(model.SupplierTableName).
			Where(squirrel.Eq{"code": code}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *supplierRepository) IsExistBySupplierTypeId(ctx context.Context, supplierTypeId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("*").
			From(model.SupplierTableName).
			Where(squirrel.Eq{"supplier_type_id": supplierTypeId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *supplierRepository) Update(ctx context.Context, supplier *model.Supplier) error {
	return defaultUpdate(r.db, ctx, supplier, "*", nil)
}

func (r *supplierRepository) Delete(ctx context.Context, supplier *model.Supplier) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, supplier, nil)
}

func (r *supplierRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.SupplierTableName)
}
