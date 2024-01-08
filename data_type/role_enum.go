package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=Role -output=role_enum_gen.go -swagoutput=../tool/swag/enum_gen/role_enum_gen.go -custom
type Role int // @name RoleEnum

const (
	RoleSuperAdmin Role = iota + 1 // Super Admin
	RoleInventory                  // Inventory
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
	}

	return []Permission{}
}

func GetRoleSuperAdminPermissions() []Permission {
	return []Permission{
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
	}
}

func GetRoleInventoryPermissions() []Permission {
	return []Permission{
		// supplier type
		PermissionSupplierTypeCreate,
		PermissionSupplierTypeFetch,
		PermissionSupplierTypeGet,
		PermissionSupplierTypeUpdate,
		PermissionSupplierTypeDelete,
	}
}
