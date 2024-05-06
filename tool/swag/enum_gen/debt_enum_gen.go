// Code generated by "go run myapp/tool/stringer -linecomment -type=DebtStatus,DebtSource -output=debt_enum_gen.go -swagoutput=../tool/swag/enum_gen/debt_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["DebtStatusEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"UNPAID",
				"CANCELED",
				"PAID",
				"RETURNED",
			},
		},
	}

	PostSwaggerDefinitions["DebtSourceEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"PRODUCT_RECEIVE",
			},
		},
	}
}
