package dto_request

type RoleFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"updated_at"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"desc"`
} // @name RoleFetchSorts

type RoleOptionForUserFormRequest struct {
	PaginationRequest
	Sorts  RoleFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name RoleOptionForUserFormRequest
