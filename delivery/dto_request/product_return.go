package dto_request

import (
	"mime/multipart"
	"myapp/data_type"
)

type ProductReturnCreateRequest struct {
	SupplierId    string         `json:"supplier_id" validate:"required,not_empty,uuid"`
	InvoiceNumber string         `json:"invoice_number" validate:"required,not_empty"`
	Date          data_type.Date `json:"date"`
} // @name ProductReturnCreateRequest

type ProductReturnAddItemRequest struct {
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	Qty       float64 `json:"qty" validate:"required,gt=0"`

	ProductReturnId string `json:"-" swaggerignore:"true"`
} // @name ProductReturnAddItemRequest

type ProductReturnAddImageRequest struct {
	FilePath    string  `json:"file_path" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty"`

	ProductReturnId string `json:"-" swaggerignore:"true"`
} // @name ProductReturnAddImageRequest

type ProductReturnUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
} // @name ProductReturnUploadRequest

type ProductReturnFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=date created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductReturnFetchSorts

type ProductReturnFetchRequest struct {
	PaginationRequest
	Sorts      ProductReturnFetchSorts        `json:"sorts" validate:"unique=Field,dive"`
	Status     *data_type.ProductReturnStatus `json:"status" validate:"omitempty,data_type_enum"`
	SupplierId *string                        `json:"supplier_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	Phrase     *string                        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductReturnFetchRequest

type ProductReturnGetRequest struct {
	ProductReturnId string `json:"-" swaggerignore:"true"`
} // @name ProductReturnGetRequest

type ProductReturnUpdateRequest struct {
	InvoiceNumber string         `json:"invoice_number" validate:"required,not_empty"`
	Date          data_type.Date `json:"date"`

	ProductReturnId string `json:"-" swaggerignore:"true"`
} // @name ProductReturnUpdateRequest

type ProductReturnMarkCompleteRequest struct {
	ProductReturnId string `json:"-" swaggerignore:"true"`
} // @name ProductReturnMarkCompleteRequest

type ProductReturnDeleteRequest struct {
	ProductReturnId string `json:"-" swaggerignore:"true"`
} // @name ProductReturnDeleteRequest

type ProductReturnDeleteImageRequest struct {
	ProductReturnId      string `json:"-" swaggerignore:"true"`
	ProductReturnImageId string `json:"-" swaggerignore:"true"`
} // @name ProductReturnDeleteImageRequest

type ProductReturnDeleteItemRequest struct {
	ProductReturnItemId string `json:"-" swaggerignore:"true"`
	ProductReturnId     string `json:"-" swaggerignore:"true"`
} // @name ProductReturnDeleteItemRequest
