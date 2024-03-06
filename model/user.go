package model

const UserTableName = "users"

type User struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Name     string `db:"name"`
	Password string `db:"password"`
	IsActive bool   `db:"is_active"`
	Timestamp

	Permissions []Permission `db:"-"`
	Roles       []Role       `db:"-"`
	UserRoles   []UserRole   `db:"-"`
}

func (m *User) TableName() string {
	return UserTableName
}

func (m *User) TableIds() []string {
	return []string{"id"}
}

func (m *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.Id,
		"username":   m.Username,
		"name":       m.Name,
		"password":   m.Password,
		"is_active":  m.IsActive,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}

type UserQueryOption struct {
	QueryOption

	Phrase   *string
	IsActive *bool
	RoleIds  []string
}

var _ PrepareOption = &UserQueryOption{}

func (o *UserQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *UserQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
