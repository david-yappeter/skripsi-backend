package model

const PurchaseOrderItemTableName = "purchase_order_items"

type PurchaseOrderItem struct {
	Id              string  `db:"id"`
	PurchaseOrderId string  `db:"purchase_order_id"`
	ProductUnitId   string  `db:"product_unit_id"`
	UserId          string  `db:"user_id"`
	Qty             float64 `db:"qty"`
	ScaleToBase     float64 `db:"scale_to_base"`
	PricePerUnit    float64 `db:"price_per_unit"`
	Timestamp

	ProductUnit *ProductUnit `db:"-"`
	User        *User        `db:"-"`
}

func (m *PurchaseOrderItem) TableName() string {
	return PurchaseOrderItemTableName
}

func (m *PurchaseOrderItem) TableIds() []string {
	return []string{"id"}
}

func (m *PurchaseOrderItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"purchase_order_id": m.PurchaseOrderId,
		"product_unit_id":   m.ProductUnitId,
		"user_id":           m.UserId,
		"qty":               m.Qty,
		"scale_to_base":     m.ScaleToBase,
		"price_per_unit":    m.PricePerUnit,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}

func (m PurchaseOrderItem) BaseQty() float64 {
	return m.Qty * m.ScaleToBase
}
