package dto_response

import "myapp/model"

type ProductResponse struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description" extensions:"x-nullable"`
	Price       *float64 `json:"price"`
	IsActive    bool     `json:"is_active"`

	Timestamp
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

	return r
}
