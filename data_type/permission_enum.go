package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=Permission -output=permission_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_enum_gen.go -custom
type Permission int // @name PermissionEnum

const (
	// admin balance
	PermissionAdminBalanceCreate Permission = iota + 1 // ADMIN_BALANCE_CREATE
	PermissionAdminBalanceFetch                        // ADMIN_BALANCE_FETCH
	PermissionAdminBalanceGet                          // ADMIN_BALANCE_GET
	PermissionAdminBalanceUpdate                       // ADMIN_BALANCE_UPDATE
	PermissionAdminBalanceDelete                       // ADMIN_BALANCE_DELETE

	// admin supplier
	PermissionAdminSupplierCreate // ADMIN_SUPPLIER_CREATE
	PermissionAdminSupplierFetch  // ADMIN_SUPPLIER_FETCH
	PermissionAdminSupplierGet    // ADMIN_SUPPLIER_GET
	PermissionAdminSupplierUpdate // ADMIN_SUPPLIER_UPDATE
	PermissionAdminSupplierDelete // ADMIN_SUPPLIER_DELETE

	// admin supplier type
	PermissionAdminSupplierTypeCreate // ADMIN_SUPPLIER_TYPE_CREATE
	PermissionAdminSupplierTypeFetch  // ADMIN_SUPPLIER_TYPE_FETCH
	PermissionAdminSupplierTypeGet    // ADMIN_SUPPLIER_TYPE_GET
	PermissionAdminSupplierTypeUpdate // ADMIN_SUPPLIER_TYPE_UPDATE
	PermissionAdminSupplierTypeDelete // ADMIN_SUPPLIER_TYPE_DELETE

	// admin product unit
	PermissionAdminProductUnitCreate // ADMIN_PRODUCT_UNIT_CREATE
	PermissionAdminProductUnitUpload // ADMIN_PRODUCT_UNIT_UPLOAD
	PermissionAdminProductUnitGet    // ADMIN_PRODUCT_UNIT_GET
	PermissionAdminProductUnitUpdate // ADMIN_PRODUCT_UNIT_UPDATE
	PermissionAdminProductUnitDelete // ADMIN_PRODUCT_UNIT_DELETE

	// admin role
	PermissionAdminRoleOptionForUserForm // ADMIN_ROLE_OPTION_FOR_USER_FORM

	// admin user
	PermissionAdminUserCreate         // ADMIN_USER_CREATE
	PermissionAdminUserUpdate         // ADMIN_USER_UPDATE
	PermissionAdminUserUpdatePassword // ADMIN_USER_UPDATE_PASSWORD
	PermissionAdminUserUpdateActive   // ADMIN_USER_UPDATE_ACTIVE
	PermissionAdminUserUpdateInActive // ADMIN_USER_UPDATE_INACTIVE
	PermissionAdminUserAddRole        // ADMIN_USER_ADD_ROLE
	PermissionAdminUserDeleteRole     // ADMIN_USER_DELETE_ROLE

	// admin unit
	PermissionAdminUnitCreate // ADMIN_UNIT_CREATE
	PermissionAdminUnitFetch  // ADMIN_UNIT_FETCH
	PermissionAdminUnitGet    // ADMIN_UNIT_GET
	PermissionAdminUnitUpdate // ADMIN_UNIT_UPDATE
	PermissionAdminUnitDelete // ADMIN_UNIT_DELETE

	// customer
	PermissionCustomerCreate // CUSTOMER_CREATE
	PermissionCustomerFetch  // CUSTOMER_FETCH
	PermissionCustomerGet    // CUSTOMER_GET
	PermissionCustomerUpdate // CUSTOMER_UPDATE
	PermissionCustomerDelete // CUSTOMER_DELETE

	// product receive
	PermissionProductReceiveCreate      // PRODUCT_RECEIVE_CREATE
	PermissionProductReceiveUpload      // PRODUCT_RECEIVE_UPLOAD
	PermissionProductReceiveAddItem     // PRODUCT_RECEIVE_ADD_ITEM
	PermissionProductReceiveAddImage    // PRODUCT_RECEIVE_ADD_IMAGE
	PermissionProductReceiveFetch       // PRODUCT_RECEIVE_FETCH
	PermissionProductReceiveGet         // PRODUCT_RECEIVE_GET
	PermissionProductReceiveDelete      // PRODUCT_RECEIVE_DELETE
	PermissionProductReceiveDeleteItem  // PRODUCT_RECEIVE_DELETE_ITEM
	PermissionProductReceiveDeleteImage // PRODUCT_RECEIVE_DELETE_IMAGE

	// supplier
	PermissionSupplierCreate // SUPPLIER_CREATE
	PermissionSupplierFetch  // SUPPLIER_FETCH
	PermissionSupplierGet    // SUPPLIER_GET
	PermissionSupplierUpdate // SUPPLIER_UPDATE
	PermissionSupplierDelete // SUPPLIER_DELETE

	// supplier type
	PermissionSupplierTypeCreate // SUPPLIER_TYPE_CREATE
	PermissionSupplierTypeFetch  // SUPPLIER_TYPE_FETCH
	PermissionSupplierTypeGet    // SUPPLIER_TYPE_GET
	PermissionSupplierTypeUpdate // SUPPLIER_TYPE_UPDATE
	PermissionSupplierTypeDelete // SUPPLIER_TYPE_DELETE
)
