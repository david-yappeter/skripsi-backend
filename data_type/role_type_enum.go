package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=RoleType -output=role_type_enum_gen.go -swagoutput=../tool/swag/enum_gen/role_type_enum_gen.go -custom
type RoleType int // @name RoleTypeEnum

const (
	RoleTypeSuperAdmin RoleType = iota + 1 // SUPER_ADMIN
	RoleTypeDoctor                         // DOCTOR
	RoleTypeNurse                          // NURSE
	RoleTypePharmacist                     // PHARMACIST
	RoleTypeGlobal                         // GLOBAL
	RoleTypeAdmin                          // ADMIN
)

var rolesByRoleType = map[RoleType][]Role{}

func (rt RoleType) Roles() []Role {
	return rolesByRoleType[rt]
}

func (rt RoleType) IsExclusiveRoleType() bool {
	return rt == RoleTypeDoctor ||
		rt == RoleTypeNurse ||
		rt == RoleTypePharmacist
}

func init() {
	for _, role := range ListRole() {
		if _, ok := rolesByRoleType[role.RoleType()]; !ok {
			rolesByRoleType[role.RoleType()] = []Role{}
		}

		rolesByRoleType[role.RoleType()] = append(rolesByRoleType[role.RoleType()], role)
	}
}
