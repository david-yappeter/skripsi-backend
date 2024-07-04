package model

import "math"

const DeliveryOrderItemTableName = "delivery_order_items"

type DeliveryOrderItem struct {
	Id              string  `db:"id"`
	DeliveryOrderId string  `db:"delivery_order_id"`
	ProductUnitId   string  `db:"product_unit_id"`
	Qty             float64 `db:"qty"`
	ScaleToBase     float64 `db:"scale_to_base"`
	PricePerUnit    float64 `db:"price_per_unit"`
	DiscountPerUnit float64 `db:"discount_per_unit"`
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
		"qty":               m.Qty,
		"scale_to_base":     m.ScaleToBase,
		"price_per_unit":    m.PricePerUnit,
		"discount_per_unit": m.DiscountPerUnit,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}

func (m DeliveryOrderItem) BaseQty() float64 {
	return m.Qty * m.ScaleToBase
}

func (m DeliveryOrderItem) Total() float64 {
	return m.BaseQty() * math.Max(m.PricePerUnit-m.DiscountPerUnit, 0)
}
