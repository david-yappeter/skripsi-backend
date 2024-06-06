package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=ProductReturnStatus -output=product_return_enum_gen.go -swagoutput=../tool/swag/enum_gen/product_return_enum_gen.go -custom
type ProductReturnStatus int // @name ProductReturnStatusEnum

const (
	ProductReturnStatusPending   ProductReturnStatus = iota + 1 // PENDING
	ProductReturnStatusCompleted                                // COMPLETED
)
