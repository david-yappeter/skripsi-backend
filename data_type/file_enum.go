package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=FileType -output=file_enum_gen.go -swagoutput=../tool/swag/enum_gen/file_enum_gen.go -custom
type FileType int // @name FileTypeEnum

const (
	FileTypeProductImage              FileType = iota + 1 // PRODUCT_IMAGE
	FileTypeProductReceiveImage                           // PRODUCT_RECEIVE_IMAGE
	FileTypePurchaseOrderImage                            // PURCHASE_ORDER_IMAGE
	FileTypeDeliveryOrderImage                            // DELIVERY_ORDER_IMAGE
	FileTypeCustomerPaymentImage                          // CUSTOMER_PAYMENT_IMAGE
	FileTypeDebtPaymentImage                              // DEBT_PAYMENT_IMAGE
	FileTypeDeliveryOrderReturnImage                      // DELIVERY_ORDER_RETURN_IMAGE
	FileTypeProductReceiveReturnImage                     // PRODUCT_RECEIVE_RETURN_IMAGE
	FileTypeProductReturnImage                            // PRODUCT_RETURN_IMAGE
)
