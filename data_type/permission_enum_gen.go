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
	_ = x[PermissionCashierSessionDownloadReport-16]
	_ = x[PermissionCashierSessionGetCurrent-17]
	_ = x[PermissionCashierSessionEnd-18]
	_ = x[PermissionCustomerCreate-19]
	_ = x[PermissionCustomerFetch-20]
	_ = x[PermissionCustomerGet-21]
	_ = x[PermissionCustomerUpdate-22]
	_ = x[PermissionCustomerDelete-23]
	_ = x[PermissionCustomerOptionForDeliveryOrderForm-24]
	_ = x[PermissionCustomerDebtUploadImage-25]
	_ = x[PermissionCustomerDebtFetch-26]
	_ = x[PermissionCustomerDebtGet-27]
	_ = x[PermissionCustomerDebtPayment-28]
	_ = x[PermissionCustomerTypeCreate-29]
	_ = x[PermissionCustomerTypeFetch-30]
	_ = x[PermissionCustomerTypeGet-31]
	_ = x[PermissionCustomerTypeUpdate-32]
	_ = x[PermissionCustomerTypeDelete-33]
	_ = x[PermissionCustomerTypeAddDiscount-34]
	_ = x[PermissionCustomerTypeUpdateDiscount-35]
	_ = x[PermissionCustomerTypeDeleteDiscount-36]
	_ = x[PermissionCustomerTypeOptionForCustomerForm-37]
	_ = x[PermissionDeliveryOrderCreate-38]
	_ = x[PermissionDeliveryOrderUpload-39]
	_ = x[PermissionDeliveryOrderAddItem-40]
	_ = x[PermissionDeliveryOrderAddImage-41]
	_ = x[PermissionDeliveryOrderAddDriver-42]
	_ = x[PermissionDeliveryOrderFetch-43]
	_ = x[PermissionDeliveryOrderGet-44]
	_ = x[PermissionDeliveryOrderMarkOngoing-45]
	_ = x[PermissionDeliveryOrderCancel-46]
	_ = x[PermissionDeliveryOrderMarkCompleted-47]
	_ = x[PermissionDeliveryOrderDeliveryLocation-48]
	_ = x[PermissionDeliveryOrderDelete-49]
	_ = x[PermissionDeliveryOrderDeleteItem-50]
	_ = x[PermissionDeliveryOrderDeleteImage-51]
	_ = x[PermissionDeliveryOrderDeleteDriver-52]
	_ = x[PermissionProductCreate-53]
	_ = x[PermissionProductUpload-54]
	_ = x[PermissionProductFetch-55]
	_ = x[PermissionProductGet-56]
	_ = x[PermissionProductUpdate-57]
	_ = x[PermissionProductDelete-58]
	_ = x[PermissionProductOptionForProductReceiveForm-59]
	_ = x[PermissionProductOptionForDeliveryOrderForm-60]
	_ = x[PermissionProductOptionForCustomerTypeForm-61]
	_ = x[PermissionProductOptionForCartAddItemForm-62]
	_ = x[PermissionProductDiscountCreate-63]
	_ = x[PermissionProductDiscountUpdate-64]
	_ = x[PermissionProductDiscountDelete-65]
	_ = x[PermissionProductReceiveCreate-66]
	_ = x[PermissionProductReceiveUpload-67]
	_ = x[PermissionProductReceiveAddItem-68]
	_ = x[PermissionProductReceiveAddImage-69]
	_ = x[PermissionProductReceiveCancel-70]
	_ = x[PermissionProductReceiveMarkComplete-71]
	_ = x[PermissionProductReceiveFetch-72]
	_ = x[PermissionProductReceiveGet-73]
	_ = x[PermissionProductReceiveDelete-74]
	_ = x[PermissionProductReceiveDeleteItem-75]
	_ = x[PermissionProductReceiveDeleteImage-76]
	_ = x[PermissionProductStockFetch-77]
	_ = x[PermissionProductStockGet-78]
	_ = x[PermissionProductStockDownloadReport-79]
	_ = x[PermissionProductStockAdjustment-80]
	_ = x[PermissionProductUnitCreate-81]
	_ = x[PermissionProductUnitGet-82]
	_ = x[PermissionProductUnitUpdate-83]
	_ = x[PermissionProductUnitDelete-84]
	_ = x[PermissionProductUnitOptionForProductReceiveForm-85]
	_ = x[PermissionProductUnitOptionForDeliveryOrderForm-86]
	_ = x[PermissionRoleOptionForUserForm-87]
	_ = x[PermissionShopOrderFetch-88]
	_ = x[PermissionShopOrderGet-89]
	_ = x[PermissionSupplierCreate-90]
	_ = x[PermissionSupplierFetch-91]
	_ = x[PermissionSupplierGet-92]
	_ = x[PermissionSupplierUpdate-93]
	_ = x[PermissionSupplierDelete-94]
	_ = x[PermissionSupplierOptionForProductReceiveForm-95]
	_ = x[PermissionSupplierTypeCreate-96]
	_ = x[PermissionSupplierTypeFetch-97]
	_ = x[PermissionSupplierTypeGet-98]
	_ = x[PermissionSupplierTypeUpdate-99]
	_ = x[PermissionSupplierTypeDelete-100]
	_ = x[PermissionSupplierTypeOptionForSupplierForm-101]
	_ = x[PermissionTiktokProductCreate-102]
	_ = x[PermissionTiktokProductUploadImage-103]
	_ = x[PermissionTiktokProductFetchBrands-104]
	_ = x[PermissionTiktokProductFetchCategories-105]
	_ = x[PermissionTiktokProductGetCategoryRules-106]
	_ = x[PermissionTiktokProductGetCategoryAttributes-107]
	_ = x[PermissionTiktokProductGet-108]
	_ = x[PermissionTiktokProductUpdate-109]
	_ = x[PermissionTiktokProductRecommendedCategory-110]
	_ = x[PermissionTiktokProductActivate-111]
	_ = x[PermissionTiktokProductDeactivate-112]
	_ = x[PermissionTransactionCheckoutCart-113]
	_ = x[PermissionUserCreate-114]
	_ = x[PermissionUserFetch-115]
	_ = x[PermissionUserGet-116]
	_ = x[PermissionUserUpdate-117]
	_ = x[PermissionUserUpdatePassword-118]
	_ = x[PermissionUserUpdateActive-119]
	_ = x[PermissionUserUpdateInActive-120]
	_ = x[PermissionUserAddRole-121]
	_ = x[PermissionUserDeleteRole-122]
	_ = x[PermissionUserOptionForCashierSessionFilter-123]
	_ = x[PermissionUnitCreate-124]
	_ = x[PermissionUnitFetch-125]
	_ = x[PermissionUnitGet-126]
	_ = x[PermissionUnitUpdate-127]
	_ = x[PermissionUnitDelete-128]
	_ = x[PermissionUnitOptionForProductUnitForm-129]
	_ = x[PermissionUnitOptionForProductUnitToUnitForm-130]
}

const _Permission_nameReadable = "BALANCE_CREATE, BALANCE_FETCH, BALANCE_GET, BALANCE_UPDATE, BALANCE_DELETE, CART_GET_ACTIVE, CART_ADD_ITEM, CART_UPDATE_ITEM, CART_DELETE_ITEM, CART_SET_ACTIVE, CART_SET_IN_ACTIVE, CART_DELETE, CASHIER_SESSION_FETCH, CASHIER_SESSION_START, CASHIER_SESSION_GET, CASHIER_SESSION_DOWNLOAD_REPORT, CASHIER_SESSION_GET_CURRENT, CASHIER_SESSION_END, CUSTOMER_CREATE, CUSTOMER_FETCH, CUSTOMER_GET, CUSTOMER_UPDATE, CUSTOMER_DELETE, CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM, CUSTOMER_DEBT_UPLOAD_IMAGE, CUSTOMER_DEBT_FETCH, CUSTOMER_DEBT_GET, CUSTOMER_DEBT_PAYMENT, CUSTOMER_TYPE_CREATE, CUSTOMER_TYPE_FETCH, CUSTOMER_TYPE_GET, CUSTOMER_TYPE_UPDATE, CUSTOMER_TYPE_DELETE, CUSTOMER_TYPE_ADD_DISCOUNT, CUSTOMER_TYPE_UPDATE_DISCOUNT, CUSTOMER_TYPE_DELETE_DISCOUNT, CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM, DELIVERY_ORDER_CREATE, DELIVERY_ORDER_UPLOAD, DELIVERY_ORDER_ADD_ITEM, DELIVERY_ORDER_ADD_IMAGE, DELIVERY_ORDER_ADD_DRIVER, DELIVERY_ORDER_FETCH, DELIVERY_ORDER_GET, DELIVERY_ORDER_MARK_ONGOING, DELIVERY_ORDER_CANCEL, DELIVERY_ORDER_MARK_COMPLETED, DELIVERY_ORDER_DELIVERY_LOCATION, DELIVERY_ORDER_DELETE, DELIVERY_ORDER_DELETE_ITEM, DELIVERY_ORDER_DELETE_IMAGE, DELIVERY_ORDER_DELETE_DRIVER, PRODUCT_CREATE, PRODUCT_UPLOAD, PRODUCT_FETCH, PRODUCT_GET, PRODUCT_UPDATE, PRODUCT_DELETE, PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORM, PRODUCT_OPTION_FOR_DELIVERY_ORDER_FORM, PRODUCT_OPTION_FOR_CUSTOMER_TYPE_FORM, PRODUCT_OPTION_FOR_CART_ADD_ITEM_FORM, PRODUCT_DISCOUNT_CREATE, PRODUCT_DISCOUNT_UPDATE, PRODUCT_DISCOUNT_DELETE, PRODUCT_RECEIVE_CREATE, PRODUCT_RECEIVE_UPLOAD, PRODUCT_RECEIVE_ADD_ITEM, PRODUCT_RECEIVE_ADD_IMAGE, PRODUCT_RECEIVE_CANCEL, PRODUCT_RECEIVE_MARK_COMPLETE, PRODUCT_RECEIVE_FETCH, PRODUCT_RECEIVE_GET, PRODUCT_RECEIVE_DELETE, PRODUCT_RECEIVE_DELETE_ITEM, PRODUCT_RECEIVE_DELETE_IMAGE, PRODUCT_STOCK_FETCH, PRODUCT_STOCK_GET, PRODUCT_STOCK_DOWNLOAD_REPORT, PRODUCT_STOCK_ADJUSTMENT, PRODUCT_UNIT_CREATE, PRODUCT_UNIT_GET, PRODUCT_UNIT_UPDATE, PRODUCT_UNIT_DELETE, PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORM, PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORM, ROLE_OPTION_FOR_USER_FORM, SHOP_ORDER_FETCH, SHOP_ORDER_GET, SUPPLIER_CREATE, SUPPLIER_FETCH, SUPPLIER_GET, SUPPLIER_UPDATE, SUPPLIER_DELETE, SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM, SUPPLIER_TYPE_CREATE, SUPPLIER_TYPE_FETCH, SUPPLIER_TYPE_GET, SUPPLIER_TYPE_UPDATE, SUPPLIER_TYPE_DELETE, SUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORM, TIKTOK_PRODUCT_CREATE, TIKTOK_PRODUCT_UPLOAD_IMAGE, TIKTOK_PRODUCT_FETCH_BRANDS, TIKTOK_PRODUCT_FETCH_CATEGORIES, TIKTOK_PRODUCT_GET_RULES, TIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTES, TIKTOK_PRODUCT_GET, TIKTOK_PRODUCT_UPDATE, TIKTOK_PRODUCT_RECOMMENDED_CATEGORY, TIKTOK_PRODUCT_RECOMMENDED_ACTIVATE, TIKTOK_PRODUCT_RECOMMENDED_DEACTIVATE, TRANSACTION_CHECKOUT_CART, USER_CREATE, USER_FETCH, USER_GET, USER_UPDATE, USER_UPDATE_PASSWORD, USER_UPDATE_ACTIVE, USER_UPDATE_INACTIVE, USER_ADD_ROLE, USER_DELETE_ROLE, USER_OPTION_FOR_CASHIER_SESSION_FILTER, UNIT_CREATE, UNIT_FETCH, UNIT_GET, UNIT_UPDATE, UNIT_DELETE, UNIT_OPTION_FOR_PRODUCT_UNIT_FORM, UNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORM"

const _Permission_name = "BALANCE_CREATEBALANCE_FETCHBALANCE_GETBALANCE_UPDATEBALANCE_DELETECART_GET_ACTIVECART_ADD_ITEMCART_UPDATE_ITEMCART_DELETE_ITEMCART_SET_ACTIVECART_SET_IN_ACTIVECART_DELETECASHIER_SESSION_FETCHCASHIER_SESSION_STARTCASHIER_SESSION_GETCASHIER_SESSION_DOWNLOAD_REPORTCASHIER_SESSION_GET_CURRENTCASHIER_SESSION_ENDCUSTOMER_CREATECUSTOMER_FETCHCUSTOMER_GETCUSTOMER_UPDATECUSTOMER_DELETECUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORMCUSTOMER_DEBT_UPLOAD_IMAGECUSTOMER_DEBT_FETCHCUSTOMER_DEBT_GETCUSTOMER_DEBT_PAYMENTCUSTOMER_TYPE_CREATECUSTOMER_TYPE_FETCHCUSTOMER_TYPE_GETCUSTOMER_TYPE_UPDATECUSTOMER_TYPE_DELETECUSTOMER_TYPE_ADD_DISCOUNTCUSTOMER_TYPE_UPDATE_DISCOUNTCUSTOMER_TYPE_DELETE_DISCOUNTCUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORMDELIVERY_ORDER_CREATEDELIVERY_ORDER_UPLOADDELIVERY_ORDER_ADD_ITEMDELIVERY_ORDER_ADD_IMAGEDELIVERY_ORDER_ADD_DRIVERDELIVERY_ORDER_FETCHDELIVERY_ORDER_GETDELIVERY_ORDER_MARK_ONGOINGDELIVERY_ORDER_CANCELDELIVERY_ORDER_MARK_COMPLETEDDELIVERY_ORDER_DELIVERY_LOCATIONDELIVERY_ORDER_DELETEDELIVERY_ORDER_DELETE_ITEMDELIVERY_ORDER_DELETE_IMAGEDELIVERY_ORDER_DELETE_DRIVERPRODUCT_CREATEPRODUCT_UPLOADPRODUCT_FETCHPRODUCT_GETPRODUCT_UPDATEPRODUCT_DELETEPRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORMPRODUCT_OPTION_FOR_DELIVERY_ORDER_FORMPRODUCT_OPTION_FOR_CUSTOMER_TYPE_FORMPRODUCT_OPTION_FOR_CART_ADD_ITEM_FORMPRODUCT_DISCOUNT_CREATEPRODUCT_DISCOUNT_UPDATEPRODUCT_DISCOUNT_DELETEPRODUCT_RECEIVE_CREATEPRODUCT_RECEIVE_UPLOADPRODUCT_RECEIVE_ADD_ITEMPRODUCT_RECEIVE_ADD_IMAGEPRODUCT_RECEIVE_CANCELPRODUCT_RECEIVE_MARK_COMPLETEPRODUCT_RECEIVE_FETCHPRODUCT_RECEIVE_GETPRODUCT_RECEIVE_DELETEPRODUCT_RECEIVE_DELETE_ITEMPRODUCT_RECEIVE_DELETE_IMAGEPRODUCT_STOCK_FETCHPRODUCT_STOCK_GETPRODUCT_STOCK_DOWNLOAD_REPORTPRODUCT_STOCK_ADJUSTMENTPRODUCT_UNIT_CREATEPRODUCT_UNIT_GETPRODUCT_UNIT_UPDATEPRODUCT_UNIT_DELETEPRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORMPRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORMROLE_OPTION_FOR_USER_FORMSHOP_ORDER_FETCHSHOP_ORDER_GETSUPPLIER_CREATESUPPLIER_FETCHSUPPLIER_GETSUPPLIER_UPDATESUPPLIER_DELETESUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORMSUPPLIER_TYPE_CREATESUPPLIER_TYPE_FETCHSUPPLIER_TYPE_GETSUPPLIER_TYPE_UPDATESUPPLIER_TYPE_DELETESUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORMTIKTOK_PRODUCT_CREATETIKTOK_PRODUCT_UPLOAD_IMAGETIKTOK_PRODUCT_FETCH_BRANDSTIKTOK_PRODUCT_FETCH_CATEGORIESTIKTOK_PRODUCT_GET_RULESTIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTESTIKTOK_PRODUCT_GETTIKTOK_PRODUCT_UPDATETIKTOK_PRODUCT_RECOMMENDED_CATEGORYTIKTOK_PRODUCT_RECOMMENDED_ACTIVATETIKTOK_PRODUCT_RECOMMENDED_DEACTIVATETRANSACTION_CHECKOUT_CARTUSER_CREATEUSER_FETCHUSER_GETUSER_UPDATEUSER_UPDATE_PASSWORDUSER_UPDATE_ACTIVEUSER_UPDATE_INACTIVEUSER_ADD_ROLEUSER_DELETE_ROLEUSER_OPTION_FOR_CASHIER_SESSION_FILTERUNIT_CREATEUNIT_FETCHUNIT_GETUNIT_UPDATEUNIT_DELETEUNIT_OPTION_FOR_PRODUCT_UNIT_FORMUNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORM"

var _Permission_index = [...]uint16{0, 14, 27, 38, 52, 66, 81, 94, 110, 126, 141, 159, 170, 191, 212, 231, 262, 289, 308, 323, 337, 349, 364, 379, 418, 444, 463, 480, 501, 521, 540, 557, 577, 597, 623, 652, 681, 719, 740, 761, 784, 808, 833, 853, 871, 898, 919, 948, 980, 1001, 1027, 1054, 1082, 1096, 1110, 1123, 1134, 1148, 1162, 1201, 1239, 1276, 1313, 1336, 1359, 1382, 1404, 1426, 1450, 1475, 1497, 1526, 1547, 1566, 1588, 1615, 1643, 1662, 1679, 1708, 1732, 1751, 1767, 1786, 1805, 1849, 1892, 1917, 1933, 1947, 1962, 1976, 1988, 2003, 2018, 2058, 2078, 2097, 2114, 2134, 2154, 2192, 2213, 2240, 2267, 2298, 2322, 2360, 2378, 2399, 2434, 2469, 2506, 2531, 2542, 2552, 2560, 2571, 2591, 2609, 2629, 2642, 2658, 2696, 2707, 2717, 2725, 2736, 2747, 2780, 2821}

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
	case "DELIVERY_ORDER_GET":
		*i = PermissionDeliveryOrderGet
	case "DELIVERY_ORDER_MARK_ONGOING":
		*i = PermissionDeliveryOrderMarkOngoing
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
	case "PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORM":
		*i = PermissionProductOptionForProductReceiveForm
	case "PRODUCT_OPTION_FOR_DELIVERY_ORDER_FORM":
		*i = PermissionProductOptionForDeliveryOrderForm
	case "PRODUCT_OPTION_FOR_CUSTOMER_TYPE_FORM":
		*i = PermissionProductOptionForCustomerTypeForm
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
	case "PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORM":
		*i = PermissionProductUnitOptionForProductReceiveForm
	case "PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORM":
		*i = PermissionProductUnitOptionForDeliveryOrderForm
	case "ROLE_OPTION_FOR_USER_FORM":
		*i = PermissionRoleOptionForUserForm
	case "SHOP_ORDER_FETCH":
		*i = PermissionShopOrderFetch
	case "SHOP_ORDER_GET":
		*i = PermissionShopOrderGet
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
		PermissionCashierSessionDownloadReport,
		PermissionCashierSessionGetCurrent,
		PermissionCashierSessionEnd,
		PermissionCustomerCreate,
		PermissionCustomerFetch,
		PermissionCustomerGet,
		PermissionCustomerUpdate,
		PermissionCustomerDelete,
		PermissionCustomerOptionForDeliveryOrderForm,
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
		PermissionDeliveryOrderGet,
		PermissionDeliveryOrderMarkOngoing,
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
		PermissionProductOptionForProductReceiveForm,
		PermissionProductOptionForDeliveryOrderForm,
		PermissionProductOptionForCustomerTypeForm,
		PermissionProductOptionForCartAddItemForm,
		PermissionProductDiscountCreate,
		PermissionProductDiscountUpdate,
		PermissionProductDiscountDelete,
		PermissionProductReceiveCreate,
		PermissionProductReceiveUpload,
		PermissionProductReceiveAddItem,
		PermissionProductReceiveAddImage,
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
		PermissionProductUnitOptionForProductReceiveForm,
		PermissionProductUnitOptionForDeliveryOrderForm,
		PermissionRoleOptionForUserForm,
		PermissionShopOrderFetch,
		PermissionShopOrderGet,
		PermissionSupplierCreate,
		PermissionSupplierFetch,
		PermissionSupplierGet,
		PermissionSupplierUpdate,
		PermissionSupplierDelete,
		PermissionSupplierOptionForProductReceiveForm,
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
		PermissionUnitCreate,
		PermissionUnitFetch,
		PermissionUnitGet,
		PermissionUnitUpdate,
		PermissionUnitDelete,
		PermissionUnitOptionForProductUnitForm,
		PermissionUnitOptionForProductUnitToUnitForm,
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
		PermissionCashierSessionDownloadReport.String(),
		PermissionCashierSessionGetCurrent.String(),
		PermissionCashierSessionEnd.String(),
		PermissionCustomerCreate.String(),
		PermissionCustomerFetch.String(),
		PermissionCustomerGet.String(),
		PermissionCustomerUpdate.String(),
		PermissionCustomerDelete.String(),
		PermissionCustomerOptionForDeliveryOrderForm.String(),
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
		PermissionDeliveryOrderGet.String(),
		PermissionDeliveryOrderMarkOngoing.String(),
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
		PermissionProductOptionForProductReceiveForm.String(),
		PermissionProductOptionForDeliveryOrderForm.String(),
		PermissionProductOptionForCustomerTypeForm.String(),
		PermissionProductOptionForCartAddItemForm.String(),
		PermissionProductDiscountCreate.String(),
		PermissionProductDiscountUpdate.String(),
		PermissionProductDiscountDelete.String(),
		PermissionProductReceiveCreate.String(),
		PermissionProductReceiveUpload.String(),
		PermissionProductReceiveAddItem.String(),
		PermissionProductReceiveAddImage.String(),
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
		PermissionProductUnitOptionForProductReceiveForm.String(),
		PermissionProductUnitOptionForDeliveryOrderForm.String(),
		PermissionRoleOptionForUserForm.String(),
		PermissionShopOrderFetch.String(),
		PermissionShopOrderGet.String(),
		PermissionSupplierCreate.String(),
		PermissionSupplierFetch.String(),
		PermissionSupplierGet.String(),
		PermissionSupplierUpdate.String(),
		PermissionSupplierDelete.String(),
		PermissionSupplierOptionForProductReceiveForm.String(),
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
		PermissionUnitCreate.String(),
		PermissionUnitFetch.String(),
		PermissionUnitGet.String(),
		PermissionUnitUpdate.String(),
		PermissionUnitDelete.String(),
		PermissionUnitOptionForProductUnitForm.String(),
		PermissionUnitOptionForProductUnitToUnitForm.String(),
	}
}
