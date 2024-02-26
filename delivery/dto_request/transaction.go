package dto_request

import "myapp/data_type"

type TransactionCheckoutCartRequest struct {
	PaymentType     data_type.TransactionPaymentType `json:"payment_type" validate:"required,data_type_enum"`
	ReferenceNumber *string                          `json:"reference_number" validate:"omitempty,not_empty"`
} // @name TransactionCheckoutCartRequest
