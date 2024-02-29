package model

import "myapp/data_type"

const DebtPaymentTableName = "debt_payments"

type DebtPayment struct {
	Id          string             `db:"id"`
	UserId      string             `db:"user_id"`
	ImageFileId string             `db:"image_file_id"`
	DebtId      string             `db:"debt_id"`
	Amount      float64            `db:"amount"`
	Description *string            `db:"description"`
	PaidAt      data_type.DateTime `db:"paid_at"`

	Timestamp
}

func (m *DebtPayment) TableName() string {
	return DebtPaymentTableName
}

func (m *DebtPayment) TableIds() []string {
	return []string{"id"}
}

func (m *DebtPayment) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.Id,
		"user_id":       m.UserId,
		"image_file_id": m.ImageFileId,
		"debt_id":       m.DebtId,
		"amount":        m.Amount,
		"description":   m.Description,
		"paid_at":       m.PaidAt,
		"created_at":    m.CreatedAt,
		"updated_at":    m.UpdatedAt,
	}
}

type DebtPaymentQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &DebtPaymentQueryOption{}

func (o *DebtPaymentQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"dp.*"}
	}
}

func (o *DebtPaymentQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
