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
	_ = x[PermissionAdminUserCreate-1]
	_ = x[PermissionAdminUserUpdate-2]
	_ = x[PermissionAdminUserUpdatePassword-3]
	_ = x[PermissionAdminUserUpdateActive-4]
	_ = x[PermissionAdminUserUpdateInActive-5]
	_ = x[PermissionAdminUnitCreate-6]
	_ = x[PermissionAdminUnitFetch-7]
	_ = x[PermissionAdminUnitGet-8]
	_ = x[PermissionAdminUnitUpdate-9]
	_ = x[PermissionAdminUnitDelete-10]
	_ = x[PermissionAdminSupplierCreate-11]
	_ = x[PermissionAdminSupplierFetch-12]
	_ = x[PermissionAdminSupplierGet-13]
	_ = x[PermissionAdminSupplierUpdate-14]
	_ = x[PermissionAdminSupplierDelete-15]
	_ = x[PermissionAdminSupplierTypeCreate-16]
	_ = x[PermissionAdminSupplierTypeFetch-17]
	_ = x[PermissionAdminSupplierTypeGet-18]
	_ = x[PermissionAdminSupplierTypeUpdate-19]
	_ = x[PermissionAdminSupplierTypeDelete-20]
	_ = x[PermissionAdminProductUnitCreate-21]
	_ = x[PermissionAdminProductUnitUpload-22]
	_ = x[PermissionAdminProductUnitGet-23]
	_ = x[PermissionAdminProductUnitUpdate-24]
	_ = x[PermissionAdminProductUnitDelete-25]
	_ = x[PermissionCustomerCreate-26]
	_ = x[PermissionCustomerFetch-27]
	_ = x[PermissionCustomerGet-28]
	_ = x[PermissionCustomerUpdate-29]
	_ = x[PermissionCustomerDelete-30]
	_ = x[PermissionSupplierTypeCreate-31]
	_ = x[PermissionSupplierTypeFetch-32]
	_ = x[PermissionSupplierTypeGet-33]
	_ = x[PermissionSupplierTypeUpdate-34]
	_ = x[PermissionSupplierTypeDelete-35]
}

const _Permission_nameReadable = "ADMIN_USER_CREATE, ADMIN_USER_UPDATE, ADMIN_USER_UPDATE_PASSWORD, ADMIN_USER_UPDATE_ACTIVE, ADMIN_USER_UPDATE_INACTIVE, ADMIN_UNIT_CREATE, ADMIN_UNIT_FETCH, ADMIN_UNIT_GET, ADMIN_UNIT_UPDATE, ADMIN_UNIT_DELETE, ADMIN_SUPPLIER_CREATE, ADMIN_SUPPLIER_FETCH, ADMIN_SUPPLIER_GET, ADMIN_SUPPLIER_UPDATE, ADMIN_SUPPLIER_DELETE, ADMIN_SUPPLIER_TYPE_CREATE, ADMIN_SUPPLIER_TYPE_FETCH, ADMIN_SUPPLIER_TYPE_GET, ADMIN_SUPPLIER_TYPE_UPDATE, ADMIN_SUPPLIER_TYPE_DELETE, PRODUCT_ADMIN_UNIT_CREATE, PRODUCT_ADMIN_UNIT_UPLOAD, PRODUCT_ADMIN_UNIT_GET, PRODUCT_ADMIN_UNIT_UPDATE, PRODUCT_ADMIN_UNIT_DELETE, CUSTOMER_CREATE, CUSTOMER_FETCH, CUSTOMER_GET, CUSTOMER_UPDATE, CUSTOMER_DELETE, SUPPLIER_TYPE_CREATE, SUPPLIER_TYPE_FETCH, SUPPLIER_TYPE_GET, SUPPLIER_TYPE_UPDATE, SUPPLIER_TYPE_DELETE"

const _Permission_name = "ADMIN_USER_CREATEADMIN_USER_UPDATEADMIN_USER_UPDATE_PASSWORDADMIN_USER_UPDATE_ACTIVEADMIN_USER_UPDATE_INACTIVEADMIN_UNIT_CREATEADMIN_UNIT_FETCHADMIN_UNIT_GETADMIN_UNIT_UPDATEADMIN_UNIT_DELETEADMIN_SUPPLIER_CREATEADMIN_SUPPLIER_FETCHADMIN_SUPPLIER_GETADMIN_SUPPLIER_UPDATEADMIN_SUPPLIER_DELETEADMIN_SUPPLIER_TYPE_CREATEADMIN_SUPPLIER_TYPE_FETCHADMIN_SUPPLIER_TYPE_GETADMIN_SUPPLIER_TYPE_UPDATEADMIN_SUPPLIER_TYPE_DELETEPRODUCT_ADMIN_UNIT_CREATEPRODUCT_ADMIN_UNIT_UPLOADPRODUCT_ADMIN_UNIT_GETPRODUCT_ADMIN_UNIT_UPDATEPRODUCT_ADMIN_UNIT_DELETECUSTOMER_CREATECUSTOMER_FETCHCUSTOMER_GETCUSTOMER_UPDATECUSTOMER_DELETESUPPLIER_TYPE_CREATESUPPLIER_TYPE_FETCHSUPPLIER_TYPE_GETSUPPLIER_TYPE_UPDATESUPPLIER_TYPE_DELETE"

var _Permission_index = [...]uint16{0, 17, 34, 60, 84, 110, 127, 143, 157, 174, 191, 212, 232, 250, 271, 292, 318, 343, 366, 392, 418, 443, 468, 490, 515, 540, 555, 569, 581, 596, 611, 631, 650, 667, 687, 707}

func (i *Permission) determine(s string) {
	switch s {
	case "ADMIN_USER_CREATE":
		*i = PermissionAdminUserCreate
	case "ADMIN_USER_UPDATE":
		*i = PermissionAdminUserUpdate
	case "ADMIN_USER_UPDATE_PASSWORD":
		*i = PermissionAdminUserUpdatePassword
	case "ADMIN_USER_UPDATE_ACTIVE":
		*i = PermissionAdminUserUpdateActive
	case "ADMIN_USER_UPDATE_INACTIVE":
		*i = PermissionAdminUserUpdateInActive
	case "ADMIN_UNIT_CREATE":
		*i = PermissionAdminUnitCreate
	case "ADMIN_UNIT_FETCH":
		*i = PermissionAdminUnitFetch
	case "ADMIN_UNIT_GET":
		*i = PermissionAdminUnitGet
	case "ADMIN_UNIT_UPDATE":
		*i = PermissionAdminUnitUpdate
	case "ADMIN_UNIT_DELETE":
		*i = PermissionAdminUnitDelete
	case "ADMIN_SUPPLIER_CREATE":
		*i = PermissionAdminSupplierCreate
	case "ADMIN_SUPPLIER_FETCH":
		*i = PermissionAdminSupplierFetch
	case "ADMIN_SUPPLIER_GET":
		*i = PermissionAdminSupplierGet
	case "ADMIN_SUPPLIER_UPDATE":
		*i = PermissionAdminSupplierUpdate
	case "ADMIN_SUPPLIER_DELETE":
		*i = PermissionAdminSupplierDelete
	case "ADMIN_SUPPLIER_TYPE_CREATE":
		*i = PermissionAdminSupplierTypeCreate
	case "ADMIN_SUPPLIER_TYPE_FETCH":
		*i = PermissionAdminSupplierTypeFetch
	case "ADMIN_SUPPLIER_TYPE_GET":
		*i = PermissionAdminSupplierTypeGet
	case "ADMIN_SUPPLIER_TYPE_UPDATE":
		*i = PermissionAdminSupplierTypeUpdate
	case "ADMIN_SUPPLIER_TYPE_DELETE":
		*i = PermissionAdminSupplierTypeDelete
	case "PRODUCT_ADMIN_UNIT_CREATE":
		*i = PermissionAdminProductUnitCreate
	case "PRODUCT_ADMIN_UNIT_UPLOAD":
		*i = PermissionAdminProductUnitUpload
	case "PRODUCT_ADMIN_UNIT_GET":
		*i = PermissionAdminProductUnitGet
	case "PRODUCT_ADMIN_UNIT_UPDATE":
		*i = PermissionAdminProductUnitUpdate
	case "PRODUCT_ADMIN_UNIT_DELETE":
		*i = PermissionAdminProductUnitDelete
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
		PermissionAdminUserCreate,
		PermissionAdminUserUpdate,
		PermissionAdminUserUpdatePassword,
		PermissionAdminUserUpdateActive,
		PermissionAdminUserUpdateInActive,
		PermissionAdminUnitCreate,
		PermissionAdminUnitFetch,
		PermissionAdminUnitGet,
		PermissionAdminUnitUpdate,
		PermissionAdminUnitDelete,
		PermissionAdminSupplierCreate,
		PermissionAdminSupplierFetch,
		PermissionAdminSupplierGet,
		PermissionAdminSupplierUpdate,
		PermissionAdminSupplierDelete,
		PermissionAdminSupplierTypeCreate,
		PermissionAdminSupplierTypeFetch,
		PermissionAdminSupplierTypeGet,
		PermissionAdminSupplierTypeUpdate,
		PermissionAdminSupplierTypeDelete,
		PermissionAdminProductUnitCreate,
		PermissionAdminProductUnitUpload,
		PermissionAdminProductUnitGet,
		PermissionAdminProductUnitUpdate,
		PermissionAdminProductUnitDelete,
		PermissionCustomerCreate,
		PermissionCustomerFetch,
		PermissionCustomerGet,
		PermissionCustomerUpdate,
		PermissionCustomerDelete,
		PermissionSupplierTypeCreate,
		PermissionSupplierTypeFetch,
		PermissionSupplierTypeGet,
		PermissionSupplierTypeUpdate,
		PermissionSupplierTypeDelete,
	}
}

func ListPermissionString() []string {
	return []string{
		PermissionAdminUserCreate.String(),
		PermissionAdminUserUpdate.String(),
		PermissionAdminUserUpdatePassword.String(),
		PermissionAdminUserUpdateActive.String(),
		PermissionAdminUserUpdateInActive.String(),
		PermissionAdminUnitCreate.String(),
		PermissionAdminUnitFetch.String(),
		PermissionAdminUnitGet.String(),
		PermissionAdminUnitUpdate.String(),
		PermissionAdminUnitDelete.String(),
		PermissionAdminSupplierCreate.String(),
		PermissionAdminSupplierFetch.String(),
		PermissionAdminSupplierGet.String(),
		PermissionAdminSupplierUpdate.String(),
		PermissionAdminSupplierDelete.String(),
		PermissionAdminSupplierTypeCreate.String(),
		PermissionAdminSupplierTypeFetch.String(),
		PermissionAdminSupplierTypeGet.String(),
		PermissionAdminSupplierTypeUpdate.String(),
		PermissionAdminSupplierTypeDelete.String(),
		PermissionAdminProductUnitCreate.String(),
		PermissionAdminProductUnitUpload.String(),
		PermissionAdminProductUnitGet.String(),
		PermissionAdminProductUnitUpdate.String(),
		PermissionAdminProductUnitDelete.String(),
		PermissionCustomerCreate.String(),
		PermissionCustomerFetch.String(),
		PermissionCustomerGet.String(),
		PermissionCustomerUpdate.String(),
		PermissionCustomerDelete.String(),
		PermissionSupplierTypeCreate.String(),
		PermissionSupplierTypeFetch.String(),
		PermissionSupplierTypeGet.String(),
		PermissionSupplierTypeUpdate.String(),
		PermissionSupplierTypeDelete.String(),
	}
}
