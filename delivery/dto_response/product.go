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

	ProductUnits  []ProductUnitResponse    `json:"product_units" extensions:"x-nullable"`
	Discount      *ProductDiscountResponse `json:"discount" extensions:"x-nullable"`
	Stock         *ProductStockResponse    `json:"stock" extensions:"x-nullable"`
	ImageFile     *FileResponse            `json:"image_file" extensions:"x-nullable"`
	TiktokProduct *TiktokProductResponse   `json:"tiktok_product" extensions:"x-nullable"`

	IsLoss *bool `json:"is_loss" extensions:"x-nullable"`
} // @name ProductResponse

func NewProductResponse(product model.Product) ProductResponse {
	r := ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsActive:    product.IsActive,
		Timestamp:   Timestamp(product.Timestamp),
		IsLoss:      product.IsLoss(),
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

	if product.ProductDiscount != nil {
		r.Discount = util.Pointer(NewProductDiscountResponse(*product.ProductDiscount))
	}

	if product.TiktokProduct != nil {
		r.TiktokProduct = util.Pointer(NewTiktokProductResponse(*product.TiktokProduct))
	}

	return r
}
