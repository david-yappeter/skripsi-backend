package dto_response

import (
	"myapp/model"
	"myapp/util"
)

type CustomerTypeDiscountResponse struct {
	Id                 string   `json:"id"`
	ProductId          string   `json:"product_id"`
	CustomerTypeId     string   `json:"customer_type_id"`
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage"`
	DiscountAmount     *float64 `json:"discount_amount"`

	Timestamp

	Product *ProductResponse `json:"product" extensions:"x-nullable"`
} // @name CustomerTypeDiscountResponse

func NewCustomerTypeDiscountResponse(customerTypeDiscount model.CustomerTypeDiscount) CustomerTypeDiscountResponse {
	r := CustomerTypeDiscountResponse{
		Id:                 customerTypeDiscount.Id,
		ProductId:          customerTypeDiscount.ProductId,
		CustomerTypeId:     customerTypeDiscount.CustomerTypeId,
		IsActive:           customerTypeDiscount.IsActive,
		DiscountPercentage: customerTypeDiscount.DiscountPercentage,
		DiscountAmount:     customerTypeDiscount.DiscountAmount,
		Timestamp:          Timestamp(customerTypeDiscount.Timestamp),
	}

	if customerTypeDiscount.Product != nil {
		r.Product = util.Pointer(NewProductResponse(*customerTypeDiscount.Product))
	}

	return r
}
