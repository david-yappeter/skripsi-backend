package model

const ProductStockTableName = "product_stocks"

type ProductStock struct {
	Id        string  `db:"id"`
	ProductId string  `db:"product_id"`
	Qty       float64 `db:"qty"`

	Timestamp
}

func (m *ProductStock) TableName() string {
	return ProductStockTableName
}

func (m *ProductStock) TableIds() []string {
	return []string{"id"}
}

func (m *ProductStock) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.Id,
		"product_id": m.ProductId,
		"qty":        m.Qty,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}

type ProductStockQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &ProductStockQueryOption{}

func (o *ProductStockQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *ProductStockQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
