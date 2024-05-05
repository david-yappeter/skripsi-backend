package dto_response

import (
	"myapp/model"
)

type DeliveryOrderReturnResponse struct {
	Id              string  `json:"id"`
	DeliveryOrderId string  `json:"delivery_order_id"`
	Description     *string `json:"description" extensions:"x-nullable"`

	Timestamp

	CreatedBy *UserResponse                      `json:"created_by" extensions:"x-nullable"`
	Images    []DeliveryOrderReturnImageResponse `json:"images" extensions:"x-nullable"`
} // @name DeliveryOrderReturnResponse

func NewDeliveryOrderReturnResponse(deliveryOrderReturn model.DeliveryOrderReturn) DeliveryOrderReturnResponse {
	r := DeliveryOrderReturnResponse{
		Id:              deliveryOrderReturn.Id,
		DeliveryOrderId: deliveryOrderReturn.DeliveryOrderId,
		Description:     deliveryOrderReturn.Description,
		Timestamp:       Timestamp(deliveryOrderReturn.Timestamp),
	}

	if deliveryOrderReturn.User != nil {
		r.CreatedBy = NewUserResponseP(*deliveryOrderReturn.User)
	}

	for _, deliveryOrderImage := range deliveryOrderReturn.DeliveryOrderReturnImages {
		r.Images = append(r.Images, NewDeliveryOrderReturnImageResponse(deliveryOrderImage))
	}

	return r
}
