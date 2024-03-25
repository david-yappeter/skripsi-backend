package dto_request

import (
	"mime/multipart"
	"myapp/data_type"
)

type ProductReceiveCreateRequest struct {
	SupplierId    string         `json:"supplier_id" validate:"required,not_empty,uuid"`
	InvoiceNumber string         `json:"invoice_number" validate:"required,not_empty"`
	Date          data_type.Date `json:"date"`
} // @name ProductReceiveCreateRequest

type ProductReceiveAddItemRequest struct {
	ProductId    string  `json:"product_id" validate:"required,not_empty,uuid"`
	UnitId       string  `json:"unit_id" validate:"required,not_empty,uuid"`
	Qty          float64 `json:"qty" validate:"required,gt=0"`
	PricePerUnit float64 `json:"price_per_unit" validate:"required,gt=0"`

	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveAddItemRequest

type ProductReceiveAddImageRequest struct {
	FilePath string `json:"file_path" validate:"required,not_empty"`

	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveAddImageRequest

type ProductReceiveUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`

	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveUploadRequest

type ProductReceiveFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductReceiveFetchSorts

type ProductReceiveFetchRequest struct {
	PaginationRequest
	Sorts      ProductReceiveFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	SupplierId *string                  `json:"supplier_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	Phrase     *string                  `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductReceiveFetchRequest

type ProductReceiveGetRequest struct {
	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveGetRequest

type ProductReceiveUpdateRequest struct {
	InvoiceNumber string         `json:"invoice_number" validate:"required,not_empty"`
	Date          data_type.Date `json:"date"`

	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveUpdateRequest

type ProductReceiveCancelRequest struct {
	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveCancelRequest

type ProductReceiveMarkCompleteRequest struct {
	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveMarkCompleteRequest

type ProductReceiveDeleteRequest struct {
	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveDeleteRequest

type ProductReceiveDeleteImageRequest struct {
	ProductReceiveId string `json:"-" swaggerignore:"true"`
	FileId           string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveDeleteImageRequest

type ProductReceiveDeleteItemRequest struct {
	ProductReceiveId     string `json:"-" swaggerignore:"true"`
	ProductReceiveItemId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveDeleteItemRequest
