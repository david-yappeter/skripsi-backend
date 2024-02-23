package model

import "myapp/data_type"

const TransactionTableName = "transactions"

type Transaction struct {
	Id               string                      `db:"id"`
	CashierSessionId string                      `db:"cashier_session_id"`
	Status           data_type.TransactionStatus `db:"status"`
	Total            float64                     `db:"total"`
	PaymentAt        data_type.NullDateTime      `db:"payment_at"`

	Timestamp
}

func (m *Transaction) TableName() string {
	return TransactionTableName
}

func (m *Transaction) TableIds() []string {
	return []string{"id"}
}

func (m *Transaction) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                 m.Id,
		"cashier_session_id": m.CashierSessionId,
		"status":             m.Status,
		"total":              m.Total,
		"payment_at":         m.PaymentAt,
		"created_at":         m.CreatedAt,
		"updated_at":         m.UpdatedAt,
	}
}

type TransactionQueryOption struct {
	QueryOption

	CashierSessionId *string
	Status           *data_type.TransactionStatus
}

var _ PrepareOption = &TransactionQueryOption{}

func (o *TransactionQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"t.*"}
	}
}

func (o *TransactionQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
