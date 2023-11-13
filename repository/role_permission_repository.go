package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type RolePermissionRepository interface {
	// create
	InsertMany(ctx context.Context, rolePermissions []model.RolePermission, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context) (int, error)
	IsExist(ctx context.Context, userId string, permissionId string) (bool, error)
	FetchByRoleIds(ctx context.Context, roleIds []string) ([]model.RolePermission, error)

	// delete
	Delete(ctx context.Context, rolePermission *model.RolePermission, options ...data_type.RepositoryOption) error
	DeleteByPermissionId(ctx context.Context, permissionId string) error
	Truncate(ctx context.Context) error
}

type rolePermissionRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewRolePermissionRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) RolePermissionRepository {
	return &rolePermissionRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *rolePermissionRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.RolePermission, error) {
	rolePermissions := []model.RolePermission{}
	if err := fetch(r.db, ctx, &rolePermissions, stmt); err != nil {
		return nil, err
	}

	return rolePermissions, nil
}

func (r *rolePermissionRepository) InsertMany(ctx context.Context, rolePermissions []model.RolePermission, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range rolePermissions {
		arr = append(arr, &rolePermissions[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *rolePermissionRepository) Count(ctx context.Context) (int, error) {
	stmt := stmtBuilder.Select("COUNT(*) as count").
		From(model.RolePermissionTableName)

	return count(r.db, ctx, stmt)
}

func (r *rolePermissionRepository) IsExist(ctx context.Context, userId string, permissionId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("*").
			From(fmt.Sprintf("%s rp", model.RolePermissionTableName)).
			InnerJoin(fmt.Sprintf("%s ur ON rp.role_id = ur.role_id", model.UserRoleTableName)).
			Where("ur.user_id = ? AND permission_id = ?", userId, permissionId).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *rolePermissionRepository) FetchByRoleIds(ctx context.Context, roleIds []string) ([]model.RolePermission, error) {
	if len(roleIds) == 0 {
		return []model.RolePermission{}, nil
	}

	stmt := stmtBuilder.Select("*").
		From(model.RolePermissionTableName).
		Where(squirrel.Eq{"role_id": roleIds})

	return r.fetch(ctx, stmt)
}

func (r *rolePermissionRepository) Delete(ctx context.Context, rolePermission *model.RolePermission, options ...data_type.RepositoryOption) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, rolePermission, nil, options...)
}

// don't need audit log because only used for sync permission command
func (r *rolePermissionRepository) DeleteByPermissionId(ctx context.Context, permissionId string) error {
	whereStmt := squirrel.Eq{
		"permission_id": permissionId,
	}

	return destroy(r.db, ctx, model.RolePermissionTableName, whereStmt)
}

func (r *rolePermissionRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.RolePermissionTableName)
}
