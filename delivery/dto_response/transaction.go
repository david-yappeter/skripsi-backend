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

	Items    []TransactionItemResponse    `json:"items" extensions:"x-nullable"`
	Payments []TransactionPaymentResponse `json:"payments" extensions:"x-nullable"`
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

	for _, transactionItem := range transaction.TransactionItems {
		r.Items = append(r.Items, NewTransactionItemResponse(transactionItem))
	}

	for _, transactionPayment := range transaction.TransactionPayments {
		r.Payments = append(r.Payments, NewTransactionPaymentResponse(transactionPayment))
	}

	return r
}

type TransactionSummaryResponse struct {
	Date            data_type.Date `json:"date"`
	TotalGrossSales float64        `json:"total_gross_sales"`
	TotalNetSales   float64        `json:"total_net_sales"`

	Transactions []TransactionResponse `json:"transactions"`
}

func NewTransactionSummaryResponse(transactionSummary model.TransactionSummary) TransactionSummaryResponse {
	r := TransactionSummaryResponse{
		Date:            transactionSummary.Date,
		TotalGrossSales: transactionSummary.TotalGrossSales,
		TotalNetSales:   transactionSummary.TotalNetSales,
	}

	for _, transaction := range transactionSummary.Transactions {
		r.Transactions = append(r.Transactions, NewTransactionResponse(transaction))
	}

	return r
}
