package model

import "myapp/data_type"

const ShopOrderTableName = "shop_orders"

type ShopOrder struct {
	Id                   string                          `db:"id"`
	TrackingNumber       string                          `db:"tracking_number"`
	PlatformType         data_type.ShopOrderPlatformType `db:"platform_type"`
	TrackingStatus       string                          `db:"tracking_status"`
	RecipientName        string                          `db:"recipient_name"`
	RecipientFullAddress string                          `db:"recipient_full_address"`
	RecipientPhoneNumber string                          `db:"recipient_phone_number"`
	OriginalPrice        float64                         `db:"original_price"`
	SalePrice            float64                         `db:"sale_price"`

	Timestamp
}

func (m *ShopOrder) TableName() string {
	return ShopOrderTableName
}

func (m *ShopOrder) TableIds() []string {
	return []string{"id"}
}

func (m *ShopOrder) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                     m.Id,
		"tracking_number":        m.TrackingNumber,
		"platform_type":          m.PlatformType,
		"tracking_status":        m.TrackingStatus,
		"recipient_name":         m.RecipientName,
		"recipient_full_address": m.RecipientFullAddress,
		"recipient_phone_number": m.RecipientPhoneNumber,
		"original_price":         m.OriginalPrice,
		"sale_price":             m.SalePrice,
		"created_at":             m.CreatedAt,
		"updated_at":             m.UpdatedAt,
	}
}

type ShopOrderQueryOption struct {
	QueryOption

	Phrase *string
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
