package model

const SupplierTypeTableName = "supplier_types"

type SupplierType struct {
	Id          string  `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`

	Timestamp
}

func (m *SupplierType) TableName() string {
	return SupplierTypeTableName
}

func (m *SupplierType) TableIds() []string {
	return []string{"id"}
}

func (m *SupplierType) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.Id,
		"name":        m.Name,
		"description": m.Description,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}

type SupplierTypeQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &SupplierTypeQueryOption{}

func (o *SupplierTypeQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *SupplierTypeQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
