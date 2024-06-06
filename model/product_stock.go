package model

const ProductStockTableName = "product_stocks"

type ProductStock struct {
	Id            string  `db:"id"`
	ProductId     string  `db:"product_id"`
	Qty           float64 `db:"qty"`
	BaseCostPrice float64 `db:"base_cost_price"`
	Timestamp

	Product *Product `db:"-"`
}

func (m *ProductStock) TableName() string {
	return ProductStockTableName
}

func (m *ProductStock) TableIds() []string {
	return []string{"id"}
}

func (m *ProductStock) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              m.Id,
		"product_id":      m.ProductId,
		"qty":             m.Qty,
		"base_cost_price": m.BaseCostPrice,
		"created_at":      m.CreatedAt,
		"updated_at":      m.UpdatedAt,
	}
}

func (m ProductStock) RecalculateBaseCostPrice(qtyChange float64, baseCostPrice float64) float64 {
	if qtyChange <= 0 {
		panic("invalid qty change")
	}

	return ((m.BaseCostPrice * m.Qty) + (baseCostPrice * qtyChange)) / (qtyChange + m.Qty)
}

type ProductStockQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &ProductStockQueryOption{}

func (o *ProductStockQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"ps.*"}
	}
}

func (o *ProductStockQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
