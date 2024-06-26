package dto_response

import "myapp/model"

type SupplierResponse struct {
	Id             string  `json:"id"`
	SupplierTypeId string  `json:"supplier_type_id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	IsActive       bool    `json:"is_active"`
	Address        string  `json:"address"`
	Phone          string  `json:"phone"`
	Email          *string `json:"email" extensions:"x-nullable"`
	Description    *string `json:"description" extensions:"x-nullable"`

	Timestamp

	SupplierType *SupplierTypeResponse `json:"supplier_type" extensions:"x-nullable"`
} // @name SupplierResponse

func NewSupplierResponse(supplier model.Supplier) SupplierResponse {
	r := SupplierResponse{
		Id:             supplier.Id,
		SupplierTypeId: supplier.SupplierTypeId,
		Code:           supplier.Code,
		Name:           supplier.Name,
		IsActive:       supplier.IsActive,
		Address:        supplier.Address,
		Phone:          supplier.Phone,
		Email:          supplier.Email,
		Description:    supplier.Description,
		Timestamp:      Timestamp(supplier.Timestamp),
	}

	if supplier.SupplierType != nil {
		r.SupplierType = NewSupplierTypeResponseP(*supplier.SupplierType)
	}

	return r
}

func NewSupplierResponseP(supplier model.Supplier) *SupplierResponse {
	r := NewSupplierResponse(supplier)
	return &r
}

type SupplierDebtSummaryResponse struct {
	SupplierId   string  `json:"supplier_id"`
	SupplierName string  `json:"supplier_name"`
	TotalDebt    float64 `json:"total_debt"`
}

func NewSupplierDebtSummaryResponse(supplierDebtSummary model.SupplierDebtSummary) SupplierDebtSummaryResponse {
	r := SupplierDebtSummaryResponse{
		SupplierId:   supplierDebtSummary.SupplierId,
		SupplierName: supplierDebtSummary.SupplierName,
		TotalDebt:    supplierDebtSummary.TotalDebt,
	}
	return r
}
