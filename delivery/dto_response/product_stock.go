package dto_response

import "myapp/model"

type ProductStockResponse struct {
	Id        string  `json:"id"`
	ProductId string  `json:"product_id"`
	Qty       float64 `json:"qty"`

	Timestamp
} // @name ProductStockResponse

func NewProductStockResponse(productStock model.ProductStock) ProductStockResponse {
	r := ProductStockResponse{
		Id:        productStock.Id,
		ProductId: productStock.ProductId,
		Qty:       productStock.Qty,
		Timestamp: Timestamp(productStock.Timestamp),
	}

	return r
}
