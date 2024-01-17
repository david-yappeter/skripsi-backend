// Code generated by "go run myapp/tool/stringer -linecomment -type=FileType -output=file_enum_gen.go -swagoutput=../tool/swag/enum_gen/file_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["FileTypeEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"PRODUCT_UNIT_IMAGE",
				"PRODUCT_RECEIVE_IMAGE",
			},
		},
	}
}
