package dto_request

type ProductStockAdjustmentFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=created_at updated_at" example:"created_at"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductStockAdjustmentFetchSorts

type ProductStockAdjustmentFetchRequest struct {
	PaginationRequest
	Sorts          ProductStockAdjustmentFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	UserId         *string                          `json:"user_id" validate:"omitempty,not_empty,uuid"`
	ProductStockId *string                          `json:"product_stock_id" validate:"omitempty,not_empty,uuid"`
} // @name ProductStockAdjustmentFetchRequest
