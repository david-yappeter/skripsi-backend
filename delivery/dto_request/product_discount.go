package dto_request

type ProductDiscountCreateRequest struct {
	ProductId          string   `json:"product_id" validate:"required,not_empty,uuid"`
	MinimumQty         float64  `json:"minimum_qty" validate:"required,gt=0"`
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage" validate:"omitempty,gt=0,lte=100"`
	DiscountAmount     *float64 `json:"discount_amount" validate:"required_without=DiscountPercentage,omitempty,excluded_with=DiscountPercentage,gt=0"`
} // @name ProductDiscountCreateRequest

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
