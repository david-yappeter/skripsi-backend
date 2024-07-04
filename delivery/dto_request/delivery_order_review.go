package dto_request

import "myapp/data_type"

type DeliveryOrderReviewCreateGuestRequest struct {
	DeliveryOrderId string                            `json:"delivery_order_id" validate:"required,not_empty,uuid"`
	Type            data_type.DeliveryOrderReviewType `json:"type" validate:"required,data_type_enum"`
	Description     string                            `json:"description" validate:"required,not_empty"`
} // @name DeliveryOrderReviewCreateGuestRequest

type DeliveryOrderReviewFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=date created_at updated_at" example:"date"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name DeliveryOrderReviewFetchSorts

type DeliveryOrderReviewFetchRequest struct {
	PaginationRequest
	Sorts DeliveryOrderReviewFetchSorts `json:"sorts" validate:"unique=Field,dive"`

	Type *data_type.DeliveryOrderReviewType `json:"type" validate:"omitempty,data_type_enum"`
} // @name DeliveryOrderReviewFetchRequest

type DeliveryOrderReviewGetRequest struct {
	DeliveryOrderReviewId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderReviewGetRequest

type DeliveryOrderReviewIsExistByDeliveryOrderRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderReviewIsExistByDeliveryOrderRequest
