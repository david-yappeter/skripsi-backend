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
	Weight          float64                                         `json:"weight" validate:"required,gt=0"`
	WeightUnit      data_type.TiktokProductPackageWeight            `json:"weight_unit" validate:"required,data_type_enum"`
	Attributes      []gotiktok.CreateProductRequestProductAttribute `json:"attributes"`
	SizeChartUri    *string                                         `json:"size_chart_uri" validate:"omitempty,not_empty"`
} // @name TiktokProductCreateRequest

type TiktokProductUploadImageRequest struct {
	File multipart.FileHeader `json:"file" validate:"required"`
} // @name TiktokProductUploadImageRequest

type TiktokProductFetchBrandsRequest struct {
	NextPageToken *string `json:"next_page_token" validate:"omitempty,not_empty"`
	Phrase        *string `json:"phrase" validate:"omitempty,not_empty"`
	CategoryId    *string `json:"category_id" validate:"omitempty,not_empty"`
} // @name TiktokProductFetchBrandsRequest

type TiktokProductGetCategoryRulesRequest struct {
	CategoryId string `json:"-" swaggerignore:"true"`
} // @name TiktokProductGetCategoryRulesRequest

type TiktokProductGetCategoryAttributesRequest struct {
	CategoryId string `json:"-" swaggerignore:"true"`
} // @name TiktokProductGetCategoryAttributesRequest

type TiktokProductGetRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`
} // @name TiktokProductGetRequest

type TiktokProductUpdateRequest struct {
	Title           string                                          `json:"title" validate:"required,not_empty,min=25"`
	Description     string                                          `json:"description" validate:"required,not_empty"`
	CategoryId      string                                          `json:"category_id" validate:"required,not_empty"`
	BrandId         *string                                         `json:"brand_id" validate:"omitempty,not_empty"`
	ImagesUri       []string                                        `json:"images_uri" validate:"min=1"`
	DimensionHeight *float64                                        `json:"dimension_height" validate:"required_with=DimensionUnit,omitempty,gt=0"`
	DimensionWidth  *float64                                        `json:"dimension_width" validate:"required_with=DimensionUnit,omitempty,gt=0"`
	DimensionLength *float64                                        `json:"dimension_length" validate:"required_with=DimensionUnit,omitempty,gt=0"`
	DimensionUnit   *data_type.TiktokProductDimensionUnit           `json:"dimension_unit" validate:"omitempty,data_type_enum"`
	Weight          float64                                         `json:"weight" validate:"required,gt=0"`
	WeightUnit      data_type.TiktokProductPackageWeight            `json:"weight_unit" validate:"required,data_type_enum"`
	Attributes      []gotiktok.UpdateProductRequestProductAttribute `json:"attributes"`
	SizeChartUri    *string                                         `json:"size_chart_uri" validate:"omitempty,not_empty"`

	ProductId string `json:"-" swaggerignore:"true"`
} // @name TiktokProductUpdateRequest

type TiktokProductRecommendedCategoryRequest struct {
	ProductTitle string   `json:"product_title" validate:"required,not_empty"`
	Description  *string  `json:"description" validate:"omitempty,not_empty"`
	ImagesUri    []string `json:"images_uri"`
} // @name TiktokProductRecommendCategoryRequest

type TiktokProductActivateRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`
} // @name TiktokProductActivateRequest

type TiktokProductDeactivateRequest struct {
	ProductId string `json:"-" swaggerignore:"true"`
} // @name TiktokProductDeactivateRequest
