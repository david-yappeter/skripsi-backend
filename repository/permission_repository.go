package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type PermissionRepository interface {
	// create
	InsertMany(ctx context.Context, permissions []model.Permission, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context) (int, error)
	Fetch(ctx context.Context) ([]model.Permission, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.Permission, error)
	FetchByRoleIds(ctx context.Context, roleIds []string) ([]model.Permission, error)
	FetchByTitles(ctx context.Context, titles []data_type.Permission) ([]model.Permission, error)
	GetByTitle(ctx context.Context, title data_type.Permission) (*model.Permission, error)
	IsExistByUserIdAndPermissionTypes(ctx context.Context, userId string, permissionTypeEnums []data_type.PermissionType) (bool, error)

	// delete
	Delete(ctx context.Context, permission *model.Permission, options ...data_type.RepositoryOption) error
	Truncate(ctx context.Context) error
}

type permissionRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewPermissionRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) PermissionRepository {
	return &permissionRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *permissionRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Permission, error) {
	permissions := []model.Permission{}
	if err := fetch(r.db, ctx, &permissions, stmt); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *permissionRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Permission, error) {
	permission := model.Permission{}
	if err := get(r.db, ctx, &permission, stmt); err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *permissionRepository) InsertMany(ctx context.Context, permissions []model.Permission, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range permissions {
		arr = append(arr, &permissions[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *permissionRepository) Count(ctx context.Context) (int, error) {
	stmt := stmtBuilder.Select("COUNT(*) as count").
		From(model.PermissionTableName)

	return count(r.db, ctx, stmt)
}

func (r *permissionRepository) Fetch(ctx context.Context) ([]model.Permission, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PermissionTableName)

	return r.fetch(ctx, stmt)
}

func (r *permissionRepository) FetchByIds(ctx context.Context, ids []string) ([]model.Permission, error) {
	if len(ids) == 0 {
		return []model.Permission{}, nil
	}

	stmt := stmtBuilder.Select("*").
		From(model.PermissionTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *permissionRepository) FetchByRoleIds(ctx context.Context, roleIds []string) ([]model.Permission, error) {
	if len(roleIds) == 0 {
		return []model.Permission{}, nil
	}

	stmt := stmtBuilder.Select("p.*").
		From(fmt.Sprintf("%s p", model.PermissionTableName)).
		JoinClause(
			squirrel.ConcatExpr(
				fmt.Sprintf("INNER JOIN %s rp ON p.id = rp.permission_id AND ", model.RolePermissionTableName),
				squirrel.Eq{"rp.role_id": roleIds},
			),
		)

	return r.fetch(ctx, stmt)
}

func (r *permissionRepository) FetchByTitles(ctx context.Context, titles []data_type.Permission) ([]model.Permission, error) {
	if len(titles) == 0 {
		return []model.Permission{}, nil
	}

	stmt := stmtBuilder.Select("*").
		From(model.PermissionTableName).
		Where(squirrel.Eq{"title": titles})

	return r.fetch(ctx, stmt)
}

func (r *permissionRepository) GetByTitle(ctx context.Context, title data_type.Permission) (*model.Permission, error) {
	stmt := stmtBuilder.Select("*").
		From(model.PermissionTableName).
		Where(squirrel.Eq{"title": title})

	return r.get(ctx, stmt)
}

func (r *permissionRepository) IsExistByUserIdAndPermissionTypes(ctx context.Context, userId string, permissionTypeEnums []data_type.PermissionType) (bool, error) {
	if len(permissionTypeEnums) == 0 {
		return false, nil
	}

	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(fmt.Sprintf("%s p", model.PermissionTableName)).
			InnerJoin(fmt.Sprintf("%s rp ON p.id = rp.permission_id", model.RolePermissionTableName)).
			InnerJoin(fmt.Sprintf("%s ur ON rp.role_id = ur.role_id", model.UserRoleTableName)).
			Where("ur.user_id = ?", userId).
			Where(squirrel.Eq{"p.type": permissionTypeEnums}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *permissionRepository) Delete(ctx context.Context, permission *model.Permission, options ...data_type.RepositoryOption) error {
	return defaultDestroy(r.db, r.loggerStack, ctx, permission, nil, options...)
}

func (r *permissionRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.PermissionTableName)
}
