package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=ShopOrderPlatformType,ShopOrderTrackingStatus -output=shop_order_enum_gen.go -swagoutput=../tool/swag/enum_gen/shop_order_enum_gen.go -custom
type ShopOrderPlatformType int // @name ShopOrderPlatformTypeEnum

const (
	ShopOrderPlatformTypeTiktokShop ShopOrderPlatformType = iota + 1 // TIKTOK_SHOP
	ShopOrderPlatformTypeShopee                                      // SHOPEE
)

type ShopOrderTrackingStatus int // @name ShopOrderTrackingStatusEnum

const (
	ShopOrderTrackingStatusUnpaid             ShopOrderTrackingStatus = iota + 1 // UNPAID
	ShopOrderTrackingStatusAwaitingShipment                                      // AWAITING_SHIPMENT
	ShopOrderTrackingStatusAwaitingCollection                                    // AWAITING_COLLECTION
	ShopOrderTrackingStatusPartiallyShipping                                     // PARTIALLY_SHIPPING
	ShopOrderTrackingStatusShipping                                              // SHIPPING
	ShopOrderTrackingStatusCancel                                                // CANCEL
	ShopOrderTrackingStatusDelivered                                             // DELIVERED
	ShopOrderTrackingStatusCompleted                                             // COMPLETED
	ShopOrderTrackingStatusWillReturn                                            // WILL_RETURN
	ShopOrderTrackingStatusReturned                                              // RETURNED
)
