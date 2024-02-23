package model

import "myapp/data_type"

const TransactionPaymentTableName = "transaction_payments"

type TransactionPayment struct {
	Id              string                           `db:"id"`
	TransactionId   string                           `db:"transaction_id"`
	PaymentType     data_type.TransactionPaymentType `db:"payment_type"`
	ReferenceNumber *string                          `db:"reference_number"`
	Total           float64                          `db:"total"`

	Timestamp
}

func (m *TransactionPayment) TableName() string {
	return TransactionPaymentTableName
}

func (m *TransactionPayment) TableIds() []string {
	return []string{"id"}
}

func (m *TransactionPayment) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               m.Id,
		"transaction_id":   m.TransactionId,
		"payment_type":     m.PaymentType,
		"reference_number": m.ReferenceNumber,
		"total":            m.Total,
		"created_at":       m.CreatedAt,
		"updated_at":       m.UpdatedAt,
	}
}
