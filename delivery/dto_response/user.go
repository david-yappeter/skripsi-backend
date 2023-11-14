package dto_response

import "myapp/model"

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
} // @name UserResponse

func NewUserResponse(user model.User) UserResponse {
	r := UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		IsActive: user.IsActive,
	}

	return r
}
