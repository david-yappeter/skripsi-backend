package repository

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type UserRepository interface {
	// create
	Insert(ctx context.Context, user *model.User) error
	InsertMany(ctx context.Context, users []model.User, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.UserQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.UserQueryOption) ([]model.User, error)
	FetchByIds(ctx context.Context, ids []string) ([]model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByUsernameAndIsActive(ctx context.Context, username string, isActive bool) (*model.User, error)
	Get(ctx context.Context, id string) (*model.User, error)
	GetByIdAndIsActive(ctx context.Context, id string, isActive bool) (*model.User, error)
	IsExistByUsername(ctx context.Context, username string) (bool, error)

	// update
	Update(ctx context.Context, user *model.User) error

	// delete
	Truncate(ctx context.Context) error
}

type userRepository struct {
	db infrastructure.DBTX
}

func NewUserRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.User, error) {
	users := []model.User{}
	if err := fetch(r.db, ctx, &users, stmt); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.User, error) {
	user := model.User{}
	if err := get(r.db, ctx, &user, stmt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) prepareQuery(option model.UserQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s u", model.UserTableName))

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.Or{
			squirrel.ILike{"u.name": phrase},
			squirrel.ILike{"u.username": phrase},
		})
	}

	if option.IsActive != nil {
		stmt = stmt.Where(squirrel.Eq{"u.is_active": option.IsActive})
	}

	if len(option.RoleIds) > 0 {
		stmt = stmt.Where(
			stmtBuilder.Select("1").
				From(fmt.Sprintf("%s ur", model.UserRoleTableName)).
				Where("u.id = ur.user_id").
				Where(squirrel.Eq{"ur.role_id": option.RoleIds}).
				Prefix("EXISTS (").Suffix(")"),
		)
	}

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *userRepository) Insert(ctx context.Context, user *model.User) error {
	return defaultInsert(r.db, ctx, user, "*")
}

func (r *userRepository) InsertMany(ctx context.Context, users []model.User, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range users {
		arr = append(arr, &users[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *userRepository) Count(ctx context.Context, options ...model.UserQueryOption) (int, error) {
	option := model.UserQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *userRepository) Fetch(ctx context.Context, options ...model.UserQueryOption) ([]model.User, error) {
	option := model.UserQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *userRepository) FetchByIds(ctx context.Context, ids []string) ([]model.User, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserTableName).
		Where(squirrel.Eq{"id": ids})

	return r.fetch(ctx, stmt)
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserTableName).
		Where(squirrel.ILike{"username": username}).
		Limit(1)

	return r.get(ctx, stmt)
}

func (r *userRepository) GetByUsernameAndIsActive(ctx context.Context, username string, isActive bool) (*model.User, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserTableName).
		Where(squirrel.ILike{"username": username}).
		Where(squirrel.Eq{"is_active": isActive}).
		Limit(1)

	return r.get(ctx, stmt)
}

func (r *userRepository) Get(ctx context.Context, id string) (*model.User, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserTableName).
		Where("id = ?", id)

	return r.get(ctx, stmt)
}

func (r *userRepository) GetByIdAndIsActive(ctx context.Context, id string, isActive bool) (*model.User, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserTableName).
		Where(squirrel.ILike{"id": id}).
		Where(squirrel.Eq{"is_active": isActive}).
		Limit(1)

	return r.get(ctx, stmt)
}

func (r *userRepository) IsExistByUsername(ctx context.Context, username string) (bool, error) {
	user, err := r.GetByUsername(ctx, username)
	if err != nil && err != constant.ErrNoData {
		return false, err
	}

	isExist := user != nil

	return isExist, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return defaultUpdate(r.db, ctx, user, "*", nil)
}

func (r *userRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.UserTableName)
}
