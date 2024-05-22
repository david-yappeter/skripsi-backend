package dto_response

import (
	"myapp/model"
)

type DeliveryOrderPositionResponse struct {
	Id              string  `json:"id"`
	DeliveryOrderId string  `json:"delivery_order_id"`
	DriverUserId    string  `json:"driver_user_id"`
	Latitude        float64 `json:"latitude" validate:"min=-90,max=90"`
	Longitude       float64 `json:"longitude" validate:"min=-180,max=180"`
	Bearing         float64 `json:"bearing"`

	Timestamp
} // @name DeliveryOrderPositionResponse

func NewDeliveryOrderPositionResponse(deliveryOrderPosition model.DeliveryOrderPosition) DeliveryOrderPositionResponse {
	r := DeliveryOrderPositionResponse{
		Id:              deliveryOrderPosition.Id,
		DeliveryOrderId: deliveryOrderPosition.DeliveryOrderId,
		DriverUserId:    deliveryOrderPosition.DriverUserId,
		Latitude:        deliveryOrderPosition.Latitude,
		Longitude:       deliveryOrderPosition.Longitude,
		Bearing:         deliveryOrderPosition.Bearing,
		Timestamp:       Timestamp(deliveryOrderPosition.Timestamp),
	}

	return r
}

func NewDeliveryOrderPositionResponseP(supplierType model.DeliveryOrderPosition) *DeliveryOrderPositionResponse {
	r := NewDeliveryOrderPositionResponse(supplierType)

	return &r
}
