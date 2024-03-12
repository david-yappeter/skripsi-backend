// Code generated by "go run myapp/tool/stringer -linecomment -type=ShopOrderPlatformType,ShopOrderTrackingStatus -output=shop_order_enum_gen.go -swagoutput=../tool/swag/enum_gen/shop_order_enum_gen.go -custom"; DO NOT EDIT.

package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ShopOrderPlatformTypeTiktokShop-1]
	_ = x[ShopOrderPlatformTypeShopee-2]
}

const _ShopOrderPlatformType_nameReadable = "TIKTOK_SHOP, SHOPEE"

const _ShopOrderPlatformType_name = "TIKTOK_SHOPSHOPEE"

var _ShopOrderPlatformType_index = [...]uint8{0, 11, 17}

func (i *ShopOrderPlatformType) Determine(s string) {
	switch s {
	case "TIKTOK_SHOP":
		*i = ShopOrderPlatformTypeTiktokShop
	case "SHOPEE":
		*i = ShopOrderPlatformTypeShopee
	default:
		*i = 0
	}
}

func (i ShopOrderPlatformType) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i ShopOrderPlatformType) GetValidValuesString() string {
	return _ShopOrderPlatformType_nameReadable
}

func (i ShopOrderPlatformType) String() string {
	i -= 1
	if i < 0 || i >= ShopOrderPlatformType(len(_ShopOrderPlatformType_index)-1) {
		return "ShopOrderPlatformType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _ShopOrderPlatformType_name[_ShopOrderPlatformType_index[i]:_ShopOrderPlatformType_index[i+1]]
}

func (i ShopOrderPlatformType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *ShopOrderPlatformType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *ShopOrderPlatformType) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *ShopOrderPlatformType) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i ShopOrderPlatformType) Value() (driver.Value, error) {
	return i.String(), nil
}

func ShopOrderPlatformTypeP(v ShopOrderPlatformType) *ShopOrderPlatformType {
	return &v
}

func ListShopOrderPlatformType() []ShopOrderPlatformType {
	return []ShopOrderPlatformType{
		ShopOrderPlatformTypeTiktokShop,
		ShopOrderPlatformTypeShopee,
	}
}

func ListShopOrderPlatformTypeString() []string {
	return []string{
		ShopOrderPlatformTypeTiktokShop.String(),
		ShopOrderPlatformTypeShopee.String(),
	}
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ShopOrderTrackingStatusUnpaid-1]
	_ = x[ShopOrderTrackingStatusAwaitingShipment-2]
	_ = x[ShopOrderTrackingStatusAwaitingCollection-3]
	_ = x[ShopOrderTrackingStatusPartiallyShipping-4]
	_ = x[ShopOrderTrackingStatusShipping-5]
	_ = x[ShopOrderTrackingStatusCancel-6]
	_ = x[ShopOrderTrackingStatusDelivered-7]
	_ = x[ShopOrderTrackingStatusCompleted-8]
	_ = x[ShopOrderTrackingStatusWillReturn-9]
	_ = x[ShopOrderTrackingStatusReturned-10]
}

const _ShopOrderTrackingStatus_nameReadable = "UNPAID, AWAITING_SHIPMENT, AWAITING_COLLECTION, PARTIALLY_SHIPPING, SHIPPING, CANCEL, DELIVERED, COMPLETED, WILL_RETURN, RETURNED"

const _ShopOrderTrackingStatus_name = "UNPAIDAWAITING_SHIPMENTAWAITING_COLLECTIONPARTIALLY_SHIPPINGSHIPPINGCANCELDELIVEREDCOMPLETEDWILL_RETURNRETURNED"

var _ShopOrderTrackingStatus_index = [...]uint8{0, 6, 23, 42, 60, 68, 74, 83, 92, 103, 111}

func (i *ShopOrderTrackingStatus) Determine(s string) {
	switch s {
	case "UNPAID":
		*i = ShopOrderTrackingStatusUnpaid
	case "AWAITING_SHIPMENT":
		*i = ShopOrderTrackingStatusAwaitingShipment
	case "AWAITING_COLLECTION":
		*i = ShopOrderTrackingStatusAwaitingCollection
	case "PARTIALLY_SHIPPING":
		*i = ShopOrderTrackingStatusPartiallyShipping
	case "SHIPPING":
		*i = ShopOrderTrackingStatusShipping
	case "CANCEL":
		*i = ShopOrderTrackingStatusCancel
	case "DELIVERED":
		*i = ShopOrderTrackingStatusDelivered
	case "COMPLETED":
		*i = ShopOrderTrackingStatusCompleted
	case "WILL_RETURN":
		*i = ShopOrderTrackingStatusWillReturn
	case "RETURNED":
		*i = ShopOrderTrackingStatusReturned
	default:
		*i = 0
	}
}

func (i ShopOrderTrackingStatus) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i ShopOrderTrackingStatus) GetValidValuesString() string {
	return _ShopOrderTrackingStatus_nameReadable
}

func (i ShopOrderTrackingStatus) String() string {
	i -= 1
	if i < 0 || i >= ShopOrderTrackingStatus(len(_ShopOrderTrackingStatus_index)-1) {
		return "ShopOrderTrackingStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _ShopOrderTrackingStatus_name[_ShopOrderTrackingStatus_index[i]:_ShopOrderTrackingStatus_index[i+1]]
}

func (i ShopOrderTrackingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *ShopOrderTrackingStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *ShopOrderTrackingStatus) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *ShopOrderTrackingStatus) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i ShopOrderTrackingStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func ShopOrderTrackingStatusP(v ShopOrderTrackingStatus) *ShopOrderTrackingStatus {
	return &v
}

func ListShopOrderTrackingStatus() []ShopOrderTrackingStatus {
	return []ShopOrderTrackingStatus{
		ShopOrderTrackingStatusUnpaid,
		ShopOrderTrackingStatusAwaitingShipment,
		ShopOrderTrackingStatusAwaitingCollection,
		ShopOrderTrackingStatusPartiallyShipping,
		ShopOrderTrackingStatusShipping,
		ShopOrderTrackingStatusCancel,
		ShopOrderTrackingStatusDelivered,
		ShopOrderTrackingStatusCompleted,
		ShopOrderTrackingStatusWillReturn,
		ShopOrderTrackingStatusReturned,
	}
}

func ListShopOrderTrackingStatusString() []string {
	return []string{
		ShopOrderTrackingStatusUnpaid.String(),
		ShopOrderTrackingStatusAwaitingShipment.String(),
		ShopOrderTrackingStatusAwaitingCollection.String(),
		ShopOrderTrackingStatusPartiallyShipping.String(),
		ShopOrderTrackingStatusShipping.String(),
		ShopOrderTrackingStatusCancel.String(),
		ShopOrderTrackingStatusDelivered.String(),
		ShopOrderTrackingStatusCompleted.String(),
		ShopOrderTrackingStatusWillReturn.String(),
		ShopOrderTrackingStatusReturned.String(),
	}
}