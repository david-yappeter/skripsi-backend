package model

const CustomerTypeDiscountTableName = "customer_type_discounts"

type CustomerTypeDiscount struct {
	Id                 string   `db:"id"`
	ProductId          string   `db:"product_id"`
	CustomerTypeId     string   `db:"customer_type_id"`
	IsActive           bool     `db:"is_active"`
	DiscountPercentage *float64 `db:"discount_percentage"`
	DiscountAmount     *float64 `db:"discount_amount"`

	Timestamp

	Product *Product `db:"-"`
}

func (m *CustomerTypeDiscount) TableName() string {
	return CustomerTypeDiscountTableName
}

func (m *CustomerTypeDiscount) TableIds() []string {
	return []string{"id"}
}

func (m *CustomerTypeDiscount) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                  m.Id,
		"product_id":          m.ProductId,
		"customer_type_id":    m.CustomerTypeId,
		"is_active":           m.IsActive,
		"discount_percentage": m.DiscountPercentage,
		"discount_amount":     m.DiscountAmount,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
	}
}

type CustomerTypeDiscountQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &CustomerTypeDiscountQueryOption{}

func (o *CustomerTypeDiscountQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"ctd.*"}
	}
}

func (o *CustomerTypeDiscountQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
