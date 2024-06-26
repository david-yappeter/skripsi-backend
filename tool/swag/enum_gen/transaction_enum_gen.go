// Code generated by "go run myapp/tool/stringer -linecomment -type=TransactionStatus -output=transaction_enum_gen.go -swagoutput=../tool/swag/enum_gen/transaction_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["TransactionStatusEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"UNPAID",
				"PAID",
			},
		},
	}
}
