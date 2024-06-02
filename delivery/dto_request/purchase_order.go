package dto_request

import (
	"mime/multipart"
	"myapp/data_type"
)

type PurchaseOrderCreateRequest struct {
	SupplierId    string         `json:"supplier_id" validate:"required,not_empty,uuid"`
	InvoiceNumber string         `json:"invoice_number" validate:"required,not_empty"`
	Date          data_type.Date `json:"date"`
} // @name PurchaseOrderCreateRequest

type PurchaseOrderAddItemRequest struct {
	ProductId    string  `json:"product_id" validate:"required,not_empty,uuid"`
	UnitId       string  `json:"unit_id" validate:"required,not_empty,uuid"`
	Qty          float64 `json:"qty" validate:"required,gt=0"`
	PricePerUnit float64 `json:"price_per_unit" validate:"required,gt=0"`

	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderAddItemRequest

type PurchaseOrderAddImageRequest struct {
	FilePath    string  `json:"file_path" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty"`

	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderAddImageRequest

type PurchaseOrderUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
} // @name PurchaseOrderUploadRequest

type PurchaseOrderFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=date created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name PurchaseOrderFetchSorts

type PurchaseOrderFetchRequest struct {
	PaginationRequest
	Sorts      PurchaseOrderFetchSorts        `json:"sorts" validate:"unique=Field,dive"`
	Status     *data_type.PurchaseOrderStatus `json:"status" validate:"omitempty,data_type_enum"`
	SupplierId *string                        `json:"supplier_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	Phrase     *string                        `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name PurchaseOrderFetchRequest

type PurchaseOrderGetRequest struct {
	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderGetRequest

type PurchaseOrderUpdateRequest struct {
	InvoiceNumber string         `json:"invoice_number" validate:"required,not_empty"`
	Date          data_type.Date `json:"date"`

	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderUpdateRequest

type PurchaseOrderOngoingRequest struct {
	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderOngoingRequest

type PurchaseOrderCancelRequest struct {
	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderCancelRequest

type PurchaseOrderMarkCompleteRequest struct {
	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderMarkCompleteRequest

type PurchaseOrderDeleteRequest struct {
	PurchaseOrderId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderDeleteRequest

type PurchaseOrderDeleteImageRequest struct {
	PurchaseOrderId      string `json:"-" swaggerignore:"true"`
	PurchaseOrderImageId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderDeleteImageRequest

type PurchaseOrderDeleteItemRequest struct {
	PurchaseOrderId     string `json:"-" swaggerignore:"true"`
	PurchaseOrderItemId string `json:"-" swaggerignore:"true"`
} // @name PurchaseOrderDeleteItemRequest
