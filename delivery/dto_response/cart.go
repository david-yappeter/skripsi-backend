package dto_response

import (
	"math"
	"myapp/model"
)

type CartResponse struct {
	Id               string  `json:"id"`
	CashierSessionId string  `json:"cashier_session_id"`
	Name             *string `json:"name" extensions:"x-nullable"`
	IsActive         bool    `json:"is_active"`
	Timestamp

	Subtotal       float64                 `json:"subtotal"`
	TotalDiscount  float64                 `json:"total_discount"`
	GrandTotal     float64                 `json:"grand_total"`
	CashierSession *CashierSessionResponse `json:"cashier_session" extensions:"x-nullable"`
	Items          []CartItemResponse      `json:"items" extensions:"x-nullable"`
} // @name CartResponse

func NewCartResponse(cart model.Cart) CartResponse {
	r := CartResponse{
		Id:               cart.Id,
		CashierSessionId: cart.CashierSessionId,
		Name:             cart.Name,
		IsActive:         cart.IsActive,
		Timestamp:        Timestamp(cart.Timestamp),
		Subtotal:         cart.Subtotal(),
		TotalDiscount:    cart.TotalDiscount(),
	}

	r.GrandTotal = math.Max(r.Subtotal-r.TotalDiscount, 0)

	if cart.CashierSession != nil {
		r.CashierSession = NewCashierSessionResponseP(*cart.CashierSession)
	}

	if len(cart.CartItems) > 0 {
		for _, cartItem := range cart.CartItems {
			r.Items = append(r.Items, NewCartItemResponse(cartItem))
		}
	}

	return r
}

func NewCartResponseP(cart model.Cart) *CartResponse {
	r := NewCartResponse(cart)

	return &r
}
