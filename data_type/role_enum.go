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
}

func (r Role) RoleType() RoleType {
	return roleTypeByRole[r]
}

func (r Role) Permissions() []Permission {
	switch r {
	case RoleSuperAdmin:
		return GetRoleSuperAdminPermissions()
	case RoleInventory:
		return GetRoleInventoryPermissions()
	case RoleCashier:
		return GetRoleCashier()
	case RoleDriver:
		return GetRoleDriver()
	}

	return []Permission{}
}

func GetRoleSuperAdminPermissions() []Permission {
	return []Permission{
		// admin balance
		PermissionAdminBalanceCreate,
		PermissionAdminBalanceFetch,
		PermissionAdminBalanceGet,
		PermissionAdminBalanceUpdate,
		PermissionAdminBalanceDelete,

		// admin user
		PermissionAdminUserCreate,
		PermissionAdminUserUpdate,
		PermissionAdminUserUpdatePassword,
		PermissionAdminUserUpdateActive,
		PermissionAdminUserUpdateInActive,

		// admin unit
		PermissionAdminUnitCreate,
		PermissionAdminUnitFetch,
		PermissionAdminUnitGet,
		PermissionAdminUnitUpdate,
		PermissionAdminUnitDelete,

		// admin supplier
		PermissionAdminSupplierCreate,
		PermissionAdminSupplierFetch,
		PermissionAdminSupplierGet,
		PermissionAdminSupplierUpdate,
		PermissionAdminSupplierDelete,

		// admin supplier type
		PermissionAdminSupplierTypeCreate,
		PermissionAdminSupplierTypeFetch,
		PermissionAdminSupplierTypeGet,
		PermissionAdminSupplierTypeUpdate,
		PermissionAdminSupplierTypeDelete,

		// admin product unit
		PermissionAdminProductUnitCreate,
		PermissionAdminProductUnitUpload,
		PermissionAdminProductUnitGet,
		PermissionAdminProductUnitUpdate,
		PermissionAdminProductUnitDelete,
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
	}
}

func GetRoleCashier() []Permission {
	return []Permission{}
}

func GetRoleDriver() []Permission {
	return []Permission{}
}
