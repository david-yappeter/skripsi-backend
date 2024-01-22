// Code generated by "go run myapp/tool/stringer -linecomment -type=Permission -output=permission_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["PermissionEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"BALANCE_CREATE",
				"BALANCE_FETCH",
				"BALANCE_GET",
				"BALANCE_UPDATE",
				"BALANCE_DELETE",
				"CUSTOMER_CREATE",
				"CUSTOMER_FETCH",
				"CUSTOMER_GET",
				"CUSTOMER_UPDATE",
				"CUSTOMER_DELETE",
				"DELIVERY_ORDER_CREATE",
				"DELIVERY_ORDER_UPLOAD",
				"DELIVERY_ORDER_ADD_ITEM",
				"DELIVERY_ORDER_ADD_IMAGE",
				"DELIVERY_ORDER_FETCH",
				"DELIVERY_ORDER_GET",
				"DELIVERY_ORDER_DELETE",
				"DELIVERY_ORDER_DELETE_ITEM",
				"DELIVERY_ORDER_DELETE_IMAGE",
				"PRODUCT_CREATE",
				"PRODUCT_FETCH",
				"PRODUCT_GET",
				"PRODUCT_UPDATE",
				"PRODUCT_DELETE",
				"PRODUCT_RECEIVE_CREATE",
				"PRODUCT_RECEIVE_UPLOAD",
				"PRODUCT_RECEIVE_ADD_ITEM",
				"PRODUCT_RECEIVE_ADD_IMAGE",
				"PRODUCT_RECEIVE_FETCH",
				"PRODUCT_RECEIVE_GET",
				"PRODUCT_RECEIVE_DELETE",
				"PRODUCT_RECEIVE_DELETE_ITEM",
				"PRODUCT_RECEIVE_DELETE_IMAGE",
				"PRODUCT_UNIT_CREATE",
				"PRODUCT_UNIT_UPLOAD",
				"PRODUCT_UNIT_GET",
				"PRODUCT_UNIT_UPDATE",
				"PRODUCT_UNIT_DELETE",
				"ROLE_OPTION_FOR_USER_FORM",
				"SUPPLIER_CREATE",
				"SUPPLIER_FETCH",
				"SUPPLIER_GET",
				"SUPPLIER_UPDATE",
				"SUPPLIER_DELETE",
				"SUPPLIER_TYPE_CREATE",
				"SUPPLIER_TYPE_FETCH",
				"SUPPLIER_TYPE_GET",
				"SUPPLIER_TYPE_UPDATE",
				"SUPPLIER_TYPE_DELETE",
				"USER_CREATE",
				"USER_UPDATE",
				"USER_UPDATE_PASSWORD",
				"USER_UPDATE_ACTIVE",
				"USER_UPDATE_INACTIVE",
				"USER_ADD_ROLE",
				"USER_DELETE_ROLE",
				"UNIT_CREATE",
				"UNIT_FETCH",
				"UNIT_GET",
				"UNIT_UPDATE",
				"UNIT_DELETE",
			},
		},
	}
}
