// Code generated by "go run myapp/tool/stringer -linecomment -type=TransactionStatus -output=transaction_enum_gen.go -swagoutput=../tool/swag/enum_gen/transaction_enum_gen.go -custom"; DO NOT EDIT.

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
	_ = x[TransactionStatusUnpaid-1]
	_ = x[TransactionStatusPaid-2]
}

const _TransactionStatus_nameReadable = "UNPAID, PAID"

const _TransactionStatus_name = "UNPAIDPAID"

var _TransactionStatus_index = [...]uint8{0, 6, 10}

func (i *TransactionStatus) Determine(s string) {
	switch s {
	case "UNPAID":
		*i = TransactionStatusUnpaid
	case "PAID":
		*i = TransactionStatusPaid
	default:
		*i = 0
	}
}

func (i TransactionStatus) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i TransactionStatus) GetValidValuesString() string {
	return _TransactionStatus_nameReadable
}

func (i TransactionStatus) String() string {
	i -= 1
	if i < 0 || i >= TransactionStatus(len(_TransactionStatus_index)-1) {
		return "TransactionStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _TransactionStatus_name[_TransactionStatus_index[i]:_TransactionStatus_index[i+1]]
}

func (i TransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *TransactionStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *TransactionStatus) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *TransactionStatus) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i TransactionStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func TransactionStatusP(v TransactionStatus) *TransactionStatus {
	return &v
}

func ListTransactionStatus() []TransactionStatus {
	return []TransactionStatus{
		TransactionStatusUnpaid,
		TransactionStatusPaid,
	}
}

func ListTransactionStatusString() []string {
	return []string{
		TransactionStatusUnpaid.String(),
		TransactionStatusPaid.String(),
	}
}
