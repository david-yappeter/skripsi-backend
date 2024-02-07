// Code generated by "go run myapp/tool/stringer -linecomment -type=CustomerDebtStatus,CustomerDebtDebtSource -output=customer_debt_enum_gen.go -swagoutput=../tool/swag/enum_gen/customer_debt_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["CustomerDebtStatusEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"UNPAID",
				"CANCELED",
				"PAID",
			},
		},
	}

	PostSwaggerDefinitions["CustomerDebtDebtSourceEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"DELIVERY_ORDER",
			},
		},
	}
}