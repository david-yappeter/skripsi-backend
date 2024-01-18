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
	GetByUserIdAndRoleId(ctx context.Context, userId string, roleId string) (*model.UserRole, error)
	GetByUserIdAndRoleName(ctx context.Context, userId string, roleName data_type.Role) (*model.UserRole, error)
	IsExistByUserIdAndRoleNames(ctx context.Context, userId string, roleNames []data_type.Role) (bool, error)

	// delete
	Delete(ctx context.Context, userRole *model.UserRole) error
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

func (r *userRoleRepository) GetByUserIdAndRoleId(ctx context.Context, userId string, roleId string) (*model.UserRole, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserRoleTableName).
		Where(squirrel.Eq{
			"user_id": userId,
			"role_id": roleId,
		})

	return r.get(ctx, stmt)
}

func (r *userRoleRepository) GetByUserIdAndRoleName(ctx context.Context, userId string, roleNames data_type.Role) (*model.UserRole, error) {
	stmt := stmtBuilder.Select("ur.*").
		From(fmt.Sprintf("%s ur", model.UserRoleTableName)).
		InnerJoin(fmt.Sprintf("%s r ON r.id = ur.role_id", model.RoleTableName)).
		Where(squirrel.And{
			squirrel.Eq{"ur.user_id": userId},
			squirrel.Eq{"r.name": roleNames},
		})

	return r.get(ctx, stmt)
}

func (r *userRoleRepository) IsExistByUserIdAndRoleNames(ctx context.Context, userId string, roleNames []data_type.Role) (bool, error) {
	if len(roleNames) == 0 {
		return false, nil
	}

	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("ur.*").
			From(fmt.Sprintf("%s ur", model.UserRoleTableName)).
			InnerJoin(fmt.Sprintf("%s r ON r.id = ur.role_id", model.RoleTableName)).
			Where(squirrel.And{
				squirrel.Eq{"ur.user_id": userId},
				squirrel.Eq{"r.name": roleNames},
			}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *userRoleRepository) Delete(ctx context.Context, userRole *model.UserRole) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, userRole, nil)
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
