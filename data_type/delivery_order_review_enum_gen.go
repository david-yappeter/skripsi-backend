// Code generated by "go run myapp/tool/stringer -linecomment -type=DeliveryOrderReviewType -output=delivery_order_review_enum_gen.go -swagoutput=../tool/swag/enum_gen/delivery_order_review_enum_gen.go -custom"; DO NOT EDIT.

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
	_ = x[DeliveryOrderReviewTypeProduct-1]
	_ = x[DeliveryOrderReviewTypeDelivery-2]
}

const _DeliveryOrderReviewType_nameReadable = "PRODUCT, DELIVERY"

const _DeliveryOrderReviewType_name = "PRODUCTDELIVERY"

var _DeliveryOrderReviewType_index = [...]uint8{0, 7, 15}

func (i *DeliveryOrderReviewType) Determine(s string) {
	switch s {
	case "PRODUCT":
		*i = DeliveryOrderReviewTypeProduct
	case "DELIVERY":
		*i = DeliveryOrderReviewTypeDelivery
	default:
		*i = 0
	}
}

func (i DeliveryOrderReviewType) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i DeliveryOrderReviewType) GetValidValuesString() string {
	return _DeliveryOrderReviewType_nameReadable
}

func (i DeliveryOrderReviewType) String() string {
	i -= 1
	if i < 0 || i >= DeliveryOrderReviewType(len(_DeliveryOrderReviewType_index)-1) {
		return "DeliveryOrderReviewType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _DeliveryOrderReviewType_name[_DeliveryOrderReviewType_index[i]:_DeliveryOrderReviewType_index[i+1]]
}

func (i DeliveryOrderReviewType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *DeliveryOrderReviewType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *DeliveryOrderReviewType) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *DeliveryOrderReviewType) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i DeliveryOrderReviewType) Value() (driver.Value, error) {
	return i.String(), nil
}

func DeliveryOrderReviewTypeP(v DeliveryOrderReviewType) *DeliveryOrderReviewType {
	return &v
}

func ListDeliveryOrderReviewType() []DeliveryOrderReviewType {
	return []DeliveryOrderReviewType{
		DeliveryOrderReviewTypeProduct,
		DeliveryOrderReviewTypeDelivery,
	}
}

func ListDeliveryOrderReviewTypeString() []string {
	return []string{
		DeliveryOrderReviewTypeProduct.String(),
		DeliveryOrderReviewTypeDelivery.String(),
	}
}
