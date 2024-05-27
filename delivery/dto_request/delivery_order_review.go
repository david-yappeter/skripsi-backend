package dto_request

type DeliveryOrderReviewCreateGuestRequest struct {
	DeliveryOrderId string  `json:"delivery_order_id" validate:"required,not_empty,uuid"`
	StarRating      int     `json:"star_rating" validate:"min=1,max=5"`
	Description     *string `json:"description" validate:"omitempty,not_empty"`
} // @name DeliveryOrderReviewCreateGuestRequest

type DeliveryOrderReviewFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=date created_at updated_at" example:"date"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name DeliveryOrderReviewFetchSorts

type DeliveryOrderReviewFetchRequest struct {
	PaginationRequest
	Sorts DeliveryOrderReviewFetchSorts `json:"sorts" validate:"unique=Field,dive"`

	StarRating *int `json:"star_rating" validate:"omitempty,min=1,max=5"`
} // @name DeliveryOrderReviewFetchRequest

type DeliveryOrderReviewGetRequest struct {
	DeliveryOrderReviewId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderReviewGetRequest

type DeliveryOrderReviewIsExistByDeliveryOrderRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderReviewIsExistByDeliveryOrderRequest
