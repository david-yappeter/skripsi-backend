package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type PurchaseOrderResponse struct {
	Id                  string                        `json:"id"`
	SupplierId          string                        `json:"supplier_id"`
	UserId              string                        `json:"user_id"`
	InvoiceNumber       string                        `json:"invoice_number"`
	Date                data_type.Date                `json:"date"`
	Status              data_type.PurchaseOrderStatus `json:"status"`
	TotalEstimatedPrice float64                       `json:"total_estimated_price"`

	Items    []PurchaseOrderItemResponse  `json:"items" extensions:"x-nullable"`
	Images   []PurchaseOrderImageResponse `json:"images" extensions:"x-nullable"`
	Supplier *SupplierResponse            `json:"supplier" extensions:"x-nullable"`
	Timestamp
} // @name PurchaseOrderResponse

func NewPurchaseOrderResponse(purchaseOrder model.PurchaseOrder) PurchaseOrderResponse {
	r := PurchaseOrderResponse{
		Id:                  purchaseOrder.Id,
		SupplierId:          purchaseOrder.SupplierId,
		UserId:              purchaseOrder.UserId,
		InvoiceNumber:       purchaseOrder.InvoiceNumber,
		Date:                purchaseOrder.Date,
		Status:              purchaseOrder.Status,
		TotalEstimatedPrice: purchaseOrder.TotalEstimatedPrice,
		Timestamp:           Timestamp(purchaseOrder.Timestamp),
		Items:               []PurchaseOrderItemResponse{},
	}

	if purchaseOrder.Supplier != nil {
		r.Supplier = NewSupplierResponseP(*purchaseOrder.Supplier)
	}

	for _, purchaseOrderImage := range purchaseOrder.PurchaseOrderImages {
		r.Images = append(r.Images, NewPurchaseOrderImageResponse(purchaseOrderImage))
	}

	for _, purchaseOrderItem := range purchaseOrder.PurchaseOrderItems {
		r.Items = append(r.Items, NewPurchaseOrderItemResponse(purchaseOrderItem))
	}

	return r
}

func NewPurchaseOrderResponseP(supplierType model.PurchaseOrder) *PurchaseOrderResponse {
	r := NewPurchaseOrderResponse(supplierType)

	return &r
}
