package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type TiktokConfigRepository interface {
	// insert
	InsertMany(ctx context.Context, tiktokConfigs []model.TiktokConfig, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.TiktokConfigQueryOption) (int, error)
	Get(ctx context.Context) (*model.TiktokConfig, error)

	// update
	Update(ctx context.Context, tiktokConfig *model.TiktokConfig) error
}

type tiktokConfigRepository struct {
	db infrastructure.DBTX
}

func NewTiktokConfigRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) TiktokConfigRepository {
	return &tiktokConfigRepository{
		db: db,
	}
}

func (r *tiktokConfigRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.TiktokConfig, error) {
	tiktokConfig := model.TiktokConfig{}
	if err := get(r.db, ctx, &tiktokConfig, stmt); err != nil {
		return nil, err
	}

	return &tiktokConfig, nil
}

func (r *tiktokConfigRepository) prepareQuery(option model.TiktokConfigQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s tc", model.TiktokConfigTableName))

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *tiktokConfigRepository) InsertMany(ctx context.Context, users []model.TiktokConfig, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range users {
		arr = append(arr, &users[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *tiktokConfigRepository) Count(ctx context.Context, options ...model.TiktokConfigQueryOption) (int, error) {
	option := model.TiktokConfigQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *tiktokConfigRepository) Get(ctx context.Context) (*model.TiktokConfig, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TiktokConfigTableName).
		Limit(1)

	return r.get(ctx, stmt)
}

func (r *tiktokConfigRepository) Update(ctx context.Context, tiktokConfig *model.TiktokConfig) error {
	return defaultUpdate(r.db, ctx, tiktokConfig, "*", nil)
}
