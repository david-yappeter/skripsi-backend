package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type TransactionResponse struct {
	Id               string                      `json:"id"`
	CashierSessionId string                      `json:"cashier_session_id"`
	Status           data_type.TransactionStatus `json:"status"`
	Total            float64                     `json:"total"`
	PaymentAt        data_type.NullDateTime      `json:"payment_at"`

	Timestamp
} // @name TransactionResponse

func NewTransactionResponse(transaction model.Transaction) TransactionResponse {
	r := TransactionResponse{
		Id:               transaction.Id,
		CashierSessionId: transaction.CashierSessionId,
		Status:           transaction.Status,
		Total:            transaction.Total,
		PaymentAt:        transaction.PaymentAt,
		Timestamp:        Timestamp(transaction.Timestamp),
	}

	return r
}
