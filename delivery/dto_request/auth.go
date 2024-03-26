package dto_request

type AuthUsernameLoginRequest struct {
	Username string `json:"username" validate:"required,not_empty,max=255"`
	Password string `json:"password" validate:"required,not_empty,max=255"`
} // @name AuthUsernameLoginRequest

type AuthUsernameLoginDriverRequest struct {
	Username string `json:"username" validate:"required,not_empty,max=255"`
	Password string `json:"password" validate:"required,not_empty,max=255"`
} // @name AuthUsernameLoginDriverRequest
