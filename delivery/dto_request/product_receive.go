package dto_request

import (
	"mime/multipart"
	"myapp/data_type"
)

type ProductReceiveAddImageRequest struct {
	FilePath    string  `json:"file_path" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty"`

	ProductReceiveId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveAddImageRequest

type ProductReceiveUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
} // @name ProductReceiveUploadRequest

type ProductReceiveFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=date created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductReceiveFetchSorts

type ProductReceiveFetchRequest struct {
	PaginationRequest
	Sorts      ProductReceiveFetchSorts        `json:"sorts" validate:"unique=Field,dive"`
	Status     *data_type.ProductReceiveStatus `json:"status" validate:"omitempty,data_type_enum"`
	SupplierId *string                         `json:"supplier_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	Phrase     *string                         `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
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

type ProductReceiveUpdateItemRequest struct {
	QtyEligible float64 `json:"qty_eligible" validate:"required,gte=0"`

	ProductReceiveId     string `json:"-" swaggerignore:"true"`
	ProductReceiveItemId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveUpdateItemRequest

type ProductReceiveDeleteImageRequest struct {
	ProductReceiveId      string `json:"-" swaggerignore:"true"`
	ProductReceiveImageId string `json:"-" swaggerignore:"true"`
} // @name ProductReceiveDeleteImageRequest
