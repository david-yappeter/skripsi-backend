package repository

import (
	"context"
	"fmt"
	"myapp/infrastructure"

	"github.com/jmoiron/sqlx"
)

type RepositoryManager interface {
	Transaction(
		ctx context.Context,
		fn func(tx infrastructure.DBTX, loggerStack infrastructure.LoggerStack) error,
	) error

	PermissionRepository() PermissionRepository
	RolePermissionRepository() RolePermissionRepository
	RoleRepository() RoleRepository
	UserAccessTokenRepository() UserAccessTokenRepository
	UserRepository() UserRepository
}

type repositoryManager struct {
	db          *sqlx.DB
	loggerStack infrastructure.LoggerStack

	permissionRepository      PermissionRepository
	rolePermissionRepository  RolePermissionRepository
	roleRepository            RoleRepository
	userAccessTokenRepository UserAccessTokenRepository
	userRepository            UserRepository
}

func (r *repositoryManager) Transaction(
	ctx context.Context,
	fn func(tx infrastructure.DBTX, loggerStack infrastructure.LoggerStack) error,
) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return translateSqlError(err)
	}

	if err := fn(tx, r.loggerStack); err != nil {
		err = translateSqlError(err)
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf(
				"transaction error: %v"+"\n"+
					"rollback err: %v",
				err,
				rbErr,
			)
		}

		return err
	}

	return translateSqlError(tx.Commit())
}

func (r *repositoryManager) PermissionRepository() PermissionRepository {
	return r.permissionRepository
}

func (r *repositoryManager) RolePermissionRepository() RolePermissionRepository {
	return r.rolePermissionRepository
}
func (r *repositoryManager) RoleRepository() RoleRepository {
	return r.roleRepository
}

func (r *repositoryManager) UserAccessTokenRepository() UserAccessTokenRepository {
	return r.userAccessTokenRepository
}

func (r *repositoryManager) UserRepository() UserRepository {
	return r.userRepository
}

func NewRepositoryManager(infrastructureManager infrastructure.InfrastructureManager) RepositoryManager {
	db := infrastructureManager.GetDB()
	loggerStack := infrastructureManager.GetLoggerStack()

	return &repositoryManager{
		db:          db,
		loggerStack: loggerStack,

		roleRepository: NewRoleRepository(
			db,
			loggerStack,
		),
		userAccessTokenRepository: NewUserAccessTokenRepository(
			db,
			loggerStack,
		),
		userRepository: NewUserRepository(
			db,
			loggerStack,
		),
	}
}
