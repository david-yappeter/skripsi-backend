package model

const DeliveryOrderItemTableName = "delivery_order_items"

type DeliveryOrderItem struct {
	Id              string  `db:"id"`
	DeliveryOrderId string  `db:"delivery_order_id"`
	ProductUnitId   string  `db:"product_unit_id"`
	UserId          string  `db:"user_id"`
	Qty             float64 `db:"qty"`
	PricePerUnit    float64 `db:"price_per_unit"`
	Timestamp

	ProductUnit            *ProductUnit            `db:"-"`
	DeliveryOrderItemCosts []DeliveryOrderItemCost `db:"-"`
}

func (m *DeliveryOrderItem) TableName() string {
	return DeliveryOrderItemTableName
}

func (m *DeliveryOrderItem) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"delivery_order_id": m.DeliveryOrderId,
		"product_unit_id":   m.ProductUnitId,
		"user_id":           m.UserId,
		"qty":               m.Qty,
		"price_per_unit":    m.PricePerUnit,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
