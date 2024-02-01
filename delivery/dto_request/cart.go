package dto_request

type CartCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CartCreateRequest

type CartFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CartFetchSorts

type CartFetchRequest struct {
	PaginationRequest
	Sorts  CartFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CartFetchRequest

type CartGetRequest struct {
	CartId string `json:"-" swaggerignore:"true"`
} // @name CartGetRequest

type CartUpdateRequest struct {
	CartId string `json:"-" swaggerignore:"true"`

	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CartUpdateRequest

type CartDeleteRequest struct {
	CartId string `json:"-" swaggerignore:"true"`
} // @name CartDeleteRequest
