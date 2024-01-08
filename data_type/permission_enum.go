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

	// admin supplier type
	PermissionAdminSupplierTypeCreate // ADMIN_SUPPLIER_TYPE_CREATE
	PermissionAdminSupplierTypeFetch  // ADMIN_SUPPLIER_TYPE_FETCH
	PermissionAdminSupplierTypeGet    // ADMIN_SUPPLIER_TYPE_GET
	PermissionAdminSupplierTypeUpdate // ADMIN_SUPPLIER_TYPE_UPDATE
	PermissionAdminSupplierTypeDelete // ADMIN_SUPPLIER_TYPE_DELETE

	// supplier type
	PermissionSupplierTypeCreate // SUPPLIER_TYPE_CREATE
	PermissionSupplierTypeFetch  // SUPPLIER_TYPE_FETCH
	PermissionSupplierTypeGet    // SUPPLIER_TYPE_GET
	PermissionSupplierTypeUpdate // SUPPLIER_TYPE_UPDATE
	PermissionSupplierTypeDelete // SUPPLIER_TYPE_DELETE
)
