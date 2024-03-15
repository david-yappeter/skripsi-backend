package repository

import (
	"context"
	"database/sql"
	"fmt"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/jmoiron/sqlx"
)

type RepositoryManager interface {
	Transaction(
		ctx context.Context,
		fn func(ctx context.Context) error,
	) error

	BalanceRepository() BalanceRepository
	CartRepository() CartRepository
	CartItemRepository() CartItemRepository
	CashierSessionRepository() CashierSessionRepository
	CustomerDebtRepository() CustomerDebtRepository
	CustomerPaymentRepository() CustomerPaymentRepository
	CustomerRepository() CustomerRepository
	CustomerTypeDiscountRepository() CustomerTypeDiscountRepository
	CustomerTypeRepository() CustomerTypeRepository
	DebtPaymentRepository() DebtPaymentRepository
	DebtRepository() DebtRepository
	DeliveryOrderRepository() DeliveryOrderRepository
	DeliveryOrderDriverRepository() DeliveryOrderDriverRepository
	DeliveryOrderImageRepository() DeliveryOrderImageRepository
	DeliveryOrderItemCostRepository() DeliveryOrderItemCostRepository
	DeliveryOrderItemRepository() DeliveryOrderItemRepository
	FileRepository() FileRepository
	PermissionRepository() PermissionRepository
	ProductDiscountRepository() ProductDiscountRepository
	ProductReceiveRepository() ProductReceiveRepository
	ProductReceiveItemRepository() ProductReceiveItemRepository
	ProductReceiveImageRepository() ProductReceiveImageRepository
	ProductRepository() ProductRepository
	ProductStockMutationRepository() ProductStockMutationRepository
	ProductStockRepository() ProductStockRepository
	ProductUnitRepository() ProductUnitRepository
	RolePermissionRepository() RolePermissionRepository
	RoleRepository() RoleRepository
	ShopOrderRepository() ShopOrderRepository
	ShopOrderItemRepository() ShopOrderItemRepository
	SupplierRepository() SupplierRepository
	SupplierTypeRepository() SupplierTypeRepository
	TiktokConfigRepository() TiktokConfigRepository
	TiktokProductRepository() TiktokProductRepository
	TransactionItemCostRepository() TransactionItemCostRepository
	TransactionItemRepository() TransactionItemRepository
	TransactionPaymentRepository() TransactionPaymentRepository
	TransactionRepository() TransactionRepository
	UnitRepository() UnitRepository
	UserAccessTokenRepository() UserAccessTokenRepository
	UserRepository() UserRepository
	UserRoleRepository() UserRoleRepository
}

type repositoryManager struct {
	db          *sqlx.DB
	loggerStack infrastructure.LoggerStack

	balanceRepository               BalanceRepository
	cartRepository                  CartRepository
	cartItemRepository              CartItemRepository
	cashierSessionRepository        CashierSessionRepository
	customerDebtRepository          CustomerDebtRepository
	customerPaymentRepository       CustomerPaymentRepository
	customerRepository              CustomerRepository
	customerTypeDiscountRepository  CustomerTypeDiscountRepository
	customerTypeRepository          CustomerTypeRepository
	debtPaymentRepository           DebtPaymentRepository
	debtRepository                  DebtRepository
	deliveryOrderRepository         DeliveryOrderRepository
	deliveryOrderDriverRepository   DeliveryOrderDriverRepository
	deliveryOrderImageRepository    DeliveryOrderImageRepository
	deliveryOrderItemCostRepository DeliveryOrderItemCostRepository
	deliveryOrderItemRepository     DeliveryOrderItemRepository
	fileRepository                  FileRepository
	permissionRepository            PermissionRepository
	productDiscountRepository       ProductDiscountRepository
	productReceiveRepository        ProductReceiveRepository
	productReceiveItemRepository    ProductReceiveItemRepository
	productReceiveImageRepository   ProductReceiveImageRepository
	productRepository               ProductRepository
	productStockMutationRepository  ProductStockMutationRepository
	productStockRepository          ProductStockRepository
	productUnitRepository           ProductUnitRepository
	rolePermissionRepository        RolePermissionRepository
	roleRepository                  RoleRepository
	shopOrderRepository             ShopOrderRepository
	shopOrderItemRepository         ShopOrderItemRepository
	supplierRepository              SupplierRepository
	supplierTypeRepository          SupplierTypeRepository
	tiktokConfigRepository          TiktokConfigRepository
	tiktokProductRepository         TiktokProductRepository
	transactionItemCostRepository   TransactionItemCostRepository
	transactionItemRepository       TransactionItemRepository
	transactionPaymentRepository    TransactionPaymentRepository
	transactionRepository           TransactionRepository
	unitRepository                  UnitRepository
	userAccessTokenRepository       UserAccessTokenRepository
	userRepository                  UserRepository
	userRoleRepository              UserRoleRepository
}

func (r *repositoryManager) Transaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) (err error) {
	var tx *sqlx.Tx

	defer func() {
		if err != nil && tx != nil {
			if rbErr := tx.Rollback(); rbErr != nil && rbErr != sql.ErrTxDone {
				err = fmt.Errorf("%v\nrollback err: %v", err, rbErr)
			}
		}
	}()

	tx, err = r.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	ctx, err = model.SetDbtxCtx(ctx, tx)
	if err != nil {
		return
	}

	if err = fn(ctx); err != nil {
		return
	}

	return tx.Commit()
}

func (r *repositoryManager) BalanceRepository() BalanceRepository {
	return r.balanceRepository
}

func (r *repositoryManager) CartRepository() CartRepository {
	return r.cartRepository
}

func (r *repositoryManager) CartItemRepository() CartItemRepository {
	return r.cartItemRepository
}

func (r *repositoryManager) CashierSessionRepository() CashierSessionRepository {
	return r.cashierSessionRepository
}

func (r *repositoryManager) CustomerDebtRepository() CustomerDebtRepository {
	return r.customerDebtRepository
}

func (r *repositoryManager) CustomerPaymentRepository() CustomerPaymentRepository {
	return r.customerPaymentRepository
}

func (r *repositoryManager) CustomerRepository() CustomerRepository {
	return r.customerRepository
}

func (r *repositoryManager) CustomerTypeDiscountRepository() CustomerTypeDiscountRepository {
	return r.customerTypeDiscountRepository
}

func (r *repositoryManager) CustomerTypeRepository() CustomerTypeRepository {
	return r.customerTypeRepository
}

func (r *repositoryManager) DebtPaymentRepository() DebtPaymentRepository {
	return r.debtPaymentRepository
}

func (r *repositoryManager) DebtRepository() DebtRepository {
	return r.debtRepository
}

func (r *repositoryManager) DeliveryOrderRepository() DeliveryOrderRepository {
	return r.deliveryOrderRepository
}

func (r *repositoryManager) DeliveryOrderDriverRepository() DeliveryOrderDriverRepository {
	return r.deliveryOrderDriverRepository
}

func (r *repositoryManager) DeliveryOrderImageRepository() DeliveryOrderImageRepository {
	return r.deliveryOrderImageRepository
}

func (r *repositoryManager) DeliveryOrderItemCostRepository() DeliveryOrderItemCostRepository {
	return r.deliveryOrderItemCostRepository
}

func (r *repositoryManager) DeliveryOrderItemRepository() DeliveryOrderItemRepository {
	return r.deliveryOrderItemRepository
}

func (r *repositoryManager) FileRepository() FileRepository {
	return r.fileRepository
}

func (r *repositoryManager) PermissionRepository() PermissionRepository {
	return r.permissionRepository
}

func (r *repositoryManager) ProductDiscountRepository() ProductDiscountRepository {
	return r.productDiscountRepository
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

func (r *repositoryManager) ProductStockMutationRepository() ProductStockMutationRepository {
	return r.productStockMutationRepository
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

func (r *repositoryManager) ShopOrderRepository() ShopOrderRepository {
	return r.shopOrderRepository
}

func (r *repositoryManager) ShopOrderItemRepository() ShopOrderItemRepository {
	return r.shopOrderItemRepository
}

func (r *repositoryManager) SupplierRepository() SupplierRepository {
	return r.supplierRepository
}

func (r *repositoryManager) SupplierTypeRepository() SupplierTypeRepository {
	return r.supplierTypeRepository
}

func (r *repositoryManager) TiktokConfigRepository() TiktokConfigRepository {
	return r.tiktokConfigRepository
}

func (r *repositoryManager) TiktokProductRepository() TiktokProductRepository {
	return r.tiktokProductRepository
}

func (r *repositoryManager) TransactionItemCostRepository() TransactionItemCostRepository {
	return r.transactionItemCostRepository
}

func (r *repositoryManager) TransactionItemRepository() TransactionItemRepository {
	return r.transactionItemRepository
}

func (r *repositoryManager) TransactionPaymentRepository() TransactionPaymentRepository {
	return r.transactionPaymentRepository
}

func (r *repositoryManager) TransactionRepository() TransactionRepository {
	return r.transactionRepository
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
		cartRepository: NewCartRepository(
			db,
			loggerStack,
		),
		cartItemRepository: NewCartItemRepository(
			db,
			loggerStack,
		),
		cashierSessionRepository: NewCashierSessionRepository(
			db,
			loggerStack,
		),
		customerDebtRepository: NewCustomerDebtRepository(
			db,
			loggerStack,
		),
		customerPaymentRepository: NewCustomerPaymentRepository(
			db,
			loggerStack,
		),
		customerRepository: NewCustomerRepository(
			db,
			loggerStack,
		),
		customerTypeDiscountRepository: NewCustomerTypeDiscountRepository(
			db,
			loggerStack,
		),
		customerTypeRepository: NewCustomerTypeRepository(
			db,
			loggerStack,
		),
		debtPaymentRepository: NewDebtPaymentRepository(
			db,
			loggerStack,
		),
		debtRepository: NewDebtRepository(
			db,
			loggerStack,
		),
		deliveryOrderRepository: NewDeliveryOrderRepository(
			db,
			loggerStack,
		),
		deliveryOrderDriverRepository: NewDeliveryOrderDriverRepository(
			db,
			loggerStack,
		),
		deliveryOrderImageRepository: NewDeliveryOrderImageRepository(
			db,
			loggerStack,
		),
		deliveryOrderItemCostRepository: NewDeliveryOrderItemCostRepository(
			db,
			loggerStack,
		),
		deliveryOrderItemRepository: NewDeliveryOrderItemRepository(
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
		productDiscountRepository: NewProductDiscountRepository(
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
		productStockMutationRepository: NewProductStockMutationRepository(
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
		shopOrderRepository: NewShopOrderRepository(
			db,
			loggerStack,
		),
		shopOrderItemRepository: NewShopOrderItemRepository(
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
		tiktokConfigRepository: NewTiktokConfigRepository(
			db,
			loggerStack,
		),
		tiktokProductRepository: NewTiktokProductRepository(
			db,
			loggerStack,
		),
		transactionItemCostRepository: NewTransactionItemCostRepository(
			db,
			loggerStack,
		),
		transactionItemRepository: NewTransactionItemRepository(
			db,
			loggerStack,
		),
		transactionPaymentRepository: NewTransactionPaymentRepository(
			db,
			loggerStack,
		),
		transactionRepository: NewTransactionRepository(
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
