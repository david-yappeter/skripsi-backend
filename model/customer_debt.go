package model

import "myapp/data_type"

const CustomerDebtTableName = "customer_debts"

type CustomerDebt struct {
	Id              string             `db:"id"`
	CustomerId      string             `db:"customer_id"`
	DueDate         data_type.NullDate `db:"due_date"`
	Status          string             `db:"status"`
	Amount          float64            `db:"amount"`
	RemainingAmount float64            `db:"remaining_amount"`

	Timestamp
}

func (m *CustomerDebt) TableName() string {
	return CustomerDebtTableName
}

func (m *CustomerDebt) TableIds() []string {
	return []string{"id"}
}

func (m *CustomerDebt) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               m.Id,
		"customer_id":      m.CustomerId,
		"due_date":         m.DueDate,
		"status":           m.Status,
		"amount":           m.Amount,
		"remaining_amount": m.RemainingAmount,
		"created_at":       m.CreatedAt,
		"updated_at":       m.UpdatedAt,
	}
}

type CustomerDebtQueryOption struct {
	QueryOption

	CustomerId *string
	Phrase     *string
}

var _ PrepareOption = &CustomerDebtQueryOption{}

func (o *CustomerDebtQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"cd.*"}
	}
}

func (o *CustomerDebtQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
