package dto_request

import "myapp/data_type"

type ProductReceiveCreateRequest struct {
	SupplierId string         `json:"supplier_id" validate:"required,not_empty,uuid"`
	Date       data_type.Date `json:"date"`
} // @name ProductReceiveCreateRequest

type ProductReceiveAddItemRequest struct {
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	UnitId    string  `json:"unit_id" validate:"required,not_empty,uuid"`
	Qty       float64 `json:"qty" validate:"required,gt=0"`

	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveAddItemRequest

type ProductReceiveFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductReceiveFetchSorts

type ProductReceiveFetchRequest struct {
	PaginationRequest
	Sorts  ProductReceiveFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                  `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductReceiveFetchRequest

type ProductReceiveGetRequest struct {
	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveGetRequest

type ProductReceiveDeleteRequest struct {
	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveDeleteRequest
