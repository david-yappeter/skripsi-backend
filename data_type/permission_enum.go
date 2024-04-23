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
	PermissionCashierSessionFetch            // CASHIER_SESSION_FETCH
	PermissionCashierSessionStart            // CASHIER_SESSION_START
	PermissionCashierSessionGet              // CASHIER_SESSION_GET
	PermissionCashierSessionFetchTransaction // CASHIER_SESSION_FETCH_TRANSACTION
	PermissionCashierSessionDownloadReport   // CASHIER_SESSION_DOWNLOAD_REPORT
	PermissionCashierSessionGetCurrent       // CASHIER_SESSION_GET_CURRENT
	PermissionCashierSessionEnd              // CASHIER_SESSION_END

	// customer
	PermissionCustomerCreate                       // CUSTOMER_CREATE
	PermissionCustomerFetch                        // CUSTOMER_FETCH
	PermissionCustomerGet                          // CUSTOMER_GET
	PermissionCustomerUpdate                       // CUSTOMER_UPDATE
	PermissionCustomerDelete                       // CUSTOMER_DELETE
	PermissionCustomerOptionForDeliveryOrderForm   // CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM
	PermissionCustomerOptionForDeliveryOrderFilter // CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FILTER

	// customer debt
	PermissionCustomerDebtUploadImage // CUSTOMER_DEBT_UPLOAD_IMAGE
	PermissionCustomerDebtFetch       // CUSTOMER_DEBT_FETCH
	PermissionCustomerDebtGet         // CUSTOMER_DEBT_GET
	PermissionCustomerDebtPayment     // CUSTOMER_DEBT_PAYMENT

	// customer type
	PermissionCustomerTypeCreate                                           // CUSTOMER_TYPE_CREATE
	PermissionCustomerTypeFetch                                            // CUSTOMER_TYPE_FETCH
	PermissionCustomerTypeGet                                              // CUSTOMER_TYPE_GET
	PermissionCustomerTypeUpdate                                           // CUSTOMER_TYPE_UPDATE
	PermissionCustomerTypeDelete                                           // CUSTOMER_TYPE_DELETE
	PermissionCustomerTypeAddDiscount                                      // CUSTOMER_TYPE_ADD_DISCOUNT
	PermissionCustomerTypeUpdateDiscount                                   // CUSTOMER_TYPE_UPDATE_DISCOUNT
	PermissionCustomerTypeDeleteDiscount                                   // CUSTOMER_TYPE_DELETE_DISCOUNT
	PermissionCustomerTypeOptionForCustomerForm                            // CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM
	PermissionCustomerTypeOptionForWhatsappProductPriceChangeBroadcastForm // CUSTOMER_TYPE_OPTION_FOR_WHATSAPP_PRODUCT_PRICE_CHANGE_BROADCAST_FORM

	// debt
	PermissionDebtUploadImage // DEBT_UPLOAD_IMAGE
	PermissionDebtFetch       // DEBT_FETCH
	PermissionDebtGet         // DEBT_GET
	PermissionDebtPayment     // DEBT_PAYMENT

	// delivery order
	PermissionDeliveryOrderCreate           // DELIVERY_ORDER_CREATE
	PermissionDeliveryOrderUpload           // DELIVERY_ORDER_UPLOAD
	PermissionDeliveryOrderAddItem          // DELIVERY_ORDER_ADD_ITEM
	PermissionDeliveryOrderAddImage         // DELIVERY_ORDER_ADD_IMAGE
	PermissionDeliveryOrderAddDriver        // DELIVERY_ORDER_ADD_DRIVER
	PermissionDeliveryOrderFetch            // DELIVERY_ORDER_FETCH
	PermissionDeliveryOrderFetchDriver      // DELIVERY_ORDER_FETCH_DRIVER
	PermissionDeliveryOrderGet              // DELIVERY_ORDER_GET
	PermissionDeliveryOrderActiveForDriver  // DELIVERY_ORDER_ACTIVE_FOR_DRIVER
	PermissionDeliveryOrderMarkOngoing      // DELIVERY_ORDER_MARK_ONGOING
	PermissionDeliveryOrderDelivering       // DELIVERY_ORDER_DELIVERING
	PermissionDeliveryOrderUpdate           // DELIVERY_ORDER_UPDATE
	PermissionDeliveryOrderCancel           // DELIVERY_ORDER_CANCEL
	PermissionDeliveryOrderMarkCompleted    // DELIVERY_ORDER_MARK_COMPLETED
	PermissionDeliveryOrderDeliveryLocation // DELIVERY_ORDER_DELIVERY_LOCATION
	PermissionDeliveryOrderDelete           // DELIVERY_ORDER_DELETE
	PermissionDeliveryOrderDeleteItem       // DELIVERY_ORDER_DELETE_ITEM
	PermissionDeliveryOrderDeleteImage      // DELIVERY_ORDER_DELETE_IMAGE
	PermissionDeliveryOrderDeleteDriver     // DELIVERY_ORDER_DELETE_DRIVER

	// delivery order review
	PermissionDeliveryOrderReviewFetch // DELIVERY_ORDER_REVIEW_FETCH
	PermissionDeliveryOrderReviewGet   // DELIVERY_ORDER_REVIEW_GET

	// product
	PermissionProductCreate                            // PRODUCT_CREATE
	PermissionProductUpload                            // PRODUCT_UPLOAD
	PermissionProductFetch                             // PRODUCT_FETCH
	PermissionProductGet                               // PRODUCT_GET
	PermissionProductUpdate                            // PRODUCT_UPDATE
	PermissionProductDelete                            // PRODUCT_DELETE
	PermissionProductOptionForProductReceiveItemForm   // PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM
	PermissionProductOptionForDeliveryOrderItemForm    // PRODUCT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM
	PermissionProductOptionForCustomerTypeDiscountForm // PRODUCT_OPTION_FOR_CUSTOMER_TYPE_DISCOUNT_FORM
	PermissionProductOptionForCartAddItemForm          // PRODUCT_OPTION_FOR_CART_ADD_ITEM_FORM

	// product discount
	PermissionProductDiscountCreate // PRODUCT_DISCOUNT_CREATE
	PermissionProductDiscountUpdate // PRODUCT_DISCOUNT_UPDATE
	PermissionProductDiscountDelete // PRODUCT_DISCOUNT_DELETE

	// product receive
	PermissionProductReceiveCreate       // PRODUCT_RECEIVE_CREATE
	PermissionProductReceiveUpload       // PRODUCT_RECEIVE_UPLOAD
	PermissionProductReceiveAddItem      // PRODUCT_RECEIVE_ADD_ITEM
	PermissionProductReceiveAddImage     // PRODUCT_RECEIVE_ADD_IMAGE
	PermissionProductReceiveUpdate       // PRODUCT_RECEIVE_UPDATE
	PermissionProductReceiveCancel       // PRODUCT_RECEIVE_CANCEL
	PermissionProductReceiveMarkComplete // PRODUCT_RECEIVE_MARK_COMPLETE
	PermissionProductReceiveFetch        // PRODUCT_RECEIVE_FETCH
	PermissionProductReceiveGet          // PRODUCT_RECEIVE_GET
	PermissionProductReceiveDelete       // PRODUCT_RECEIVE_DELETE
	PermissionProductReceiveDeleteItem   // PRODUCT_RECEIVE_DELETE_ITEM
	PermissionProductReceiveDeleteImage  // PRODUCT_RECEIVE_DELETE_IMAGE

	// product stock
	PermissionProductStockFetch          // PRODUCT_STOCK_FETCH
	PermissionProductStockGet            // PRODUCT_STOCK_GET
	PermissionProductStockDownloadReport // PRODUCT_STOCK_DOWNLOAD_REPORT
	PermissionProductStockAdjustment     // PRODUCT_STOCK_ADJUSTMENT

	// product unit
	PermissionProductUnitCreate                          // PRODUCT_UNIT_CREATE
	PermissionProductUnitGet                             // PRODUCT_UNIT_GET
	PermissionProductUnitUpdate                          // PRODUCT_UNIT_UPDATE
	PermissionProductUnitDelete                          // PRODUCT_UNIT_DELETE
	PermissionProductUnitOptionForProductReceiveItemForm // PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM
	PermissionProductUnitOptionForDeliveryOrderItemForm  // PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM

	// role
	PermissionRoleOptionForUserForm // ROLE_OPTION_FOR_USER_FORM

	// shop order
	PermissionShopOrderFetch // SHOP_ORDER_FETCH
	PermissionShopOrderGet   // SHOP_ORDER_GET

	// ssr
	PermissionSsrWhatsappLogin // SSR_WHATSAPP_LOGIN

	// supplier
	PermissionSupplierCreate                        // SUPPLIER_CREATE
	PermissionSupplierFetch                         // SUPPLIER_FETCH
	PermissionSupplierGet                           // SUPPLIER_GET
	PermissionSupplierUpdate                        // SUPPLIER_UPDATE
	PermissionSupplierDelete                        // SUPPLIER_DELETE
	PermissionSupplierOptionForProductReceiveForm   // SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM
	PermissionSupplierOptionForProductReceiveFilter // SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FILTER

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
	PermissionTransactionGet          // TRANSACTION_GET

	// user
	PermissionUserCreate                           // USER_CREATE
	PermissionUserFetch                            // USER_FETCH
	PermissionUserGet                              // USER_GET
	PermissionUserUpdate                           // USER_UPDATE
	PermissionUserUpdatePassword                   // USER_UPDATE_PASSWORD
	PermissionUserUpdateActive                     // USER_UPDATE_ACTIVE
	PermissionUserUpdateInActive                   // USER_UPDATE_INACTIVE
	PermissionUserAddRole                          // USER_ADD_ROLE
	PermissionUserDeleteRole                       // USER_DELETE_ROLE
	PermissionUserOptionForCashierSessionFilter    // USER_OPTION_FOR_CASHIER_SESSION_FILTER
	PermissionUserOptionForDeliveryOrderDriverForm // USER_OPTION_FOR_DELIVERY_ORDER_DRIVER_FORM

	// unit
	PermissionUnitCreate                         // UNIT_CREATE
	PermissionUnitFetch                          // UNIT_FETCH
	PermissionUnitGet                            // UNIT_GET
	PermissionUnitUpdate                         // UNIT_UPDATE
	PermissionUnitDelete                         // UNIT_DELETE
	PermissionUnitOptionForProductUnitForm       // UNIT_OPTION_FOR_PRODUCT_UNIT_FORM
	PermissionUnitOptionForProductUnitToUnitForm // UNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORM

	// whatsapp
	PermissionWhatsappIsLoggedIn                    // WHATSAPP_IS_LOGGED_IN
	PermissionWhatsappProductPriceChangeBroadcast   // WHATSAPP_PRODUCT_PRICE_CHANGE_BROADCAST
	PermissionWhatsappCustomerTypeDiscountBroadcast // WHATSAPP_CUSTOMER_TYPE_DISCOUNT_BROADCAST
	PermissionWhatsappLogout                        // WHATSAPP_LOGOUT
)
