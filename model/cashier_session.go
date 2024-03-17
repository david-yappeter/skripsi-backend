package model

import "myapp/data_type"

const CashierSessionTableName = "cashier_sessions"

type CashierSession struct {
	Id           string                         `db:"id"`
	UserId       string                         `db:"user_id"`
	Status       data_type.CashierSessionStatus `db:"status"`
	StartingCash float64                        `db:"starting_cash"`
	EndingCash   *float64                       `db:"ending_cash"`
	StartedAt    data_type.DateTime             `db:"started_at"`
	EndedAt      data_type.NullDateTime         `db:"ended_at"`
	Timestamp

	User *User `db:"-"`
}

func (m *CashierSession) TableName() string {
	return CashierSessionTableName
}

func (m *CashierSession) TableIds() []string {
	return []string{"id"}
}

func (m *CashierSession) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.Id,
		"user_id":       m.UserId,
		"status":        m.Status,
		"starting_cash": m.StartingCash,
		"ending_cash":   m.EndingCash,
		"started_at":    m.StartedAt,
		"ended_at":      m.EndedAt,
		"created_at":    m.CreatedAt,
		"updated_at":    m.UpdatedAt,
	}
}

type CashierSessionQueryOption struct {
	QueryOption

	StartedAtLte data_type.NullDateTime
	EndedAtGte   data_type.NullDateTime
	UserId       *string
	Status       *data_type.CashierSessionStatus
	Phrase       *string
}

var _ PrepareOption = &CashierSessionQueryOption{}

func (o *CashierSessionQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"cs.*"}
	}
}

func (o *CashierSessionQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
