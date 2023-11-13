package model

import "myapp/data_type"

const UserAccessTokenTableName = "user_access_tokens"

type UserAccessToken struct {
	Id        string             `db:"id"`
	UserId    string             `db:"user_id"`
	Revoked   bool               `db:"revoked"`
	ExpiredAt data_type.DateTime `db:"expired_at"`
	Timestamp
}

func (m *UserAccessToken) TableName() string {
	return UserAccessTokenTableName
}

func (m *UserAccessToken) TableIds() []string {
	return []string{"id"}
}

func (m *UserAccessToken) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.Id,
		"user_id":    m.UserId,
		"revoked":    m.Revoked,
		"expired_at": m.ExpiredAt,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}
