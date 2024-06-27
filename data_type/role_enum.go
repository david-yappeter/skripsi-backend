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
		// // balance
		// PermissionBalanceCreate,
		// PermissionBalanceFetch,
		// PermissionBalanceGet,
		// PermissionBalanceUpdate,
		// PermissionBalanceDelete,

		// cashier session
		PermissionCashierSessionFetch,
		PermissionCashierSessionGet,
		PermissionCashierSessionFetchTransaction,
		PermissionCashierSessionDownloadReport,
		// PermissionCashierSessionGetCurrent,

		// dashboard
		PermissionDashboardSummarizeDebt,
		PermissionDashboardSummarizeTransaction,

		// shop order
		PermissionShopOrderFetch,
		PermissionShopOrderGet,

		// ssr
		PermissionSsrWhatsappLogin,

		// // supplier
		// PermissionSupplierCreate,
		// PermissionSupplierFetch,
		// PermissionSupplierGet,
		// PermissionSupplierUpdate,
		// PermissionSupplierDelete,

		// // supplier type
		// PermissionSupplierTypeCreate,
		// PermissionSupplierTypeFetch,
		// PermissionSupplierTypeGet,
		// PermissionSupplierTypeUpdate,
		// PermissionSupplierTypeDelete,
		// PermissionSupplierTypeOptionForSupplierForm,

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
		PermissionUserOptionForDeliveryOrderDriverForm,

		// // unit
		// PermissionUnitCreate,
		// PermissionUnitFetch,
		// PermissionUnitGet,
		// PermissionUnitUpdate,
		// PermissionUnitDelete,
		// PermissionUnitOptionForProductUnitForm,
		// PermissionUnitOptionForProductUnitToUnitForm,

		// whatsapp
		PermissionWhatsappIsLoggedIn,
		PermissionWhatsappLogout,
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
		PermissionCustomerOptionForDeliveryOrderFilter,

		// customer type
		PermissionCustomerTypeCreate,
		PermissionCustomerTypeFetch,
		PermissionCustomerTypeGet,
		PermissionCustomerTypeUpdate,
		PermissionCustomerTypeDelete,
		PermissionCustomerTypeAddDiscount,
		PermissionCustomerTypeUpdateDiscount,
		PermissionCustomerTypeDeleteDiscount,
		PermissionCustomerTypeOptionForCustomerForm,
		PermissionCustomerTypeOptionForWhatsappProductPriceChangeBroadcastForm,

		// dashboard
		PermissionDashboardSummarizeDebt,
		PermissionDashboardSummarizeTransaction,

		// customer debt
		PermissionCustomerDebtUploadImage,
		PermissionCustomerDebtFetch,
		PermissionCustomerDebtGet,
		PermissionCustomerDebtPayment,

		// debt
		PermissionDebtUploadImage,
		PermissionDebtFetch,
		PermissionDebtGet,
		PermissionDebtPayment,

		// delivery order
		PermissionDeliveryOrderCreate,
		PermissionDeliveryOrderUpload,
		PermissionDeliveryOrderAddItem,
		PermissionDeliveryOrderAddImage,
		PermissionDeliveryOrderAddDriver,
		PermissionDeliveryOrderFetch,
		PermissionDeliveryOrderGet,
		PermissionDeliveryOrderMarkOngoing,
		PermissionDeliveryOrderUpdate,
		PermissionDeliveryOrderCancel,
		PermissionDeliveryOrderReturned,
		PermissionDeliveryOrderDelete,
		PermissionDeliveryOrderDeleteItem,
		PermissionDeliveryOrderDeleteImage,
		PermissionDeliveryOrderDeleteDriver,

		// delivery order review
		PermissionDeliveryOrderReviewFetch,
		PermissionDeliveryOrderReviewGet,

		// product
		PermissionProductCreate,
		PermissionProductUpload,
		PermissionProductFetch,
		PermissionProductGet,
		PermissionProductUpdate,
		PermissionProductDelete,
		PermissionProductOptionForProductReceiveItemForm,
		PermissionProductOptionForDeliveryOrderItemForm,
		PermissionProductOptionForCustomerTypeDiscountForm,
		PermissionProductOptionForCartAddItemForm,
		PermissionProductOptionForProductDiscountForm,

		// product discount
		PermissionProductDiscountCreate,
		PermissionProductDiscountFetch,
		PermissionProductDiscountGet,
		PermissionProductDiscountUpdate,
		PermissionProductDiscountDelete,

		// product receive
		PermissionProductReceiveUpload,
		PermissionProductReceiveAddImage,
		PermissionProductReceiveUpdate,
		PermissionProductReceiveCancel,
		PermissionProductReceiveMarkComplete,
		PermissionProductReceiveFetch,
		PermissionProductReceiveGet,
		PermissionProductReceiveUpdateItem,
		PermissionProductReceiveDelete,
		PermissionProductReceiveDeleteImage,

		// product return
		PermissionProductReturnCreate,
		PermissionProductReturnUpload,
		PermissionProductReturnAddItem,
		PermissionProductReturnAddImage,
		PermissionProductReturnUpdate,
		PermissionProductReturnCancel,
		PermissionProductReturnMarkComplete,
		PermissionProductReturnFetch,
		PermissionProductReturnGet,
		PermissionProductReturnDelete,
		PermissionProductReturnDeleteItem,
		PermissionProductReturnDeleteImage,

		// product stock adjustment
		PermissionProductStockAdjustmentFetch,

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
		PermissionProductUnitOptionForProductReceiveItemForm,
		PermissionProductUnitOptionForDeliveryOrderItemForm,
		PermissionProductOptionForProductReceiveItemForm,

		// purchase order
		PermissionPurchaseOrderCreate,
		PermissionPurchaseOrderUpload,
		PermissionPurchaseOrderAddItem,
		PermissionPurchaseOrderAddImage,
		PermissionPurchaseOrderUpdate,
		PermissionPurchaseOrderCancel,
		PermissionPurchaseOrderOngoing,
		PermissionPurchaseOrderMarkComplete,
		PermissionPurchaseOrderFetch,
		PermissionPurchaseOrderGet,
		PermissionPurchaseOrderDelete,
		PermissionPurchaseOrderDeleteItem,
		PermissionPurchaseOrderDeleteImage,

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
		PermissionSupplierOptionForProductReceiveFilter,

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

		// user
		PermissionUserOptionForProductStockAdjustmentFilter,
		PermissionUserOptionForCashierSessionFilter,
		PermissionUserOptionForDeliveryOrderDriverForm,

		// whatsapp
		PermissionWhatsappProductPriceChangeBroadcast,
		PermissionWhatsappCustomerTypeDiscountBroadcast,
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
		PermissionCashierSessionFetchTransaction,
		PermissionCashierSessionDownloadReport,
		PermissionCashierSessionGetCurrent,
		PermissionCashierSessionEnd,

		// product
		PermissionProductOptionForCartAddItemForm,

		// transaction
		PermissionTransactionCheckoutCart,
		PermissionTransactionGet,
	}
}

func GetRoleDriver() []Permission {
	return []Permission{
		// delivery order
		PermissionDeliveryOrderFetchDriver,
		PermissionDeliveryOrderGet,
		PermissionDeliveryOrderDelivering,
		PermissionDeliveryOrderCancel,
		PermissionDeliveryOrderMarkCompleted,
		PermissionDeliveryOrderActiveForDriver,
		PermissionDeliveryOrderDeliveryLocation,
	}
}
