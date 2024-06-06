// Code generated by "go run myapp/tool/stringer -linecomment -type=ProductReturnStatus -output=product_return_enum_gen.go -swagoutput=../tool/swag/enum_gen/product_return_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["ProductReturnStatusEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"PENDING",
				"COMPLETED",
			},
		},
	}
}
