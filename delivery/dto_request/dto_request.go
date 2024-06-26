package dto_request

type PaginationRequest struct {
	Page  *int `json:"page" validate:"required_with=Limit,omitempty,gte=1" example:"1"`
	Limit *int `json:"limit" validate:"required_with=Page,omitempty,gte=1,lte=100" example:"100"`
} // @name PaginationRequest

type TiktokWebhookBaseRequest[T comparable] struct {
	Type      int    `json:"type"`
	ShopId    string `json:"shop_id"`
	Timestamp int64  `json:"timestamp"`
	Data      T      `json:"data"`
}
