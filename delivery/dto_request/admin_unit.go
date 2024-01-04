package dto_request

type AdminUnitCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminUnitCreateRequest

type AdminUnitFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name AdminUnitFetchSorts

type AdminUnitFetchRequest struct {
	PaginationRequest
	Sorts  AdminUnitFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string             `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminUnitFetchRequest

type AdminUnitGetRequest struct {
	Id string `json:"-" swaggerignore:"true"`
} // @name AdminUnitGetRequest

type AdminUnitUpdateRequest struct {
	Id string `json:"-" swaggerignore:"true"`

	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminUnitUpdateRequest

type AdminUnitDeleteRequest struct {
	Id string `json:"-" swaggerignore:"true"`
} // @name AdminUnitDeleteRequest
