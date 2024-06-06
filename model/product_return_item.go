package model

const ProductReturnItemTableName = "product_return_items"

type ProductReturnItem struct {
	Id              string  `db:"id"`
	ProductReturnId string  `db:"product_return_id"`
	ProductUnitId   string  `db:"product_unit_id"`
	Qty             float64 `db:"qty"`
	ScaleToBase     float64 `db:"scale_to_base"`
	BaseCostPrice   float64 `db:"base_cost_price"`
	Timestamp

	User        *User        `db:"-"`
	ProductUnit *ProductUnit `db:"-"`
}

func (m *ProductReturnItem) TableName() string {
	return ProductReturnItemTableName
}

func (m *ProductReturnItem) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReturnItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"product_return_id": m.ProductReturnId,
		"product_unit_id":   m.ProductUnitId,
		"qty":               m.Qty,
		"scale_to_base":     m.ScaleToBase,
		"base_cost_price":   m.BaseCostPrice,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}

func (m ProductReturnItem) BaseQty() float64 {
	return m.Qty * m.ScaleToBase
}
