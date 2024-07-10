package dto_response

import "myapp/model"

type ProductStockResponse struct {
	Id            string  `json:"id"`
	ProductId     string  `json:"product_id"`
	Qty           float64 `json:"qty"`
	BaseCostPrice float64 `json:"base_cost_price"`

	Timestamp
} // @name ProductStockResponse

func NewProductStockResponse(productStock model.ProductStock) ProductStockResponse {
	r := ProductStockResponse{
		Id:            productStock.Id,
		ProductId:     productStock.ProductId,
		Qty:           productStock.Qty,
		BaseCostPrice: productStock.BaseCostPrice,
		Timestamp:     Timestamp(productStock.Timestamp),
	}

	return r
}
