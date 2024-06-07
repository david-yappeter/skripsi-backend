package dto_response

import (
	"myapp/model"
)

type PurchaseOrderItemResponse struct {
	Id              string  `json:"id"`
	PurchaseOrderId string  `json:"purchase_order_id"`
	ProductUnitId   string  `json:"product_unit_id"`
	Qty             float64 `json:"qty"`
	PricePerUnit    float64 `json:"price_per_unit"`
	Timestamp

	ProductUnit *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
} // @name PurchaseOrderItemResponse

func NewPurchaseOrderItemResponse(purchaseOrderItem model.PurchaseOrderItem) PurchaseOrderItemResponse {
	r := PurchaseOrderItemResponse{
		Id:              purchaseOrderItem.Id,
		PurchaseOrderId: purchaseOrderItem.PurchaseOrderId,
		ProductUnitId:   purchaseOrderItem.ProductUnitId,
		Qty:             purchaseOrderItem.Qty,
		PricePerUnit:    purchaseOrderItem.PricePerUnit,
		Timestamp:       Timestamp(purchaseOrderItem.Timestamp),
	}

	if purchaseOrderItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*purchaseOrderItem.ProductUnit)
	}

	return r
}

func NewPurchaseOrderItemResponseP(purchaseOrderImage model.PurchaseOrderItem) *PurchaseOrderItemResponse {
	r := NewPurchaseOrderItemResponse(purchaseOrderImage)

	return &r
}
