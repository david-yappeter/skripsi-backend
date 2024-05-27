package dto_response

import "myapp/model"

type ProductStockAdjustmentResponse struct {
	Id             string  `json:"id"`
	UserId         string  `json:"user_id"`
	ProductStockId string  `json:"product_stock_id"`
	PreviousQty    float64 `json:"previous_qty"`
	UpdatedQty     float64 `json:"updated_qty"`

	Timestamp

	User *UserResponse `json:"user" extensions:"x-nullable"`
} // @name ProductStockAdjustmentResponse

func NewProductStockAdjustmentResponse(productStockAdjustment model.ProductStockAdjustment) ProductStockAdjustmentResponse {
	r := ProductStockAdjustmentResponse{
		Id:             productStockAdjustment.Id,
		UserId:         productStockAdjustment.UserId,
		ProductStockId: productStockAdjustment.ProductStockId,
		PreviousQty:    productStockAdjustment.PreviousQty,
		UpdatedQty:     productStockAdjustment.UpdatedQty,
		Timestamp:      Timestamp(productStockAdjustment.Timestamp),
	}

	if productStockAdjustment.User != nil {
		r.User = NewUserResponseP(*productStockAdjustment.User)
	}

	return r
}
