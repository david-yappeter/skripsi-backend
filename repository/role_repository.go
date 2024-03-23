package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type RoleRepository interface {
	// create
	InsertMany(ctx context.Context, roles []model.Role) error

	// read
	Count(ctx context.Context, options ...model.RoleQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.RoleQueryOption) ([]model.Role, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.Role, error)
	FetchByNames(ctx context.Context, roleEnums []data_type.Role) ([]model.Role, error)
	FetchByUserId(ctx context.Context, id string) ([]model.Role, error)
	FetchIdsByNames(ctx context.Context, roleEnums []data_type.Role) ([]string, error)
	FetchIdsByUserId(ctx context.Context, id string) ([]string, error)
	GetById(ctx context.Context, id string) (*model.Role, error)
	GetByName(ctx context.Context, name data_type.Role) (*model.Role, error)
	IsIdsExist(ctx context.Context, ids []string) (bool, error)

	// delete
	Truncate(ctx context.Context) error
}

type roleRepository struct {
	db          infrastructure.DBTX
	loggerStack infrastructure.LoggerStack
}

func NewRoleRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) RoleRepository {
	return &roleRepository{
		db:          db,
		loggerStack: loggerStack,
	}
}

func (r *roleRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.Role, error) {
	roles := []model.Role{}
	if err := fetch(r.db, ctx, &roles, stmt); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.Role, error) {
	role := model.Role{}
	if err := get(r.db, ctx, &role, stmt); err != nil {
		return nil, err
	}

	return &role, nil
}

func (r roleRepository) prepareQuery(option model.RoleQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().From(model.RoleTableName)

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.ILike{"name": phrase})
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *roleRepository) InsertMany(ctx context.Context, roles []model.Role) error {
	arr := []model.BaseModel{}
	for i := range roles {
		arr = append(arr, &roles[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *roleRepository) Count(ctx context.Context, options ...model.RoleQueryOption) (int, error) {
	option := model.RoleQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *roleRepository) Fetch(ctx context.Context, options ...model.RoleQueryOption) ([]model.Role, error) {
	option := model.RoleQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *roleRepository) FetchByIds(ctx context.Context, ids []string) ([]model.Role, error) {
	if len(ids) == 0 {
		return []model.Role{}, nil
	}

	stmt := stmtBuilder.Select("*").
		From(model.RoleTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *roleRepository) FetchByNames(ctx context.Context, roleEnums []data_type.Role) ([]model.Role, error) {
	if len(roleEnums) == 0 {
		return []model.Role{}, nil
	}

	stmt := stmtBuilder.Select("*").
		From(model.RoleTableName).
		Where(squirrel.Eq{"name": roleEnums})

	return r.fetch(ctx, stmt)
}

func (r *roleRepository) FetchByUserId(ctx context.Context, id string) ([]model.Role, error) {
	stmt := stmtBuilder.Select("r.*").
		From(fmt.Sprintf("%s r ", model.RoleTableName)).
		InnerJoin(fmt.Sprintf("%s ur ON r.id = ur.role_id", model.UserRoleTableName)).
		Where("ur.user_id = ?", id)

	return r.fetch(ctx, stmt)
}

func (r *roleRepository) FetchIdsByNames(ctx context.Context, roleEnums []data_type.Role) ([]string, error) {
	roles, err := r.FetchByNames(ctx, roleEnums)
	if err != nil {
		return nil, err
	}

	roleIds := []string{}
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}

	return roleIds, nil
}

func (r *roleRepository) FetchIdsByUserId(ctx context.Context, id string) ([]string, error) {
	roles, err := r.FetchByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	roleIds := []string{}
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}

	return roleIds, nil
}

func (r *roleRepository) GetById(ctx context.Context, id string) (*model.Role, error) {
	stmt := stmtBuilder.Select("*").
		From(model.RoleTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *roleRepository) GetByName(ctx context.Context, name data_type.Role) (*model.Role, error) {
	stmt := stmtBuilder.Select("*").
		From(model.RoleTableName).
		Where(squirrel.Eq{"name": name})

	return r.get(ctx, stmt)
}

func (r *roleRepository) IsIdsExist(ctx context.Context, ids []string) (bool, error) {
	roles, err := r.FetchByIds(ctx, ids)
	if err != nil {
		return false, err
	}

	checker := map[string]struct{}{}
	for _, id := range ids {
		checker[id] = struct{}{}
	}
	for _, v := range roles {
		delete(checker, v.Id)
	}
	isExist := len(checker) == 0

	return isExist, nil
}

func (r *roleRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.RoleTableName)
}
