package model

const TransactionItemTableName = "transaction_items"

type TransactionItem struct {
	Id              string   `db:"id"`
	TransactionId   string   `db:"transaction_id"`
	ProductUnitId   string   `db:"product_unit_id"`
	Qty             float64  `db:"qty"`
	PricePerUnit    float64  `db:"price_per_unit"`
	DiscountPerUnit *float64 `db:"discount_per_unit"`
	Timestamp

	ProductUnit *ProductUnit `db:"-"`
}

func (m *TransactionItem) TableName() string {
	return TransactionItemTableName
}

func (m *TransactionItem) TableIds() []string {
	return []string{"id"}
}

func (m *TransactionItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"transaction_id":    m.TransactionId,
		"product_unit_id":   m.ProductUnitId,
		"qty":               m.Qty,
		"price_per_unit":    m.PricePerUnit,
		"discount_per_unit": m.DiscountPerUnit,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
