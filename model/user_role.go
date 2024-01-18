package model

const UserRoleTableName = "user_roles"

type UserRole struct {
	UserId string `db:"user_id"`
	RoleId string `db:"role_id"`
	Timestamp

	Role *Role `db:"-"`
}

func (m *UserRole) TableName() string {
	return UserRoleTableName
}

func (m *UserRole) TableIds() []string {
	return []string{"user_id", "role_id"}
}

func (m *UserRole) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id":    m.UserId,
		"role_id":    m.RoleId,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}
