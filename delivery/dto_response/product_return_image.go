package dto_response

import (
	"myapp/model"
)

type ProductReturnImageResponse struct {
	Id              string  `json:"id"`
	ProductReturnId string  `json:"product_return_id"`
	FileId          string  `json:"file_id"`
	Description     *string `json:"description" extensions:"x-nullable"`

	Timestamp

	File *FileResponse `json:"file" extensions:"x-nullable"`
} // @name ProductReturnImageResponse

func NewProductReturnImageResponse(productReturnImage model.ProductReturnImage) ProductReturnImageResponse {
	r := ProductReturnImageResponse{
		Id:              productReturnImage.Id,
		ProductReturnId: productReturnImage.ProductReturnId,
		FileId:          productReturnImage.FileId,
		Description:     productReturnImage.Description,
		Timestamp:       Timestamp(productReturnImage.Timestamp),
	}

	if productReturnImage.File != nil {
		r.File = NewFileResponseP(*productReturnImage.File)
	}

	return r
}

func NewProductReturnImageResponseP(productReturnImage model.ProductReturnImage) *ProductReturnImageResponse {
	r := NewProductReturnImageResponse(productReturnImage)

	return &r
}
