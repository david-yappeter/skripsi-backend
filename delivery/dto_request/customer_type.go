package dto_request

type CustomerTypeCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeCreateRequest

type CustomerTypeFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CustomerTypeFetchSorts

type CustomerTypeFetchRequest struct {
	PaginationRequest
	Sorts  CustomerTypeFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeFetchRequest

type CustomerTypeGetRequest struct {
	CustomerTypeId string `json:"-" swaggerignore:"true"`
} // @name CustomerTypeGetRequest

type CustomerTypeUpdateRequest struct {
	CustomerTypeId string `json:"-" swaggerignore:"true"`

	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeUpdateRequest

type CustomerTypeDeleteRequest struct {
	CustomerTypeId string `json:"-" swaggerignore:"true"`
} // @name CustomerTypeDeleteRequest
