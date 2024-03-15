package dto_request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,not_empty,max=255"`
	Username string `json:"username" validate:"required,not_empty,max=255"`
	Password string `json:"password" validate:"required,not_empty,max=255"`
	IsActive bool   `json:"is_active"`
} // @name UserCreateRequest

type UserFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=username is_active created_at updated_at" example:"username"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name UserFetchSorts

type UserFetchRequest struct {
	PaginationRequest
	Sorts    UserFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase   *string        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
	IsActive *bool          `json:"is_active" extensions:"x-nullable"`
	RoleIds  []string       `json:"role_ids" extensions:"x-nullable"`
} // @name UserFetchRequest

type UserGetRequest struct {
	UserId string `json:"-" swaggerignore:"true"`
} // @name UserGetRequest

type UserUpdateRequest struct {
	Name     string `json:"name" validate:"required,not_empty,max=255"`
	IsActive bool   `json:"is_active"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserUpdateRequest

type UserUpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,not_empty,max=255"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserUpdatePasswordRequest

type UserAddRoleRequest struct {
	RoleId string `json:"role_id" validate:"required,not_empty,uuid"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserAddRoleRequest

type UserDeleteRoleRequest struct {
	RoleId string `json:"role_id" validate:"required,not_empty,uuid"`

	UserId string `json:"-" swaggerignore:"true"`
} // @name UserDeleteRoleRequest
