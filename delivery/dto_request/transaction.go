package dto_request

import "myapp/data_type"

type TransactionCheckoutCartRequest struct {
	PaymentType     data_type.TransactionPaymentType `json:"payment_type" validate:"data_type_enum"`
	CashPaid        *float64                         `json:"cash_paid" validate:"omitempty,gte=0"`
	ReferenceNumber *string                          `json:"reference_number" validate:"omitempty,not_empty"`
} // @name TransactionCheckoutCartRequest
