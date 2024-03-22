package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=Role -output=role_enum_gen.go -swagoutput=../tool/swag/enum_gen/role_enum_gen.go -custom
type Role int // @name RoleEnum

const (
	RoleSuperAdmin Role = iota + 1 // Super Admin
	RoleInventory                  // Inventory
	RoleCashier                    // Cashier
	RoleDriver                     // Driver
)

var roleTypeByRole = map[Role]RoleType{
	RoleSuperAdmin: RoleTypeSuperAdmin,
	RoleInventory:  RoleTypeGlobal,
	RoleCashier:    RoleTypeGlobal,
	RoleDriver:     RoleTypeGlobal,
}

func (r Role) RoleType() RoleType {
	return roleTypeByRole[r]
}

func (r Role) Permissions() []Permission {
	switch r {
	case RoleSuperAdmin:
		return GetRoleSuperPermissions()
	case RoleInventory:
		return GetRoleInventoryPermissions()
	case RoleCashier:
		return GetRoleCashier()
	case RoleDriver:
		return GetRoleDriver()
	}

	return []Permission{}
}

func GetRoleSuperPermissions() []Permission {
	return []Permission{
		// balance
		PermissionBalanceCreate,
		PermissionBalanceFetch,
		PermissionBalanceGet,
		PermissionBalanceUpdate,
		PermissionBalanceDelete,

		// cashier session
		PermissionCashierSessionFetch,

		// shop order
		PermissionShopOrderFetch,
		PermissionShopOrderGet,

		// supplier
		PermissionSupplierCreate,
		PermissionSupplierFetch,
		PermissionSupplierGet,
		PermissionSupplierUpdate,
		PermissionSupplierDelete,

		// supplier type
		PermissionSupplierTypeCreate,
		PermissionSupplierTypeFetch,
		PermissionSupplierTypeGet,
		PermissionSupplierTypeUpdate,
		PermissionSupplierTypeDelete,
		PermissionSupplierTypeOptionForSupplierForm,

		// role
		PermissionRoleOptionForUserForm,

		// user
		PermissionUserCreate,
		PermissionUserFetch,
		PermissionUserGet,
		PermissionUserUpdate,
		PermissionUserUpdatePassword,
		PermissionUserUpdateActive,
		PermissionUserUpdateInActive,
		PermissionUserAddRole,
		PermissionUserDeleteRole,
		PermissionUserOptionForCashierSessionFilter,

		// unit
		PermissionUnitCreate,
		PermissionUnitFetch,
		PermissionUnitGet,
		PermissionUnitUpdate,
		PermissionUnitDelete,
		PermissionUnitOptionForProductUnitForm,
		PermissionUnitOptionForProductUnitToUnitForm,
	}
}

func GetRoleInventoryPermissions() []Permission {
	return []Permission{
		// customer
		PermissionCustomerCreate,
		PermissionCustomerFetch,
		PermissionCustomerGet,
		PermissionCustomerUpdate,
		PermissionCustomerDelete,
		PermissionCustomerOptionForDeliveryOrderForm,

		// customer type
		PermissionCustomerTypeCreate,
		PermissionCustomerTypeFetch,
		PermissionCustomerTypeGet,
		PermissionCustomerTypeUpdate,
		PermissionCustomerTypeDelete,
		PermissionCustomerTypeOptionForCustomerForm,
		PermissionCustomerTypeAddDiscount,
		PermissionCustomerTypeUpdateDiscount,
		PermissionCustomerTypeDeleteDiscount,

		// customer debt
		PermissionCustomerDebtUploadImage,
		PermissionCustomerDebtFetch,
		PermissionCustomerDebtGet,
		PermissionCustomerDebtPayment,

		// delivery order
		PermissionDeliveryOrderCreate,
		PermissionDeliveryOrderUpload,
		PermissionDeliveryOrderAddItem,
		PermissionDeliveryOrderAddImage,
		PermissionDeliveryOrderAddDriver,
		PermissionDeliveryOrderFetch,
		PermissionDeliveryOrderGet,
		PermissionDeliveryOrderMarkOngoing,
		PermissionDeliveryOrderCancel,
		PermissionDeliveryOrderMarkCompleted,
		PermissionDeliveryOrderDelete,
		PermissionDeliveryOrderDeleteItem,
		PermissionDeliveryOrderDeleteImage,
		PermissionDeliveryOrderDeleteDriver,

		// product
		PermissionProductCreate,
		PermissionProductUpload,
		PermissionProductFetch,
		PermissionProductGet,
		PermissionProductUpdate,
		PermissionProductDelete,
		PermissionProductOptionForProductReceiveForm,
		PermissionProductOptionForDeliveryOrderForm,
		PermissionProductOptionForCustomerTypeDiscountForm,
		PermissionProductOptionForCartAddItemForm,

		// product discount
		PermissionProductDiscountCreate,
		PermissionProductDiscountUpdate,
		PermissionProductDiscountDelete,

		// product receive
		PermissionProductReceiveCreate,
		PermissionProductReceiveUpload,
		PermissionProductReceiveAddItem,
		PermissionProductReceiveAddImage,
		PermissionProductReceiveCancel,
		PermissionProductReceiveMarkComplete,
		PermissionProductReceiveFetch,
		PermissionProductReceiveGet,
		PermissionProductReceiveDelete,
		PermissionProductReceiveDeleteItem,
		PermissionProductReceiveDeleteImage,

		// product stock
		PermissionProductStockFetch,
		PermissionProductStockGet,
		PermissionProductStockDownloadReport,
		PermissionProductStockAdjustment,

		// product_unit
		PermissionProductUnitCreate,
		PermissionProductUnitGet,
		PermissionProductUnitUpdate,
		PermissionProductUnitDelete,
		PermissionProductUnitOptionForDeliveryOrderForm,
		PermissionProductUnitOptionForProductReceiveForm,

		// supplier type
		PermissionSupplierTypeCreate,
		PermissionSupplierTypeFetch,
		PermissionSupplierTypeGet,
		PermissionSupplierTypeUpdate,
		PermissionSupplierTypeDelete,
		PermissionSupplierTypeOptionForSupplierForm,

		// shop order
		PermissionShopOrderFetch,
		PermissionShopOrderGet,

		// supplier
		PermissionSupplierCreate,
		PermissionSupplierFetch,
		PermissionSupplierGet,
		PermissionSupplierUpdate,
		PermissionSupplierDelete,
		PermissionSupplierOptionForProductReceiveForm,

		// tiktok product
		PermissionTiktokProductCreate,
		PermissionTiktokProductUploadImage,
		PermissionTiktokProductFetchBrands,
		PermissionTiktokProductFetchCategories,
		PermissionTiktokProductGetCategoryRules,
		PermissionTiktokProductGetCategoryAttributes,
		PermissionTiktokProductGet,
		PermissionTiktokProductUpdate,
		PermissionTiktokProductRecommendedCategory,
		PermissionTiktokProductActivate,
		PermissionTiktokProductDeactivate,

		// unit
		PermissionUnitCreate,
		PermissionUnitFetch,
		PermissionUnitGet,
		PermissionUnitUpdate,
		PermissionUnitDelete,
		PermissionUnitOptionForProductUnitForm,
		PermissionUnitOptionForProductUnitToUnitForm,
	}
}

func GetRoleCashier() []Permission {
	return []Permission{
		// cart
		PermissionCartGetActive,
		PermissionCartAddItem,
		PermissionCartUpdateItem,
		PermissionCartDeleteItem,
		PermissionCartSetActive,
		PermissionCartSetInActive,
		PermissionCartDelete,

		// cashier session
		PermissionCashierSessionFetch,
		PermissionCashierSessionStart,
		PermissionCashierSessionGet,
		PermissionCashierSessionDownloadReport,
		PermissionCashierSessionGetCurrent,
		PermissionCashierSessionEnd,

		// transaction
		PermissionTransactionCheckoutCart,
	}
}

func GetRoleDriver() []Permission {
	return []Permission{}
}
