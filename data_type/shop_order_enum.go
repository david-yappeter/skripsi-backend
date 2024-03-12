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

func (s ShopOrderTrackingStatus) NextValidStatuses() []ShopOrderTrackingStatus {
	switch s {
	case ShopOrderTrackingStatusUnpaid:
		return []ShopOrderTrackingStatus{
			ShopOrderTrackingStatusAwaitingShipment,
			ShopOrderTrackingStatusCancel,
		}
	case ShopOrderTrackingStatusAwaitingShipment:
		return []ShopOrderTrackingStatus{
			ShopOrderTrackingStatusAwaitingCollection,
			ShopOrderTrackingStatusPartiallyShipping,
			ShopOrderTrackingStatusCancel,
		}
	case ShopOrderTrackingStatusPartiallyShipping:
		return []ShopOrderTrackingStatus{
			ShopOrderTrackingStatusAwaitingCollection,
			ShopOrderTrackingStatusCancel,
		}
	case ShopOrderTrackingStatusAwaitingCollection:
		return []ShopOrderTrackingStatus{
			ShopOrderTrackingStatusShipping,
			ShopOrderTrackingStatusCancel,
		}
	case ShopOrderTrackingStatusShipping:
		return []ShopOrderTrackingStatus{
			ShopOrderTrackingStatusDelivered,
			ShopOrderTrackingStatusCompleted,
			ShopOrderTrackingStatusWillReturn,
			ShopOrderTrackingStatusCancel,
		}
	case ShopOrderTrackingStatusDelivered:
		return []ShopOrderTrackingStatus{
			ShopOrderTrackingStatusCompleted,
		}
	case ShopOrderTrackingStatusCompleted:
		return []ShopOrderTrackingStatus{}

	case ShopOrderTrackingStatusWillReturn:
		return []ShopOrderTrackingStatus{
			ShopOrderTrackingStatusReturned,
			ShopOrderTrackingStatusCompleted,
		}
	case ShopOrderTrackingStatusReturned:
		return []ShopOrderTrackingStatus{}
	case ShopOrderTrackingStatusCancel:
		return []ShopOrderTrackingStatus{}
	// case ShopOrderTrackingStatusLost:
	// 	return []ShopOrderTrackingStatus{}
	default:
		panic("invalid order tracking status")
	}
}

func (s ShopOrderTrackingStatus) IsNextStatusValid(nextStatus ShopOrderTrackingStatus) bool {
	statuses := s.NextValidStatuses()

	isExist := false

	for i := range statuses {
		if statuses[i] == nextStatus {
			isExist = true
			break
		}
	}

	return isExist
}

func DetermineOrderTrackingStatusByString(status string) ShopOrderTrackingStatus {
	switch status {
	case "UNPAID":
		return ShopOrderTrackingStatusUnpaid
	case "AWAITING_SHIPMENT", "READY_TO_SHIP", "RETRY_SHIP":
		return ShopOrderTrackingStatusAwaitingShipment
	case "AWAITING_COLLECTION", "PROCESSED":
		return ShopOrderTrackingStatusAwaitingCollection
	case "PARTIALLY_SHIPPING":
		return ShopOrderTrackingStatusPartiallyShipping
	case "SHIPPING", "IN_TRANSIT", "SHIPPED":
		return ShopOrderTrackingStatusShipping
	case "CANCEL", "CANCELLED", "IN_CANCEL":
		return ShopOrderTrackingStatusCancel
	case "DELIVERED", "TO_CONFIRM_RECEIVE":
		return ShopOrderTrackingStatusDelivered
	case "COMPLETED":
		return ShopOrderTrackingStatusCompleted
	case "WILL_RETURN", "TO_RETURN":
		return ShopOrderTrackingStatusWillReturn
	case "RETURNED":
		return ShopOrderTrackingStatusReturned
	default:
		panic("invalid order tracking status")
	}
}
