package model

import "myapp/data_type"

const TiktokProductTableName = "tiktok_products"

type TiktokPlatformProduct struct {
	Id              string                                `db:"-"`
	Status          string                                `db:"-"`
	Title           string                                `db:"-"`
	Description     string                                `db:"-"`
	Category        TiktokCategory                        `db:"-"`
	Brand           *TiktokBrand                          `db:"-"`
	DimensionHeight *int                                  `db:"-"`
	DimensionWidth  *int                                  `db:"-"`
	DimensionLength *int                                  `db:"-"`
	DimensionUnit   *data_type.TiktokProductDimensionUnit `db:"-"`
	WeightValue     *float64                              `db:"-"`
	WeightUnit      *data_type.TiktokProductPackageWeight `db:"-"`

	Images     []TiktokPlatformImage     `db:"-"`
	Attributes []TiktokPlatformAttribute `db:"-"`
}

type TiktokProduct struct {
	TiktokProductId string                        `db:"tiktok_product_id"`
	ProductId       string                        `db:"product_id"`
	Status          data_type.TiktokProductStatus `db:"status"`

	Timestamp
}

func (m *TiktokProduct) TableName() string {
	return TiktokProductTableName
}

func (m *TiktokProduct) TableIds() []string {
	return []string{"id"}
}

func (m *TiktokProduct) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tiktok_product_id": m.TiktokProductId,
		"product_id":        m.ProductId,
		"status":            m.Status,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}

type TiktokProductQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &TiktokProductQueryOption{}

func (o *TiktokProductQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *TiktokProductQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
