// Code generated by "go run myapp/tool/stringer -linecomment -type=Permission -output=permission_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_enum_gen.go -custom"; DO NOT EDIT.

package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PermissionBalanceCreate-1]
	_ = x[PermissionBalanceFetch-2]
	_ = x[PermissionBalanceGet-3]
	_ = x[PermissionBalanceUpdate-4]
	_ = x[PermissionBalanceDelete-5]
	_ = x[PermissionCartGetActive-6]
	_ = x[PermissionCartAddItem-7]
	_ = x[PermissionCartUpdateItem-8]
	_ = x[PermissionCartDeleteItem-9]
	_ = x[PermissionCartSetActive-10]
	_ = x[PermissionCartSetInActive-11]
	_ = x[PermissionCartDelete-12]
	_ = x[PermissionCashierSessionFetch-13]
	_ = x[PermissionCashierSessionStart-14]
	_ = x[PermissionCashierSessionGet-15]
	_ = x[PermissionCashierSessionFetchTransaction-16]
	_ = x[PermissionCashierSessionDownloadReport-17]
	_ = x[PermissionCashierSessionGetCurrent-18]
	_ = x[PermissionCashierSessionEnd-19]
	_ = x[PermissionCustomerCreate-20]
	_ = x[PermissionCustomerFetch-21]
	_ = x[PermissionCustomerGet-22]
	_ = x[PermissionCustomerUpdate-23]
	_ = x[PermissionCustomerDelete-24]
	_ = x[PermissionCustomerOptionForDeliveryOrderForm-25]
	_ = x[PermissionCustomerOptionForDeliveryOrderFilter-26]
	_ = x[PermissionCustomerDebtUploadImage-27]
	_ = x[PermissionCustomerDebtFetch-28]
	_ = x[PermissionCustomerDebtGet-29]
	_ = x[PermissionCustomerDebtPayment-30]
	_ = x[PermissionCustomerTypeCreate-31]
	_ = x[PermissionCustomerTypeFetch-32]
	_ = x[PermissionCustomerTypeGet-33]
	_ = x[PermissionCustomerTypeUpdate-34]
	_ = x[PermissionCustomerTypeDelete-35]
	_ = x[PermissionCustomerTypeAddDiscount-36]
	_ = x[PermissionCustomerTypeUpdateDiscount-37]
	_ = x[PermissionCustomerTypeDeleteDiscount-38]
	_ = x[PermissionCustomerTypeOptionForCustomerForm-39]
	_ = x[PermissionDeliveryOrderCreate-40]
	_ = x[PermissionDeliveryOrderUpload-41]
	_ = x[PermissionDeliveryOrderAddItem-42]
	_ = x[PermissionDeliveryOrderAddImage-43]
	_ = x[PermissionDeliveryOrderAddDriver-44]
	_ = x[PermissionDeliveryOrderFetch-45]
	_ = x[PermissionDeliveryOrderFetchDriver-46]
	_ = x[PermissionDeliveryOrderGet-47]
	_ = x[PermissionDeliveryOrderActiveForDriver-48]
	_ = x[PermissionDeliveryOrderMarkOngoing-49]
	_ = x[PermissionDeliveryOrderDelivering-50]
	_ = x[PermissionDeliveryOrderCancel-51]
	_ = x[PermissionDeliveryOrderMarkCompleted-52]
	_ = x[PermissionDeliveryOrderDeliveryLocation-53]
	_ = x[PermissionDeliveryOrderDelete-54]
	_ = x[PermissionDeliveryOrderDeleteItem-55]
	_ = x[PermissionDeliveryOrderDeleteImage-56]
	_ = x[PermissionDeliveryOrderDeleteDriver-57]
	_ = x[PermissionProductCreate-58]
	_ = x[PermissionProductUpload-59]
	_ = x[PermissionProductFetch-60]
	_ = x[PermissionProductGet-61]
	_ = x[PermissionProductUpdate-62]
	_ = x[PermissionProductDelete-63]
	_ = x[PermissionProductOptionForProductReceiveItemForm-64]
	_ = x[PermissionProductOptionForDeliveryOrderItemForm-65]
	_ = x[PermissionProductOptionForCustomerTypeDiscountForm-66]
	_ = x[PermissionProductOptionForCartAddItemForm-67]
	_ = x[PermissionProductDiscountCreate-68]
	_ = x[PermissionProductDiscountUpdate-69]
	_ = x[PermissionProductDiscountDelete-70]
	_ = x[PermissionProductReceiveCreate-71]
	_ = x[PermissionProductReceiveUpload-72]
	_ = x[PermissionProductReceiveAddItem-73]
	_ = x[PermissionProductReceiveAddImage-74]
	_ = x[PermissionProductReceiveUpdate-75]
	_ = x[PermissionProductReceiveCancel-76]
	_ = x[PermissionProductReceiveMarkComplete-77]
	_ = x[PermissionProductReceiveFetch-78]
	_ = x[PermissionProductReceiveGet-79]
	_ = x[PermissionProductReceiveDelete-80]
	_ = x[PermissionProductReceiveDeleteItem-81]
	_ = x[PermissionProductReceiveDeleteImage-82]
	_ = x[PermissionProductStockFetch-83]
	_ = x[PermissionProductStockGet-84]
	_ = x[PermissionProductStockDownloadReport-85]
	_ = x[PermissionProductStockAdjustment-86]
	_ = x[PermissionProductUnitCreate-87]
	_ = x[PermissionProductUnitGet-88]
	_ = x[PermissionProductUnitUpdate-89]
	_ = x[PermissionProductUnitDelete-90]
	_ = x[PermissionProductUnitOptionForProductReceiveItemForm-91]
	_ = x[PermissionProductUnitOptionForDeliveryOrderItemForm-92]
	_ = x[PermissionRoleOptionForUserForm-93]
	_ = x[PermissionShopOrderFetch-94]
	_ = x[PermissionShopOrderGet-95]
	_ = x[PermissionSsrWhatsappLogin-96]
	_ = x[PermissionSupplierCreate-97]
	_ = x[PermissionSupplierFetch-98]
	_ = x[PermissionSupplierGet-99]
	_ = x[PermissionSupplierUpdate-100]
	_ = x[PermissionSupplierDelete-101]
	_ = x[PermissionSupplierOptionForProductReceiveForm-102]
	_ = x[PermissionSupplierOptionForProductReceiveFilter-103]
	_ = x[PermissionSupplierTypeCreate-104]
	_ = x[PermissionSupplierTypeFetch-105]
	_ = x[PermissionSupplierTypeGet-106]
	_ = x[PermissionSupplierTypeUpdate-107]
	_ = x[PermissionSupplierTypeDelete-108]
	_ = x[PermissionSupplierTypeOptionForSupplierForm-109]
	_ = x[PermissionTiktokProductCreate-110]
	_ = x[PermissionTiktokProductUploadImage-111]
	_ = x[PermissionTiktokProductFetchBrands-112]
	_ = x[PermissionTiktokProductFetchCategories-113]
	_ = x[PermissionTiktokProductGetCategoryRules-114]
	_ = x[PermissionTiktokProductGetCategoryAttributes-115]
	_ = x[PermissionTiktokProductGet-116]
	_ = x[PermissionTiktokProductUpdate-117]
	_ = x[PermissionTiktokProductRecommendedCategory-118]
	_ = x[PermissionTiktokProductActivate-119]
	_ = x[PermissionTiktokProductDeactivate-120]
	_ = x[PermissionTransactionCheckoutCart-121]
	_ = x[PermissionTransactionGet-122]
	_ = x[PermissionUserCreate-123]
	_ = x[PermissionUserFetch-124]
	_ = x[PermissionUserGet-125]
	_ = x[PermissionUserUpdate-126]
	_ = x[PermissionUserUpdatePassword-127]
	_ = x[PermissionUserUpdateActive-128]
	_ = x[PermissionUserUpdateInActive-129]
	_ = x[PermissionUserAddRole-130]
	_ = x[PermissionUserDeleteRole-131]
	_ = x[PermissionUserOptionForCashierSessionFilter-132]
	_ = x[PermissionUserOptionForDeliveryOrderDriverForm-133]
	_ = x[PermissionUnitCreate-134]
	_ = x[PermissionUnitFetch-135]
	_ = x[PermissionUnitGet-136]
	_ = x[PermissionUnitUpdate-137]
	_ = x[PermissionUnitDelete-138]
	_ = x[PermissionUnitOptionForProductUnitForm-139]
	_ = x[PermissionUnitOptionForProductUnitToUnitForm-140]
	_ = x[PermissionWhatsappIsLoggedIn-141]
	_ = x[PermissionWhatsappLogout-142]
}

const _Permission_nameReadable = "BALANCE_CREATE, BALANCE_FETCH, BALANCE_GET, BALANCE_UPDATE, BALANCE_DELETE, CART_GET_ACTIVE, CART_ADD_ITEM, CART_UPDATE_ITEM, CART_DELETE_ITEM, CART_SET_ACTIVE, CART_SET_IN_ACTIVE, CART_DELETE, CASHIER_SESSION_FETCH, CASHIER_SESSION_START, CASHIER_SESSION_GET, CASHIER_SESSION_FETCH_TRANSACTION, CASHIER_SESSION_DOWNLOAD_REPORT, CASHIER_SESSION_GET_CURRENT, CASHIER_SESSION_END, CUSTOMER_CREATE, CUSTOMER_FETCH, CUSTOMER_GET, CUSTOMER_UPDATE, CUSTOMER_DELETE, CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM, CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FILTER, CUSTOMER_DEBT_UPLOAD_IMAGE, CUSTOMER_DEBT_FETCH, CUSTOMER_DEBT_GET, CUSTOMER_DEBT_PAYMENT, CUSTOMER_TYPE_CREATE, CUSTOMER_TYPE_FETCH, CUSTOMER_TYPE_GET, CUSTOMER_TYPE_UPDATE, CUSTOMER_TYPE_DELETE, CUSTOMER_TYPE_ADD_DISCOUNT, CUSTOMER_TYPE_UPDATE_DISCOUNT, CUSTOMER_TYPE_DELETE_DISCOUNT, CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM, DELIVERY_ORDER_CREATE, DELIVERY_ORDER_UPLOAD, DELIVERY_ORDER_ADD_ITEM, DELIVERY_ORDER_ADD_IMAGE, DELIVERY_ORDER_ADD_DRIVER, DELIVERY_ORDER_FETCH, DELIVERY_ORDER_FETCH_DRIVER, DELIVERY_ORDER_GET, DELIVERY_ORDER_ACTIVE_FOR_DRIVER, DELIVERY_ORDER_MARK_ONGOING, DELIVERY_ORDER_DELIVERING, DELIVERY_ORDER_CANCEL, DELIVERY_ORDER_MARK_COMPLETED, DELIVERY_ORDER_DELIVERY_LOCATION, DELIVERY_ORDER_DELETE, DELIVERY_ORDER_DELETE_ITEM, DELIVERY_ORDER_DELETE_IMAGE, DELIVERY_ORDER_DELETE_DRIVER, PRODUCT_CREATE, PRODUCT_UPLOAD, PRODUCT_FETCH, PRODUCT_GET, PRODUCT_UPDATE, PRODUCT_DELETE, PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM, PRODUCT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM, PRODUCT_OPTION_FOR_CUSTOMER_TYPE_DISCOUNT_FORM, PRODUCT_OPTION_FOR_CART_ADD_ITEM_FORM, PRODUCT_DISCOUNT_CREATE, PRODUCT_DISCOUNT_UPDATE, PRODUCT_DISCOUNT_DELETE, PRODUCT_RECEIVE_CREATE, PRODUCT_RECEIVE_UPLOAD, PRODUCT_RECEIVE_ADD_ITEM, PRODUCT_RECEIVE_ADD_IMAGE, PRODUCT_RECEIVE_UPDATE, PRODUCT_RECEIVE_CANCEL, PRODUCT_RECEIVE_MARK_COMPLETE, PRODUCT_RECEIVE_FETCH, PRODUCT_RECEIVE_GET, PRODUCT_RECEIVE_DELETE, PRODUCT_RECEIVE_DELETE_ITEM, PRODUCT_RECEIVE_DELETE_IMAGE, PRODUCT_STOCK_FETCH, PRODUCT_STOCK_GET, PRODUCT_STOCK_DOWNLOAD_REPORT, PRODUCT_STOCK_ADJUSTMENT, PRODUCT_UNIT_CREATE, PRODUCT_UNIT_GET, PRODUCT_UNIT_UPDATE, PRODUCT_UNIT_DELETE, PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM, PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM, ROLE_OPTION_FOR_USER_FORM, SHOP_ORDER_FETCH, SHOP_ORDER_GET, SSR_WHATSAPP_LOGIN, SUPPLIER_CREATE, SUPPLIER_FETCH, SUPPLIER_GET, SUPPLIER_UPDATE, SUPPLIER_DELETE, SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM, SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FILTER, SUPPLIER_TYPE_CREATE, SUPPLIER_TYPE_FETCH, SUPPLIER_TYPE_GET, SUPPLIER_TYPE_UPDATE, SUPPLIER_TYPE_DELETE, SUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORM, TIKTOK_PRODUCT_CREATE, TIKTOK_PRODUCT_UPLOAD_IMAGE, TIKTOK_PRODUCT_FETCH_BRANDS, TIKTOK_PRODUCT_FETCH_CATEGORIES, TIKTOK_PRODUCT_GET_RULES, TIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTES, TIKTOK_PRODUCT_GET, TIKTOK_PRODUCT_UPDATE, TIKTOK_PRODUCT_RECOMMENDED_CATEGORY, TIKTOK_PRODUCT_RECOMMENDED_ACTIVATE, TIKTOK_PRODUCT_RECOMMENDED_DEACTIVATE, TRANSACTION_CHECKOUT_CART, TRANSACTION_GET, USER_CREATE, USER_FETCH, USER_GET, USER_UPDATE, USER_UPDATE_PASSWORD, USER_UPDATE_ACTIVE, USER_UPDATE_INACTIVE, USER_ADD_ROLE, USER_DELETE_ROLE, USER_OPTION_FOR_CASHIER_SESSION_FILTER, USER_OPTION_FOR_DELIVERY_ORDER_DRIVER_FORM, UNIT_CREATE, UNIT_FETCH, UNIT_GET, UNIT_UPDATE, UNIT_DELETE, UNIT_OPTION_FOR_PRODUCT_UNIT_FORM, UNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORM, WHATSAPP_IS_LOGGED_IN, WHATSAPP_LOGOUT"

const _Permission_name = "BALANCE_CREATEBALANCE_FETCHBALANCE_GETBALANCE_UPDATEBALANCE_DELETECART_GET_ACTIVECART_ADD_ITEMCART_UPDATE_ITEMCART_DELETE_ITEMCART_SET_ACTIVECART_SET_IN_ACTIVECART_DELETECASHIER_SESSION_FETCHCASHIER_SESSION_STARTCASHIER_SESSION_GETCASHIER_SESSION_FETCH_TRANSACTIONCASHIER_SESSION_DOWNLOAD_REPORTCASHIER_SESSION_GET_CURRENTCASHIER_SESSION_ENDCUSTOMER_CREATECUSTOMER_FETCHCUSTOMER_GETCUSTOMER_UPDATECUSTOMER_DELETECUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORMCUSTOMER_OPTION_FOR_DELIVERY_ORDER_FILTERCUSTOMER_DEBT_UPLOAD_IMAGECUSTOMER_DEBT_FETCHCUSTOMER_DEBT_GETCUSTOMER_DEBT_PAYMENTCUSTOMER_TYPE_CREATECUSTOMER_TYPE_FETCHCUSTOMER_TYPE_GETCUSTOMER_TYPE_UPDATECUSTOMER_TYPE_DELETECUSTOMER_TYPE_ADD_DISCOUNTCUSTOMER_TYPE_UPDATE_DISCOUNTCUSTOMER_TYPE_DELETE_DISCOUNTCUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORMDELIVERY_ORDER_CREATEDELIVERY_ORDER_UPLOADDELIVERY_ORDER_ADD_ITEMDELIVERY_ORDER_ADD_IMAGEDELIVERY_ORDER_ADD_DRIVERDELIVERY_ORDER_FETCHDELIVERY_ORDER_FETCH_DRIVERDELIVERY_ORDER_GETDELIVERY_ORDER_ACTIVE_FOR_DRIVERDELIVERY_ORDER_MARK_ONGOINGDELIVERY_ORDER_DELIVERINGDELIVERY_ORDER_CANCELDELIVERY_ORDER_MARK_COMPLETEDDELIVERY_ORDER_DELIVERY_LOCATIONDELIVERY_ORDER_DELETEDELIVERY_ORDER_DELETE_ITEMDELIVERY_ORDER_DELETE_IMAGEDELIVERY_ORDER_DELETE_DRIVERPRODUCT_CREATEPRODUCT_UPLOADPRODUCT_FETCHPRODUCT_GETPRODUCT_UPDATEPRODUCT_DELETEPRODUCT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORMPRODUCT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORMPRODUCT_OPTION_FOR_CUSTOMER_TYPE_DISCOUNT_FORMPRODUCT_OPTION_FOR_CART_ADD_ITEM_FORMPRODUCT_DISCOUNT_CREATEPRODUCT_DISCOUNT_UPDATEPRODUCT_DISCOUNT_DELETEPRODUCT_RECEIVE_CREATEPRODUCT_RECEIVE_UPLOADPRODUCT_RECEIVE_ADD_ITEMPRODUCT_RECEIVE_ADD_IMAGEPRODUCT_RECEIVE_UPDATEPRODUCT_RECEIVE_CANCELPRODUCT_RECEIVE_MARK_COMPLETEPRODUCT_RECEIVE_FETCHPRODUCT_RECEIVE_GETPRODUCT_RECEIVE_DELETEPRODUCT_RECEIVE_DELETE_ITEMPRODUCT_RECEIVE_DELETE_IMAGEPRODUCT_STOCK_FETCHPRODUCT_STOCK_GETPRODUCT_STOCK_DOWNLOAD_REPORTPRODUCT_STOCK_ADJUSTMENTPRODUCT_UNIT_CREATEPRODUCT_UNIT_GETPRODUCT_UNIT_UPDATEPRODUCT_UNIT_DELETEPRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORMPRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORMROLE_OPTION_FOR_USER_FORMSHOP_ORDER_FETCHSHOP_ORDER_GETSSR_WHATSAPP_LOGINSUPPLIER_CREATESUPPLIER_FETCHSUPPLIER_GETSUPPLIER_UPDATESUPPLIER_DELETESUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORMSUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FILTERSUPPLIER_TYPE_CREATESUPPLIER_TYPE_FETCHSUPPLIER_TYPE_GETSUPPLIER_TYPE_UPDATESUPPLIER_TYPE_DELETESUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORMTIKTOK_PRODUCT_CREATETIKTOK_PRODUCT_UPLOAD_IMAGETIKTOK_PRODUCT_FETCH_BRANDSTIKTOK_PRODUCT_FETCH_CATEGORIESTIKTOK_PRODUCT_GET_RULESTIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTESTIKTOK_PRODUCT_GETTIKTOK_PRODUCT_UPDATETIKTOK_PRODUCT_RECOMMENDED_CATEGORYTIKTOK_PRODUCT_RECOMMENDED_ACTIVATETIKTOK_PRODUCT_RECOMMENDED_DEACTIVATETRANSACTION_CHECKOUT_CARTTRANSACTION_GETUSER_CREATEUSER_FETCHUSER_GETUSER_UPDATEUSER_UPDATE_PASSWORDUSER_UPDATE_ACTIVEUSER_UPDATE_INACTIVEUSER_ADD_ROLEUSER_DELETE_ROLEUSER_OPTION_FOR_CASHIER_SESSION_FILTERUSER_OPTION_FOR_DELIVERY_ORDER_DRIVER_FORMUNIT_CREATEUNIT_FETCHUNIT_GETUNIT_UPDATEUNIT_DELETEUNIT_OPTION_FOR_PRODUCT_UNIT_FORMUNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORMWHATSAPP_IS_LOGGED_INWHATSAPP_LOGOUT"

var _Permission_index = [...]uint16{0, 14, 27, 38, 52, 66, 81, 94, 110, 126, 141, 159, 170, 191, 212, 231, 264, 295, 322, 341, 356, 370, 382, 397, 412, 451, 492, 518, 537, 554, 575, 595, 614, 631, 651, 671, 697, 726, 755, 793, 814, 835, 858, 882, 907, 927, 954, 972, 1004, 1031, 1056, 1077, 1106, 1138, 1159, 1185, 1212, 1240, 1254, 1268, 1281, 1292, 1306, 1320, 1364, 1407, 1453, 1490, 1513, 1536, 1559, 1581, 1603, 1627, 1652, 1674, 1696, 1725, 1746, 1765, 1787, 1814, 1842, 1861, 1878, 1907, 1931, 1950, 1966, 1985, 2004, 2053, 2101, 2126, 2142, 2156, 2174, 2189, 2203, 2215, 2230, 2245, 2285, 2327, 2347, 2366, 2383, 2403, 2423, 2461, 2482, 2509, 2536, 2567, 2591, 2629, 2647, 2668, 2703, 2738, 2775, 2800, 2815, 2826, 2836, 2844, 2855, 2875, 2893, 2913, 2926, 2942, 2980, 3022, 3033, 3043, 3051, 3062, 3073, 3106, 3147, 3168, 3183}

func (i *Permission) Determine(s string) {
	switch s {
	case "BALANCE_CREATE":
		*i = PermissionBalanceCreate
	case "BALANCE_FETCH":
		*i = PermissionBalanceFetch
	case "BALANCE_GET":
		*i = PermissionBalanceGet
	case "BALANCE_UPDATE":
		*i = PermissionBalanceUpdate
	case "BALANCE_DELETE":
		*i = PermissionBalanceDelete
	case "CART_GET_ACTIVE":
		*i = PermissionCartGetActive
	case "CART_ADD_ITEM":
		*i = PermissionCartAddItem
	case "CART_UPDATE_ITEM":
		*i = PermissionCartUpdateItem
	case "CART_DELETE_ITEM":
		*i = PermissionCartDeleteItem
	case "CART_SET_ACTIVE":
		*i = PermissionCartSetActive
	case "CART_SET_IN_ACTIVE":
		*i = PermissionCartSetInActive
	case "CART_DELETE":
		*i = PermissionCartDelete
	case "CASHIER_SESSION_FETCH":
		*i = PermissionCashierSessionFetch
	case "CASHIER_SESSION_START":
		*i = PermissionCashierSessionStart
	case "CASHIER_SESSION_GET":
		*i = PermissionCashierSessionGet
	case "CASHIER_SESSION_FETCH_TRANSACTION":
		*i = PermissionCashierSessionFetchTransaction
	case "CASHIER_SESSION_DOWNLOAD_REPORT":
		*i = PermissionCashierSessionDownloadReport
	case "CASHIER_SESSION_GET_CURRENT":
		*i = PermissionCashierSessionGetCurrent
	case "CASHIER_SESSION_END":
		*i = PermissionCashierSessionEnd
	case "CUSTOMER_CREATE":
		*i = PermissionCustomerCreate
	case "CUSTOMER_FETCH":
		*i = PermissionCustomerFetch
	case "CUSTOMER_GET":
		*i = PermissionCustomerGet
	case "CUSTOMER_UPDATE":
		*i = PermissionCustomerUpdate
	case "CUSTOMER_DELETE":
		*i = PermissionCustomerDelete
	case "CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM":
		*i = PermissionCustomerOptionForDeliveryOrderForm
	case "CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FILTER":
		*i = PermissionCustomerOptionForDeliveryOrderFilter
	case "CUSTOMER_DEBT_UPLOAD_IMAGE":
		*i = PermissionCustomerDebtUploadImage
	case "CUSTOMER_DEBT_FETCH":
		*i = PermissionCustomerDebtFetch
	case "CUSTOMER_DEBT_GET":
		*i = PermissionCustomerDebtGet
	case "CUSTOMER_DEBT_PAYMENT":
		*i = PermissionCustomerDebtPayment
	case "CUSTOMER_TYPE_CREATE":
		*i = PermissionCustomerTypeCreate
	case "CUSTOMER_TYPE_FETCH":
		*i = PermissionCustomerTypeFetch
	case "CUSTOMER_TYPE_GET":
		*i = PermissionCustomerTypeGet
	case "CUSTOMER_TYPE_UPDATE":
		*i = PermissionCustomerTypeUpdate
	case "CUSTOMER_TYPE_DELETE":
		*i = PermissionCustomerTypeDelete
	case "CUSTOMER_TYPE_ADD_DISCOUNT":
		*i = PermissionCustomerTypeAddDiscount
	case "CUSTOMER_TYPE_UPDATE_DISCOUNT":
		*i = PermissionCustomerTypeUpdateDiscount
	case "CUSTOMER_TYPE_DELETE_DISCOUNT":
		*i = PermissionCustomerTypeDeleteDiscount
	case "CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM":
		*i = PermissionCustomerTypeOptionForCustomerForm
	case "DELIVERY_ORDER_CREATE":
		*i = PermissionDeliveryOrderCreate
	case "DELIVERY_ORDER_UPLOAD":
		*i = PermissionDeliveryOrderUpload
	case "DELIVERY_ORDER_ADD_ITEM":
		*i = PermissionDeliveryOrderAddItem
	case "DELIVERY_ORDER_ADD_IMAGE":
		*i = PermissionDeliveryOrderAddImage
	case "DELIVERY_ORDER_ADD_DRIVER":
		*i = PermissionDeliveryOrderAddDriver
	case "DELIVERY_ORDER_FETCH":
		*i = PermissionDeliveryOrderFetch
	case "DELIVERY_ORDER_FETCH_DRIVER":
		*i = PermissionDeliveryOrderFetchDriver
	case "DELIVERY_ORDER_GET":
		*i = PermissionDeliveryOrderGet
	case "DELIVERY_ORDER_ACTIVE_FOR_DRIVER":
		*i = PermissionDeliveryOrderActiveForDriver
	case "DELIVERY_ORDER_MARK_ONGOING":
		*i = PermissionDeliveryOrderMarkOngoing
	case "DELIVERY_ORDER_DELIVERING":
		*i = PermissionDeliveryOrderDelivering
	case "DELIVERY_ORDER_CANCEL":
		*i = PermissionDeliveryOrderCancel
	case "DELIVERY_ORDER_MARK_COMPLETED":
		*i = PermissionDeliveryOrderMarkCompleted
	case "DELIVERY_ORDER_DELIVERY_LOCATION":
		*i = PermissionDeliveryOrderDeliveryLocation
	case "DELIVERY_ORDER_DELETE":
		*i = PermissionDeliveryOrderDelete
	case "DELIVERY_ORDER_DELETE_ITEM":
		*i = PermissionDeliveryOrderDeleteItem
	case "DELIVERY_ORDER_DELETE_IMAGE":
		*i = PermissionDeliveryOrderDeleteImage
	case "DELIVERY_ORDER_DELETE_DRIVER":
		*i = PermissionDeliveryOrderDeleteDriver
	case "PRODUCT_CREATE":
		*i = PermissionProductCreate
	case "PRODUCT_UPLOAD":
		*i = PermissionProductUpload
	case "PRODUCT_FETCH":
		*i = PermissionProductFetch
	case "PRODUCT_GET":
		*i = PermissionProductGet
	case "PRODUCT_UPDATE":
		*i = PermissionProductUpdate
	case "PRODUCT_DELETE":
		*i = PermissionProductDelete
	case "PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM":
		*i = PermissionProductOptionForProductReceiveItemForm
	case "PRODUCT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM":
		*i = PermissionProductOptionForDeliveryOrderItemForm
	case "PRODUCT_OPTION_FOR_CUSTOMER_TYPE_DISCOUNT_FORM":
		*i = PermissionProductOptionForCustomerTypeDiscountForm
	case "PRODUCT_OPTION_FOR_CART_ADD_ITEM_FORM":
		*i = PermissionProductOptionForCartAddItemForm
	case "PRODUCT_DISCOUNT_CREATE":
		*i = PermissionProductDiscountCreate
	case "PRODUCT_DISCOUNT_UPDATE":
		*i = PermissionProductDiscountUpdate
	case "PRODUCT_DISCOUNT_DELETE":
		*i = PermissionProductDiscountDelete
	case "PRODUCT_RECEIVE_CREATE":
		*i = PermissionProductReceiveCreate
	case "PRODUCT_RECEIVE_UPLOAD":
		*i = PermissionProductReceiveUpload
	case "PRODUCT_RECEIVE_ADD_ITEM":
		*i = PermissionProductReceiveAddItem
	case "PRODUCT_RECEIVE_ADD_IMAGE":
		*i = PermissionProductReceiveAddImage
	case "PRODUCT_RECEIVE_UPDATE":
		*i = PermissionProductReceiveUpdate
	case "PRODUCT_RECEIVE_CANCEL":
		*i = PermissionProductReceiveCancel
	case "PRODUCT_RECEIVE_MARK_COMPLETE":
		*i = PermissionProductReceiveMarkComplete
	case "PRODUCT_RECEIVE_FETCH":
		*i = PermissionProductReceiveFetch
	case "PRODUCT_RECEIVE_GET":
		*i = PermissionProductReceiveGet
	case "PRODUCT_RECEIVE_DELETE":
		*i = PermissionProductReceiveDelete
	case "PRODUCT_RECEIVE_DELETE_ITEM":
		*i = PermissionProductReceiveDeleteItem
	case "PRODUCT_RECEIVE_DELETE_IMAGE":
		*i = PermissionProductReceiveDeleteImage
	case "PRODUCT_STOCK_FETCH":
		*i = PermissionProductStockFetch
	case "PRODUCT_STOCK_GET":
		*i = PermissionProductStockGet
	case "PRODUCT_STOCK_DOWNLOAD_REPORT":
		*i = PermissionProductStockDownloadReport
	case "PRODUCT_STOCK_ADJUSTMENT":
		*i = PermissionProductStockAdjustment
	case "PRODUCT_UNIT_CREATE":
		*i = PermissionProductUnitCreate
	case "PRODUCT_UNIT_GET":
		*i = PermissionProductUnitGet
	case "PRODUCT_UNIT_UPDATE":
		*i = PermissionProductUnitUpdate
	case "PRODUCT_UNIT_DELETE":
		*i = PermissionProductUnitDelete
	case "PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM":
		*i = PermissionProductUnitOptionForProductReceiveItemForm
	case "PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM":
		*i = PermissionProductUnitOptionForDeliveryOrderItemForm
	case "ROLE_OPTION_FOR_USER_FORM":
		*i = PermissionRoleOptionForUserForm
	case "SHOP_ORDER_FETCH":
		*i = PermissionShopOrderFetch
	case "SHOP_ORDER_GET":
		*i = PermissionShopOrderGet
	case "SSR_WHATSAPP_LOGIN":
		*i = PermissionSsrWhatsappLogin
	case "SUPPLIER_CREATE":
		*i = PermissionSupplierCreate
	case "SUPPLIER_FETCH":
		*i = PermissionSupplierFetch
	case "SUPPLIER_GET":
		*i = PermissionSupplierGet
	case "SUPPLIER_UPDATE":
		*i = PermissionSupplierUpdate
	case "SUPPLIER_DELETE":
		*i = PermissionSupplierDelete
	case "SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM":
		*i = PermissionSupplierOptionForProductReceiveForm
	case "SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FILTER":
		*i = PermissionSupplierOptionForProductReceiveFilter
	case "SUPPLIER_TYPE_CREATE":
		*i = PermissionSupplierTypeCreate
	case "SUPPLIER_TYPE_FETCH":
		*i = PermissionSupplierTypeFetch
	case "SUPPLIER_TYPE_GET":
		*i = PermissionSupplierTypeGet
	case "SUPPLIER_TYPE_UPDATE":
		*i = PermissionSupplierTypeUpdate
	case "SUPPLIER_TYPE_DELETE":
		*i = PermissionSupplierTypeDelete
	case "SUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORM":
		*i = PermissionSupplierTypeOptionForSupplierForm
	case "TIKTOK_PRODUCT_CREATE":
		*i = PermissionTiktokProductCreate
	case "TIKTOK_PRODUCT_UPLOAD_IMAGE":
		*i = PermissionTiktokProductUploadImage
	case "TIKTOK_PRODUCT_FETCH_BRANDS":
		*i = PermissionTiktokProductFetchBrands
	case "TIKTOK_PRODUCT_FETCH_CATEGORIES":
		*i = PermissionTiktokProductFetchCategories
	case "TIKTOK_PRODUCT_GET_RULES":
		*i = PermissionTiktokProductGetCategoryRules
	case "TIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTES":
		*i = PermissionTiktokProductGetCategoryAttributes
	case "TIKTOK_PRODUCT_GET":
		*i = PermissionTiktokProductGet
	case "TIKTOK_PRODUCT_UPDATE":
		*i = PermissionTiktokProductUpdate
	case "TIKTOK_PRODUCT_RECOMMENDED_CATEGORY":
		*i = PermissionTiktokProductRecommendedCategory
	case "TIKTOK_PRODUCT_RECOMMENDED_ACTIVATE":
		*i = PermissionTiktokProductActivate
	case "TIKTOK_PRODUCT_RECOMMENDED_DEACTIVATE":
		*i = PermissionTiktokProductDeactivate
	case "TRANSACTION_CHECKOUT_CART":
		*i = PermissionTransactionCheckoutCart
	case "TRANSACTION_GET":
		*i = PermissionTransactionGet
	case "USER_CREATE":
		*i = PermissionUserCreate
	case "USER_FETCH":
		*i = PermissionUserFetch
	case "USER_GET":
		*i = PermissionUserGet
	case "USER_UPDATE":
		*i = PermissionUserUpdate
	case "USER_UPDATE_PASSWORD":
		*i = PermissionUserUpdatePassword
	case "USER_UPDATE_ACTIVE":
		*i = PermissionUserUpdateActive
	case "USER_UPDATE_INACTIVE":
		*i = PermissionUserUpdateInActive
	case "USER_ADD_ROLE":
		*i = PermissionUserAddRole
	case "USER_DELETE_ROLE":
		*i = PermissionUserDeleteRole
	case "USER_OPTION_FOR_CASHIER_SESSION_FILTER":
		*i = PermissionUserOptionForCashierSessionFilter
	case "USER_OPTION_FOR_DELIVERY_ORDER_DRIVER_FORM":
		*i = PermissionUserOptionForDeliveryOrderDriverForm
	case "UNIT_CREATE":
		*i = PermissionUnitCreate
	case "UNIT_FETCH":
		*i = PermissionUnitFetch
	case "UNIT_GET":
		*i = PermissionUnitGet
	case "UNIT_UPDATE":
		*i = PermissionUnitUpdate
	case "UNIT_DELETE":
		*i = PermissionUnitDelete
	case "UNIT_OPTION_FOR_PRODUCT_UNIT_FORM":
		*i = PermissionUnitOptionForProductUnitForm
	case "UNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORM":
		*i = PermissionUnitOptionForProductUnitToUnitForm
	case "WHATSAPP_IS_LOGGED_IN":
		*i = PermissionWhatsappIsLoggedIn
	case "WHATSAPP_LOGOUT":
		*i = PermissionWhatsappLogout
	default:
		*i = 0
	}
}

func (i Permission) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i Permission) GetValidValuesString() string {
	return _Permission_nameReadable
}

func (i Permission) String() string {
	i -= 1
	if i < 0 || i >= Permission(len(_Permission_index)-1) {
		return "Permission(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _Permission_name[_Permission_index[i]:_Permission_index[i+1]]
}

func (i Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Permission) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *Permission) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *Permission) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i Permission) Value() (driver.Value, error) {
	return i.String(), nil
}

func PermissionP(v Permission) *Permission {
	return &v
}

func ListPermission() []Permission {
	return []Permission{
		PermissionBalanceCreate,
		PermissionBalanceFetch,
		PermissionBalanceGet,
		PermissionBalanceUpdate,
		PermissionBalanceDelete,
		PermissionCartGetActive,
		PermissionCartAddItem,
		PermissionCartUpdateItem,
		PermissionCartDeleteItem,
		PermissionCartSetActive,
		PermissionCartSetInActive,
		PermissionCartDelete,
		PermissionCashierSessionFetch,
		PermissionCashierSessionStart,
		PermissionCashierSessionGet,
		PermissionCashierSessionFetchTransaction,
		PermissionCashierSessionDownloadReport,
		PermissionCashierSessionGetCurrent,
		PermissionCashierSessionEnd,
		PermissionCustomerCreate,
		PermissionCustomerFetch,
		PermissionCustomerGet,
		PermissionCustomerUpdate,
		PermissionCustomerDelete,
		PermissionCustomerOptionForDeliveryOrderForm,
		PermissionCustomerOptionForDeliveryOrderFilter,
		PermissionCustomerDebtUploadImage,
		PermissionCustomerDebtFetch,
		PermissionCustomerDebtGet,
		PermissionCustomerDebtPayment,
		PermissionCustomerTypeCreate,
		PermissionCustomerTypeFetch,
		PermissionCustomerTypeGet,
		PermissionCustomerTypeUpdate,
		PermissionCustomerTypeDelete,
		PermissionCustomerTypeAddDiscount,
		PermissionCustomerTypeUpdateDiscount,
		PermissionCustomerTypeDeleteDiscount,
		PermissionCustomerTypeOptionForCustomerForm,
		PermissionDeliveryOrderCreate,
		PermissionDeliveryOrderUpload,
		PermissionDeliveryOrderAddItem,
		PermissionDeliveryOrderAddImage,
		PermissionDeliveryOrderAddDriver,
		PermissionDeliveryOrderFetch,
		PermissionDeliveryOrderFetchDriver,
		PermissionDeliveryOrderGet,
		PermissionDeliveryOrderActiveForDriver,
		PermissionDeliveryOrderMarkOngoing,
		PermissionDeliveryOrderDelivering,
		PermissionDeliveryOrderCancel,
		PermissionDeliveryOrderMarkCompleted,
		PermissionDeliveryOrderDeliveryLocation,
		PermissionDeliveryOrderDelete,
		PermissionDeliveryOrderDeleteItem,
		PermissionDeliveryOrderDeleteImage,
		PermissionDeliveryOrderDeleteDriver,
		PermissionProductCreate,
		PermissionProductUpload,
		PermissionProductFetch,
		PermissionProductGet,
		PermissionProductUpdate,
		PermissionProductDelete,
		PermissionProductOptionForProductReceiveItemForm,
		PermissionProductOptionForDeliveryOrderItemForm,
		PermissionProductOptionForCustomerTypeDiscountForm,
		PermissionProductOptionForCartAddItemForm,
		PermissionProductDiscountCreate,
		PermissionProductDiscountUpdate,
		PermissionProductDiscountDelete,
		PermissionProductReceiveCreate,
		PermissionProductReceiveUpload,
		PermissionProductReceiveAddItem,
		PermissionProductReceiveAddImage,
		PermissionProductReceiveUpdate,
		PermissionProductReceiveCancel,
		PermissionProductReceiveMarkComplete,
		PermissionProductReceiveFetch,
		PermissionProductReceiveGet,
		PermissionProductReceiveDelete,
		PermissionProductReceiveDeleteItem,
		PermissionProductReceiveDeleteImage,
		PermissionProductStockFetch,
		PermissionProductStockGet,
		PermissionProductStockDownloadReport,
		PermissionProductStockAdjustment,
		PermissionProductUnitCreate,
		PermissionProductUnitGet,
		PermissionProductUnitUpdate,
		PermissionProductUnitDelete,
		PermissionProductUnitOptionForProductReceiveItemForm,
		PermissionProductUnitOptionForDeliveryOrderItemForm,
		PermissionRoleOptionForUserForm,
		PermissionShopOrderFetch,
		PermissionShopOrderGet,
		PermissionSsrWhatsappLogin,
		PermissionSupplierCreate,
		PermissionSupplierFetch,
		PermissionSupplierGet,
		PermissionSupplierUpdate,
		PermissionSupplierDelete,
		PermissionSupplierOptionForProductReceiveForm,
		PermissionSupplierOptionForProductReceiveFilter,
		PermissionSupplierTypeCreate,
		PermissionSupplierTypeFetch,
		PermissionSupplierTypeGet,
		PermissionSupplierTypeUpdate,
		PermissionSupplierTypeDelete,
		PermissionSupplierTypeOptionForSupplierForm,
		PermissionTiktokProductCreate,
		PermissionTiktokProductUploadImage,
		PermissionTiktokProductFetchBrands,
		PermissionTiktokProductFetchCategories,
		PermissionTiktokProductGetCategoryRules,
		PermissionTiktokProductGetCategoryAttributes,
		PermissionTiktokProductGet,
		PermissionTiktokProductUpdate,
		PermissionTiktokProductRecommendedCategory,
		PermissionTiktokProductActivate,
		PermissionTiktokProductDeactivate,
		PermissionTransactionCheckoutCart,
		PermissionTransactionGet,
		PermissionUserCreate,
		PermissionUserFetch,
		PermissionUserGet,
		PermissionUserUpdate,
		PermissionUserUpdatePassword,
		PermissionUserUpdateActive,
		PermissionUserUpdateInActive,
		PermissionUserAddRole,
		PermissionUserDeleteRole,
		PermissionUserOptionForCashierSessionFilter,
		PermissionUserOptionForDeliveryOrderDriverForm,
		PermissionUnitCreate,
		PermissionUnitFetch,
		PermissionUnitGet,
		PermissionUnitUpdate,
		PermissionUnitDelete,
		PermissionUnitOptionForProductUnitForm,
		PermissionUnitOptionForProductUnitToUnitForm,
		PermissionWhatsappIsLoggedIn,
		PermissionWhatsappLogout,
	}
}

func ListPermissionString() []string {
	return []string{
		PermissionBalanceCreate.String(),
		PermissionBalanceFetch.String(),
		PermissionBalanceGet.String(),
		PermissionBalanceUpdate.String(),
		PermissionBalanceDelete.String(),
		PermissionCartGetActive.String(),
		PermissionCartAddItem.String(),
		PermissionCartUpdateItem.String(),
		PermissionCartDeleteItem.String(),
		PermissionCartSetActive.String(),
		PermissionCartSetInActive.String(),
		PermissionCartDelete.String(),
		PermissionCashierSessionFetch.String(),
		PermissionCashierSessionStart.String(),
		PermissionCashierSessionGet.String(),
		PermissionCashierSessionFetchTransaction.String(),
		PermissionCashierSessionDownloadReport.String(),
		PermissionCashierSessionGetCurrent.String(),
		PermissionCashierSessionEnd.String(),
		PermissionCustomerCreate.String(),
		PermissionCustomerFetch.String(),
		PermissionCustomerGet.String(),
		PermissionCustomerUpdate.String(),
		PermissionCustomerDelete.String(),
		PermissionCustomerOptionForDeliveryOrderForm.String(),
		PermissionCustomerOptionForDeliveryOrderFilter.String(),
		PermissionCustomerDebtUploadImage.String(),
		PermissionCustomerDebtFetch.String(),
		PermissionCustomerDebtGet.String(),
		PermissionCustomerDebtPayment.String(),
		PermissionCustomerTypeCreate.String(),
		PermissionCustomerTypeFetch.String(),
		PermissionCustomerTypeGet.String(),
		PermissionCustomerTypeUpdate.String(),
		PermissionCustomerTypeDelete.String(),
		PermissionCustomerTypeAddDiscount.String(),
		PermissionCustomerTypeUpdateDiscount.String(),
		PermissionCustomerTypeDeleteDiscount.String(),
		PermissionCustomerTypeOptionForCustomerForm.String(),
		PermissionDeliveryOrderCreate.String(),
		PermissionDeliveryOrderUpload.String(),
		PermissionDeliveryOrderAddItem.String(),
		PermissionDeliveryOrderAddImage.String(),
		PermissionDeliveryOrderAddDriver.String(),
		PermissionDeliveryOrderFetch.String(),
		PermissionDeliveryOrderFetchDriver.String(),
		PermissionDeliveryOrderGet.String(),
		PermissionDeliveryOrderActiveForDriver.String(),
		PermissionDeliveryOrderMarkOngoing.String(),
		PermissionDeliveryOrderDelivering.String(),
		PermissionDeliveryOrderCancel.String(),
		PermissionDeliveryOrderMarkCompleted.String(),
		PermissionDeliveryOrderDeliveryLocation.String(),
		PermissionDeliveryOrderDelete.String(),
		PermissionDeliveryOrderDeleteItem.String(),
		PermissionDeliveryOrderDeleteImage.String(),
		PermissionDeliveryOrderDeleteDriver.String(),
		PermissionProductCreate.String(),
		PermissionProductUpload.String(),
		PermissionProductFetch.String(),
		PermissionProductGet.String(),
		PermissionProductUpdate.String(),
		PermissionProductDelete.String(),
		PermissionProductOptionForProductReceiveItemForm.String(),
		PermissionProductOptionForDeliveryOrderItemForm.String(),
		PermissionProductOptionForCustomerTypeDiscountForm.String(),
		PermissionProductOptionForCartAddItemForm.String(),
		PermissionProductDiscountCreate.String(),
		PermissionProductDiscountUpdate.String(),
		PermissionProductDiscountDelete.String(),
		PermissionProductReceiveCreate.String(),
		PermissionProductReceiveUpload.String(),
		PermissionProductReceiveAddItem.String(),
		PermissionProductReceiveAddImage.String(),
		PermissionProductReceiveUpdate.String(),
		PermissionProductReceiveCancel.String(),
		PermissionProductReceiveMarkComplete.String(),
		PermissionProductReceiveFetch.String(),
		PermissionProductReceiveGet.String(),
		PermissionProductReceiveDelete.String(),
		PermissionProductReceiveDeleteItem.String(),
		PermissionProductReceiveDeleteImage.String(),
		PermissionProductStockFetch.String(),
		PermissionProductStockGet.String(),
		PermissionProductStockDownloadReport.String(),
		PermissionProductStockAdjustment.String(),
		PermissionProductUnitCreate.String(),
		PermissionProductUnitGet.String(),
		PermissionProductUnitUpdate.String(),
		PermissionProductUnitDelete.String(),
		PermissionProductUnitOptionForProductReceiveItemForm.String(),
		PermissionProductUnitOptionForDeliveryOrderItemForm.String(),
		PermissionRoleOptionForUserForm.String(),
		PermissionShopOrderFetch.String(),
		PermissionShopOrderGet.String(),
		PermissionSsrWhatsappLogin.String(),
		PermissionSupplierCreate.String(),
		PermissionSupplierFetch.String(),
		PermissionSupplierGet.String(),
		PermissionSupplierUpdate.String(),
		PermissionSupplierDelete.String(),
		PermissionSupplierOptionForProductReceiveForm.String(),
		PermissionSupplierOptionForProductReceiveFilter.String(),
		PermissionSupplierTypeCreate.String(),
		PermissionSupplierTypeFetch.String(),
		PermissionSupplierTypeGet.String(),
		PermissionSupplierTypeUpdate.String(),
		PermissionSupplierTypeDelete.String(),
		PermissionSupplierTypeOptionForSupplierForm.String(),
		PermissionTiktokProductCreate.String(),
		PermissionTiktokProductUploadImage.String(),
		PermissionTiktokProductFetchBrands.String(),
		PermissionTiktokProductFetchCategories.String(),
		PermissionTiktokProductGetCategoryRules.String(),
		PermissionTiktokProductGetCategoryAttributes.String(),
		PermissionTiktokProductGet.String(),
		PermissionTiktokProductUpdate.String(),
		PermissionTiktokProductRecommendedCategory.String(),
		PermissionTiktokProductActivate.String(),
		PermissionTiktokProductDeactivate.String(),
		PermissionTransactionCheckoutCart.String(),
		PermissionTransactionGet.String(),
		PermissionUserCreate.String(),
		PermissionUserFetch.String(),
		PermissionUserGet.String(),
		PermissionUserUpdate.String(),
		PermissionUserUpdatePassword.String(),
		PermissionUserUpdateActive.String(),
		PermissionUserUpdateInActive.String(),
		PermissionUserAddRole.String(),
		PermissionUserDeleteRole.String(),
		PermissionUserOptionForCashierSessionFilter.String(),
		PermissionUserOptionForDeliveryOrderDriverForm.String(),
		PermissionUnitCreate.String(),
		PermissionUnitFetch.String(),
		PermissionUnitGet.String(),
		PermissionUnitUpdate.String(),
		PermissionUnitDelete.String(),
		PermissionUnitOptionForProductUnitForm.String(),
		PermissionUnitOptionForProductUnitToUnitForm.String(),
		PermissionWhatsappIsLoggedIn.String(),
		PermissionWhatsappLogout.String(),
	}
}
