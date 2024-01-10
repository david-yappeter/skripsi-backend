package dto_request

import "mime/multipart"

type AdminProductUnitCreateRequest struct {
	ToUnitId      *string `json:"to_unit_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	ImageFilePath *string `json:"image_file_path" validate:"omitempty,not_empty"`
	UnitId        string  `json:"unit_id" validate:"required,not_empty,uuid"`
	ProductId     string  `json:"product_id" validate:"required,not_empty,uuid"`
	Scale         float64 `json:"scale" validate:"required,not_empty,gte=1"`
} // @name AdminProductUnitCreateRequest

type AdminProductUnitUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
}

type AdminProductUnitGetRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name AdminProductUnitGetRequest

type AdminProductUnitUpdateRequest struct {
	ToUnitId      *string `json:"to_unit_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	ImageFilePath *string `json:"image_file_path" validate:"omitempty,not_empty"`
	UnitId        string  `json:"unit_id" validate:"required,not_empty,uuid"`
	ProductId     string  `json:"product_id" validate:"required,not_empty,uuid"`
	Scale         float64 `json:"scale" validate:"required,not_empty,gte=1"`

	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name AdminProductUnitUpdateRequest

type AdminProductUnitDeleteRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name AdminProductUnitDeleteRequest
