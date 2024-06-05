// Code generated by "go run myapp/tool/stringer -linecomment -type=PurchaseOrderStatus -output=purchase_order_enum_gen.go -swagoutput=../tool/swag/enum_gen/purchase_order_enum_gen.go -custom"; DO NOT EDIT.

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
	_ = x[PurchaseOrderStatusPending-1]
	_ = x[PurchaseOrderStatusOngoing-2]
	_ = x[PurchaseOrderStatusCanceled-3]
	_ = x[PurchaseOrderStatusCompleted-4]
}

const _PurchaseOrderStatus_nameReadable = "PENDING, ONGOING, CANCELED, COMPLETED"

const _PurchaseOrderStatus_name = "PENDINGONGOINGCANCELEDCOMPLETED"

var _PurchaseOrderStatus_index = [...]uint8{0, 7, 14, 22, 31}

func (i *PurchaseOrderStatus) Determine(s string) {
	switch s {
	case "PENDING":
		*i = PurchaseOrderStatusPending
	case "ONGOING":
		*i = PurchaseOrderStatusOngoing
	case "CANCELED":
		*i = PurchaseOrderStatusCanceled
	case "COMPLETED":
		*i = PurchaseOrderStatusCompleted
	default:
		*i = 0
	}
}

func (i PurchaseOrderStatus) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i PurchaseOrderStatus) GetValidValuesString() string {
	return _PurchaseOrderStatus_nameReadable
}

func (i PurchaseOrderStatus) String() string {
	i -= 1
	if i < 0 || i >= PurchaseOrderStatus(len(_PurchaseOrderStatus_index)-1) {
		return "PurchaseOrderStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _PurchaseOrderStatus_name[_PurchaseOrderStatus_index[i]:_PurchaseOrderStatus_index[i+1]]
}

func (i PurchaseOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *PurchaseOrderStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *PurchaseOrderStatus) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *PurchaseOrderStatus) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i PurchaseOrderStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func PurchaseOrderStatusP(v PurchaseOrderStatus) *PurchaseOrderStatus {
	return &v
}

func ListPurchaseOrderStatus() []PurchaseOrderStatus {
	return []PurchaseOrderStatus{
		PurchaseOrderStatusPending,
		PurchaseOrderStatusOngoing,
		PurchaseOrderStatusCanceled,
		PurchaseOrderStatusCompleted,
	}
}

func ListPurchaseOrderStatusString() []string {
	return []string{
		PurchaseOrderStatusPending.String(),
		PurchaseOrderStatusOngoing.String(),
		PurchaseOrderStatusCanceled.String(),
		PurchaseOrderStatusCompleted.String(),
	}
}