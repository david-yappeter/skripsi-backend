package model

const CartTableName = "carts"

type Cart struct {
	Id               string  `db:"id"`
	CashierSessionId string  `db:"cashier_session_id"`
	Name             *string `db:"name"`
	IsActive         bool    `db:"is_active"`
	Timestamp

	CashierSession *CashierSession `db:"-"`
	CartItems      []CartItem      `db:"-"`
}

func (m *Cart) TableName() string {
	return CartTableName
}

func (m *Cart) TableIds() []string {
	return []string{"id"}
}

func (m *Cart) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                 m.Id,
		"cashier_session_id": m.CashierSessionId,
		"name":               m.Name,
		"is_active":          m.IsActive,
		"created_at":         m.CreatedAt,
		"updated_at":         m.UpdatedAt,
	}
}

func (m *Cart) Subtotal() float64 {
	subtotal := 0.0

	for _, cartItem := range m.CartItems {
		subtotal += cartItem.Subtotal()
	}

	return subtotal
}

func (m *Cart) TotalDiscount() float64 {
	discount := 0.0

	for _, cartItem := range m.CartItems {
		discount += cartItem.TotalDiscount()
	}

	return discount
}

type CartQueryOption struct {
	QueryOption

	CashierSessionId *string
	IsActive         *bool
	Phrase           *string
}

var _ PrepareOption = &CartQueryOption{}

func (o *CartQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"c.*"}
	}
}

func (o *CartQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
