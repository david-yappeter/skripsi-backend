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
	_ = x[PermissionDeliveryOrderCreate-26]
	_ = x[PermissionDeliveryOrderUpload-27]
	_ = x[PermissionDeliveryOrderAddItem-28]
	_ = x[PermissionDeliveryOrderAddImage-29]
	_ = x[PermissionDeliveryOrderAddDriver-30]
	_ = x[PermissionDeliveryOrderFetch-31]
	_ = x[PermissionDeliveryOrderGet-32]
	_ = x[PermissionDeliveryOrderMarkOngoing-33]
	_ = x[PermissionDeliveryOrderCancel-34]
	_ = x[PermissionDeliveryOrderMarkCompleted-35]
	_ = x[PermissionDeliveryOrderDelete-36]
	_ = x[PermissionDeliveryOrderDeleteItem-37]
	_ = x[PermissionDeliveryOrderDeleteImage-38]
	_ = x[PermissionDeliveryOrderDeleteDriver-39]
	_ = x[PermissionProductCreate-40]
	_ = x[PermissionProductFetch-41]
	_ = x[PermissionProductGet-42]
	_ = x[PermissionProductUpdate-43]
	_ = x[PermissionProductDelete-44]
	_ = x[PermissionProductOptionForProductReceiveForm-45]
	_ = x[PermissionProductOptionForDeliveryOrderForm-46]
	_ = x[PermissionProductReceiveCreate-47]
	_ = x[PermissionProductReceiveUpload-48]
	_ = x[PermissionProductReceiveAddItem-49]
	_ = x[PermissionProductReceiveAddImage-50]
	_ = x[PermissionProductReceiveCancel-51]
	_ = x[PermissionProductReceiveMarkComplete-52]
	_ = x[PermissionProductReceiveFetch-53]
	_ = x[PermissionProductReceiveGet-54]
	_ = x[PermissionProductReceiveDelete-55]
	_ = x[PermissionProductReceiveDeleteItem-56]
	_ = x[PermissionProductReceiveDeleteImage-57]
	_ = x[PermissionProductStockFetch-58]
	_ = x[PermissionProductStockGet-59]
	_ = x[PermissionProductStockAdjustment-60]
	_ = x[PermissionProductUnitCreate-61]
	_ = x[PermissionProductUnitUpload-62]
	_ = x[PermissionProductUnitGet-63]
	_ = x[PermissionProductUnitUpdate-64]
	_ = x[PermissionProductUnitDelete-65]
	_ = x[PermissionProductUnitOptionForProductReceiveForm-66]
	_ = x[PermissionProductUnitOptionForDeliveryOrderForm-67]
	_ = x[PermissionRoleOptionForUserForm-68]
	_ = x[PermissionSupplierCreate-69]
	_ = x[PermissionSupplierFetch-70]
	_ = x[PermissionSupplierGet-71]
	_ = x[PermissionSupplierUpdate-72]
	_ = x[PermissionSupplierDelete-73]
	_ = x[PermissionSupplierOptionForProductReceiveForm-74]
	_ = x[PermissionSupplierTypeCreate-75]
	_ = x[PermissionSupplierTypeFetch-76]
	_ = x[PermissionSupplierTypeGet-77]
	_ = x[PermissionSupplierTypeUpdate-78]
	_ = x[PermissionSupplierTypeDelete-79]
	_ = x[PermissionTiktokProductCreate-80]
	_ = x[PermissionUserCreate-81]
	_ = x[PermissionUserUpdate-82]
	_ = x[PermissionUserUpdatePassword-83]
	_ = x[PermissionUserUpdateActive-84]
	_ = x[PermissionUserUpdateInActive-85]
	_ = x[PermissionUserAddRole-86]
	_ = x[PermissionUserDeleteRole-87]
	_ = x[PermissionUnitCreate-88]
	_ = x[PermissionUnitFetch-89]
	_ = x[PermissionUnitGet-90]
	_ = x[PermissionUnitUpdate-91]
	_ = x[PermissionUnitDelete-92]
}

const _Permission_nameReadable = "BALANCE_CREATE, BALANCE_FETCH, BALANCE_GET, BALANCE_UPDATE, BALANCE_DELETE, CART_GET_ACTIVE, CART_ADD_ITEM, CART_UPDATE_ITEM, CART_DELETE_ITEM, CART_SET_ACTIVE, CART_SET_IN_ACTIVE, CART_DELETE, CASHIER_SESSION_START, CASHIER_SESSION_GET_CURRENT, CASHIER_SESSION_END, CUSTOMER_CREATE, CUSTOMER_FETCH, CUSTOMER_GET, CUSTOMER_UPDATE, CUSTOMER_DELETE, CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM, CUSTOMER_DEBT_UPLOAD_IMAGE, CUSTOMER_DEBT_FETCH, CUSTOMER_DEBT_GET, CUSTOMER_DEBT_PAYMENT, DELIVERY_ORDER_CREATE, DELIVERY_ORDER_UPLOAD, DELIVERY_ORDER_ADD_ITEM, DELIVERY_ORDER_ADD_IMAGE, DELIVERY_ORDER_ADD_DRIVER, DELIVERY_ORDER_FETCH, DELIVERY_ORDER_GET, DELIVERY_ORDER_MARK_ONGOING, DELIVERY_ORDER_CANCEL, DELIVERY_ORDER_MARK_COMPLETED, DELIVERY_ORDER_DELETE, DELIVERY_ORDER_DELETE_ITEM, DELIVERY_ORDER_DELETE_IMAGE, DELIVERY_ORDER_DELETE_DRIVER, PRODUCT_CREATE, PRODUCT_FETCH, PRODUCT_GET, PRODUCT_UPDATE, PRODUCT_DELETE, PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORM, PRODUCT_OPTION_FOR_DELIVERY_ORDER_FORM, PRODUCT_RECEIVE_CREATE, PRODUCT_RECEIVE_UPLOAD, PRODUCT_RECEIVE_ADD_ITEM, PRODUCT_RECEIVE_ADD_IMAGE, PRODUCT_RECEIVE_CANCEL, PRODUCT_RECEIVE_MARK_COMPLETE, PRODUCT_RECEIVE_FETCH, PRODUCT_RECEIVE_GET, PRODUCT_RECEIVE_DELETE, PRODUCT_RECEIVE_DELETE_ITEM, PRODUCT_RECEIVE_DELETE_IMAGE, PRODUCT_STOCK_FETCH, PRODUCT_STOCK_GET, PRODUCT_STOCK_ADJUSTMENT, PRODUCT_UNIT_CREATE, PRODUCT_UNIT_UPLOAD, PRODUCT_UNIT_GET, PRODUCT_UNIT_UPDATE, PRODUCT_UNIT_DELETE, PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORM, PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORM, ROLE_OPTION_FOR_USER_FORM, SUPPLIER_CREATE, SUPPLIER_FETCH, SUPPLIER_GET, SUPPLIER_UPDATE, SUPPLIER_DELETE, SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM, SUPPLIER_TYPE_CREATE, SUPPLIER_TYPE_FETCH, SUPPLIER_TYPE_GET, SUPPLIER_TYPE_UPDATE, SUPPLIER_TYPE_DELETE, TIKTOK_PRODUCT_CREATE, USER_CREATE, USER_UPDATE, USER_UPDATE_PASSWORD, USER_UPDATE_ACTIVE, USER_UPDATE_INACTIVE, USER_ADD_ROLE, USER_DELETE_ROLE, UNIT_CREATE, UNIT_FETCH, UNIT_GET, UNIT_UPDATE, UNIT_DELETE"

const _Permission_name = "BALANCE_CREATEBALANCE_FETCHBALANCE_GETBALANCE_UPDATEBALANCE_DELETECART_GET_ACTIVECART_ADD_ITEMCART_UPDATE_ITEMCART_DELETE_ITEMCART_SET_ACTIVECART_SET_IN_ACTIVECART_DELETECASHIER_SESSION_STARTCASHIER_SESSION_GET_CURRENTCASHIER_SESSION_ENDCUSTOMER_CREATECUSTOMER_FETCHCUSTOMER_GETCUSTOMER_UPDATECUSTOMER_DELETECUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORMCUSTOMER_DEBT_UPLOAD_IMAGECUSTOMER_DEBT_FETCHCUSTOMER_DEBT_GETCUSTOMER_DEBT_PAYMENTDELIVERY_ORDER_CREATEDELIVERY_ORDER_UPLOADDELIVERY_ORDER_ADD_ITEMDELIVERY_ORDER_ADD_IMAGEDELIVERY_ORDER_ADD_DRIVERDELIVERY_ORDER_FETCHDELIVERY_ORDER_GETDELIVERY_ORDER_MARK_ONGOINGDELIVERY_ORDER_CANCELDELIVERY_ORDER_MARK_COMPLETEDDELIVERY_ORDER_DELETEDELIVERY_ORDER_DELETE_ITEMDELIVERY_ORDER_DELETE_IMAGEDELIVERY_ORDER_DELETE_DRIVERPRODUCT_CREATEPRODUCT_FETCHPRODUCT_GETPRODUCT_UPDATEPRODUCT_DELETEPRODUCT_OPTION_FOR_PRODUCT_RECEIVE_FORMPRODUCT_OPTION_FOR_DELIVERY_ORDER_FORMPRODUCT_RECEIVE_CREATEPRODUCT_RECEIVE_UPLOADPRODUCT_RECEIVE_ADD_ITEMPRODUCT_RECEIVE_ADD_IMAGEPRODUCT_RECEIVE_CANCELPRODUCT_RECEIVE_MARK_COMPLETEPRODUCT_RECEIVE_FETCHPRODUCT_RECEIVE_GETPRODUCT_RECEIVE_DELETEPRODUCT_RECEIVE_DELETE_ITEMPRODUCT_RECEIVE_DELETE_IMAGEPRODUCT_STOCK_FETCHPRODUCT_STOCK_GETPRODUCT_STOCK_ADJUSTMENTPRODUCT_UNIT_CREATEPRODUCT_UNIT_UPLOADPRODUCT_UNIT_GETPRODUCT_UNIT_UPDATEPRODUCT_UNIT_DELETEPRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_FORMPRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_FORMROLE_OPTION_FOR_USER_FORMSUPPLIER_CREATESUPPLIER_FETCHSUPPLIER_GETSUPPLIER_UPDATESUPPLIER_DELETESUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORMSUPPLIER_TYPE_CREATESUPPLIER_TYPE_FETCHSUPPLIER_TYPE_GETSUPPLIER_TYPE_UPDATESUPPLIER_TYPE_DELETETIKTOK_PRODUCT_CREATEUSER_CREATEUSER_UPDATEUSER_UPDATE_PASSWORDUSER_UPDATE_ACTIVEUSER_UPDATE_INACTIVEUSER_ADD_ROLEUSER_DELETE_ROLEUNIT_CREATEUNIT_FETCHUNIT_GETUNIT_UPDATEUNIT_DELETE"

var _Permission_index = [...]uint16{0, 14, 27, 38, 52, 66, 81, 94, 110, 126, 141, 159, 170, 191, 218, 237, 252, 266, 278, 293, 308, 347, 373, 392, 409, 430, 451, 472, 495, 519, 544, 564, 582, 609, 630, 659, 680, 706, 733, 761, 775, 788, 799, 813, 827, 866, 904, 926, 948, 972, 997, 1019, 1048, 1069, 1088, 1110, 1137, 1165, 1184, 1201, 1225, 1244, 1263, 1279, 1298, 1317, 1361, 1404, 1429, 1444, 1458, 1470, 1485, 1500, 1540, 1560, 1579, 1596, 1616, 1636, 1657, 1668, 1679, 1699, 1717, 1737, 1750, 1766, 1777, 1787, 1795, 1806, 1817}

func (i *Permission) determine(s string) {
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
	case "PRODUCT_UNIT_UPLOAD":
		*i = PermissionProductUnitUpload
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
	case "TIKTOK_PRODUCT_CREATE":
		*i = PermissionTiktokProductCreate
	case "USER_CREATE":
		*i = PermissionUserCreate
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

	i.determine(s)

	return nil
}

func (i *Permission) UnmarshalText(b []byte) error {
	i.determine(string(b))

	return nil
}

func (i *Permission) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.determine(s)
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
		PermissionProductUnitUpload,
		PermissionProductUnitGet,
		PermissionProductUnitUpdate,
		PermissionProductUnitDelete,
		PermissionProductUnitOptionForProductReceiveForm,
		PermissionProductUnitOptionForDeliveryOrderForm,
		PermissionRoleOptionForUserForm,
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
		PermissionTiktokProductCreate,
		PermissionUserCreate,
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
		PermissionProductUnitUpload.String(),
		PermissionProductUnitGet.String(),
		PermissionProductUnitUpdate.String(),
		PermissionProductUnitDelete.String(),
		PermissionProductUnitOptionForProductReceiveForm.String(),
		PermissionProductUnitOptionForDeliveryOrderForm.String(),
		PermissionRoleOptionForUserForm.String(),
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
		PermissionTiktokProductCreate.String(),
		PermissionUserCreate.String(),
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
	}
}
