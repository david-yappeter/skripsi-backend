package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type TransactionItemRepository interface {
	// create
	Insert(ctx context.Context, transactionItem *model.TransactionItem) error
	InsertMany(ctx context.Context, transactionItems []model.TransactionItem, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.TransactionItemQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.TransactionItemQueryOption) ([]model.TransactionItem, error)
	Get(ctx context.Context, id string) (*model.TransactionItem, error)

	// update
	Update(ctx context.Context, transactionItem *model.TransactionItem) error

	// delete
	Delete(ctx context.Context, transactionItem *model.TransactionItem) error
	Truncate(ctx context.Context) error
}

type transactionItemRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewTransactionItemRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) TransactionItemRepository {
	return &transactionItemRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *transactionItemRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.TransactionItem, error) {
	transactionItems := []model.TransactionItem{}
	if err := fetch(r.db, ctx, &transactionItems, stmt); err != nil {
		return nil, err
	}

	return transactionItems, nil
}

func (r *transactionItemRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.TransactionItem, error) {
	transactionItem := model.TransactionItem{}
	if err := get(r.db, ctx, &transactionItem, stmt); err != nil {
		return nil, err
	}

	return &transactionItem, nil
}

func (r *transactionItemRepository) prepareQuery(option model.TransactionItemQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.TransactionItemTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
		})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *transactionItemRepository) Insert(ctx context.Context, transactionItem *model.TransactionItem) error {
	return defaultInsert(r.db, ctx, transactionItem, "*")
}

func (r *transactionItemRepository) InsertMany(ctx context.Context, transactionItems []model.TransactionItem, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range transactionItems {
		arr = append(arr, &transactionItems[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *transactionItemRepository) Count(ctx context.Context, options ...model.TransactionItemQueryOption) (int, error) {
	option := model.TransactionItemQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *transactionItemRepository) Fetch(ctx context.Context, options ...model.TransactionItemQueryOption) ([]model.TransactionItem, error) {
	option := model.TransactionItemQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *transactionItemRepository) Get(ctx context.Context, id string) (*model.TransactionItem, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TransactionItemTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *transactionItemRepository) Update(ctx context.Context, transactionItem *model.TransactionItem) error {
	return defaultUpdate(r.db, ctx, transactionItem, "*", nil)
}

func (r *transactionItemRepository) Delete(ctx context.Context, transactionItem *model.TransactionItem) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, transactionItem, nil)
}

func (r *transactionItemRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.TransactionItemTableName)
}
