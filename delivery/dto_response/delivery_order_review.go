package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type DeliveryOrderReviewResponse struct {
	Id              string                            `json:"id"`
	DeliveryOrderId string                            `json:"delivery_order_id"`
	Type            data_type.DeliveryOrderReviewType `json:"type"`
	Description     *string                           `json:"description"`

	Timestamp
	DeliveryOrder *DeliveryOrderResponse `json:"delivery_order"`
} // @name DeliveryOrderReviewResponse

func NewDeliveryOrderReviewResponse(deliveryOrderReview model.DeliveryOrderReview) DeliveryOrderReviewResponse {
	r := DeliveryOrderReviewResponse{
		Id:              deliveryOrderReview.Id,
		DeliveryOrderId: deliveryOrderReview.DeliveryOrderId,
		Type:            deliveryOrderReview.Type,
		Description:     deliveryOrderReview.Description,
		Timestamp:       Timestamp(deliveryOrderReview.Timestamp),
		DeliveryOrder:   &DeliveryOrderResponse{},
	}

	if deliveryOrderReview.DeliveryOrder != nil {
		r.DeliveryOrder = NewDeliveryOrderResponseP(*deliveryOrderReview.DeliveryOrder)
	}

	return r
}
