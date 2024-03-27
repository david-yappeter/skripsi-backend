package dto_response

import (
	"myapp/model"
	"myapp/util"
)

type DeliveryOrderItemResponse struct {
	Id              string   `json:"id"`
	DeliveryOrderId string   `json:"delivery_order_id"`
	ProductUnitId   string   `json:"product_unit_id"`
	Qty             float64  `json:"qty"`
	PricePerUnit    float64  `json:"price_per_base_unit"`
	PriceTotal      *float64 `json:"price_total"`
	DiscountPerUnit float64  `json:"discount_per_unit"`
	Timestamp

	ProductUnit *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
	CreatedBy   *UserResponse        `json:"created_by" extensions:"x-nullable"`
} // @name DeliveryOrderItemResponse

func NewDeliveryOrderItemResponse(deliveryOrderItem model.DeliveryOrderItem) DeliveryOrderItemResponse {
	r := DeliveryOrderItemResponse{
		Id:              deliveryOrderItem.Id,
		DeliveryOrderId: deliveryOrderItem.DeliveryOrderId,
		ProductUnitId:   deliveryOrderItem.ProductUnitId,
		Qty:             deliveryOrderItem.Qty,
		PricePerUnit:    deliveryOrderItem.PricePerUnit,
		DiscountPerUnit: deliveryOrderItem.DiscountPerUnit,
		Timestamp:       Timestamp(deliveryOrderItem.Timestamp),
	}

	if deliveryOrderItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*deliveryOrderItem.ProductUnit)
		r.PriceTotal = util.Float64P(deliveryOrderItem.ProductUnit.ScaleToBase * deliveryOrderItem.Qty * deliveryOrderItem.PricePerUnit)
	}

	return r
}

func NewDeliveryOrderItemResponseP(productReceiveImage model.DeliveryOrderItem) *DeliveryOrderItemResponse {
	r := NewDeliveryOrderItemResponse(productReceiveImage)

	return &r
}
