package model

import "myapp/data_type"

const DebtTableName = "debts"

type Debt struct {
	Id                   string               `db:"id"`
	DebtSource           data_type.DebtSource `db:"debt_source"`
	DebtSourceIdentifier string               `db:"debt_source_identifier"`
	DueDate              data_type.NullDate   `db:"due_date"`
	Status               data_type.DebtStatus `db:"status"`
	Amount               float64              `db:"amount"`
	RemainingAmount      float64              `db:"remaining_amount"`

	Timestamp
	DebtPayments   []DebtPayment   `db:"-"`
	Supplier       *Supplier       `db:"-"`
	ProductReceive *ProductReceive `db:"-"` // if DebtSource is DebtSourceProductReceive
}

func (m *Debt) TableName() string {
	return DebtTableName
}

func (m *Debt) TableIds() []string {
	return []string{"id"}
}

func (m *Debt) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                     m.Id,
		"debt_source":            m.DebtSource,
		"debt_source_identifier": m.DebtSourceIdentifier,
		"due_date":               m.DueDate,
		"status":                 m.Status,
		"amount":                 m.Amount,
		"remaining_amount":       m.RemainingAmount,
		"created_at":             m.CreatedAt,
		"updated_at":             m.UpdatedAt,
	}
}

type DebtQueryOption struct {
	QueryOption

	Id     *string
	Status *data_type.DebtStatus
	Phrase *string
}

var _ PrepareOption = &DebtQueryOption{}

func (o *DebtQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"d.*"}
	}
}

func (o *DebtQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
