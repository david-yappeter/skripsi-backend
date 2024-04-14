package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type SequenceRepository interface {
	// create
	Insert(ctx context.Context, sequence *model.Sequence) error
	InsertMany(ctx context.Context, sequences []model.Sequence, options ...data_type.RepositoryOption) error

	// read
	GetLatestByUniqueIdentifier(ctx context.Context, uniqueIdentifier string) (*model.Sequence, error)
}

type sequenceRepository struct {
	db infrastructure.DBTX
}

func NewSequenceRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) SequenceRepository {
	return &sequenceRepository{
		db: db,
	}
}

func (r *sequenceRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Sequence, error) {
	sequences := []model.Sequence{}
	if err := fetch(r.db, ctx, &sequences, stmt); err != nil {
		return nil, err
	}

	return sequences, nil
}

func (r *sequenceRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Sequence, error) {
	sequence := model.Sequence{}
	if err := get(r.db, ctx, &sequence, stmt); err != nil {
		return nil, err
	}

	return &sequence, nil
}

func (r *sequenceRepository) Insert(ctx context.Context, sequence *model.Sequence) error {
	return defaultInsert(r.db, ctx, sequence, "*")
}

func (r *sequenceRepository) InsertMany(ctx context.Context, sequences []model.Sequence, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range sequences {
		arr = append(arr, &sequences[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *sequenceRepository) GetLatestByUniqueIdentifier(ctx context.Context, uniqueIdentifier string) (*model.Sequence, error) {
	stmt := stmtBuilder.Select("*").
		Where(squirrel.Eq{"unique_identifier": uniqueIdentifier}).
		Limit(1).
		OrderBy("sequence DESC")

	return r.get(ctx, stmt)
}
