package dto_request

import (
	"mime/multipart"
	"myapp/data_type"
)

type DeliveryOrderCreateRequest struct {
	CustomerId string         `json:"customer_id" validate:"required,not_empty,uuid"`
	Date       data_type.Date `json:"date"`
} // @name DeliveryOrderCreateRequest

type DeliveryOrderAddItemRequest struct {
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	UnitId    string  `json:"unit_id" validate:"required,not_empty,uuid"`
	Qty       float64 `json:"qty" validate:"required,gt=0"`

	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderAddItemRequest

type DeliveryOrderAddImageRequest struct {
	FilePath    string  `json:"file_path" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty"`

	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderAddImageRequest

type DeliveryOrderAddDriverRequest struct {
	DriverUserId string `json:"driver_user_id" validate:"required,not_empty,uuid"`

	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderAddDriverRequest

type DeliveryOrderUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
} // @name DeliveryOrderUploadRequest

type DeliveryOrderFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=date created_at updated_at" example:"date"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name DeliveryOrderFetchSorts

type DeliveryOrderFetchRequest struct {
	PaginationRequest
	Sorts      DeliveryOrderFetchSorts        `json:"sorts" validate:"unique=Field,dive"`
	CustomerId *string                        `json:"customer_id" validate:"omitempty,not_empty,uuid"`
	Status     *data_type.DeliveryOrderStatus `json:"status" validate:"omitempty,data_type_enum"`
	Phrase     *string                        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name DeliveryOrderFetchRequest

type DeliveryOrderFetchDriverRequest struct {
	PaginationRequest
	Status *data_type.DeliveryOrderStatus `json:"status" validate:"omitempty,dive,data_type_enum"`
} // @name DeliveryOrderFetchDriverRequest

type DeliveryOrderGetRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderGetRequest

type LatestDeliveryLocationRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name LatestDeliveryLocationRequest

type DeliveryOrderMarkOngoingRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderMarkOngoingRequest

type DeliveryOrderDeliveringRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeliveringRequest

type DeliveryOrderCancelRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderCancelRequest

type DeliveryOrderMarkCompletedRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderMarkCompletedRequest

type DeliveryOrderDeliveryLocationRequest struct {
	Latitude  float64 `json:"latitude" validate:"latitude"`
	Longitude float64 `json:"longitude" validate:"longitude"`
	Bearing   float64 `json:"bearing" validate:"min=0,max=360"`

	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderMarkCompletedRequest

type DeliveryOrderDeleteRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeleteRequest

type DeliveryOrderDeleteImageRequest struct {
	FileId          string `json:"-" swaggerignore:"true"`
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeleteImageRequest

type DeliveryOrderDeleteItemRequest struct {
	DeliveryOrderItemId string `json:"-" swaggerignore:"true"`
	DeliveryOrderId     string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeleteItemRequest

type DeliveryOrderDeleteDriverRequest struct {
	DriverUserId    string `json:"-" swaggerignore:"true"`
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeleteDriverRequest
