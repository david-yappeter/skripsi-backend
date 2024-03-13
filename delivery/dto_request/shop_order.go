package dto_request

import "myapp/data_type"

type ShopOrderFetchRequest struct {
	PaginationRequest
	Phrase         *string                            `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
	PlatformType   *data_type.ShopOrderPlatformType   `json:"platform_type" validate:"omitempty,data_type_enum"`
	TrackingStatus *data_type.ShopOrderTrackingStatus `json:"tracking_status" validate:"omitempty,data_type_enum"`
	WithItems      bool                               `json:"with_items"`
} // @name ShopOrderFetchRequest

type ShopOrderGetRequest struct {
	ShopOrderId string `json:"-" swaggerignore:"true"`
} // @name ShopOrderGetRequest
