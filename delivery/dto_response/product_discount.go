package dto_response

import (
	"myapp/model"
	"myapp/util"
)

type ProductDiscountResponse struct {
	Id                 string   `json:"id"`
	ProductId          string   `json:"product_id"`
	MinimumQty         float64  `json:"minimum_qty"`
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage"`
	DiscountAmount     *float64 `json:"discount_amount"`

	Timestamp

	Product *ProductResponse `json:"product"`
} // @name ProductDiscountResponse

func NewProductDiscountResponse(productDiscount model.ProductDiscount) ProductDiscountResponse {
	r := ProductDiscountResponse{
		Id:                 productDiscount.Id,
		ProductId:          productDiscount.ProductId,
		MinimumQty:         productDiscount.MinimumQty,
		IsActive:           productDiscount.IsActive,
		DiscountPercentage: productDiscount.DiscountPercentage,
		DiscountAmount:     productDiscount.DiscountAmount,
		Timestamp:          Timestamp(productDiscount.Timestamp),
	}

	if productDiscount.Product != nil {
		r.Product = util.Pointer(NewProductResponse(*productDiscount.Product))
	}

	return r
}
