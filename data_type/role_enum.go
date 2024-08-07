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
		// customer
		PermissionCustomerCreate,
		PermissionCustomerFetch,
		PermissionCustomerGet,
		PermissionCustomerUpdate,
		PermissionCustomerDelete,
		PermissionCustomerOptionForDeliveryOrderForm,
		PermissionCustomerOptionForDeliveryOrderFilter,
		PermissionCustomerOptionForCustomerDebtReportForm,
		PermissionCustomerOptionForWhatsappCustomerDebtBroadcastForm,

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

		// customer type discount
		PermissionCustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastForm,

		// cashier session
		PermissionCashierSessionFetch,
		PermissionCashierSessionGet,
		PermissionCashierSessionFetchTransaction,
		PermissionCashierSessionDownloadReport,

		// dashboard
		PermissionDashboardSummarizeDebt,
		PermissionDashboardSummarizeTransaction,

		// delivery order review
		PermissionDeliveryOrderReviewFetch,
		PermissionDeliveryOrderReviewGet,

		// shop order
		PermissionShopOrderFetch,
		PermissionShopOrderGet,

		// ssr
		PermissionSsrWhatsappLogin,

		// supplier
		PermissionSupplierCreate,
		PermissionSupplierFetch,
		PermissionSupplierGet,
		PermissionSupplierUpdate,
		PermissionSupplierDelete,
		PermissionSupplierOptionForProductReceiveForm,
		PermissionSupplierOptionForProductReceiveFilter,

		// supplier type
		PermissionSupplierTypeCreate,
		PermissionSupplierTypeFetch,
		PermissionSupplierTypeGet,
		PermissionSupplierTypeUpdate,
		PermissionSupplierTypeDelete,
		PermissionSupplierTypeOptionForSupplierForm,

		// product
		PermissionProductOptionForProductReceiveItemForm,
		PermissionProductOptionForDeliveryOrderItemForm,
		PermissionProductOptionForCustomerTypeDiscountForm,
		PermissionProductOptionForCartAddItemForm,
		PermissionProductOptionForProductDiscountForm,

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
		PermissionUserOptionForProductStockAdjustmentFilter,
		PermissionUserOptionForCashierSessionFilter,
		PermissionUserOptionForDeliveryOrderDriverForm,

		// whatsapp
		PermissionWhatsappIsLoggedIn,
		PermissionWhatsappLogout,
		PermissionWhatsappProductPriceChangeBroadcast,
		PermissionWhatsappCustomerDebtBroadcast,
		PermissionWhatsappCustomerTypeDiscountBroadcast,
		PermissionWhatsappCustomerTypeDiscountManyProductBroadcast,
	}
}

func GetRoleInventoryPermissions() []Permission {
	return []Permission{
		// customer debt
		PermissionCustomerDebtUploadImage,
		PermissionCustomerDebtDownloadReport,
		PermissionCustomerDebtFetch,
		PermissionCustomerDebtGet,
		PermissionCustomerDebtPayment,

		// customer
		PermissionCustomerOptionForDeliveryOrderForm,
		PermissionCustomerOptionForDeliveryOrderFilter,
		PermissionCustomerOptionForCustomerDebtReportForm,
		PermissionCustomerOptionForWhatsappCustomerDebtBroadcastForm,

		// customer type
		PermissionCustomerTypeOptionForWhatsappProductPriceChangeBroadcastForm,

		// dashboard
		PermissionDashboardSummarizeDebt,
		PermissionDashboardSummarizeTransaction,

		// debt
		PermissionDebtUploadImage,
		PermissionDebtFetch,
		PermissionDebtDownloadReport,
		PermissionDebtGet,
		PermissionDebtPayment,

		// delivery order
		PermissionDeliveryOrderCreate,
		PermissionDeliveryOrderDownloadReport,
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

		// supplier
		PermissionSupplierOptionForProductReceiveForm,
		PermissionSupplierOptionForProductReceiveFilter,

		// shop order
		PermissionShopOrderFetch,
		PermissionShopOrderGet,

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
		PermissionWhatsappCustomerDebtBroadcast,
		PermissionWhatsappCustomerTypeDiscountBroadcast,
		PermissionWhatsappCustomerTypeDiscountManyProductBroadcast,
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
		PermissionCashierSessionFetchForCurrentUser,
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
		PermissionTransactionReprint,
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
