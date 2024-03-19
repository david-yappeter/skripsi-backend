package dto_request

type UnitCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name UnitCreateRequest

type UnitFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name UnitFetchSorts

type UnitFetchRequest struct {
	PaginationRequest
	Sorts  UnitFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name UnitFetchRequest

type UnitGetRequest struct {
	UnitId string `json:"-" swaggerignore:"true"`
} // @name UnitGetRequest

type UnitUpdateRequest struct {
	UnitId string `json:"-" swaggerignore:"true"`

	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name UnitUpdateRequest

type UnitDeleteRequest struct {
	UnitId string `json:"-" swaggerignore:"true"`
} // @name UnitDeleteRequest

type UnitOptionForProductUnitFormRequest struct {
	PaginationRequest
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	Phrase    *string `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name UnitOptionForProductUnitFormRequest

type UnitOptionForProductUnitToUnitFormRequest struct {
	PaginationRequest
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	Phrase    *string `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name UnitOptionForProductUnitToUnitFormRequest
