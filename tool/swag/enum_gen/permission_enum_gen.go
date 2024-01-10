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
				"ADMIN_USER_CREATE",
				"ADMIN_USER_UPDATE",
				"ADMIN_USER_UPDATE_PASSWORD",
				"ADMIN_USER_UPDATE_ACTIVE",
				"ADMIN_USER_UPDATE_INACTIVE",
				"ADMIN_UNIT_CREATE",
				"ADMIN_UNIT_FETCH",
				"ADMIN_UNIT_GET",
				"ADMIN_UNIT_UPDATE",
				"ADMIN_UNIT_DELETE",
				"ADMIN_SUPPLIER_CREATE",
				"ADMIN_SUPPLIER_FETCH",
				"ADMIN_SUPPLIER_GET",
				"ADMIN_SUPPLIER_UPDATE",
				"ADMIN_SUPPLIER_DELETE",
				"ADMIN_SUPPLIER_TYPE_CREATE",
				"ADMIN_SUPPLIER_TYPE_FETCH",
				"ADMIN_SUPPLIER_TYPE_GET",
				"ADMIN_SUPPLIER_TYPE_UPDATE",
				"ADMIN_SUPPLIER_TYPE_DELETE",
				"PRODUCT_ADMIN_UNIT_CREATE",
				"PRODUCT_ADMIN_UNIT_UPLOAD",
				"PRODUCT_ADMIN_UNIT_GET",
				"PRODUCT_ADMIN_UNIT_UPDATE",
				"PRODUCT_ADMIN_UNIT_DELETE",
				"CUSTOMER_CREATE",
				"CUSTOMER_FETCH",
				"CUSTOMER_GET",
				"CUSTOMER_UPDATE",
				"CUSTOMER_DELETE",
				"SUPPLIER_TYPE_CREATE",
				"SUPPLIER_TYPE_FETCH",
				"SUPPLIER_TYPE_GET",
				"SUPPLIER_TYPE_UPDATE",
				"SUPPLIER_TYPE_DELETE",
			},
		},
	}
}
