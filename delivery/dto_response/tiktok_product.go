package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type TiktokPlatformProductResponse struct {
	Id              string                                `json:"id"`
	Status          string                                `json:"status"`
	Title           string                                `json:"title"`
	Description     string                                `json:"description"`
	DimensionHeight *int                                  `json:"dimension_height"`
	DimensionWidth  *int                                  `json:"dimension_width"`
	DimensionLength *int                                  `json:"dimension_length"`
	DimensionUnit   *data_type.TiktokProductDimensionUnit `json:"dimension_unit"`
	WeightValue     *float64                              `json:"weight_value"`
	WeightUnit      *data_type.TiktokProductPackageWeight `json:"weight_unit"`
	Category        TiktokCategoryResponse                `json:"category"`
	Brand           *TiktokBrandResponse                  `json:"brand"`
	Images          []TiktokPlatformImageResponse         `json:"images"`
	Attributes      []TiktokPlatformAttributeResponse     `json:"attributes"`
} // @name TiktokPlatformProductResponse

func NewTiktokPlatformProductResponse(tiktokPlatformProduct model.TiktokPlatformProduct) TiktokPlatformProductResponse {
	r := TiktokPlatformProductResponse{
		Id:              tiktokPlatformProduct.Id,
		Status:          tiktokPlatformProduct.Status,
		Title:           tiktokPlatformProduct.Title,
		Description:     tiktokPlatformProduct.Description,
		DimensionHeight: tiktokPlatformProduct.DimensionHeight,
		DimensionWidth:  tiktokPlatformProduct.DimensionWidth,
		DimensionLength: tiktokPlatformProduct.DimensionLength,
		DimensionUnit:   tiktokPlatformProduct.DimensionUnit,
		WeightValue:     tiktokPlatformProduct.WeightValue,
		WeightUnit:      tiktokPlatformProduct.WeightUnit,
		Category:        NewTiktokCategoryResponse(tiktokPlatformProduct.Category),
		Brand:           nil,
		Images:          []TiktokPlatformImageResponse{},
		Attributes:      []TiktokPlatformAttributeResponse{},
	}

	if tiktokPlatformProduct.Brand != nil {
		r.Brand = &TiktokBrandResponse{
			Id:   tiktokPlatformProduct.Brand.Id,
			Name: tiktokPlatformProduct.Brand.Name,
		}
	}

	for _, image := range tiktokPlatformProduct.Images {
		r.Images = append(r.Images, NewTiktokPlatformImageResponse(image))
	}

	for _, attribute := range tiktokPlatformProduct.Attributes {
		r.Attributes = append(r.Attributes, NewTiktokPlatformAttributeResponse(attribute))
	}

	return r
}
