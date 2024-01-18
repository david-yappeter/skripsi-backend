package dto_request

type AdminRoleFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"updated_at"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"desc"`
} // @name AdminRoleFetchSorts

type AdminRoleOptionForUserFormRequest struct {
	PaginationRequest
	Sorts  AdminRoleFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string             `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminRoleOptionForUserFormRequest
