package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=WebhookTiktokOrderStatusChange -output=webhook_enum_gen.go -swagoutput=../tool/swag/enum_gen/webhook_enum_gen.go -custom
type WebhookTiktokOrderStatusChange int // @name WebhookTiktokOrderStatusChangeEnum

const (
	WebhookTiktokOrderStatusChangeUnpaid             WebhookTiktokOrderStatusChange = iota + 1 // UNPAID
	WebhookTiktokOrderStatusChangeOnHold                                                       // ON_HOLD
	WebhookTiktokOrderStatusChangeAwaitingShipment                                             // AWAITING_SHIPMENT
	WebhookTiktokOrderStatusChangeAwaitingCollection                                           // AWAITING_COLLECTION
	WebhookTiktokOrderStatusChangeAwaitingCancel                                               // CANCEL
	WebhookTiktokOrderStatusChangeAwaitingInTransit                                            // IN_TRANSIT
	WebhookTiktokOrderStatusChangeAwaitingDelivered                                            // DELIVERED
	WebhookTiktokOrderStatusChangeAwaitingCompleted                                            // COMPLETED
)
