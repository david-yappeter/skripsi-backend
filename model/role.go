package model

import (
	"myapp/data_type"
)

const RoleTableName = "roles"

type Role struct {
	Id          string         `db:"id"`
	Name        data_type.Role `db:"name"`
	Description *string        `db:"description"`
	Timestamp

	// system
	RolePermissions []RolePermission `db:"-"`
}

func (m *Role) TableName() string {
	return RoleTableName
}

func (m *Role) TableIds() []string {
	return []string{"id"}
}

func (m *Role) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.Id,
		"name":        m.Name,
		"description": m.Description,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}

func (r Role) RoleType() data_type.RoleType {
	return r.Name.RoleType()
}

type RoleQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &RoleQueryOption{}

func (o *RoleQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
