package dto_response

import (
	"myapp/model"
	"myapp/util"
)

type ProductUnitResponse struct {
	Id          string  `json:"id"`
	ToUnitId    *string `json:"to_unit_id" extensions:"x-nullable"`
	UnitId      string  `json:"unit_id"`
	ProductId   string  `json:"product_id"`
	Scale       float64 `json:"scale"`
	ScaleToBase float64 `json:"scale_to_base"`

	Timestamp

	Product *ProductResponse `json:"product" extensions:"x-nullable"`
	Unit    *UnitResponse    `json:"unit" extensions:"x-nullable"`
} // @name ProductUnitResponse

func NewProductUnitResponse(productUnit model.ProductUnit) ProductUnitResponse {
	r := ProductUnitResponse{
		Id:          productUnit.Id,
		ToUnitId:    productUnit.ToUnitId,
		UnitId:      productUnit.UnitId,
		ProductId:   productUnit.ProductId,
		Scale:       productUnit.Scale,
		ScaleToBase: productUnit.ScaleToBase,
		Timestamp:   Timestamp(productUnit.Timestamp),
	}

	if productUnit.Product != nil {
		r.Product = util.Pointer(NewProductResponse(*productUnit.Product))
	}

	if productUnit.Unit != nil {
		r.Unit = util.Pointer(NewUnitResponse(*productUnit.Unit))
	}

	return r
}

func NewProductUnitResponseP(productUnit model.ProductUnit) *ProductUnitResponse {
	r := NewProductUnitResponse(productUnit)

	return &r
}
