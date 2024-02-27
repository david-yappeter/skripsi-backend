package model

import "myapp/data_type"

const ProductStockMutationTableName = "product_stock_mutations"

type ProductStockMutation struct {
	Id            string                             `db:"id"`
	ProductUnitId string                             `db:"product_unit_id"`
	Type          data_type.ProductStockMutationType `db:"type"`
	IdentifierId  string                             `db:"identifier_id"`
	Qty           float64                            `db:"qty"`
	ScaleToBase   float64                            `db:"scale_to_base"`
	BaseQtyLeft   float64                            `db:"base_qty_left"`
	BaseCostPrice float64                            `db:"base_cost_price"`
	MutatedAt     data_type.DateTime                 `db:"mutated_at"`

	Timestamp
}

func (m *ProductStockMutation) TableName() string {
	return ProductStockMutationTableName
}

func (m *ProductStockMutation) TableIds() []string {
	return []string{"id"}
}

func (m *ProductStockMutation) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              m.Id,
		"product_unit_id": m.ProductUnitId,
		"type":            m.Type,
		"identifier_id":   m.IdentifierId,
		"qty":             m.Qty,
		"scale_to_base":   m.ScaleToBase,
		"base_qty_left":   m.BaseQtyLeft,
		"base_cost_price": m.BaseCostPrice,
		"mutated_at":      m.MutatedAt,
		"created_at":      m.CreatedAt,
		"updated_at":      m.UpdatedAt,
	}
}

type ProductStockMutationQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &ProductStockMutationQueryOption{}

func (o *ProductStockMutationQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"psm.*"}
	}
}

func (o *ProductStockMutationQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
