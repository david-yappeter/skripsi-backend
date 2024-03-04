package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type ProductReceiveResponse struct {
	Id            string                         `json:"id"`
	SupplierId    string                         `json:"supplier_id"`
	UserId        string                         `json:"user_id"`
	InvoiceNumber string                         `json:"invoice_number"`
	Date          data_type.Date                 `json:"date"`
	Status        data_type.ProductReceiveStatus `json:"status"`
	TotalPrice    float64                        `json:"total_price"`

	Items    []ProductReceiveItemResponse `json:"items"`
	Supplier *SupplierResponse            `json:"supplier" extensions:"x-nullable"`
	Timestamp
} // @name ProductReceiveResponse

func NewProductReceiveResponse(productReceive model.ProductReceive) ProductReceiveResponse {
	r := ProductReceiveResponse{
		Id:            productReceive.Id,
		SupplierId:    productReceive.SupplierId,
		UserId:        productReceive.UserId,
		InvoiceNumber: productReceive.InvoiceNumber,
		Date:          productReceive.Date,
		Status:        productReceive.Status,
		TotalPrice:    productReceive.TotalPrice,
		Timestamp:     Timestamp(productReceive.Timestamp),
		Items:         []ProductReceiveItemResponse{},
	}

	if productReceive.Supplier != nil {
		r.Supplier = NewSupplierResponseP(*productReceive.Supplier)
	}

	for _, productReceiveItem := range productReceive.ProductReceiveItems {
		r.Items = append(r.Items, NewProductReceiveItemResponse(productReceiveItem))
	}

	return r
}

func NewProductReceiveResponseP(supplierType model.ProductReceive) *ProductReceiveResponse {
	r := NewProductReceiveResponse(supplierType)

	return &r
}
