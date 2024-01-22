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

	return r
}

func NewDeliveryOrderResponseP(supplierType model.DeliveryOrder) *DeliveryOrderResponse {
	r := NewDeliveryOrderResponse(supplierType)

	return &r
}
