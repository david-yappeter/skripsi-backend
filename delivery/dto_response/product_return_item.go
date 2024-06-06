package dto_response

import (
	"myapp/model"
)

type ProductReturnItemResponse struct {
	Id              string  `json:"id"`
	ProductReturnId string  `json:"product_return_id"`
	ProductUnitId   string  `json:"product_unit_id"`
	Qty             float64 `json:"qty"`
	ScaleToBase     float64 `json:"scale_to_base"`
	BaseCostPrice   float64 `json:"base_cost_price"`
	Timestamp

	ProductUnit *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
	CreatedBy   *UserResponse        `json:"created_by" extensions:"x-nullable"`
} // @name ProductReturnItemResponse

func NewProductReturnItemResponse(productReturnItem model.ProductReturnItem) ProductReturnItemResponse {
	r := ProductReturnItemResponse{
		Id:              productReturnItem.Id,
		ProductReturnId: productReturnItem.ProductReturnId,
		ProductUnitId:   productReturnItem.ProductUnitId,
		Qty:             productReturnItem.Qty,
		ScaleToBase:     productReturnItem.ScaleToBase,
		BaseCostPrice:   productReturnItem.BaseCostPrice,
		Timestamp:       Timestamp(productReturnItem.Timestamp),
	}

	if productReturnItem.User != nil {
		r.CreatedBy = NewUserResponseP(*productReturnItem.User)
	}

	if productReturnItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*productReturnItem.ProductUnit)
	}

	return r
}

func NewProductReturnItemResponseP(productReturnImage model.ProductReturnItem) *ProductReturnItemResponse {
	r := NewProductReturnItemResponse(productReturnImage)

	return &r
}
