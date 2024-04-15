package model

const CustomerTypeTableName = "customer_types"

type CustomerType struct {
	Id          string  `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`

	Timestamp

	CustomerTypeDiscounts []CustomerTypeDiscount `db:"-"`
	Customers             []Customer             `db:"-"`
}

func (m *CustomerType) TableName() string {
	return CustomerTypeTableName
}

func (m *CustomerType) TableIds() []string {
	return []string{"id"}
}

func (m *CustomerType) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.Id,
		"name":        m.Name,
		"description": m.Description,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}

type CustomerTypeQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &CustomerTypeQueryOption{}

func (o *CustomerTypeQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"ct.*"}
	}
}

func (o *CustomerTypeQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
