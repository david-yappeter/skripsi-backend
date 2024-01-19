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

	BalanceRepository() BalanceRepository
	CustomerRepository() CustomerRepository
	FileRepository() FileRepository
	PermissionRepository() PermissionRepository
	ProductReceiveRepository() ProductReceiveRepository
	ProductReceiveItemRepository() ProductReceiveItemRepository
	ProductReceiveImageRepository() ProductReceiveImageRepository
	ProductRepository() ProductRepository
	ProductStockRepository() ProductStockRepository
	ProductUnitRepository() ProductUnitRepository
	RolePermissionRepository() RolePermissionRepository
	RoleRepository() RoleRepository
	SupplierRepository() SupplierRepository
	SupplierTypeRepository() SupplierTypeRepository
	UnitRepository() UnitRepository
	UserAccessTokenRepository() UserAccessTokenRepository
	UserRepository() UserRepository
	UserRoleRepository() UserRoleRepository
}

type repositoryManager struct {
	db          *sqlx.DB
	loggerStack infrastructure.LoggerStack

	balanceRepository             BalanceRepository
	customerRepository            CustomerRepository
	fileRepository                FileRepository
	permissionRepository          PermissionRepository
	productReceiveRepository      ProductReceiveRepository
	productReceiveItemRepository  ProductReceiveItemRepository
	productReceiveImageRepository ProductReceiveImageRepository
	productRepository             ProductRepository
	productStockRepository        ProductStockRepository
	productUnitRepository         ProductUnitRepository
	rolePermissionRepository      RolePermissionRepository
	roleRepository                RoleRepository
	supplierRepository            SupplierRepository
	supplierTypeRepository        SupplierTypeRepository
	unitRepository                UnitRepository
	userAccessTokenRepository     UserAccessTokenRepository
	userRepository                UserRepository
	userRoleRepository            UserRoleRepository
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

func (r *repositoryManager) BalanceRepository() BalanceRepository {
	return r.balanceRepository
}

func (r *repositoryManager) CustomerRepository() CustomerRepository {
	return r.customerRepository
}

func (r *repositoryManager) FileRepository() FileRepository {
	return r.fileRepository
}

func (r *repositoryManager) PermissionRepository() PermissionRepository {
	return r.permissionRepository
}

func (r *repositoryManager) ProductReceiveRepository() ProductReceiveRepository {
	return r.productReceiveRepository
}

func (r *repositoryManager) ProductReceiveItemRepository() ProductReceiveItemRepository {
	return r.productReceiveItemRepository
}

func (r *repositoryManager) ProductReceiveImageRepository() ProductReceiveImageRepository {
	return r.productReceiveImageRepository
}

func (r *repositoryManager) ProductRepository() ProductRepository {
	return r.productRepository
}

func (r *repositoryManager) ProductStockRepository() ProductStockRepository {
	return r.productStockRepository
}

func (r *repositoryManager) ProductUnitRepository() ProductUnitRepository {
	return r.productUnitRepository
}

func (r *repositoryManager) RolePermissionRepository() RolePermissionRepository {
	return r.rolePermissionRepository
}

func (r *repositoryManager) RoleRepository() RoleRepository {
	return r.roleRepository
}

func (r *repositoryManager) SupplierRepository() SupplierRepository {
	return r.supplierRepository
}

func (r *repositoryManager) SupplierTypeRepository() SupplierTypeRepository {
	return r.supplierTypeRepository
}

func (r *repositoryManager) UnitRepository() UnitRepository {
	return r.unitRepository
}

func (r *repositoryManager) UserAccessTokenRepository() UserAccessTokenRepository {
	return r.userAccessTokenRepository
}

func (r *repositoryManager) UserRepository() UserRepository {
	return r.userRepository
}
func (r *repositoryManager) UserRoleRepository() UserRoleRepository {
	return r.userRoleRepository
}

func NewRepositoryManager(infrastructureManager infrastructure.InfrastructureManager) RepositoryManager {
	db := infrastructureManager.GetDB()
	loggerStack := infrastructureManager.GetLoggerStack()

	return &repositoryManager{
		db:          db,
		loggerStack: loggerStack,

		balanceRepository: NewBalanceRepository(
			db,
			loggerStack,
		),
		customerRepository: NewCustomerRepository(
			db,
			loggerStack,
		),
		fileRepository: NewFileRepository(
			db,
			loggerStack,
		),
		permissionRepository: NewPermissionRepository(
			db,
			loggerStack,
		),
		productReceiveRepository: NewProductReceiveRepository(
			db,
			loggerStack,
		),
		productReceiveItemRepository: NewProductReceiveItemRepository(
			db,
			loggerStack,
		),
		productReceiveImageRepository: NewProductReceiveImageRepository(
			db,
			loggerStack,
		),
		productRepository: NewProductRepository(
			db,
			loggerStack,
		),
		productStockRepository: NewProductStockRepository(
			db,
			loggerStack,
		),
		productUnitRepository: NewProductUnitRepository(
			db,
			loggerStack,
		),
		rolePermissionRepository: NewRolePermissionRepository(
			db,
			loggerStack,
		),
		roleRepository: NewRoleRepository(
			db,
			loggerStack,
		),
		supplierRepository: NewSupplierRepository(
			db,
			loggerStack,
		),
		supplierTypeRepository: NewSupplierTypeRepository(
			db,
			loggerStack,
		),
		unitRepository: NewUnitRepository(
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
		userRoleRepository: NewUserRoleRepository(
			db,
			loggerStack,
		),
	}
}
