package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type SupplierTypeRepository interface {
	// create
	Insert(ctx context.Context, supplierType *model.SupplierType) error
	InsertMany(ctx context.Context, supplierTypes []model.SupplierType, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.SupplierTypeQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.SupplierTypeQueryOption) ([]model.SupplierType, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.SupplierType, error)
	Get(ctx context.Context, id string) (*model.SupplierType, error)
	IsExistByName(ctx context.Context, name string) (bool, error)

	// update
	Update(ctx context.Context, supplierType *model.SupplierType) error

	// delete
	Delete(ctx context.Context, supplierType *model.SupplierType) error
	Truncate(ctx context.Context) error
}

type supplierTypeRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewSupplierTypeRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) SupplierTypeRepository {
	return &supplierTypeRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *supplierTypeRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.SupplierType, error) {
	supplierTypes := []model.SupplierType{}
	if err := fetch(r.db, ctx, &supplierTypes, stmt); err != nil {
		return nil, err
	}

	return supplierTypes, nil
}

func (r *supplierTypeRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.SupplierType, error) {
	supplierType := model.SupplierType{}
	if err := get(r.db, ctx, &supplierType, stmt); err != nil {
		return nil, err
	}

	return &supplierType, nil
}

func (r *supplierTypeRepository) prepareQuery(option model.SupplierTypeQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.SupplierTypeTableName))

	andStatements := squirrel.And{}

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		andStatements = append(
			andStatements,
			squirrel.Or{
				squirrel.ILike{"u.name": phrase},
			},
		)
	}

	stmt = stmt.Where(andStatements)

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *supplierTypeRepository) Insert(ctx context.Context, supplierType *model.SupplierType) error {
	return defaultInsert(r.db, ctx, supplierType, "*")
}

func (r *supplierTypeRepository) InsertMany(ctx context.Context, supplierTypes []model.SupplierType, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range supplierTypes {
		arr = append(arr, &supplierTypes[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *supplierTypeRepository) Count(ctx context.Context, options ...model.SupplierTypeQueryOption) (int, error) {
	option := model.SupplierTypeQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *supplierTypeRepository) Fetch(ctx context.Context, options ...model.SupplierTypeQueryOption) ([]model.SupplierType, error) {
	option := model.SupplierTypeQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *supplierTypeRepository) FetchByIds(ctx context.Context, ids []string) ([]model.SupplierType, error) {
	stmt := stmtBuilder.Select("*").
		From(model.SupplierTypeTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *supplierTypeRepository) Get(ctx context.Context, id string) (*model.SupplierType, error) {
	stmt := stmtBuilder.Select("*").
		From(model.SupplierTypeTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *supplierTypeRepository) IsExistByName(ctx context.Context, name string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.SupplierTypeTableName).
			Where(squirrel.Eq{"name": name}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *supplierTypeRepository) Update(ctx context.Context, supplierType *model.SupplierType) error {
	return defaultUpdate(r.db, ctx, supplierType, "*", nil)
}

func (r *supplierTypeRepository) Delete(ctx context.Context, supplierType *model.SupplierType) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, supplierType, nil)
}

func (r *supplierTypeRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.SupplierTypeTableName)
}
