package model

import (
	"myapp/data_type"
)

const RoleTableName = "roles"

type Role struct {
	Id          string         `db:"id"`
	Title       data_type.Role `db:"title"`
	Description string         `db:"description"`
	IsActive    bool           `db:"is_active"`
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
		"title":       m.Title,
		"description": m.Description,
		"is_active":   m.IsActive,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}

func (r Role) RoleType() data_type.RoleType {
	return r.Title.RoleType()
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
