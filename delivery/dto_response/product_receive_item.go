package dto_response

import (
	"myapp/model"
)

type ProductReceiveItemResponse struct {
	Id               string  `json:"id"`
	ProductReceiveId string  `json:"product_receive_id"`
	ProductUnitId    string  `json:"product_unit_id"`
	QtyEligible      float64 `json:"qty_eligible"`
	Qty              float64 `json:"qty"`
	PricePerUnit     float64 `json:"price_per_unit"`
	Timestamp

	ProductUnit *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
	CreatedBy   *UserResponse        `json:"created_by" extensions:"x-nullable"`
} // @name ProductReceiveItemResponse

func NewProductReceiveItemResponse(productReceiveItem model.ProductReceiveItem) ProductReceiveItemResponse {
	r := ProductReceiveItemResponse{
		Id:               productReceiveItem.Id,
		ProductReceiveId: productReceiveItem.ProductReceiveId,
		ProductUnitId:    productReceiveItem.ProductUnitId,
		QtyEligible:      productReceiveItem.QtyEligible,
		Qty:              productReceiveItem.QtyReceived,
		PricePerUnit:     productReceiveItem.PricePerUnit,
		Timestamp:        Timestamp(productReceiveItem.Timestamp),
	}

	if productReceiveItem.User != nil {
		r.CreatedBy = NewUserResponseP(*productReceiveItem.User)
	}

	if productReceiveItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*productReceiveItem.ProductUnit)
	}

	return r
}

func NewProductReceiveItemResponseP(productReceiveImage model.ProductReceiveItem) *ProductReceiveItemResponse {
	r := NewProductReceiveItemResponse(productReceiveImage)

	return &r
}
