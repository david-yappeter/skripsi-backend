package dto_response

import "myapp/model"

type ProductUnitResponse struct {
	Id          string  `json:"id"`
	ToUnitId    *string `json:"to_unit_id" extensions:"x-nullable"`
	ImageFileId *string `json:"image_file_id" extensions:"x-nullable"`
	UnitId      string  `json:"unit_id"`
	ProductId   string  `json:"product_id"`
	Scale       float64 `json:"scale"`
	ScaleToBase float64 `json:"scale_to_base"`

	Timestamp
} // @name ProductUnitResponse

func NewProductUnitResponse(productUnit model.ProductUnit) ProductUnitResponse {
	r := ProductUnitResponse{
		Id:          productUnit.Id,
		ToUnitId:    productUnit.ToUnitId,
		ImageFileId: productUnit.ImageFileId,
		UnitId:      productUnit.UnitId,
		ProductId:   productUnit.ProductId,
		Scale:       productUnit.Scale,
		ScaleToBase: productUnit.ScaleToBase,
		Timestamp:   Timestamp(productUnit.Timestamp),
	}

	return r
}
