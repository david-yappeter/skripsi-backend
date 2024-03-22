package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type DeliveryOrderResponse struct {
	Id            string                        `json:"id"`
	CustomerId    string                        `json:"supplier_id"`
	UserId        string                        `json:"user_id"`
	InvoiceNumber string                        `json:"invoice_number"`
	Date          data_type.Date                `json:"date"`
	Status        data_type.DeliveryOrderStatus `json:"status"`
	TotalPrice    float64                       `json:"total_price"`

	Customer *CustomerResponse `json:"customer" extensions:"x-nullable"`
	Timestamp
	Items   []DeliveryOrderItemResponse  `json:"items" extensions:"x-nullable"`
	Images  []DeliveryOrderImageResponse `json:"images" extensions:"x-nullable"`
	Drivers []UserResponse               `json:"drivers" extensions:"x-nullable"`
} // @name DeliveryOrderResponse

func NewDeliveryOrderResponse(deliveryOrder model.DeliveryOrder) DeliveryOrderResponse {
	r := DeliveryOrderResponse{
		Id:            deliveryOrder.Id,
		CustomerId:    deliveryOrder.CustomerId,
		UserId:        deliveryOrder.UserId,
		InvoiceNumber: deliveryOrder.InvoiceNumber,
		Date:          deliveryOrder.Date,
		Status:        deliveryOrder.Status,
		TotalPrice:    deliveryOrder.TotalPrice,
		Timestamp:     Timestamp(deliveryOrder.Timestamp),
	}

	if deliveryOrder.Customer != nil {
		r.Customer = NewCustomerResponseP(*deliveryOrder.Customer)
	}

	for _, deliveryOrderItem := range deliveryOrder.DeliveryOrderItems {
		r.Items = append(r.Items, NewDeliveryOrderItemResponse(deliveryOrderItem))
	}

	for _, deliveryOrderImage := range deliveryOrder.DeliveryOrderImages {
		r.Images = append(r.Images, NewDeliveryOrderImageResponse(deliveryOrderImage))
	}

	for _, deliveryOrderDriver := range deliveryOrder.DeliveryOrderDrivers {
		if deliveryOrderDriver.User != nil {
			r.Drivers = append(r.Drivers, NewUserResponse(*deliveryOrderDriver.User))
		}
	}

	return r
}

func NewDeliveryOrderResponseP(supplierType model.DeliveryOrder) *DeliveryOrderResponse {
	r := NewDeliveryOrderResponse(supplierType)

	return &r
}
