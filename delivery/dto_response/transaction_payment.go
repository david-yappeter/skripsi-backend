package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type TransactionPaymentResponse struct {
	Id              string                           `json:"id"`
	TransactionId   string                           `json:"transaction_id"`
	PaymentType     data_type.TransactionPaymentType `json:"payment_type"`
	ReferenceNumber *string                          `json:"reference_number"`
	Total           float64                          `json:"total"`

	Timestamp
} // @name TransactionPaymentResponse

func NewTransactionPaymentResponse(transactionPayment model.TransactionPayment) TransactionPaymentResponse {
	r := TransactionPaymentResponse{
		Id:              transactionPayment.Id,
		TransactionId:   transactionPayment.TransactionId,
		PaymentType:     transactionPayment.PaymentType,
		ReferenceNumber: transactionPayment.ReferenceNumber,
		Total:           transactionPayment.Total,
		Timestamp:       Timestamp(transactionPayment.Timestamp),
	}

	return r
}
