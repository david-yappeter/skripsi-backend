package model

const ShopOrderItemTableName = "shop_order_items"

type ShopOrderItem struct {
	Id                string  `db:"id"`
	ShopOrderId       string  `db:"shop_order_id"`
	ProductUnitId     string  `db:"product_unit_id"`
	PlatformProductId string  `db:"platform_product_id"`
	ImageLink         *string `db:"image_link"`
	Quantity          float64 `db:"quantity"`
	OriginalPrice     float64 `db:"original_price"`
	SalePrice         float64 `db:"sale_price"`

	Timestamp

	ProductUnit *ProductUnit `db:"-"`
}

func (m *ShopOrderItem) TableName() string {
	return ShopOrderItemTableName
}

func (m *ShopOrderItem) TableIds() []string {
	return []string{"id"}
}

func (m *ShopOrderItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                  m.Id,
		"shop_order_id":       m.ShopOrderId,
		"product_unit_id":     m.ProductUnitId,
		"platform_product_id": m.PlatformProductId,
		"image_link":          m.ImageLink,
		"quantity":            m.Quantity,
		"original_price":      m.OriginalPrice,
		"sale_price":          m.SalePrice,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
	}
}

type ShopOrderItemQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &ShopOrderItemQueryOption{}

func (o *ShopOrderItemQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"soi.*"}
	}
}

func (o *ShopOrderItemQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
