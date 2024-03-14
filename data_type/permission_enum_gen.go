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
	_ = x[PermissionCashierSessionStart-13]
	_ = x[PermissionCashierSessionGetCurrent-14]
	_ = x[PermissionCashierSessionEnd-15]
	_ = x[PermissionCustomerCreate-16]
	_ = x[PermissionCustomerFetch-17]
	_ = x[PermissionCustomerGet-18]
	_ = x[PermissionCustomerUpdate-19]
	_ = x[PermissionCustomerDelete-20]
	_ = x[PermissionCustomerOptionForDeliveryOrderForm-21]
	_ = x[PermissionCustomerDebtUploadImage-22]
	_ = x[PermissionCustomerDebtFetch-23]
	_ = x[PermissionCustomerDebtGet-24]
	_ = x[PermissionCustomerDebtPayment-25]
	_ = x[PermissionCustomerTypeCreate-26]
	_ = x[PermissionCustomerTypeFetch-27]
	_ = x[PermissionCustomerTypeGet-28]
	_ = x[PermissionCustomerTypeUpdate-29]
	_ = x[PermissionCustomerTypeDelete-30]
	_ = x[PermissionCustomerTypeOptionForCustomerForm-31]
	_ = x[PermissionDeliveryOrderCreate-32]
	_ = x[PermissionDeliveryOrderUpload-33]
	_ = x[PermissionDeliveryOrderAddItem-34]
	_ = x[PermissionDeliveryOrderAddImage-35]
	_ = x[PermissionDeliveryOrderAddDriver-36]
	_ = x[PermissionDeliveryOrderFetch-37]
	_ = x[PermissionDeliveryOrderGet-38]
	_ = x[PermissionDeliveryOrderMarkOngoing-39]
	_ = x[PermissionDeliveryOrderCancel-40]
	_ = x[PermissionDeliveryOrderMarkCompleted-41]
	_ = x[PermissionDeliveryOrderDelete-42]
	_ = x[PermissionDeliveryOrderDeleteItem-43]
	_ = x[PermissionDeliveryOrderDeleteImage-44]
	_ = x[PermissionDeliveryOrderDeleteDriver-45]
	_ = x[PermissionProductCreate-46]
	_ = x[PermissionProductUpload-47]
	_ = x[PermissionProductFetch-48]
	_ = x[PermissionProductGet-49]
	_ = x[PermissionProductUpdate-50]
	_ = x[PermissionProductDelete-51]
	_ = x[PermissionProductOptionForProductReceiveForm-52]
	_ = x[PermissionProductOptionForDeliveryOrderForm-53]
	_ = x[PermissionProductReceiveCreate-54]
	_ = x[PermissionProductReceiveUpload-55]
	_ = x[PermissionProductReceiveAddItem-56]
	_ = x[PermissionProductReceiveAddImage-57]
	_ = x[PermissionProductReceiveCancel-58]
	_ = x[PermissionProductReceiveMarkComplete-59]
	_ = x[PermissionProductReceiveFetch-60]
	_ = x[PermissionProductReceiveGet-61]
	_ = x[PermissionProductReceiveDelete-62]
	_ = x[PermissionProductReceiveDeleteItem-63]
	_ = x[PermissionProductReceiveDeleteImage-64]
	_ = x[PermissionProductStockFetch-65]
	_ = x[PermissionProductStockGet-66]
	_ = x[PermissionProductStockAdjustment-67]
	_ = x[PermissionProductUnitCreate-68]
	_ = x[PermissionProductUnitGet-69]
	_ = x[PermissionProductUnitUpdate-70]
	_ = x[PermissionProductUnitDelete-71]
	_ = x[PermissionProductUnitOptionForProductReceiveForm-72]
	_ = x[PermissionProductUnitOptionForDeliveryOrderForm-73]
	_ = x[PermissionRoleOptionForUserForm-74]
	_ = x[PermissionShopOrderFetch-75]
	_ = x[PermissionShopOrderGet-76]
	_ = x[PermissionSupplierCreate-77]
	_ = x[PermissionSupplierFetch-78]
	_ = x[PermissionSupplierGet-79]
	_ = x[PermissionSupplierUpdate-80]
	_ = x[PermissionSupplierDelete-81]
	_ = x[PermissionSupplierOptionForProductReceiveForm-82]
	_ = x[PermissionSupplierTypeCreate-83]
	_ = x[PermissionSupplierTypeFetch-84]
	_ = x[PermissionSupplierTypeGet-85]
	_ = x[PermissionSupplierTypeUpdate-86]
	_ = x[PermissionSupplierTypeDelete-87]
	_ = x[PermissionSupplierTypeOptionForSupplierForm-88]
	_ = x[PermissionTiktokProductCreate-89]
	_ = x[PermissionTiktokProductUploadImage-90]
	_ = x[PermissionTiktokProductFetchBrands-91]
	_ = x[PermissionTiktokProductFetchCategories-92]
	_ = x[PermissionTiktokProductGetCategoryRules-93]
	_ = x[PermissionTiktokProductGetCategoryAttributes-94]
	_ = x[PermissionTiktokProductGet-95]
	_ = x[PermissionTiktokProductUpdate-96]
	_ = x[PermissionTiktokProductRecommendedCategory-97]
	_ = x[PermissionTiktokProductActivate-98]
	_ = x[PermissionTiktokProductDeactivate-99]
	_ = x[PermissionTransactionCheckoutCart-100]
	_ = x[PermissionUserCreate-101]
	_ = x[PermissionUserFetch-102]
	_ = x[PermissionUserGet-103]
	_ = x[PermissionUserUpdate-104]
	_ = x[PermissionUserUpdatePassword-105]
	_ = x[PermissionUserUpdateActive-106]
	_ = x[PermissionUserUpdateInActive-107]
	_ = x[PermissionUserAddRole-108]
	_ = x[PermissionUserDeleteRole-109]
	_ = x[PermissionUnitCreate-110]
	_ = x[PermissionUnitFetch-111]
	_ = x[PermissionUnitGet-112]
	_ = x[PermissionUnitUpdate-113]
	_ = x[PermissionUnitDelete-114]
	_ = x[PermissionUnitOptionForProductUnitForm-115]
}

const _Permission_nameReadable = "BALANCE_CREATE, BALANCE_FETCH, BALANCE_GET, BALANCE_UPDATE, BALANCE_DELETE, CART_GET_ACTIVE, CART_ADD_ITEM, CART_UPDATE_ITEM, CART_DELETE_ITEM, CART_SET_ACTIVE, CART_SET_IN_ACTIVE, CART_DELETE, CASHIER_SESSION_START, CASHIER_SESSION_GET_CURRENT, CASHIER_SESSION_END, CUSTOMER_CREATE, CUSTOMER_FETCH, CUSTOMER_GET, CUSTOMER_UPDATE, CUSTOMER_DELETE, CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM, CUSTOMER_DEBT_UPLOAD_IMAGE, CUSTOMER_DEBT_FETCH, CUSTOMER_DEBT_GET, CUSTOMER_DEBT_PAYMENT, CUSTOMER_TYPE_CREATE, CUSTOMER_TYPE_FETCH, CUSTOMER_TYPE_GET, CUSTOMER_TYPE_UPDATE, CUSTOMER_TYPE_DELETE, CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM, DELIVERY_ORDER_CREATE, DELIVERY_ORDER_UPLOAD, DELIVERY_ORDER_ADD_ITEM, DELIVERY_ORDER_ADD_IMAGE, DELIVERY_ORDER_ADD_DRIVER, DELIVERY_ORDER_FETCH, DELIVERY_ORDER_GET, DELIVERY_ORDER_MARK_ONGOING, DELIVERY_ORDER_CANCEL, DELIVERY_ORDER_MARK_COMPLETED, DELIVERY_ORDER_DELETE, DELIVERY_ORDER_DELETE_ITEM, DELIVERY_ORDER_DELETE_IMAGE, DELIVERY_ORDER_DELETE_DRIVER, PRODUCT_CREATE, PRODUCT_UPLOAD, PRODUCT_FETCH, PRODUCT_GET, PRODUCT_UPDATE, PRODUCT_DELETE, PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORM, PRODUCT_OPTION_FOR_DELIVERY_ORDER_FORM, PRODUCT_RECEIVE_CREATE, PRODUCT_RECEIVE_UPLOAD, PRODUCT_RECEIVE_ADD_ITEM, PRODUCT_RECEIVE_ADD_IMAGE, PRODUCT_RECEIVE_CANCEL, PRODUCT_RECEIVE_MARK_COMPLETE, PRODUCT_RECEIVE_FETCH, PRODUCT_RECEIVE_GET, PRODUCT_RECEIVE_DELETE, PRODUCT_RECEIVE_DELETE_ITEM, PRODUCT_RECEIVE_DELETE_IMAGE, PRODUCT_STOCK_FETCH, PRODUCT_STOCK_GET, PRODUCT_STOCK_ADJUSTMENT, PRODUCT_UNIT_CREATE, PRODUCT_UNIT_GET, PRODUCT_UNIT_UPDATE, PRODUCT_UNIT_DELETE, PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORM, PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORM, ROLE_OPTION_FOR_USER_FORM, SHOP_ORDER_FETCH, SHOP_ORDER_GET, SUPPLIER_CREATE, SUPPLIER_FETCH, SUPPLIER_GET, SUPPLIER_UPDATE, SUPPLIER_DELETE, SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM, SUPPLIER_TYPE_CREATE, SUPPLIER_TYPE_FETCH, SUPPLIER_TYPE_GET, SUPPLIER_TYPE_UPDATE, SUPPLIER_TYPE_DELETE, SUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORM, TIKTOK_PRODUCT_CREATE, TIKTOK_PRODUCT_UPLOAD_IMAGE, TIKTOK_PRODUCT_FETCH_BRANDS, TIKTOK_PRODUCT_FETCH_CATEGORIES, TIKTOK_PRODUCT_GET_RULES, TIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTES, TIKTOK_PRODUCT_GET, TIKTOK_PRODUCT_UPDATE, TIKTOK_PRODUCT_RECOMMENDED_CATEGORY, TIKTOK_PRODUCT_RECOMMENDED_ACTIVATE, TIKTOK_PRODUCT_RECOMMENDED_DEACTIVATE, TRANSACTION_CHECKOUT_CART, USER_CREATE, USER_FETCH, USER_GET, USER_UPDATE, USER_UPDATE_PASSWORD, USER_UPDATE_ACTIVE, USER_UPDATE_INACTIVE, USER_ADD_ROLE, USER_DELETE_ROLE, UNIT_CREATE, UNIT_FETCH, UNIT_GET, UNIT_UPDATE, UNIT_DELETE, UNIT_OPTION_FOR_PRODUCT_UNIT_FORM"

const _Permission_name = "BALANCE_CREATEBALANCE_FETCHBALANCE_GETBALANCE_UPDATEBALANCE_DELETECART_GET_ACTIVECART_ADD_ITEMCART_UPDATE_ITEMCART_DELETE_ITEMCART_SET_ACTIVECART_SET_IN_ACTIVECART_DELETECASHIER_SESSION_STARTCASHIER_SESSION_GET_CURRENTCASHIER_SESSION_ENDCUSTOMER_CREATECUSTOMER_FETCHCUSTOMER_GETCUSTOMER_UPDATECUSTOMER_DELETECUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORMCUSTOMER_DEBT_UPLOAD_IMAGECUSTOMER_DEBT_FETCHCUSTOMER_DEBT_GETCUSTOMER_DEBT_PAYMENTCUSTOMER_TYPE_CREATECUSTOMER_TYPE_FETCHCUSTOMER_TYPE_GETCUSTOMER_TYPE_UPDATECUSTOMER_TYPE_DELETECUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORMDELIVERY_ORDER_CREATEDELIVERY_ORDER_UPLOADDELIVERY_ORDER_ADD_ITEMDELIVERY_ORDER_ADD_IMAGEDELIVERY_ORDER_ADD_DRIVERDELIVERY_ORDER_FETCHDELIVERY_ORDER_GETDELIVERY_ORDER_MARK_ONGOINGDELIVERY_ORDER_CANCELDELIVERY_ORDER_MARK_COMPLETEDDELIVERY_ORDER_DELETEDELIVERY_ORDER_DELETE_ITEMDELIVERY_ORDER_DELETE_IMAGEDELIVERY_ORDER_DELETE_DRIVERPRODUCT_CREATEPRODUCT_UPLOADPRODUCT_FETCHPRODUCT_GETPRODUCT_UPDATEPRODUCT_DELETEPRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORMPRODUCT_OPTION_FOR_DELIVERY_ORDER_FORMPRODUCT_RECEIVE_CREATEPRODUCT_RECEIVE_UPLOADPRODUCT_RECEIVE_ADD_ITEMPRODUCT_RECEIVE_ADD_IMAGEPRODUCT_RECEIVE_CANCELPRODUCT_RECEIVE_MARK_COMPLETEPRODUCT_RECEIVE_FETCHPRODUCT_RECEIVE_GETPRODUCT_RECEIVE_DELETEPRODUCT_RECEIVE_DELETE_ITEMPRODUCT_RECEIVE_DELETE_IMAGEPRODUCT_STOCK_FETCHPRODUCT_STOCK_GETPRODUCT_STOCK_ADJUSTMENTPRODUCT_UNIT_CREATEPRODUCT_UNIT_GETPRODUCT_UNIT_UPDATEPRODUCT_UNIT_DELETEPRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORMPRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORMROLE_OPTION_FOR_USER_FORMSHOP_ORDER_FETCHSHOP_ORDER_GETSUPPLIER_CREATESUPPLIER_FETCHSUPPLIER_GETSUPPLIER_UPDATESUPPLIER_DELETESUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORMSUPPLIER_TYPE_CREATESUPPLIER_TYPE_FETCHSUPPLIER_TYPE_GETSUPPLIER_TYPE_UPDATESUPPLIER_TYPE_DELETESUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORMTIKTOK_PRODUCT_CREATETIKTOK_PRODUCT_UPLOAD_IMAGETIKTOK_PRODUCT_FETCH_BRANDSTIKTOK_PRODUCT_FETCH_CATEGORIESTIKTOK_PRODUCT_GET_RULESTIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTESTIKTOK_PRODUCT_GETTIKTOK_PRODUCT_UPDATETIKTOK_PRODUCT_RECOMMENDED_CATEGORYTIKTOK_PRODUCT_RECOMMENDED_ACTIVATETIKTOK_PRODUCT_RECOMMENDED_DEACTIVATETRANSACTION_CHECKOUT_CARTUSER_CREATEUSER_FETCHUSER_GETUSER_UPDATEUSER_UPDATE_PASSWORDUSER_UPDATE_ACTIVEUSER_UPDATE_INACTIVEUSER_ADD_ROLEUSER_DELETE_ROLEUNIT_CREATEUNIT_FETCHUNIT_GETUNIT_UPDATEUNIT_DELETEUNIT_OPTION_FOR_PRODUCT_UNIT_FORM"

var _Permission_index = [...]uint16{0, 14, 27, 38, 52, 66, 81, 94, 110, 126, 141, 159, 170, 191, 218, 237, 252, 266, 278, 293, 308, 347, 373, 392, 409, 430, 450, 469, 486, 506, 526, 564, 585, 606, 629, 653, 678, 698, 716, 743, 764, 793, 814, 840, 867, 895, 909, 923, 936, 947, 961, 975, 1014, 1052, 1074, 1096, 1120, 1145, 1167, 1196, 1217, 1236, 1258, 1285, 1313, 1332, 1349, 1373, 1392, 1408, 1427, 1446, 1490, 1533, 1558, 1574, 1588, 1603, 1617, 1629, 1644, 1659, 1699, 1719, 1738, 1755, 1775, 1795, 1833, 1854, 1881, 1908, 1939, 1963, 2001, 2019, 2040, 2075, 2110, 2147, 2172, 2183, 2193, 2201, 2212, 2232, 2250, 2270, 2283, 2299, 2310, 2320, 2328, 2339, 2350, 2383}

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
	case "CASHIER_SESSION_START":
		*i = PermissionCashierSessionStart
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
		PermissionCashierSessionStart,
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
		PermissionUnitCreate,
		PermissionUnitFetch,
		PermissionUnitGet,
		PermissionUnitUpdate,
		PermissionUnitDelete,
		PermissionUnitOptionForProductUnitForm,
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
		PermissionCashierSessionStart.String(),
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
		PermissionUnitCreate.String(),
		PermissionUnitFetch.String(),
		PermissionUnitGet.String(),
		PermissionUnitUpdate.String(),
		PermissionUnitDelete.String(),
		PermissionUnitOptionForProductUnitForm.String(),
	}
}
