package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type ShopeeConfigRepository interface {
	// insert
	InsertMany(ctx context.Context, shopeeConfigs []model.ShopeeConfig, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.ShopeeConfigQueryOption) (int, error)
	Get(ctx context.Context) (*model.ShopeeConfig, error)

	// update
	Update(ctx context.Context, shopeeConfig *model.ShopeeConfig) error
}

type shopeeConfigRepository struct {
	db infrastructure.DBTX
}

func NewShopeeConfigRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) ShopeeConfigRepository {
	return &shopeeConfigRepository{
		db: db,
	}
}

func (r *shopeeConfigRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.ShopeeConfig, error) {
	shopeeConfig := model.ShopeeConfig{}
	if err := get(r.db, ctx, &shopeeConfig, stmt); err != nil {
		return nil, err
	}

	return &shopeeConfig, nil
}

func (r *shopeeConfigRepository) prepareQuery(option model.ShopeeConfigQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s tc", model.ShopeeConfigTableName))

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *shopeeConfigRepository) InsertMany(ctx context.Context, users []model.ShopeeConfig, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range users {
		arr = append(arr, &users[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *shopeeConfigRepository) Count(ctx context.Context, options ...model.ShopeeConfigQueryOption) (int, error) {
	option := model.ShopeeConfigQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *shopeeConfigRepository) Get(ctx context.Context) (*model.ShopeeConfig, error) {
	stmt := stmtBuilder.Select("*").
		From(model.ShopeeConfigTableName).
		Limit(1)

	return r.get(ctx, stmt)
}

func (r *shopeeConfigRepository) Update(ctx context.Context, shopeeConfig *model.ShopeeConfig) error {
	return defaultUpdate(r.db, ctx, shopeeConfig, "*", nil)
}
