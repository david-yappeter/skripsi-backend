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
	FilePath string `json:"file_path" validate:"required,not_empty"`

	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderAddImageRequest

type DeliveryOrderUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`

	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderUploadRequest

type DeliveryOrderFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name DeliveryOrderFetchSorts

type DeliveryOrderFetchRequest struct {
	PaginationRequest
	Sorts  DeliveryOrderFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                 `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name DeliveryOrderFetchRequest

type DeliveryOrderGetRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderGetRequest

type DeliveryOrderMarkCompletedRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderMarkCompletedRequest

type DeliveryOrderDeleteRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeleteRequest

type DeliveryOrderDeleteImageRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
	FileId          string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeleteImageRequest

type DeliveryOrderDeleteItemRequest struct {
	DeliveryOrderId string `json:"-" swaggerignore:"true"`
	ProductUnitId   string `json:"-" swaggerignore:"true"`
} // @name DeliveryOrderDeleteItemRequest