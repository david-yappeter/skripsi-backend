package model

import "myapp/data_type"

const CustomerPaymentTableName = "customer_payments"

type CustomerPayment struct {
	Id             string             `db:"id"`
	UserId         string             `db:"user_id"`
	ImageFileId    string             `db:"image_file_id"`
	CustomerDebtId string             `db:"customer_debt_id"`
	Amount         float64            `db:"amount"`
	Description    *string            `db:"description"`
	PaidAt         data_type.DateTime `db:"paid_at"`

	Timestamp
}

func (m *CustomerPayment) TableName() string {
	return CustomerPaymentTableName
}

func (m *CustomerPayment) TableIds() []string {
	return []string{"id"}
}

func (m *CustomerPayment) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               m.Id,
		"user_id":          m.UserId,
		"image_file_id":    m.ImageFileId,
		"customer_debt_id": m.CustomerDebtId,
		"amount":           m.Amount,
		"description":      m.Description,
		"paid_at":          m.PaidAt,
		"created_at":       m.CreatedAt,
		"updated_at":       m.UpdatedAt,
	}
}

type CustomerPaymentQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &CustomerPaymentQueryOption{}

func (o *CustomerPaymentQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"cp.*"}
	}
}

func (o *CustomerPaymentQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
