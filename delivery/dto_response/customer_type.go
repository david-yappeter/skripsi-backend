package dto_response

import "myapp/model"

type CustomerTypeResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description" extensions:"x-nullable"`
	Timestamp

	Discounts []CustomerTypeDiscountResponse `json:"discounts" extensions:"x-nullable"`
} // @name CustomerTypeResponse

func NewCustomerTypeResponse(customerType model.CustomerType) CustomerTypeResponse {
	r := CustomerTypeResponse{
		Id:          customerType.Id,
		Name:        customerType.Name,
		Description: customerType.Description,
		Timestamp:   Timestamp(customerType.Timestamp),
	}

	for _, discount := range customerType.CustomerTypeDiscounts {
		r.Discounts = append(r.Discounts, NewCustomerTypeDiscountResponse(discount))
	}

	return r
}
