package model

const DeliveryOrderItemCostTableName = "delivery_order_item_costs"

type DeliveryOrderItemCost struct {
	Id                  string  `db:"id"`
	DeliveryOrderItemId string  `db:"delivery_order_item_id"`
	Qty                 float64 `db:"qty"`
	BaseCostPrice       float64 `db:"base_cost_price"`
	TotalCostPrice      float64 `db:"total_cost_price"`

	Timestamp
}

func (m *DeliveryOrderItemCost) TableName() string {
	return DeliveryOrderItemCostTableName
}

func (m *DeliveryOrderItemCost) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderItemCost) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                     m.Id,
		"delivery_order_item_id": m.DeliveryOrderItemId,
		"qty":                    m.Qty,
		"base_cost_price":        m.BaseCostPrice,
		"total_cost_price":       m.TotalCostPrice,
		"created_at":             m.CreatedAt,
		"updated_at":             m.UpdatedAt,
	}
}

type DeliveryOrderItemCostQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &DeliveryOrderItemCostQueryOption{}

func (o *DeliveryOrderItemCostQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"doic.*"}
	}
}

func (o *DeliveryOrderItemCostQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
