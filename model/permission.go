package model

import "myapp/data_type"

const PermissionTableName = "permissions"

type Permission struct {
	Id          string                   `db:"id"`
	Title       data_type.Permission     `db:"title"`
	Description string                   `db:"description"`
	IsActive    bool                     `db:"is_active"`
	Type        data_type.PermissionType `db:"type"`
	Timestamp
}

func (m *Permission) TableName() string {
	return PermissionTableName
}

func (m *Permission) TableIds() []string {
	return []string{"id"}
}

func (m *Permission) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.Id,
		"title":       m.Title,
		"description": m.Description,
		"is_active":   m.IsActive,
		"type":        m.Type,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}
