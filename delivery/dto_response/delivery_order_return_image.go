package dto_response

import (
	"myapp/model"
)

type DeliveryOrderReturnImageResponse struct {
	Id                    string `json:"id"`
	DeliveryOrderReturnId string `json:"delivery_order_return_id"`
	FileId                string `json:"file_id"`

	Timestamp

	File *FileResponse `json:"file" extensions:"x-nullable"`
} // @name DeliveryOrderReturnImageResponse

func NewDeliveryOrderReturnImageResponse(deliveryOrderReturnImage model.DeliveryOrderReturnImage) DeliveryOrderReturnImageResponse {
	r := DeliveryOrderReturnImageResponse{
		Id:                    deliveryOrderReturnImage.Id,
		DeliveryOrderReturnId: deliveryOrderReturnImage.DeliveryOrderReturnId,
		Timestamp:             Timestamp(deliveryOrderReturnImage.Timestamp),
	}

	if deliveryOrderReturnImage.File != nil {
		r.File = NewFileResponseP(*deliveryOrderReturnImage.File)
	}

	return r
}
