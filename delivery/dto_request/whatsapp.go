package dto_request

type WhatsappProductPriceChangeBroadcastRequest struct {
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	OldPrice  float64 `json:"old_price" validate:"gte=0.0"`
} // @name WhatsappProductPriceChangeBroadcastRequest

type WhatsappCustomerTypeDiscountBroadcastRequest struct {
	CustomerTypeDiscountId string `json:"customer_type_discount_id" validate:"required,not_empty,uuid"`
} // @name WhatsappCustomerTypeDiscountBroadcastRequest
