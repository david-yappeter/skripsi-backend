package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=Permission -output=permission_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_enum_gen.go -custom
type Permission int // @name PermissionEnum

const (
	// admin user
	PermissionAdminUserCreate         Permission = iota + 1 // ADMIN_USER_CREATE
	PermissionAdminUserUpdate                               // ADMIN_USER_UPDATE
	PermissionAdminUserUpdatePassword                       // ADMIN_USER_UPDATE_PASSWORD
	PermissionAdminUserUpdateActive                         // ADMIN_USER_UPDATE_ACTIVE
	PermissionAdminUserUpdateInActive                       // ADMIN_USER_UPDATE_INACTIVE

	// admin unit
	PermissionAdminUnitCreate // ADMIN_UNIT_CREATE
	PermissionAdminUnitFetch  // ADMIN_UNIT_FETCH
	PermissionAdminUnitGet    // ADMIN_UNIT_GET
	PermissionAdminUnitUpdate // ADMIN_UNIT_UPDATE
	PermissionAdminUnitDelete // ADMIN_UNIT_DELETE
)
