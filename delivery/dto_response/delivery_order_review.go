package dto_response

import (
	"myapp/model"
)

type DeliveryOrderReviewResponse struct {
	Id              string  `json:"id"`
	DeliveryOrderId string  `json:"delivery_order_id"`
	StarRating      int     `json:"star_rating"`
	Description     *string `json:"description"`

	Timestamp
	DeliveryOrder *DeliveryOrderResponse `json:"delivery_order"`
} // @name DeliveryOrderReviewResponse

func NewDeliveryOrderReviewResponse(deliveryOrderReview model.DeliveryOrderReview) DeliveryOrderReviewResponse {
	r := DeliveryOrderReviewResponse{
		Id:              deliveryOrderReview.Id,
		DeliveryOrderId: deliveryOrderReview.DeliveryOrderId,
		StarRating:      deliveryOrderReview.StarRating,
		Description:     deliveryOrderReview.Description,
		Timestamp:       Timestamp(deliveryOrderReview.Timestamp),
		DeliveryOrder:   &DeliveryOrderResponse{},
	}

	if deliveryOrderReview.DeliveryOrder != nil {
		r.DeliveryOrder = NewDeliveryOrderResponseP(*deliveryOrderReview.DeliveryOrder)
	}

	return r
}
