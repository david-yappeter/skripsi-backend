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

	ProductUnits []ProductUnitResponse `json:"product_units"`
	Stock        *ProductStockResponse `json:"stock" extensions:"x-nulalble"`
	ImageFile    *FileResponse         `json:"image_file" extensions:"x-nullable"`
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

	for _, productUnit := range product.ProductUnits {
		r.ProductUnits = append(r.ProductUnits, NewProductUnitResponse(productUnit))
	}

	if product.ProductStock != nil {
		r.Stock = util.Pointer(NewProductStockResponse(*product.ProductStock))
	}

	if product.ImageFile != nil {
		r.ImageFile = NewFileResponseP(*product.ImageFile)
	}

	return r
}
