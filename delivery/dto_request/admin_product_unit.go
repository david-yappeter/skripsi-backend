package dto_request

type AdminProductUnitCreateRequest struct {
	ToUnitId *string `json:"to_unit_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
} // @name AdminProductUnitCreateRequest

type AdminProductUnitFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name AdminProductUnitFetchSorts

type AdminProductUnitFetchRequest struct {
	PaginationRequest
	Sorts    AdminProductUnitFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	IsActive bool                       `json:"is_active" extensions:"x-nullable"`
	Phrase   *string                    `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminProductUnitFetchRequest

type AdminProductUnitGetRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name AdminProductUnitGetRequest

type AdminProductUnitUpdateRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`

	Name        string   `json:"name" validate:"required,not_empty"`
	Description *string  `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
	Price       *float64 `json:"price" validate:"omitempty,gt=0" extensions:"x-nullable"`
	IsActive    bool     `json:"is_active"`
} // @name AdminProductUnitUpdateRequest

type AdminProductUnitDeleteRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name AdminProductUnitDeleteRequest
