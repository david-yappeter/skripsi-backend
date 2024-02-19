package dto_request

import (
	"mime/multipart"
	"myapp/data_type"

	gotiktok "github.com/david-yappeter/go-tiktok"
)

type TiktokProductCreateRequest struct {
	ProductId       string                                          `json:"product_id" validate:"required,not_empty"`
	Title           string                                          `json:"title" validate:"required,not_empty,min=25"`
	Description     string                                          `json:"description" validate:"required,not_empty"`
	CategoryId      string                                          `json:"category_id" validate:"required,not_empty"`
	BrandId         *string                                         `json:"brand_id" validate:"omitempty,not_empty"`
	ImagesUri       []string                                        `json:"images_uri" validate:"min=1"`
	DimensionHeight *float64                                        `json:"dimension_height" validate:"required_with=DimensionUnit,omitempty,gt=0"`
	DimensionWidth  *float64                                        `json:"dimension_width" validate:"required_with=DimensionUnit,omitempty,gt=0"`
	DimensionLength *float64                                        `json:"dimension_length" validate:"required_with=DimensionUnit,omitempty,gt=0"`
	DimensionUnit   *data_type.TiktokProductDimensionUnit           `json:"dimension_unit" validate:"omitempty,data_type_enum"`
	Weight          float64                                         `json:"weight" validate:"omitempty,gt=0"`
	WeightUnit      data_type.TiktokProductPackageWeigth            `json:"weight_unit" validate:"required,data_type_enum"`
	Attributes      []gotiktok.CreateProductRequestProductAttribute `json:"attributes"`
} // @name TiktokProductCreateRequest

type TiktokProductUploadImageRequest struct {
	File multipart.FileHeader `json:"file" validate:"required"`
} // @name TiktokProductUploadImageRequest
