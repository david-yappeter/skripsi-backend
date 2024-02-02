package model

const CartItemTableName = "cart_items"

type CartItem struct {
	Id            string  `db:"id"`
	CartId        string  `db:"cart_id"`
	ProductUnitId string  `db:"product_unit_id"`
	Qty           float64 `db:"qty"`
	Timestamp

	ProductUnit *ProductUnit `db:"-"`
}

func (m *CartItem) TableName() string {
	return CartItemTableName
}

func (m *CartItem) TableIds() []string {
	return []string{"id"}
}

func (m *CartItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              m.Id,
		"cart_id":         m.CartId,
		"product_unit_id": m.ProductUnitId,
		"qty":             m.Qty,
		"created_at":      m.CreatedAt,
		"updated_at":      m.UpdatedAt,
	}
}

type CartItemQueryOption struct {
	QueryOption

	CartId *string
	Phrase *string
}

var _ PrepareOption = &CartItemQueryOption{}

func (o *CartItemQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *CartItemQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
