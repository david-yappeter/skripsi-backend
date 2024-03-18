// Code generated by "go run myapp/tool/stringer -linecomment -type=ProductStockMutationType -output=product_stock_mutation_enum_gen.go -swagoutput=../tool/swag/enum_gen/product_stock_mutation_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["ProductStockMutationTypeEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"PRODUCT_RECEIVE_ITEM",
				"DELIVERY_ORDER_ITEM_COST_CANCEL",
				"PRODUCT_STOCK_ADJUSTMENT",
			},
		},
	}
}
