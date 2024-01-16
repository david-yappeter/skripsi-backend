package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type UserRoleRepository interface {
	// create
	InsertMany(ctx context.Context, userRoles []model.UserRole) error

	// read
	Count(ctx context.Context) (int, error)
	FetchByUserId(ctx context.Context, id string) ([]model.UserRole, error)
	FetchByUserIds(ctx context.Context, ids []string) ([]model.UserRole, error)
	FetchUserIdsByRoleId(ctx context.Context, roleId string) ([]string, error)
	FetchUserIdsByRoleIdsAndUserIds(ctx context.Context, roleIds []string, userIds []string) ([]string, error)
	GetByUserIdAndRoleTitle(ctx context.Context, userId string, roleTitle data_type.Role) (*model.UserRole, error)
	IsExistByUserIdAndRoleTitles(ctx context.Context, userId string, roleTitles []data_type.Role) (bool, error)

	// delete
	DeleteManyByUserIdAndRoleIds(ctx context.Context, userId string, roleIds []string) error
	Truncate(ctx context.Context) error
}

type userRoleRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewUserRoleRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) UserRoleRepository {
	return &userRoleRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *userRoleRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.UserRole, error) {
	userRoles := []model.UserRole{}
	if err := fetch(r.db, ctx, &userRoles, stmt); err != nil {
		return nil, err
	}

	return userRoles, nil
}

func (r *userRoleRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.UserRole, error) {
	userRole := model.UserRole{}
	if err := get(r.db, ctx, &userRole, stmt); err != nil {
		return nil, err
	}

	return &userRole, nil
}

func (r *userRoleRepository) InsertMany(ctx context.Context, userRoles []model.UserRole) error {
	arr := []model.BaseModel{}
	for i := range userRoles {
		arr = append(arr, &userRoles[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *userRoleRepository) Count(ctx context.Context) (int, error) {
	stmt := stmtBuilder.Select("COUNT(*) row_count").
		From(model.UserRoleTableName)

	return count(r.db, ctx, stmt)
}

func (r *userRoleRepository) FetchByUserId(ctx context.Context, id string) ([]model.UserRole, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserRoleTableName).
		Where("user_id = ?", id)

	return r.fetch(ctx, stmt)
}

func (r *userRoleRepository) FetchByUserIds(ctx context.Context, ids []string) ([]model.UserRole, error) {
	if len(ids) == 0 {
		return []model.UserRole{}, nil
	}

	stmt := stmtBuilder.Select("*").
		From(model.UserRoleTableName).
		Where(squirrel.Eq{"user_id": ids})

	return r.fetch(ctx, stmt)
}

func (r *userRoleRepository) FetchUserIdsByRoleId(ctx context.Context, roleId string) ([]string, error) {
	stmt := stmtBuilder.Select("user_id").
		From(model.UserRoleTableName).
		Where("role_id = ?", roleId)

	userIds := []string{}
	if err := fetch(r.db, ctx, &userIds, stmt); err != nil {
		return nil, err
	}

	return userIds, nil
}

func (r *userRoleRepository) FetchUserIdsByRoleIdsAndUserIds(ctx context.Context, roleIds []string, userIds []string) ([]string, error) {
	if len(roleIds) == 0 || len(userIds) == 0 {
		return []string{}, nil
	}

	stmt := stmtBuilder.Select("user_id").
		From(model.UserRoleTableName).
		Where(squirrel.Eq{"role_id": roleIds}).
		Where(squirrel.Eq{"user_id": userIds})

	resultUserIds := []string{}
	if err := fetch(r.db, ctx, &resultUserIds, stmt); err != nil {
		return nil, err
	}

	return resultUserIds, nil
}

func (r *userRoleRepository) GetByUserIdAndRoleTitle(ctx context.Context, userId string, roleTitle data_type.Role) (*model.UserRole, error) {
	stmt := stmtBuilder.Select("ur.*").
		From(fmt.Sprintf("%s ur", model.UserRoleTableName)).
		InnerJoin(fmt.Sprintf("%s r ON r.id = ur.role_id", model.RoleTableName)).
		Where(squirrel.And{
			squirrel.Eq{"ur.user_id": userId},
			squirrel.Eq{"r.title": roleTitle},
		})

	return r.get(ctx, stmt)
}

func (r *userRoleRepository) IsExistByUserIdAndRoleTitles(ctx context.Context, userId string, roleTitles []data_type.Role) (bool, error) {
	if len(roleTitles) == 0 {
		return false, nil
	}

	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("ur.*").
			From(fmt.Sprintf("%s ur", model.UserRoleTableName)).
			InnerJoin(fmt.Sprintf("%s r ON r.id = ur.role_id", model.RoleTableName)).
			Where(squirrel.And{
				squirrel.Eq{"ur.user_id": userId},
				squirrel.Eq{"r.title": roleTitles},
			}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *userRoleRepository) DeleteManyByUserIdAndRoleIds(ctx context.Context, userId string, roleIds []string) error {
	if len(roleIds) == 0 {
		return nil
	}

	whereStmt := squirrel.Eq{
		"user_id": userId,
		"role_id": roleIds,
	}

	return destroy(r.db, ctx, model.UserRoleTableName, whereStmt)
}

func (r *userRoleRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.UserRoleTableName)
}