package dto_response

import "myapp/model"

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`

	Roles []RoleResponse `json:"roles" extensions:"x-nullable"`
} // @name UserResponse

func NewUserResponse(user model.User) UserResponse {
	r := UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		IsActive: user.IsActive,
	}

	if len(user.UserRoles) > 0 {
		for _, v := range user.UserRoles {
			if v.Role != nil {
				r.Roles = append(r.Roles, NewRoleResponse(*v.Role))
			}
		}
	} else if user.Roles != nil {
		for _, v := range user.Roles {
			r.Roles = append(r.Roles, NewRoleResponse(v))
		}
	}

	return r
}
