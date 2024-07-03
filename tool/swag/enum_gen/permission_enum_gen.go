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
				"CART_GET_ACTIVE",
				"CART_ADD_ITEM",
				"CART_UPDATE_ITEM",
				"CART_DELETE_ITEM",
				"CART_SET_ACTIVE",
				"CART_SET_IN_ACTIVE",
				"CART_DELETE",
				"CASHIER_SESSION_FETCH",
				"CASHIER_SESSION_FETCH_FOR_CURRENT_USER",
				"CASHIER_SESSION_START",
				"CASHIER_SESSION_GET",
				"CASHIER_SESSION_FETCH_TRANSACTION",
				"CASHIER_SESSION_DOWNLOAD_REPORT",
				"CASHIER_SESSION_GET_CURRENT",
				"CASHIER_SESSION_END",
				"CUSTOMER_CREATE",
				"CUSTOMER_FETCH",
				"CUSTOMER_GET",
				"CUSTOMER_UPDATE",
				"CUSTOMER_DELETE",
				"CUSTOMER_OPTION_FOR_WHATSAPP_CUSTOMER_DEBT_BROADCAST_FORM",
				"CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM",
				"CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FILTER",
				"CUSTOMER_DEBT_UPLOAD_IMAGE",
				"CUSTOMER_DEBT_FETCH",
				"CUSTOMER_DEBT_GET",
				"CUSTOMER_DEBT_PAYMENT",
				"CUSTOMER_TYPE_CREATE",
				"CUSTOMER_TYPE_FETCH",
				"CUSTOMER_TYPE_GET",
				"CUSTOMER_TYPE_UPDATE",
				"CUSTOMER_TYPE_DELETE",
				"CUSTOMER_TYPE_ADD_DISCOUNT",
				"CUSTOMER_TYPE_UPDATE_DISCOUNT",
				"CUSTOMER_TYPE_DELETE_DISCOUNT",
				"CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM",
				"CUSTOMER_TYPE_OPTION_FOR_WHATSAPP_PRODUCT_PRICE_CHANGE_BROADCAST_FORM",
				"DASHBOARD_SUMMARIZE_DEBT",
				"DASHBOARD_SUMMARIZE_TRANSACTION",
				"DEBT_UPLOAD_IMAGE",
				"DEBT_FETCH",
				"DOWNLOAD_REPORT",
				"DEBT_GET",
				"DEBT_PAYMENT",
				"DELIVERY_ORDER_CREATE",
				"DELIVERY_ORDER_UPLOAD",
				"DELIVERY_ORDER_ADD_ITEM",
				"DELIVERY_ORDER_ADD_IMAGE",
				"DELIVERY_ORDER_ADD_DRIVER",
				"DELIVERY_ORDER_FETCH",
				"DELIVERY_ORDER_FETCH_DRIVER",
				"DELIVERY_ORDER_GET",
				"DELIVERY_ORDER_ACTIVE_FOR_DRIVER",
				"DELIVERY_ORDER_MARK_ONGOING",
				"DELIVERY_ORDER_DELIVERING",
				"DELIVERY_ORDER_UPDATE",
				"DELIVERY_ORDER_CANCEL",
				"DELIVERY_ORDER_MARK_COMPLETED",
				"DELIVERY_ORDER_RETURNED",
				"DELIVERY_ORDER_DELIVERY_LOCATION",
				"DELIVERY_ORDER_DELETE",
				"DELIVERY_ORDER_DELETE_ITEM",
				"DELIVERY_ORDER_DELETE_IMAGE",
				"DELIVERY_ORDER_DELETE_DRIVER",
				"DELIVERY_ORDER_REVIEW_FETCH",
				"DELIVERY_ORDER_REVIEW_GET",
				"PRODUCT_CREATE",
				"PRODUCT_UPLOAD",
				"PRODUCT_FETCH",
				"PRODUCT_GET",
				"PRODUCT_UPDATE",
				"PRODUCT_DELETE",
				"PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM",
				"PRODUCT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM",
				"PRODUCT_OPTION_FOR_CUSTOMER_TYPE_DISCOUNT_FORM",
				"PRODUCT_OPTION_FOR_CART_ADD_ITEM_FORM",
				"PRODUCT_OPTION_FOR_PRODUCT_DISCOUNT_FORM",
				"PRODUCT_DISCOUNT_CREATE",
				"PRODUCT_DISCOUNT_FETCH",
				"PRODUCT_DISCOUNT_GET",
				"PRODUCT_DISCOUNT_UPDATE",
				"PRODUCT_DISCOUNT_DELETE",
				"PRODUCT_RECEIVE_UPLOAD",
				"PRODUCT_RECEIVE_ADD_IMAGE",
				"PRODUCT_RECEIVE_UPDATE",
				"PRODUCT_RECEIVE_CANCEL",
				"PRODUCT_RECEIVE_MARK_COMPLETE",
				"PRODUCT_RECEIVE_FETCH",
				"PRODUCT_RECEIVE_GET",
				"PRODUCT_RECEIVE_UPDATE_ITEM",
				"PRODUCT_RECEIVE_DELETE",
				"PRODUCT_RECEIVE_DELETE_IMAGE",
				"PRODUCT_RETURN_CREATE",
				"PRODUCT_RETURN_UPLOAD",
				"PRODUCT_RETURN_ADD_ITEM",
				"PRODUCT_RETURN_ADD_IMAGE",
				"PRODUCT_RETURN_UPDATE",
				"PRODUCT_RETURN_CANCEL",
				"PRODUCT_RETURN_MARK_COMPLETE",
				"PRODUCT_RETURN_FETCH",
				"PRODUCT_RETURN_GET",
				"PRODUCT_RETURN_DELETE",
				"PRODUCT_RETURN_DELETE_ITEM",
				"PRODUCT_RETURN_DELETE_IMAGE",
				"PRODUCT_STOCK_ADJUSTMENT_FETCH",
				"PRODUCT_STOCK_FETCH",
				"PRODUCT_STOCK_GET",
				"PRODUCT_STOCK_DOWNLOAD_REPORT",
				"PRODUCT_STOCK_ADJUSTMENT",
				"PRODUCT_UNIT_CREATE",
				"PRODUCT_UNIT_GET",
				"PRODUCT_UNIT_UPDATE",
				"PRODUCT_UNIT_DELETE",
				"PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM",
				"PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM",
				"PURCHASE_ORDER_CREATE",
				"PURCHASE_ORDER_UPLOAD",
				"PURCHASE_ORDER_ADD_ITEM",
				"PURCHASE_ORDER_ADD_IMAGE",
				"PURCHASE_ORDER_UPDATE",
				"PURCHASE_ORDER_CANCEL",
				"PURCHASE_ORDER_ONGOING",
				"PURCHASE_ORDER_MARK_COMPLETE",
				"PURCHASE_ORDER_FETCH",
				"PURCHASE_ORDER_GET",
				"PURCHASE_ORDER_DELETE",
				"PURCHASE_ORDER_DELETE_ITEM",
				"PURCHASE_ORDER_DELETE_IMAGE",
				"ROLE_OPTION_FOR_USER_FORM",
				"SHOP_ORDER_FETCH",
				"SHOP_ORDER_GET",
				"SSR_WHATSAPP_LOGIN",
				"SUPPLIER_CREATE",
				"SUPPLIER_FETCH",
				"SUPPLIER_GET",
				"SUPPLIER_UPDATE",
				"SUPPLIER_DELETE",
				"SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM",
				"SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FILTER",
				"SUPPLIER_TYPE_CREATE",
				"SUPPLIER_TYPE_FETCH",
				"SUPPLIER_TYPE_GET",
				"SUPPLIER_TYPE_UPDATE",
				"SUPPLIER_TYPE_DELETE",
				"SUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORM",
				"TIKTOK_PRODUCT_CREATE",
				"TIKTOK_PRODUCT_UPLOAD_IMAGE",
				"TIKTOK_PRODUCT_FETCH_BRANDS",
				"TIKTOK_PRODUCT_FETCH_CATEGORIES",
				"TIKTOK_PRODUCT_GET_RULES",
				"TIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTES",
				"TIKTOK_PRODUCT_GET",
				"TIKTOK_PRODUCT_UPDATE",
				"TIKTOK_PRODUCT_RECOMMENDED_CATEGORY",
				"TIKTOK_PRODUCT_RECOMMENDED_ACTIVATE",
				"TIKTOK_PRODUCT_RECOMMENDED_DEACTIVATE",
				"TRANSACTION_CHECKOUT_CART",
				"TRANSACTION_GET",
				"USER_CREATE",
				"USER_FETCH",
				"USER_GET",
				"USER_UPDATE",
				"USER_UPDATE_PASSWORD",
				"USER_UPDATE_ACTIVE",
				"USER_UPDATE_INACTIVE",
				"USER_ADD_ROLE",
				"USER_DELETE_ROLE",
				"USER_OPTION_FOR_CASHIER_SESSION_FILTER",
				"USER_OPTION_FOR_DELIVERY_ORDER_DRIVER_FORM",
				"USER_OPTION_FOR_PRODUCT_STOCK_ADJUSTMENT_FILTER",
				"UNIT_CREATE",
				"UNIT_FETCH",
				"UNIT_GET",
				"UNIT_UPDATE",
				"UNIT_DELETE",
				"UNIT_OPTION_FOR_PRODUCT_UNIT_FORM",
				"UNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORM",
				"WHATSAPP_IS_LOGGED_IN",
				"WHATSAPP_PRODUCT_PRICE_CHANGE_BROADCAST",
				"WHATSAPP_CUSTOMER_DEBT_BROADCAST",
				"WHATSAPP_CUSTOMER_TYPE_DISCOUNT_BROADCAST",
				"WHATSAPP_LOGOUT",
			},
		},
	}
}
