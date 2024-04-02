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

func (m CartItem) Subtotal() float64 {
	subtotal := 0.0

	if m.ProductUnit != nil {
		productUnit := m.ProductUnit

		if productUnit.Product != nil {
			product := productUnit.Product

			if product.Price != nil {
				subtotal += *productUnit.Product.Price
			}

		}
	}

	subtotal *= m.Qty

	return subtotal
}

func (m CartItem) Discount() float64 {
	discount := 0.0

	if m.ProductUnit != nil {
		productUnit := m.ProductUnit

		if productUnit.Product != nil {
			product := productUnit.Product

			if product.Price != nil {
				if product.ProductDiscount != nil && m.Qty >= product.ProductDiscount.MinimumQty {
					productDiscount := product.ProductDiscount
					if productDiscount.DiscountAmount != nil {
						discount = *product.ProductDiscount.DiscountAmount
					} else {
						discount = *product.Price * *product.ProductDiscount.DiscountPercentage / 100.0
					}
				}
			}

		}
	}

	return discount
}

func (m CartItem) TotalDiscount() float64 {
	return m.Qty * m.Discount()
}

type CartItemQueryOption struct {
	QueryOption

	CartId *string
	Phrase *string
}

var _ PrepareOption = &CartItemQueryOption{}

func (o *CartItemQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"ci.*"}
	}
}

func (o *CartItemQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
