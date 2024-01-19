package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=Permission -output=permission_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_enum_gen.go -custom
type Permission int // @name PermissionEnum

const (
	// balance
	PermissionBalanceCreate Permission = iota + 1 // BALANCE_CREATE
	PermissionBalanceFetch                        // BALANCE_FETCH
	PermissionBalanceGet                          // BALANCE_GET
	PermissionBalanceUpdate                       // BALANCE_UPDATE
	PermissionBalanceDelete                       // BALANCE_DELETE

	// product unit
	PermissionProductUnitCreate // PRODUCT_UNIT_CREATE
	PermissionProductUnitUpload // PRODUCT_UNIT_UPLOAD
	PermissionProductUnitGet    // PRODUCT_UNIT_GET
	PermissionProductUnitUpdate // PRODUCT_UNIT_UPDATE
	PermissionProductUnitDelete // PRODUCT_UNIT_DELETE

	// role
	PermissionRoleOptionForUserForm // ROLE_OPTION_FOR_USER_FORM

	// user
	PermissionUserCreate         // USER_CREATE
	PermissionUserUpdate         // USER_UPDATE
	PermissionUserUpdatePassword // USER_UPDATE_PASSWORD
	PermissionUserUpdateActive   // USER_UPDATE_ACTIVE
	PermissionUserUpdateInActive // USER_UPDATE_INACTIVE
	PermissionUserAddRole        // USER_ADD_ROLE
	PermissionUserDeleteRole     // USER_DELETE_ROLE

	// unit
	PermissionUnitCreate // UNIT_CREATE
	PermissionUnitFetch  // UNIT_FETCH
	PermissionUnitGet    // UNIT_GET
	PermissionUnitUpdate // UNIT_UPDATE
	PermissionUnitDelete // UNIT_DELETE

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
