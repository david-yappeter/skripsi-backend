package model

const ProductTableName = "products"

type Product struct {
	Id          string   `db:"id"`
	Name        string   `db:"name"`
	Description *string  `db:"description"`
	Price       *float64 `db:"price"`
	IsActive    bool     `db:"is_active"`

	Timestamp
}

func (m *Product) TableName() string {
	return ProductTableName
}

func (m *Product) TableIds() []string {
	return []string{"id"}
}

func (m *Product) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.Id,
		"name":        m.Name,
		"description": m.Description,
		"price":       m.Price,
		"is_active":   m.IsActive,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}

type ProductQueryOption struct {
	QueryOption

	IsActive *bool
	Phrase   *string
}

var _ PrepareOption = &ProductQueryOption{}

func (o *ProductQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *ProductQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
