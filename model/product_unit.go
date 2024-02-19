package model

const ProductUnitTableName = "product_units"

type ProductUnit struct {
	Id          string  `db:"id"`
	ToUnitId    *string `db:"to_unit_id"`
	UnitId      string  `db:"unit_id"`
	ProductId   string  `db:"product_id"`
	Scale       float64 `db:"scale"`
	ScaleToBase float64 `db:"scale_to_base"`
	Timestamp

	ProductStock *ProductStock `db:"-"`
}

func (m *ProductUnit) TableName() string {
	return ProductUnitTableName
}

func (m *ProductUnit) TableIds() []string {
	return []string{"id"}
}

func (m *ProductUnit) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.Id,
		"to_unit_id":    m.ToUnitId,
		"unit_id":       m.UnitId,
		"product_id":    m.ProductId,
		"scale":         m.Scale,
		"scale_to_base": m.ScaleToBase,
		"created_at":    m.CreatedAt,
		"updated_at":    m.UpdatedAt,
	}
}

type ProductUnitQueryOption struct {
	QueryOption

	ExcludeIds []string
	Phrase     *string
}

var _ PrepareOption = &ProductUnitQueryOption{}

func (o *ProductUnitQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *ProductUnitQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
