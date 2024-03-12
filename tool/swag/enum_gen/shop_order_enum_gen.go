// Code generated by "go run myapp/tool/stringer -linecomment -type=ShopOrderPlatformType,ShopOrderTrackingStatus -output=shop_order_enum_gen.go -swagoutput=../tool/swag/enum_gen/shop_order_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["ShopOrderPlatformTypeEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"TIKTOK_SHOP",
				"SHOPEE",
			},
		},
	}

	PostSwaggerDefinitions["ShopOrderTrackingStatusEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"UNPAID",
				"AWAITING_SHIPMENT",
				"AWAITING_COLLECTION",
				"PARTIALLY_SHIPPING",
				"SHIPPING",
				"CANCEL",
				"DELIVERED",
				"COMPLETED",
				"WILL_RETURN",
				"RETURNED",
			},
		},
	}
}
