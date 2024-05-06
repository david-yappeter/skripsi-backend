package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=ProductReceiveStatus -output=product_receive_enum_gen.go -swagoutput=../tool/swag/enum_gen/product_receive_enum_gen.go -custom
type ProductReceiveStatus int // @name ProductReceiveStatusEnum

const (
	ProductReceiveStatusPending   ProductReceiveStatus = iota + 1 // PENDING
	ProductReceiveStatusCanceled                                  // CANCELED
	ProductReceiveStatusCompleted                                 // COMPLETED
	ProductReceiveStatusReturned                                  // RETURNED
)
