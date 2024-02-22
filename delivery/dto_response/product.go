package dto_response

import (
	"myapp/model"
	"myapp/util"
)

type ProductResponse struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description" extensions:"x-nullable"`
	Price       *float64 `json:"price"`
	IsActive    bool     `json:"is_active"`
	Timestamp

	Stock *ProductStockResponse `json:"stock"`
} // @name ProductResponse

func NewProductResponse(product model.Product) ProductResponse {
	r := ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsActive:    product.IsActive,
		Timestamp:   Timestamp(product.Timestamp),
	}

	if product.ProductStock != nil {
		r.Stock = util.Pointer(NewProductStockResponse(*product.ProductStock))
	}

	return r
}
