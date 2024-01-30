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

		// product
		PermissionProductCreate,
		PermissionProductFetch,
		PermissionProductGet,
		PermissionProductUpdate,
		PermissionProductDelete,

		// product unit
		PermissionProductUnitCreate,
		PermissionProductUnitUpload,
		PermissionProductUnitGet,
		PermissionProductUnitUpdate,
		PermissionProductUnitDelete,

		// role
		PermissionRoleOptionForUserForm,

		// user
		PermissionUserCreate,
		PermissionUserUpdate,
		PermissionUserUpdatePassword,
		PermissionUserUpdateActive,
		PermissionUserUpdateInActive,
		PermissionUserAddRole,
		PermissionUserDeleteRole,

		// unit
		PermissionUnitCreate,
		PermissionUnitFetch,
		PermissionUnitGet,
		PermissionUnitUpdate,
		PermissionUnitDelete,
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

		// product receive
		PermissionProductReceiveCreate,
		PermissionProductReceiveUpload,
		PermissionProductReceiveAddItem,
		PermissionProductReceiveAddImage,
		PermissionProductReceiveFetch,
		PermissionProductReceiveGet,
		PermissionProductReceiveDelete,
		PermissionProductReceiveDeleteItem,
		PermissionProductReceiveDeleteImage,

		// product
		PermissionProductOptionForDeliveryOrderForm,
		PermissionProductOptionForProductReceiveForm,

		// product_unit
		PermissionProductUnitOptionForDeliveryOrderForm,
		PermissionProductUnitOptionForProductReceiveForm,

		// supplier type
		PermissionSupplierTypeCreate,
		PermissionSupplierTypeFetch,
		PermissionSupplierTypeGet,
		PermissionSupplierTypeUpdate,
		PermissionSupplierTypeDelete,

		// supplier
		PermissionSupplierCreate,
		PermissionSupplierFetch,
		PermissionSupplierGet,
		PermissionSupplierUpdate,
		PermissionSupplierDelete,
		PermissionSupplierOptionForProductReceiveForm,
	}
}

func GetRoleCashier() []Permission {
	return []Permission{}
}

func GetRoleDriver() []Permission {
	return []Permission{}
}
