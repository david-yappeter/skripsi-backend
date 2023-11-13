package repository

import (
	"context"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type UserAccessTokenRepository interface {
	// create
	Insert(ctx context.Context, userAccessToken *model.UserAccessToken) error

	// read
	FetchByUserIdAndRevokedAndGetExpiredAt(ctx context.Context, userId string, revoked bool, expiredAt data_type.DateTime) ([]model.UserAccessToken, error)
	Get(ctx context.Context, id string) (*model.UserAccessToken, error)

	// update
	UpdateRevoked(ctx context.Context, userAccessToken *model.UserAccessToken) error
	UpdateManyRevokedByIds(ctx context.Context, revoked bool, ids []string) error
}

type userAccessTokenRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewUserAccessTokenRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) UserAccessTokenRepository {
	return &userAccessTokenRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *userAccessTokenRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.UserAccessToken, error) {
	userAccessTokens := []model.UserAccessToken{}
	if err := fetch(r.db, ctx, &userAccessTokens, stmt); err != nil {
		return nil, err
	}

	return userAccessTokens, nil
}

func (r *userAccessTokenRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.UserAccessToken, error) {
	userAccessToken := model.UserAccessToken{}
	if err := get(r.db, ctx, &userAccessToken, stmt); err != nil {
		return nil, err
	}

	return &userAccessToken, nil
}

func (r *userAccessTokenRepository) Insert(ctx context.Context, userAccessToken *model.UserAccessToken) error {
	return defaultInsert(r.db, ctx, userAccessToken, "*")
}

func (r *userAccessTokenRepository) FetchByUserIdAndRevokedAndGetExpiredAt(ctx context.Context, userId string, revoked bool, expiredAt data_type.DateTime) ([]model.UserAccessToken, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserAccessTokenTableName).
		Where(
			squirrel.And{
				squirrel.Eq{"user_id": userId},
				squirrel.Eq{"revoked": revoked},
				squirrel.GtOrEq{"expired_at": expiredAt},
			},
		)

	return r.fetch(ctx, stmt)
}

func (r *userAccessTokenRepository) Get(ctx context.Context, id string) (*model.UserAccessToken, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserAccessTokenTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *userAccessTokenRepository) UpdateRevoked(ctx context.Context, userAccessToken *model.UserAccessToken) error {
	return defaultUpdate(r.db, ctx, userAccessToken, "revoked", nil)
}

func (r *userAccessTokenRepository) UpdateManyRevokedByIds(ctx context.Context, revoked bool, ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	args := map[string]interface{}{
		"revoked": revoked,
	}
	whereStmt := squirrel.Eq{"id": ids}

	return update(r.db, ctx, model.UserAccessTokenTableName, args, whereStmt)
}
