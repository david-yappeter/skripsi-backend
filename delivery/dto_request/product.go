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
	IsLoss   *bool             `json:"is_loss" extensions:"x-nullable"`
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

type ProductOptionForProductReceiveItemFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForProductReceiveItemFormSorts

type ProductOptionForProductReceiveItemFormRequest struct {
	PaginationRequest
	Sorts  ProductOptionForProductReceiveItemFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                     `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForProductReceiveItemFormRequest

type ProductOptionForDeliveryOrderItemFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForDeliveryOrderItemFormSorts

type ProductOptionForDeliveryOrderItemFormRequest struct {
	PaginationRequest
	Sorts  ProductOptionForDeliveryOrderItemFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                    `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForDeliveryOrderItemFormRequest

type ProductOptionForCustomerTypeFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForCustomerTypeFormSorts

type ProductOptionForCustomerTypeDiscountFormRequest struct {
	PaginationRequest
	Sorts          ProductOptionForCustomerTypeFormSorts `json:"sorts" validate:"unique=Field,dive"`
	CustomerTypeId string                                `json:"customer_type_id" validate:"required,not_empty,uuid"`
	Phrase         *string                               `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForCustomerTypeDiscountFormRequest

type ProductOptionForCartAddItemFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForCartAddItemFormSorts

type ProductOptionForCartAddItemFormRequest struct {
	PaginationRequest
	Sorts  ProductOptionForCartAddItemFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                              `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForCartAddItemFormRequest

type ProductOptionForProductDiscountFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductOptionForProductDiscountFormSorts

type ProductOptionForProductDiscountFormRequest struct {
	PaginationRequest
	Sorts  ProductOptionForProductDiscountFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                  `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductOptionForProductDiscountFormRequest
