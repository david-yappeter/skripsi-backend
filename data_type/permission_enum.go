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

	// cart
	PermissionCartGetActive   // CART_GET_ACTIVE
	PermissionCartAddItem     // CART_ADD_ITEM
	PermissionCartUpdateItem  // CART_UPDATE_ITEM
	PermissionCartDeleteItem  // CART_DELETE_ITEM
	PermissionCartSetActive   // CART_SET_ACTIVE
	PermissionCartSetInActive // CART_SET_IN_ACTIVE
	PermissionCartDelete      // CART_DELETE

	// cashier session
	PermissionCashierSessionStart      // CASHIER_SESSION_START
	PermissionCashierSessionGetCurrent // CASHIER_SESSION_GET_CURRENT
	PermissionCashierSessionEnd        // CASHIER_SESSION_END

	// customer
	PermissionCustomerCreate                     // CUSTOMER_CREATE
	PermissionCustomerFetch                      // CUSTOMER_FETCH
	PermissionCustomerGet                        // CUSTOMER_GET
	PermissionCustomerUpdate                     // CUSTOMER_UPDATE
	PermissionCustomerDelete                     // CUSTOMER_DELETE
	PermissionCustomerOptionForDeliveryOrderForm // CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM

	// customer debt
	PermissionCustomerDebtUploadImage // CUSTOMER_DEBT_UPLOAD_IMAGE
	PermissionCustomerDebtFetch       // CUSTOMER_DEBT_FETCH
	PermissionCustomerDebtGet         // CUSTOMER_DEBT_GET
	PermissionCustomerDebtPayment     // CUSTOMER_DEBT_PAYMENT

	// customer type
	PermissionCustomerTypeCreate                // CUSTOMER_TYPE_CREATE
	PermissionCustomerTypeFetch                 // CUSTOMER_TYPE_FETCH
	PermissionCustomerTypeGet                   // CUSTOMER_TYPE_GET
	PermissionCustomerTypeUpdate                // CUSTOMER_TYPE_UPDATE
	PermissionCustomerTypeDelete                // CUSTOMER_TYPE_DELETE
	PermissionCustomerTypeOptionForCustomerForm // CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM
	PermissionCustomerTypeAddDiscount           // CUSTOMER_TYPE_OPTION_ADD_DISCOUNT
	PermissionCustomerTypeUpdateDiscount        // CUSTOMER_TYPE_OPTION_UPDATE_DISCOUNT
	PermissionCustomerTypeDeleteDiscount        // CUSTOMER_TYPE_OPTION_DELETE_DISCOUNT

	// delivery order
	PermissionDeliveryOrderCreate        // DELIVERY_ORDER_CREATE
	PermissionDeliveryOrderUpload        // DELIVERY_ORDER_UPLOAD
	PermissionDeliveryOrderAddItem       // DELIVERY_ORDER_ADD_ITEM
	PermissionDeliveryOrderAddImage      // DELIVERY_ORDER_ADD_IMAGE
	PermissionDeliveryOrderAddDriver     // DELIVERY_ORDER_ADD_DRIVER
	PermissionDeliveryOrderFetch         // DELIVERY_ORDER_FETCH
	PermissionDeliveryOrderGet           // DELIVERY_ORDER_GET
	PermissionDeliveryOrderMarkOngoing   // DELIVERY_ORDER_MARK_ONGOING
	PermissionDeliveryOrderCancel        // DELIVERY_ORDER_CANCEL
	PermissionDeliveryOrderMarkCompleted // DELIVERY_ORDER_MARK_COMPLETED
	PermissionDeliveryOrderDelete        // DELIVERY_ORDER_DELETE
	PermissionDeliveryOrderDeleteItem    // DELIVERY_ORDER_DELETE_ITEM
	PermissionDeliveryOrderDeleteImage   // DELIVERY_ORDER_DELETE_IMAGE
	PermissionDeliveryOrderDeleteDriver  // DELIVERY_ORDER_DELETE_DRIVER

	// product
	PermissionProductCreate                      // PRODUCT_CREATE
	PermissionProductUpload                      // PRODUCT_UPLOAD
	PermissionProductFetch                       // PRODUCT_FETCH
	PermissionProductGet                         // PRODUCT_GET
	PermissionProductUpdate                      // PRODUCT_UPDATE
	PermissionProductDelete                      // PRODUCT_DELETE
	PermissionProductOptionForProductReceiveForm // PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORM
	PermissionProductOptionForDeliveryOrderForm  // PRODUCT_OPTION_FOR_DELIVERY_ORDER_FORM
	PermissionProductOptionForCustomerTypeForm   // PRODUCT_OPTION_FOR_CUSTOMER_TYPE_FORM

	// product receive
	PermissionProductReceiveCreate       // PRODUCT_RECEIVE_CREATE
	PermissionProductReceiveUpload       // PRODUCT_RECEIVE_UPLOAD
	PermissionProductReceiveAddItem      // PRODUCT_RECEIVE_ADD_ITEM
	PermissionProductReceiveAddImage     // PRODUCT_RECEIVE_ADD_IMAGE
	PermissionProductReceiveCancel       // PRODUCT_RECEIVE_CANCEL
	PermissionProductReceiveMarkComplete // PRODUCT_RECEIVE_MARK_COMPLETE
	PermissionProductReceiveFetch        // PRODUCT_RECEIVE_FETCH
	PermissionProductReceiveGet          // PRODUCT_RECEIVE_GET
	PermissionProductReceiveDelete       // PRODUCT_RECEIVE_DELETE
	PermissionProductReceiveDeleteItem   // PRODUCT_RECEIVE_DELETE_ITEM
	PermissionProductReceiveDeleteImage  // PRODUCT_RECEIVE_DELETE_IMAGE

	// product stock
	PermissionProductStockFetch      // PRODUCT_STOCK_FETCH
	PermissionProductStockGet        // PRODUCT_STOCK_GET
	PermissionProductStockAdjustment // PRODUCT_STOCK_ADJUSTMENT

	// product unit
	PermissionProductUnitCreate                      // PRODUCT_UNIT_CREATE
	PermissionProductUnitGet                         // PRODUCT_UNIT_GET
	PermissionProductUnitUpdate                      // PRODUCT_UNIT_UPDATE
	PermissionProductUnitDelete                      // PRODUCT_UNIT_DELETE
	PermissionProductUnitOptionForProductReceiveForm // PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORM
	PermissionProductUnitOptionForDeliveryOrderForm  // PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORM

	// role
	PermissionRoleOptionForUserForm // ROLE_OPTION_FOR_USER_FORM

	// shop order
	PermissionShopOrderFetch // SHOP_ORDER_FETCH
	PermissionShopOrderGet   // SHOP_ORDER_GET

	// supplier
	PermissionSupplierCreate                      // SUPPLIER_CREATE
	PermissionSupplierFetch                       // SUPPLIER_FETCH
	PermissionSupplierGet                         // SUPPLIER_GET
	PermissionSupplierUpdate                      // SUPPLIER_UPDATE
	PermissionSupplierDelete                      // SUPPLIER_DELETE
	PermissionSupplierOptionForProductReceiveForm // SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM

	// supplier type
	PermissionSupplierTypeCreate                // SUPPLIER_TYPE_CREATE
	PermissionSupplierTypeFetch                 // SUPPLIER_TYPE_FETCH
	PermissionSupplierTypeGet                   // SUPPLIER_TYPE_GET
	PermissionSupplierTypeUpdate                // SUPPLIER_TYPE_UPDATE
	PermissionSupplierTypeDelete                // SUPPLIER_TYPE_DELETE
	PermissionSupplierTypeOptionForSupplierForm // SUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORM

	// tiktok product
	PermissionTiktokProductCreate                // TIKTOK_PRODUCT_CREATE
	PermissionTiktokProductUploadImage           // TIKTOK_PRODUCT_UPLOAD_IMAGE
	PermissionTiktokProductFetchBrands           // TIKTOK_PRODUCT_FETCH_BRANDS
	PermissionTiktokProductFetchCategories       // TIKTOK_PRODUCT_FETCH_CATEGORIES
	PermissionTiktokProductGetCategoryRules      // TIKTOK_PRODUCT_GET_RULES
	PermissionTiktokProductGetCategoryAttributes // TIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTES
	PermissionTiktokProductGet                   // TIKTOK_PRODUCT_GET
	PermissionTiktokProductUpdate                // TIKTOK_PRODUCT_UPDATE
	PermissionTiktokProductRecommendedCategory   // TIKTOK_PRODUCT_RECOMMENDED_CATEGORY
	PermissionTiktokProductActivate              // TIKTOK_PRODUCT_RECOMMENDED_ACTIVATE
	PermissionTiktokProductDeactivate            // TIKTOK_PRODUCT_RECOMMENDED_DEACTIVATE

	// transaction
	PermissionTransactionCheckoutCart // TRANSACTION_CHECKOUT_CART

	// user
	PermissionUserCreate         // USER_CREATE
	PermissionUserFetch          // USER_FETCH
	PermissionUserGet            // USER_GET
	PermissionUserUpdate         // USER_UPDATE
	PermissionUserUpdatePassword // USER_UPDATE_PASSWORD
	PermissionUserUpdateActive   // USER_UPDATE_ACTIVE
	PermissionUserUpdateInActive // USER_UPDATE_INACTIVE
	PermissionUserAddRole        // USER_ADD_ROLE
	PermissionUserDeleteRole     // USER_DELETE_ROLE

	// unit
	PermissionUnitCreate                   // UNIT_CREATE
	PermissionUnitFetch                    // UNIT_FETCH
	PermissionUnitGet                      // UNIT_GET
	PermissionUnitUpdate                   // UNIT_UPDATE
	PermissionUnitDelete                   // UNIT_DELETE
	PermissionUnitOptionForProductUnitForm // UNIT_OPTION_FOR_PRODUCT_UNIT_FORM
)
