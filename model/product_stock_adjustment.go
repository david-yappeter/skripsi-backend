package model

const ProductStockAdjustmentTableName = "product_stock_adjustments"

type ProductStockAdjustment struct {
	Id             string  `db:"id"`
	UserId         string  `db:"user_id"`
	ProductStockId string  `db:"product_stock_id"`
	PreviousQty    float64 `db:"previous_qty"`
	UpdatedQty     float64 `db:"updated_qty"`
	Timestamp
}

func (m *ProductStockAdjustment) TableName() string {
	return ProductStockAdjustmentTableName
}

func (m *ProductStockAdjustment) TableIds() []string {
	return []string{"id"}
}

func (m *ProductStockAdjustment) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               m.Id,
		"user_id":          m.UserId,
		"product_stock_id": m.ProductStockId,
		"previous_qty":     m.PreviousQty,
		"updated_qty":      m.UpdatedQty,
		"created_at":       m.CreatedAt,
		"updated_at":       m.UpdatedAt,
	}
}

type ProductStockAdjustmentQueryOption struct {
	QueryOption

	UserId         *string
	ProductStockId *string
}

var _ PrepareOption = &ProductStockAdjustmentQueryOption{}

func (o *ProductStockAdjustmentQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"psa.*"}
	}
}

func (o *ProductStockAdjustmentQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "created_at", Direction: "desc"},
		}
	}
}
