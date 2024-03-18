package dto_request

type ProductStockFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductStockFetchSorts

type ProductStockFetchRequest struct {
	PaginationRequest
	Sorts  ProductStockFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductStockFetchRequest

type ProductStockGetRequest struct {
	ProductStockId string `json:"-" swaggerignore:"true"`
} // @name ProductStockGetRequest

type ProductStockAdjustmentRequest struct {
	Qty       float64  `json:"qty"`
	CostPrice *float64 `json:"cost_price" validate:"omitempty,gte=0"`

	ProductStockId string `json:"-" swaggerignore:"true"`
} // @name ProductStockAdjustmentRequest
