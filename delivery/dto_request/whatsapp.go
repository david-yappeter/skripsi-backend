package dto_request

type WhatsappCustomerDebtBroadcastRequest struct {
	CustomerId string `json:"customer_id" validate:"required,not_empty,uuid"`
} // @name WhatsappCustomerDebtBroadcastRequest

type WhatsappCustomerTypeDiscountBroadcastRequest struct {
	CustomerTypeDiscountId string `json:"customer_type_discount_id" validate:"required,not_empty,uuid"`
} // @name WhatsappCustomerTypeDiscountBroadcastRequest

type WhatsappCustomerTypeDiscountManyProductBroadcastRequest struct {
	CustomerTypeDiscountIds []string `json:"customer_type_discount_ids" validate:"required,dive,not_empty,uuid"`
} // @name WhatsappCustomerTypeDiscountManyProductBroadcastRequest

type WhatsappProductPriceChangeBroadcastRequest struct {
	ProductId      string  `json:"product_id" validate:"required,not_empty,uuid"`
	CustomerTypeId string  `json:"customer_type_id" validate:"required,not_empty,uuid"`
	OldPrice       float64 `json:"old_price" validate:"gte=0.0"`
} // @name WhatsappProductPriceChangeBroadcastRequest
