package dto_response

import (
	"myapp/model"
)

type ProductReceiveImageResponse struct {
	Id               string  `json:"id"`
	ProductReceiveId string  `json:"product_receive_id"`
	FileId           string  `json:"file_id"`
	Description      *string `json:"description" extensions:"x-nullable"`

	Timestamp

	File *FileResponse `json:"file" extensions:"x-nullable"`
} // @name ProductReceiveImageResponse

func NewProductReceiveImageResponse(productReceiveImage model.ProductReceiveImage) ProductReceiveImageResponse {
	r := ProductReceiveImageResponse{
		Id:               productReceiveImage.Id,
		ProductReceiveId: productReceiveImage.ProductReceiveId,
		FileId:           productReceiveImage.FileId,
		Description:      productReceiveImage.Description,
		Timestamp:        Timestamp(productReceiveImage.Timestamp),
	}

	if productReceiveImage.File != nil {
		r.File = NewFileResponseP(*productReceiveImage.File)
	}

	return r
}

func NewProductReceiveImageResponseP(productReceiveImage model.ProductReceiveImage) *ProductReceiveImageResponse {
	r := NewProductReceiveImageResponse(productReceiveImage)

	return &r
}
