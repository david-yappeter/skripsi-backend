package dto_response

import (
	"myapp/model"
)

type TransactionItemResponse struct {
	Id              string   `json:"id"`
	TransactionId   string   `json:"transaction_id"`
	ProductUnitId   string   `json:"product_unit_id"`
	Qty             float64  `json:"qty"`
	PricePerUnit    float64  `json:"price_per_unit"`
	DiscountPerUnit *float64 `json:"discount_per_unit" extensions:"x-nullable"`

	Timestamp

	ProductUnit *ProductUnitResponse `json:"product_unit" extensions:"x-nullable"`
} // @name TransactionItemResponse

func NewTransactionItemResponse(transactionItem model.TransactionItem) TransactionItemResponse {
	r := TransactionItemResponse{
		Id:              transactionItem.Id,
		TransactionId:   transactionItem.TransactionId,
		ProductUnitId:   transactionItem.ProductUnitId,
		Qty:             transactionItem.Qty,
		PricePerUnit:    transactionItem.PricePerUnit,
		DiscountPerUnit: transactionItem.DiscountPerUnit,
		Timestamp:       Timestamp(transactionItem.Timestamp),
	}

	if transactionItem.ProductUnit != nil {
		r.ProductUnit = NewProductUnitResponseP(*transactionItem.ProductUnit)
	}

	return r
}
