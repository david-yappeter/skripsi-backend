package dto_request

type CartAddItemRequest struct {
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	Qty       float64 `json:"qty" validate:"required,gt=0"`
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

type CartFetchInActiveRequest struct {
	Phrase *string `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CartFetchInActiveRequest

type CartSetActiveRequest struct {
	CartId string `json:"cart_id" validate:"required,not_empty,uuid"`
} // @name CartSetActiveRequest

type CartGetRequest struct {
	CartId string `json:"-" swaggerignore:"true"`
} // @name CartGetRequest

type CartUpdateItemRequest struct {
	Qty        float64 `json:"qty" validate:"required,gt=0"`
	CartItemId string  `json:"-" swaggerignore:"true"`
} // @name CartUpdateRequest

type CartDeleteRequest struct {
	CartId string `json:"cart_id" validate:"required,not_empty,uuid"`
} // @name CartDeleteRequest

type CartDeleteItemRequest struct {
	CartItemId string `json:"-" swaggerignore:"true"`
} // @name CartDeleteRequest
