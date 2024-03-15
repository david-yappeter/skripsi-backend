package dto_request

import "mime/multipart"

type ProductCreateRequest struct {
	Name          string  `json:"name" validate:"required,not_empty"`
	Description   *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
	ImageFilePath string  `json:"image_file_path" validate:"required,not_empty"`
} // @name ProductCreateRequest

type ProductUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
} // @name ProductUploadRequest

type ProductFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductFetchSorts

type ProductFetchRequest struct {
	PaginationRequest
	Sorts    ProductFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	IsActive *bool             `json:"is_active" extensions:"x-nullable"`
	Phrase   *string           `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductFetchRequest

type ProductGetRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`
} // @name ProductGetRequest

type ProductUpdateRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`

	Name          string   `json:"name" validate:"required,not_empty"`
	Description   *string  `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
	Price         *float64 `json:"price" validate:"omitempty,gt=0" extensions:"x-nullable"`
	IsActive      bool     `json:"is_active"`
	ImageFilePath *string  `json:"image_file_path" extensions:"x-nullable"`
} // @name ProductUpdateRequest

type ProductDeleteRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`
} // @name ProductDeleteRequest

type ProductOptionForProductReceiveFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForProductReceiveFormSorts

type ProductOptionForProductReceiveFormRequest struct {
	PaginationRequest
	Sorts  ProductOptionForProductReceiveFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                 `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForProductReceiveFormRequest

type ProductOptionForDeliveryOrderFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForDeliveryOrderFormSorts

type ProductOptionForDeliveryOrderFormRequest struct {
	PaginationRequest
	Sorts  ProductOptionForDeliveryOrderFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForDeliveryOrderFormRequest

type ProductOptionForCustomerTypeFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForCustomerTypeFormSorts

type ProductOptionForCustomerTypeFormRequest struct {
	PaginationRequest
	Sorts          ProductOptionForCustomerTypeFormSorts `json:"sorts" validate:"unique=Field,dive"`
	CustomerTypeId string                                `json:"customer_type_id" validate:"required,not_empty,uuid"`
	Phrase         *string                               `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForCustomerTypeFormRequest
