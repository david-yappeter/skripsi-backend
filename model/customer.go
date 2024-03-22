package model

const CustomerTableName = "customers"

type Customer struct {
	Id             string  `db:"id"`
	CustomerTypeId *string `db:"customer_type_id"`
	Name           string  `db:"name"`
	Email          string  `db:"email"`
	Address        string  `db:"address"`
	Latitude       float64 `db:"latitude"`
	Longitude      float64 `db:"longitude"`
	Phone          string  `db:"phone"`
	IsActive       bool    `db:"is_active"`

	Timestamp

	CustomerType *CustomerType `db:"-"`
}

func (m *Customer) TableName() string {
	return CustomerTableName
}

func (m *Customer) TableIds() []string {
	return []string{"id"}
}

func (m *Customer) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               m.Id,
		"customer_type_id": m.CustomerTypeId,
		"name":             m.Name,
		"email":            m.Email,
		"address":          m.Address,
		"latitude":         m.Latitude,
		"longitude":        m.Longitude,
		"phone":            m.Phone,
		"is_active":        m.IsActive,
		"created_at":       m.CreatedAt,
		"updated_at":       m.UpdatedAt,
	}
}

type CustomerQueryOption struct {
	QueryOption

	CustomerTypeId *string
	IsActive       *bool
	Phrase         *string
}

var _ PrepareOption = &CustomerQueryOption{}

func (o *CustomerQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"c.*"}
	}
}

func (o *CustomerQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
