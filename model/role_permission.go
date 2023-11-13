package model

const RolePermissionTableName = "role_permissions"

type RolePermission struct {
	RoleId       string `db:"role_id"`
	PermissionId string `db:"permission_id"`
	Timestamp

	// system
	Role       *Role       `db:"-"`
	Permission *Permission `db:"-"`
}

func (m *RolePermission) TableName() string {
	return RolePermissionTableName
}

func (m *RolePermission) TableIds() []string {
	return []string{"role_id", "permission_id"}
}

func (m *RolePermission) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"role_id":       m.RoleId,
		"permission_id": m.PermissionId,
		"created_at":    m.CreatedAt,
		"updated_at":    m.UpdatedAt,
	}
}
