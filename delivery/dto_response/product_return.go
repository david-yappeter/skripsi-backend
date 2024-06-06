package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type ProductReturnResponse struct {
	Id            string                        `json:"id"`
	SupplierId    string                        `json:"supplier_id"`
	UserId        string                        `json:"user_id"`
	InvoiceNumber string                        `json:"invoice_number"`
	Date          data_type.Date                `json:"date"`
	Status        data_type.ProductReturnStatus `json:"status"`

	Items    []ProductReturnItemResponse  `json:"items" extensions:"x-nullable"`
	Images   []ProductReturnImageResponse `json:"images" extensions:"x-nullable"`
	Supplier *SupplierResponse            `json:"supplier" extensions:"x-nullable"`
	Timestamp
} // @name ProductReturnResponse

func NewProductReturnResponse(productReceive model.ProductReturn) ProductReturnResponse {
	r := ProductReturnResponse{
		Id:            productReceive.Id,
		SupplierId:    productReceive.SupplierId,
		UserId:        productReceive.UserId,
		InvoiceNumber: productReceive.InvoiceNumber,
		Date:          productReceive.Date,
		Status:        productReceive.Status,
		Timestamp:     Timestamp(productReceive.Timestamp),
		Items:         []ProductReturnItemResponse{},
	}

	if productReceive.Supplier != nil {
		r.Supplier = NewSupplierResponseP(*productReceive.Supplier)
	}

	for _, productReceiveImage := range productReceive.ProductReturnImages {
		r.Images = append(r.Images, NewProductReturnImageResponse(productReceiveImage))
	}

	for _, productReceiveItem := range productReceive.ProductReturnItems {
		r.Items = append(r.Items, NewProductReturnItemResponse(productReceiveItem))
	}

	return r
}

func NewProductReturnResponseP(supplierType model.ProductReturn) *ProductReturnResponse {
	r := NewProductReturnResponse(supplierType)

	return &r
}
