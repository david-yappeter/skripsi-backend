// Code generated by "go run myapp/tool/stringer -linecomment -type=TiktokProductStatus,TiktokProductDimensionUnit,TiktokProductPackageWeight -output=tiktok_product_enum_gen.go -swagoutput=../tool/swag/enum_gen/tiktok_product_enum_gen.go -custom"; DO NOT EDIT.

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
	_ = x[TiktokProductStatusActive-1]
	_ = x[TiktokProductStatusInActive-2]
}

const _TiktokProductStatus_nameReadable = "ACTIVE, IN_ACTIVE"

const _TiktokProductStatus_name = "ACTIVEIN_ACTIVE"

var _TiktokProductStatus_index = [...]uint8{0, 6, 15}

func (i *TiktokProductStatus) Determine(s string) {
	switch s {
	case "ACTIVE":
		*i = TiktokProductStatusActive
	case "IN_ACTIVE":
		*i = TiktokProductStatusInActive
	default:
		*i = 0
	}
}

func (i TiktokProductStatus) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i TiktokProductStatus) GetValidValuesString() string {
	return _TiktokProductStatus_nameReadable
}

func (i TiktokProductStatus) String() string {
	i -= 1
	if i < 0 || i >= TiktokProductStatus(len(_TiktokProductStatus_index)-1) {
		return "TiktokProductStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _TiktokProductStatus_name[_TiktokProductStatus_index[i]:_TiktokProductStatus_index[i+1]]
}

func (i TiktokProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *TiktokProductStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *TiktokProductStatus) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *TiktokProductStatus) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i TiktokProductStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func TiktokProductStatusP(v TiktokProductStatus) *TiktokProductStatus {
	return &v
}

func ListTiktokProductStatus() []TiktokProductStatus {
	return []TiktokProductStatus{
		TiktokProductStatusActive,
		TiktokProductStatusInActive,
	}
}

func ListTiktokProductStatusString() []string {
	return []string{
		TiktokProductStatusActive.String(),
		TiktokProductStatusInActive.String(),
	}
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TiktokProductDimensionUnitCentimeter-1]
}

const _TiktokProductDimensionUnit_nameReadable = "CENTIMETER"

const _TiktokProductDimensionUnit_name = "CENTIMETER"

var _TiktokProductDimensionUnit_index = [...]uint8{0, 10}

func (i *TiktokProductDimensionUnit) Determine(s string) {
	switch s {
	case "CENTIMETER":
		*i = TiktokProductDimensionUnitCentimeter
	default:
		*i = 0
	}
}

func (i TiktokProductDimensionUnit) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i TiktokProductDimensionUnit) GetValidValuesString() string {
	return _TiktokProductDimensionUnit_nameReadable
}

func (i TiktokProductDimensionUnit) String() string {
	i -= 1
	if i < 0 || i >= TiktokProductDimensionUnit(len(_TiktokProductDimensionUnit_index)-1) {
		return "TiktokProductDimensionUnit(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _TiktokProductDimensionUnit_name[_TiktokProductDimensionUnit_index[i]:_TiktokProductDimensionUnit_index[i+1]]
}

func (i TiktokProductDimensionUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *TiktokProductDimensionUnit) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *TiktokProductDimensionUnit) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *TiktokProductDimensionUnit) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i TiktokProductDimensionUnit) Value() (driver.Value, error) {
	return i.String(), nil
}

func TiktokProductDimensionUnitP(v TiktokProductDimensionUnit) *TiktokProductDimensionUnit {
	return &v
}

func ListTiktokProductDimensionUnit() []TiktokProductDimensionUnit {
	return []TiktokProductDimensionUnit{
		TiktokProductDimensionUnitCentimeter,
	}
}

func ListTiktokProductDimensionUnitString() []string {
	return []string{
		TiktokProductDimensionUnitCentimeter.String(),
	}
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TiktokProductPackageWeightKilogram-1]
}

const _TiktokProductPackageWeight_nameReadable = "KILOGRAM"

const _TiktokProductPackageWeight_name = "KILOGRAM"

var _TiktokProductPackageWeight_index = [...]uint8{0, 8}

func (i *TiktokProductPackageWeight) Determine(s string) {
	switch s {
	case "KILOGRAM":
		*i = TiktokProductPackageWeightKilogram
	default:
		*i = 0
	}
}

func (i TiktokProductPackageWeight) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i TiktokProductPackageWeight) GetValidValuesString() string {
	return _TiktokProductPackageWeight_nameReadable
}

func (i TiktokProductPackageWeight) String() string {
	i -= 1
	if i < 0 || i >= TiktokProductPackageWeight(len(_TiktokProductPackageWeight_index)-1) {
		return "TiktokProductPackageWeight(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _TiktokProductPackageWeight_name[_TiktokProductPackageWeight_index[i]:_TiktokProductPackageWeight_index[i+1]]
}

func (i TiktokProductPackageWeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *TiktokProductPackageWeight) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *TiktokProductPackageWeight) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *TiktokProductPackageWeight) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i TiktokProductPackageWeight) Value() (driver.Value, error) {
	return i.String(), nil
}

func TiktokProductPackageWeightP(v TiktokProductPackageWeight) *TiktokProductPackageWeight {
	return &v
}

func ListTiktokProductPackageWeight() []TiktokProductPackageWeight {
	return []TiktokProductPackageWeight{
		TiktokProductPackageWeightKilogram,
	}
}

func ListTiktokProductPackageWeightString() []string {
	return []string{
		TiktokProductPackageWeightKilogram.String(),
	}
}
