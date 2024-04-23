package model

const ProductDiscountTableName = "product_discounts"

type ProductDiscount struct {
	Id                 string   `db:"id"`
	ProductId          string   `db:"product_id"`
	MinimumQty         float64  `db:"minimum_qty"`
	IsActive           bool     `db:"is_active"`
	DiscountPercentage *float64 `db:"discount_percentage"`
	DiscountAmount     *float64 `db:"discount_amount"`

	Timestamp

	Product *Product `db:"-"`
}

func (m *ProductDiscount) TableName() string {
	return ProductDiscountTableName
}

func (m *ProductDiscount) TableIds() []string {
	return []string{"id"}
}

func (m *ProductDiscount) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                  m.Id,
		"product_id":          m.ProductId,
		"minimum_qty":         m.MinimumQty,
		"is_active":           m.IsActive,
		"discount_percentage": m.DiscountPercentage,
		"discount_amount":     m.DiscountAmount,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
	}
}

type ProductDiscountQueryOption struct {
	QueryOption

	ProductId *string
	IsActive  *bool
	Phrase    *string
}

var _ PrepareOption = &ProductDiscountQueryOption{}

func (o *ProductDiscountQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"pd.*"}
	}
}

func (o *ProductDiscountQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
