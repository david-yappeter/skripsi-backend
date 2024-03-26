package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=DeliveryOrderStatus -output=delivery_order_enum_gen.go -swagoutput=../tool/swag/enum_gen/delivery_order_enum_gen.go -custom
type DeliveryOrderStatus int // @name DeliveryOrderStatusEnum

const (
	DeliveryOrderStatusPending    DeliveryOrderStatus = iota + 1 // PENDING
	DeliveryOrderStatusOngoing                                   // ONGOING
	DeliveryOrderStatusDelivering                                // DELIVERING
	DeliveryOrderStatusCanceled                                  // CANCELED
	DeliveryOrderStatusCompleted                                 // COMPLETED
)
