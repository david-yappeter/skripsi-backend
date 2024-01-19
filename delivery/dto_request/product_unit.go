package dto_request

import "mime/multipart"

type ProductUnitCreateRequest struct {
	ToUnitId      *string `json:"to_unit_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	ImageFilePath *string `json:"image_file_path" validate:"omitempty,not_empty"`
	UnitId        string  `json:"unit_id" validate:"required,not_empty,uuid"`
	ProductId     string  `json:"product_id" validate:"required,not_empty,uuid"`
	Scale         float64 `json:"scale" validate:"required,not_empty,gte=1"`
} // @name ProductUnitCreateRequest

type ProductUnitUploadRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
}

type ProductUnitGetRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name ProductUnitGetRequest

type ProductUnitUpdateRequest struct {
	UnitId    string `json:"unit_id" validate:"required,not_empty,uuid"`
	ProductId string `json:"product_id" validate:"required,not_empty,uuid"`

	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name ProductUnitUpdateRequest

type ProductUnitDeleteRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name ProductUnitDeleteRequest