package dto_request

import "myapp/data_type"

type WebhookOrderStatusChangeRequest struct {
	OrderId       string                                   `json:"order_id"`
	OrderStatus   data_type.WebhookTiktokOrderStatusChange `json:"order_status"`
	IsOnHoldOrder bool                                     `json:"is_hold_order"`
	UpdateTime    int64                                    `json:"update_time"`
}
