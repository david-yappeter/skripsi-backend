package dto_request

type AdminProductCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminProductCreateRequest

type AdminProductFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name AdminProductFetchSorts

type AdminProductFetchRequest struct {
	PaginationRequest
	Sorts    AdminProductFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	IsActive bool                   `json:"is_active" extensions:"x-nullable"`
	Phrase   *string                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminProductFetchRequest

type AdminProductGetRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`
} // @name AdminProductGetRequest

type AdminProductUpdateRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`

	Name        string   `json:"name" validate:"required,not_empty"`
	Description *string  `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
	Price       *float64 `json:"price" validate:"omitempty,gt=0" extensions:"x-nullable"`
	IsActive    bool     `json:"is_active"`
} // @name AdminProductUpdateRequest

type AdminProductDeleteRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`
} // @name AdminProductDeleteRequest
