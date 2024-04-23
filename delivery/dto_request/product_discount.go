package dto_request

type ProductDiscountCreateRequest struct {
	ProductId          string   `json:"product_id" validate:"required,not_empty,uuid"`
	MinimumQty         float64  `json:"minimum_qty" validate:"required,gt=0"`
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage" validate:"omitempty,gt=0,lte=100"`
	DiscountAmount     *float64 `json:"discount_amount" validate:"required_without=DiscountPercentage,omitempty,excluded_with=DiscountPercentage,gt=0"`
} // @name ProductDiscountCreateRequest

type ProductDiscountFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=date created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductDiscountFetchSorts

type ProductDiscountFetchRequest struct {
	PaginationRequest
	Sorts     ProductDiscountFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	ProductId *string                   `json:"product_id" validate:"omitempty,not_empty,uuid"`
	IsActive  *bool                     `json:"is_active"`
	Phrase    *string                   `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductDiscountFetchRequest

type ProductDiscountGetRequest struct {
	ProductDiscountId string `json:"-" swaggerignore:"true"`
} // @name ProductDiscountGetRequest

type ProductDiscountUpdateRequest struct {
	MinimumQty         float64  `json:"minimum_qty" validate:"required,gt=0"`
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage" validate:"omitempty,gt=0,lte=100"`
	DiscountAmount     *float64 `json:"discount_amount" validate:"required_without=DiscountPercentage,omitempty,excluded_with=DiscountPercentage,gt=0"`

	ProductDiscountId string `json:"-" swaggerignore:"true"`
} // @name ProductDiscountUpdateRequest

type ProductDiscountDeleteRequest struct {
	ProductDiscountId string `json:"-" swaggerignore:"true"`
} // @name ProductDiscountDeleteRequest
