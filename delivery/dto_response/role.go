package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type RoleResponse struct {
	Id          string         `json:"id"`
	Name        data_type.Role `json:"name"`
	Description *string        `json:"description" extensions:"x-nullable"`

	Timestamp
} // @name RoleResponse

func NewRoleResponse(role model.Role) RoleResponse {
	r := RoleResponse{
		Id:          role.Id,
		Name:        role.Name,
		Description: role.Description,
		Timestamp:   Timestamp(role.Timestamp),
	}

	return r
}
