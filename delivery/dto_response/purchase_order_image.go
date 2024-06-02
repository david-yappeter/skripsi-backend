package dto_response

import (
	"myapp/model"
)

type PurchaseOrderImageResponse struct {
	Id              string  `json:"id"`
	PurchaseOrderId string  `json:"purchase_order_id"`
	FileId          string  `json:"file_id"`
	Description     *string `json:"description" extensions:"x-nullable"`

	Timestamp

	File *FileResponse `json:"file" extensions:"x-nullable"`
} // @name PurchaseOrderImageResponse

func NewPurchaseOrderImageResponse(purchaseOrderImage model.PurchaseOrderImage) PurchaseOrderImageResponse {
	r := PurchaseOrderImageResponse{
		Id:              purchaseOrderImage.Id,
		PurchaseOrderId: purchaseOrderImage.PurchaseOrderId,
		FileId:          purchaseOrderImage.FileId,
		Description:     purchaseOrderImage.Description,
		Timestamp:       Timestamp(purchaseOrderImage.Timestamp),
	}

	if purchaseOrderImage.File != nil {
		r.File = NewFileResponseP(*purchaseOrderImage.File)
	}

	return r
}

func NewPurchaseOrderImageResponseP(purchaseOrderImage model.PurchaseOrderImage) *PurchaseOrderImageResponse {
	r := NewPurchaseOrderImageResponse(purchaseOrderImage)

	return &r
}
