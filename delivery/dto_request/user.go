package dto_request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,not_empty,max=255"`
	Username string `json:"username" validate:"required,not_empty,max=255"`
	Password string `json:"password" validate:"required,not_empty,max=255"`
	IsActive bool   `json:"is_active"`
} // @name UserCreateRequest

type UserUpdateRequest struct {
	Name string `json:"name" validate:"required,not_empty,max=255"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserUpdateRequest

type UserUpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,not_empty,max=255"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserUpdatePasswordRequest

type UserUpdateActiveRequest struct {
	UserId string `json:"-" swaggerignore:"true"`
} // @name UserUpdateActiveRequest

type UserUpdateInActiveRequest struct {
	UserId string `json:"-" swaggerignore:"true"`
} // @name UserUpdateInActiveRequest

type UserAddRoleRequest struct {
	RoleId string `json:"role_id" validate:"required,not_empty,uuid"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserAddRoleRequest

type UserDeleteRoleRequest struct {
	RoleId string `json:"role_id" validate:"required,not_empty,uuid"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserDeleteRoleRequest
