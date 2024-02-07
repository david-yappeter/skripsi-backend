package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=FileType -output=file_enum_gen.go -swagoutput=../tool/swag/enum_gen/file_enum_gen.go -custom
type FileType int // @name FileTypeEnum

const (
	FileTypeProductUnitImage     FileType = iota + 1 // PRODUCT_UNIT_IMAGE
	FileTypeProductReceiveImage                      // PRODUCT_RECEIVE_IMAGE
	FileTypeDeliveryOrderImage                       // DELIVERY_ORDER_IMAGE
	FileTypeCustomerPaymentImage                     // CUSTOMER_PAYMENT_IMAGE

)
