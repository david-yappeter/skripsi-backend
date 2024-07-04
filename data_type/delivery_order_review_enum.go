package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=DeliveryOrderReviewType -output=delivery_order_review_enum_gen.go -swagoutput=../tool/swag/enum_gen/delivery_order_review_enum_gen.go -custom
type DeliveryOrderReviewType int // @name DeliveryOrderReviewTypeEnum

const (
	DeliveryOrderReviewTypeProduct  DeliveryOrderReviewType = iota + 1 // PRODUCT
	DeliveryOrderReviewTypeDelivery                                    // DELIVERY
)
