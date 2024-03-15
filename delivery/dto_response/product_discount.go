package dto_response

import "myapp/model"

type ProductDiscountResponse struct {
	Id                 string   `json:"id"`
	ProductId          string   `json:"product_id"`
	MinimumQty         float64  `json:"minimum_qty"`
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage"`
	DiscountAmount     *float64 `json:"discount_amount"`

	Timestamp
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

	return r
}
