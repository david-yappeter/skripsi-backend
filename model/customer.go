package model

const CustomerTableName = "customers"

type Customer struct {
	Id       string  `db:"id"`
	Name     string  `db:"name"`
	Email    string  `db:"email"`
	Address  *string `db:"address"`
	Phone    string  `db:"phone"`
	IsActive bool    `db:"is_active"`

	Timestamp
}

func (m *Customer) TableName() string {
	return CustomerTableName
}

func (m *Customer) TableIds() []string {
	return []string{"id"}
}

func (m *Customer) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.Id,
		"name":       m.Name,
		"email":      m.Email,
		"address":    m.Address,
		"phone":      m.Phone,
		"is_active":  m.IsActive,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}

type CustomerQueryOption struct {
	QueryOption

	IsActive *bool
	Phrase   *string
}

var _ PrepareOption = &CustomerQueryOption{}

func (o *CustomerQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *CustomerQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
