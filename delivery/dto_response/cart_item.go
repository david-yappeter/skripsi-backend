package dto_response

import "myapp/model"

type CartItemResponse struct {
	Id            string  `json:"id"`
	CartId        string  `json:"cart_id"`
	ProductUnitId string  `json:"product_unit_id"`
	Qty           float64 `json:"qty"`
	Timestamp

	Subtotal      float64              `json:"subtotal"`
	TotalDiscount *float64             `json:"total_discount" extensions:"x-nullable"`
	Total         float64              `json:"total"`
	ProductUnit   *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
} // @name CartItemResponse

func NewCartItemResponse(cartItem model.CartItem) CartItemResponse {
	r := CartItemResponse{
		Id:            cartItem.Id,
		CartId:        cartItem.CartId,
		ProductUnitId: cartItem.ProductUnitId,
		Qty:           cartItem.Qty,
		Timestamp:     Timestamp(cartItem.Timestamp),
		Subtotal:      cartItem.Subtotal(),
	}

	r.Total = r.Subtotal

	totalDiscount := cartItem.TotalDiscount()
	if totalDiscount != 0 {
		r.TotalDiscount = &totalDiscount
		r.Total -= totalDiscount
	}

	if cartItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*cartItem.ProductUnit)
	}

	return r
}
