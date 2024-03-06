package model

const UnitTableName = "units"

type Unit struct {
	Id          string  `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`

	Timestamp
}

func (m *Unit) TableName() string {
	return UnitTableName
}

func (m *Unit) TableIds() []string {
	return []string{"id"}
}

func (m *Unit) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.Id,
		"name":        m.Name,
		"description": m.Description,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}

type UnitQueryOption struct {
	QueryOption

	ProductIdNotExist *string
	Phrase            *string
}

var _ PrepareOption = &UnitQueryOption{}

func (o *UnitQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *UnitQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
