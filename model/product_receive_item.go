package model

const ProductReceiveItemTableName = "product_receive_items"

type ProductReceiveItem struct {
	Id               string  `db:"id"`
	ProductReceiveId string  `db:"product_receive_id"`
	ProductUnitId    string  `db:"product_unit_id"`
	UserId           string  `db:"user_id"`
	Qty              float64 `db:"qty"`
	ScaleToBase      float64 `db:"scale_to_base"`
	PricePerUnit     float64 `db:"price_per_unit"`
	Timestamp

	ProductUnit          *ProductUnit          `db:"-"`
	User                 *User                 `db:"-"`
	ProductStockMutation *ProductStockMutation `db:"-"` /* only available when ProductReceive status 'COMPLETED' */
}

func (m *ProductReceiveItem) TableName() string {
	return ProductReceiveItemTableName
}

func (m *ProductReceiveItem) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReceiveItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                 m.Id,
		"product_receive_id": m.ProductReceiveId,
		"product_unit_id":    m.ProductUnitId,
		"user_id":            m.UserId,
		"qty":                m.Qty,
		"scale_to_base":      m.ScaleToBase,
		"price_per_unit":     m.PricePerUnit,
		"created_at":         m.CreatedAt,
		"updated_at":         m.UpdatedAt,
	}
}

func (m ProductReceiveItem) BaseQty() float64 {
	return m.Qty * m.ScaleToBase
}
