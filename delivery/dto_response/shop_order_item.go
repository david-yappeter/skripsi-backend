package dto_response

import (
	"myapp/model"
)

type ShopOrderItemResponse struct {
	Id                string  `json:"id"`
	ProductUnitId     string  `json:"product_unit_id"`
	PlatformProductId string  `json:"platform_product_id"`
	ImageLink         *string `json:"image_link"`
	Quantity          float64 `json:"quantity"`
	OriginalPrice     float64 `json:"original_price"`
	SalePrice         float64 `json:"sale_price"`

	Timestamp

	ProductUnit *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
} // @name ShopOrderItemResponse

func NewShopOrderItemResponse(shopOrderItem model.ShopOrderItem) ShopOrderItemResponse {
	r := ShopOrderItemResponse{
		Id:                shopOrderItem.Id,
		ProductUnitId:     shopOrderItem.ProductUnitId,
		PlatformProductId: shopOrderItem.PlatformProductId,
		ImageLink:         shopOrderItem.ImageLink,
		Quantity:          shopOrderItem.Quantity,
		OriginalPrice:     shopOrderItem.OriginalPrice,
		SalePrice:         shopOrderItem.SalePrice,
		Timestamp:         Timestamp(shopOrderItem.Timestamp),
	}

	if shopOrderItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*shopOrderItem.ProductUnit)
	}

	return r
}
