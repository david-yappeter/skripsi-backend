// Code generated by "go run myapp/tool/stringer -linecomment -type=ProductStockMutationType -output=product_stock_mutation_enum_gen.go -swagoutput=../tool/swag/enum_gen/product_stock_mutation_enum_gen.go -custom"; DO NOT EDIT.

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
	_ = x[ProductStockMutationTypeProductReceiveItem-1]
	_ = x[ProductStockMutationTypeDeliveryOrderItemCostCancel-2]
	_ = x[ProductStockMutationTypeProductStockAdjustment-3]
	_ = x[ProductStockMutationTypeDeliveryOrderItemReturned-4]
}

const _ProductStockMutationType_nameReadable = "PRODUCT_RECEIVE_ITEM, DELIVERY_ORDER_ITEM_COST_CANCEL, PRODUCT_STOCK_ADJUSTMENT, DELIVERY_ORDER_ITEM_RETURNED"

const _ProductStockMutationType_name = "PRODUCT_RECEIVE_ITEMDELIVERY_ORDER_ITEM_COST_CANCELPRODUCT_STOCK_ADJUSTMENTDELIVERY_ORDER_ITEM_RETURNED"

var _ProductStockMutationType_index = [...]uint8{0, 20, 51, 75, 103}

func (i *ProductStockMutationType) Determine(s string) {
	switch s {
	case "PRODUCT_RECEIVE_ITEM":
		*i = ProductStockMutationTypeProductReceiveItem
	case "DELIVERY_ORDER_ITEM_COST_CANCEL":
		*i = ProductStockMutationTypeDeliveryOrderItemCostCancel
	case "PRODUCT_STOCK_ADJUSTMENT":
		*i = ProductStockMutationTypeProductStockAdjustment
	case "DELIVERY_ORDER_ITEM_RETURNED":
		*i = ProductStockMutationTypeDeliveryOrderItemReturned
	default:
		*i = 0
	}
}

func (i ProductStockMutationType) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i ProductStockMutationType) GetValidValuesString() string {
	return _ProductStockMutationType_nameReadable
}

func (i ProductStockMutationType) String() string {
	i -= 1
	if i < 0 || i >= ProductStockMutationType(len(_ProductStockMutationType_index)-1) {
		return "ProductStockMutationType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _ProductStockMutationType_name[_ProductStockMutationType_index[i]:_ProductStockMutationType_index[i+1]]
}

func (i ProductStockMutationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *ProductStockMutationType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *ProductStockMutationType) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *ProductStockMutationType) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i ProductStockMutationType) Value() (driver.Value, error) {
	return i.String(), nil
}

func ProductStockMutationTypeP(v ProductStockMutationType) *ProductStockMutationType {
	return &v
}

func ListProductStockMutationType() []ProductStockMutationType {
	return []ProductStockMutationType{
		ProductStockMutationTypeProductReceiveItem,
		ProductStockMutationTypeDeliveryOrderItemCostCancel,
		ProductStockMutationTypeProductStockAdjustment,
		ProductStockMutationTypeDeliveryOrderItemReturned,
	}
}

func ListProductStockMutationTypeString() []string {
	return []string{
		ProductStockMutationTypeProductReceiveItem.String(),
		ProductStockMutationTypeDeliveryOrderItemCostCancel.String(),
		ProductStockMutationTypeProductStockAdjustment.String(),
		ProductStockMutationTypeDeliveryOrderItemReturned.String(),
	}
}
