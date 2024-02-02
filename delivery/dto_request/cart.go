package dto_request

type CartAddItemRequest struct {
	ProductUnitId string  `json:"product_unit_id" validate:"required,not_empty,uuid"`
	Qty           float64 `json:"qty" validate:"required,gt=0"`
} // @name CartAddItemRequest

type CartFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CartFetchSorts

type CartFetchRequest struct {
	PaginationRequest
	Sorts  CartFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CartFetchRequest

type CartFetchFetchInActiveRequest struct {
	Phrase *string `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CartFetchFetchInActiveRequest

type CartGetRequest struct {
	CartId string `json:"-" swaggerignore:"true"`
} // @name CartGetRequest

type CartUpdateItemRequest struct {
	ProductUnitId string  `json:"product_unit_id" validate:"required,not_empty,uuid"`
	Qty           float64 `json:"qty" validate:"required,gt=0"`
} // @name CartUpdateRequest

type CartDeleteItemRequest struct {
	ProductUnitId string `json:"product_unit_id" validate:"required,not_empty,uuid"`
} // @name CartDeleteRequest
