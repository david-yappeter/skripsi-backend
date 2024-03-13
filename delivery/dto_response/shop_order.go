package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type ShopOrderResponse struct {
	Id                        string                            `json:"id"`
	TrackingNumber            *string                           `json:"tracking_number"`
	PlatformIdentifier        string                            `json:"platform_identifier"`
	PlatformType              data_type.ShopOrderPlatformType   `json:"platform_type"`
	TrackingStatus            data_type.ShopOrderTrackingStatus `json:"tracking_status"`
	RecipientName             string                            `json:"recipient_name"`
	RecipientFullAddress      string                            `json:"recipient_full_address"`
	RecipientPhoneNumber      string                            `json:"recipient_phone_number"`
	ShippingFee               float64                           `json:"shipping_fee"`
	TotalOriginalProductPrice float64                           `json:"total_original_product_price"`
	Subtotal                  float64                           `json:"subtotal"`
	Tax                       float64                           `json:"tax"`
	TotalAmount               float64                           `json:"total_amount"`

	Timestamp
} // @name ShopOrderResponse

func NewShopOrderResponse(shopOrder model.ShopOrder) ShopOrderResponse {
	r := ShopOrderResponse{
		Id:                        shopOrder.Id,
		TrackingNumber:            shopOrder.TrackingNumber,
		PlatformIdentifier:        shopOrder.PlatformIdentifier,
		PlatformType:              shopOrder.PlatformType,
		TrackingStatus:            shopOrder.TrackingStatus,
		RecipientName:             shopOrder.RecipientName,
		RecipientFullAddress:      shopOrder.RecipientFullAddress,
		RecipientPhoneNumber:      shopOrder.RecipientPhoneNumber,
		ShippingFee:               shopOrder.ShippingFee,
		TotalOriginalProductPrice: shopOrder.TotalOriginalProductPrice,
		Subtotal:                  shopOrder.Subtotal,
		Tax:                       shopOrder.Tax,
		TotalAmount:               shopOrder.TotalAmount,
		Timestamp:                 Timestamp(shopOrder.Timestamp),
	}

	return r
}
