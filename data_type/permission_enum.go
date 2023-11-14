package data_type

import (
	"fmt"
)

//go:generate go run myapp/tool/stringer -linecomment -type=Permission -output=permission_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_enum_gen.go -custom
type Permission int // @name PermissionEnum

const (
	// admin user
	PermissionAdminUserCreate         Permission = iota + 1 // ADMIN_USER_CREATE
	PermissionAdminUserUpdate                               // ADMIN_USER_UPDATE
	PermissionAdminUserUpdatePassword                       // ADMIN_USER_UPDATE_PASSWORD
	PermissionAdminUserUpdateActive                         // ADMIN_USER_UPDATE_ACTIVE
	PermissionAdminUserUpdateInActive                       // ADMIN_USER_UPDATE_INACTIVE
)

var permissionTypeByPermission = map[Permission]PermissionType{}

// Don't use this for use case
func (p Permission) PermissionType() PermissionType {
	permissionType, exist := permissionTypeByPermission[p]
	if !exist {
		panic(fmt.Errorf("Permission %s is not registered in permissionTypeByPermission", p))
	}

	return permissionType
}
