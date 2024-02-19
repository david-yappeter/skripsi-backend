package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type FileRepository interface {
	// create
	Insert(ctx context.Context, file *model.File) error
	InsertMany(ctx context.Context, files []model.File, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context) (int, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.File, error)
	Get(ctx context.Context, id string) (*model.File, error)

	// update
	Update(ctx context.Context, file *model.File) error

	// delete
	Delete(ctx context.Context, file *model.File) error
	Truncate(ctx context.Context) error
}

type fileRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewFileRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) FileRepository {
	return &fileRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *fileRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.File, error) {
	files := []model.File{}
	if err := fetch(r.db, ctx, &files, stmt); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *fileRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.File, error) {
	file := model.File{}
	if err := get(r.db, ctx, &file, stmt); err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *fileRepository) Insert(ctx context.Context, file *model.File) error {
	return defaultInsert(r.db, ctx, file, "*")
}

func (r *fileRepository) InsertMany(ctx context.Context, files []model.File, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range files {
		arr = append(arr, &files[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *fileRepository) Count(ctx context.Context) (int, error) {
	stmt := stmtBuilder.Select("COUNT(*) as count").
		From(model.FileTableName)

	return count(r.db, ctx, stmt)
}

func (r *fileRepository) FetchByIds(ctx context.Context, ids []string) ([]model.File, error) {
	stmt := stmtBuilder.Select("*").
		From(model.FileTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *fileRepository) Get(ctx context.Context, id string) (*model.File, error) {
	stmt := stmtBuilder.Select("*").
		From(model.FileTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *fileRepository) Update(ctx context.Context, file *model.File) error {
	return defaultUpdate(r.db, ctx, file, "*", nil)
}

func (r *fileRepository) Delete(ctx context.Context, file *model.File) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, file, nil)
}

func (r *fileRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.FileTableName)
}
