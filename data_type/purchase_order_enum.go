package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=PurchaseOrderStatus -output=purchase_order_enum_gen.go -swagoutput=../tool/swag/enum_gen/purchase_order_enum_gen.go -custom
type PurchaseOrderStatus int // @name PurchaseOrderStatusEnum

const (
	PurchaseOrderStatusPending   PurchaseOrderStatus = iota + 1 // PENDING
	PurchaseOrderStatusOngoing                                  // ONGOING
	PurchaseOrderStatusCanceled                                 // CANCELED
	PurchaseOrderStatusCompleted                                // COMPLETED
)
