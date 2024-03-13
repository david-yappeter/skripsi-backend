package model

import "myapp/data_type"

const ShopOrderTableName = "shop_orders"

type ShopOrder struct {
	Id                        string                            `db:"id"`
	TrackingNumber            *string                           `db:"tracking_number"`
	PlatformIdentifier        string                            `db:"platform_identifier"`
	PlatformType              data_type.ShopOrderPlatformType   `db:"platform_type"`
	TrackingStatus            data_type.ShopOrderTrackingStatus `db:"tracking_status"`
	RecipientName             string                            `db:"recipient_name"`
	RecipientFullAddress      string                            `db:"recipient_full_address"`
	RecipientPhoneNumber      string                            `db:"recipient_phone_number"`
	ShippingFee               float64                           `db:"shipping_fee"`
	TotalOriginalProductPrice float64                           `db:"total_original_product_price"`
	Subtotal                  float64                           `db:"subtotal"`
	Tax                       float64                           `db:"tax"`
	TotalAmount               float64                           `db:"total_amount"`

	Timestamp

	ShopOrderItems []ShopOrderItem `db:"-"`
}

func (m *ShopOrder) TableName() string {
	return ShopOrderTableName
}

func (m *ShopOrder) TableIds() []string {
	return []string{"id"}
}

func (m *ShopOrder) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                           m.Id,
		"tracking_number":              m.TrackingNumber,
		"platform_identifier":          m.PlatformIdentifier,
		"platform_type":                m.PlatformType,
		"tracking_status":              m.TrackingStatus,
		"recipient_name":               m.RecipientName,
		"recipient_full_address":       m.RecipientFullAddress,
		"recipient_phone_number":       m.RecipientPhoneNumber,
		"shipping_fee":                 m.ShippingFee,
		"total_original_product_price": m.TotalOriginalProductPrice,
		"subtotal":                     m.Subtotal,
		"tax":                          m.Tax,
		"total_amount":                 m.TotalAmount,
		"created_at":                   m.CreatedAt,
		"updated_at":                   m.UpdatedAt,
	}
}

type ShopOrderQueryOption struct {
	QueryOption

	TrackingStatus *data_type.ShopOrderTrackingStatus
	PlatformType   *data_type.ShopOrderPlatformType
	Phrase         *string
}

var _ PrepareOption = &ShopOrderQueryOption{}

func (o *ShopOrderQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"so.*"}
	}
}

func (o *ShopOrderQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
