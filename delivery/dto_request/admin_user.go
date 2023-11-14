package dto_request

type AdminUserCreateRequest struct {
	Name     string `json:"name" validate:"required,not_empty,max=255"`
	Username string `json:"username" validate:"required,not_empty,max=255"`
	Password string `json:"password" validate:"required,not_empty,max=255"`
	IsActive bool   `json:"is_active"`
} // @name AdminUserCreateRequest

type AdminUserUpdateRequest struct {
	Name string `json:"name" validate:"required,not_empty,max=255"`

	Id string `json:"-" swaggerignore:"true"`
} // @name AdminUserUpdateRequest

type AdminUserUpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,not_empty,max=255"`

	Id string `json:"-" swaggerignore:"true"`
} // @name AdminUserUpdatePasswordRequest

type AdminUserUpdateActiveRequest struct {
	Id string `json:"-" swaggerignore:"true"`
} // @name AdminUserUpdateActiveRequest

type AdminUserUpdateInActiveRequest struct {
	Id string `json:"-" swaggerignore:"true"`
} // @name AdminUserUpdateInActiveRequest
