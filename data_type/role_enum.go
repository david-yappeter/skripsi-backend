package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=Role -output=role_enum_gen.go -swagoutput=../tool/swag/enum_gen/role_enum_gen.go -custom
type Role int // @name RoleEnum

const (
	RoleSuperAdmin Role = iota + 1 // Super Admin
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
	}

	return []Permission{}
}

func GetRoleSuperAdminPermissions() []Permission {
	return []Permission{
		PermissionAdminUserCreate,
		PermissionAdminUserUpdate,
		PermissionAdminUserUpdatePassword,
		PermissionAdminUserUpdateActive,
		PermissionAdminUserUpdateInActive,
	}
}
