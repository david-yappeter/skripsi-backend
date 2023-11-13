package dto_request

type AuthUsernameLoginRequest struct {
	Username string `json:"username" validate:"required,not_empty,max=255"`
	Password string `json:"password" validate:"required,not_empty,max=255"`
} // @name AuthUsernameLoginRequest

type AuthUsernameRegisterRequest struct {
	Name     string `json:"name" validate:"required,not_empty,max=255"`
	Username string `json:"username" validate:"required,not_empty,max=255"`
	Password string `json:"password" validate:"required,not_empty,max=255"`
	IsActive bool   `json:"is_active"`
} // @name AuthUsernameRegisterRequest