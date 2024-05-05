// Code generated by "go run myapp/tool/stringer -linecomment -type=FileType -output=file_enum_gen.go -swagoutput=../tool/swag/enum_gen/file_enum_gen.go -custom"; DO NOT EDIT.

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
	_ = x[FileTypeProductImage-1]
	_ = x[FileTypeProductReceiveImage-2]
	_ = x[FileTypeDeliveryOrderImage-3]
	_ = x[FileTypeCustomerPaymentImage-4]
	_ = x[FileTypeDebtPaymentImage-5]
	_ = x[FileTypeDeliveryOrderReturnImage-6]
	_ = x[FileTypeProductReceiveReturnImage-7]
}

const _FileType_nameReadable = "PRODUCT_IMAGE, PRODUCT_RECEIVE_IMAGE, DELIVERY_ORDER_IMAGE, CUSTOMER_PAYMENT_IMAGE, DEBT_PAYMENT_IMAGE, DELIVERY_ORDER_RETURN_IMAGE, PRODUCT_RECEIVE_RETURN_IMAGE"

const _FileType_name = "PRODUCT_IMAGEPRODUCT_RECEIVE_IMAGEDELIVERY_ORDER_IMAGECUSTOMER_PAYMENT_IMAGEDEBT_PAYMENT_IMAGEDELIVERY_ORDER_RETURN_IMAGEPRODUCT_RECEIVE_RETURN_IMAGE"

var _FileType_index = [...]uint8{0, 13, 34, 54, 76, 94, 121, 149}

func (i *FileType) Determine(s string) {
	switch s {
	case "PRODUCT_IMAGE":
		*i = FileTypeProductImage
	case "PRODUCT_RECEIVE_IMAGE":
		*i = FileTypeProductReceiveImage
	case "DELIVERY_ORDER_IMAGE":
		*i = FileTypeDeliveryOrderImage
	case "CUSTOMER_PAYMENT_IMAGE":
		*i = FileTypeCustomerPaymentImage
	case "DEBT_PAYMENT_IMAGE":
		*i = FileTypeDebtPaymentImage
	case "DELIVERY_ORDER_RETURN_IMAGE":
		*i = FileTypeDeliveryOrderReturnImage
	case "PRODUCT_RECEIVE_RETURN_IMAGE":
		*i = FileTypeProductReceiveReturnImage
	default:
		*i = 0
	}
}

func (i FileType) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i FileType) GetValidValuesString() string {
	return _FileType_nameReadable
}

func (i FileType) String() string {
	i -= 1
	if i < 0 || i >= FileType(len(_FileType_index)-1) {
		return "FileType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _FileType_name[_FileType_index[i]:_FileType_index[i+1]]
}

func (i FileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *FileType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.Determine(s)

	return nil
}

func (i *FileType) UnmarshalText(b []byte) error {
	i.Determine(string(b))

	return nil
}

func (i *FileType) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.Determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i FileType) Value() (driver.Value, error) {
	return i.String(), nil
}

func FileTypeP(v FileType) *FileType {
	return &v
}

func ListFileType() []FileType {
	return []FileType{
		FileTypeProductImage,
		FileTypeProductReceiveImage,
		FileTypeDeliveryOrderImage,
		FileTypeCustomerPaymentImage,
		FileTypeDebtPaymentImage,
		FileTypeDeliveryOrderReturnImage,
		FileTypeProductReceiveReturnImage,
	}
}

func ListFileTypeString() []string {
	return []string{
		FileTypeProductImage.String(),
		FileTypeProductReceiveImage.String(),
		FileTypeDeliveryOrderImage.String(),
		FileTypeCustomerPaymentImage.String(),
		FileTypeDebtPaymentImage.String(),
		FileTypeDeliveryOrderReturnImage.String(),
		FileTypeProductReceiveReturnImage.String(),
	}
}
