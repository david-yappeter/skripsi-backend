package dto_response

import (
	"myapp/model"
)

type DeliveryOrderImageResponse struct {
	Id              string  `json:"id"`
	DeliveryOrderId string  `json:"delivery_order_id"`
	FileId          string  `json:"file_id"`
	Description     *string `json:"description"`

	Timestamp

	File *FileResponse `json:"file" extensions:"x-nullable"`
} // @name DeliveryOrderImageResponse

func NewDeliveryOrderImageResponse(deliveryOrderImage model.DeliveryOrderImage) DeliveryOrderImageResponse {
	r := DeliveryOrderImageResponse{
		Id:              deliveryOrderImage.Id,
		DeliveryOrderId: deliveryOrderImage.DeliveryOrderId,
		FileId:          deliveryOrderImage.FileId,
		Description:     deliveryOrderImage.Description,
		Timestamp:       Timestamp(deliveryOrderImage.Timestamp),
	}

	if deliveryOrderImage.File != nil {
		r.File = NewFileResponseP(*deliveryOrderImage.File)
	}

	return r
}

func NewDeliveryOrderImageResponseP(productReceiveImage model.DeliveryOrderImage) *DeliveryOrderImageResponse {
	r := NewDeliveryOrderImageResponse(productReceiveImage)

	return &r
}
