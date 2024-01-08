package dto_response

import "myapp/model"

type SupplierTypeResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description" extensions:"x-nullable"`
	Timestamp
} // @name SupplierTypeResponse

func NewSupplierTypeResponse(supplierType model.SupplierType) SupplierTypeResponse {
	r := SupplierTypeResponse{
		Id:          supplierType.Id,
		Name:        supplierType.Name,
		Description: supplierType.Description,
		Timestamp:   Timestamp(supplierType.Timestamp),
	}

	return r
}

func NewSupplierTypeResponseP(supplierType model.SupplierType) *SupplierTypeResponse {
	r := NewSupplierTypeResponse(supplierType)

	return &r
}
