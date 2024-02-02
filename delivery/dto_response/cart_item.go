package dto_response

import "myapp/model"

type CartItemResponse struct {
	Id            string `json:"id"`
	CartId        string `json:"cart_id"`
	ProductUnitId string `json:"product_unit_id"`
	Timestamp

	ProductUnit *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
} // @name CartItemResponse

func NewCartItemResponse(cartItem model.CartItem) CartItemResponse {
	r := CartItemResponse{
		Id:            cartItem.Id,
		CartId:        cartItem.CartId,
		ProductUnitId: cartItem.ProductUnitId,
		Timestamp:     Timestamp(cartItem.Timestamp),
	}

	if cartItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*cartItem.ProductUnit)
	}

	return r
}
