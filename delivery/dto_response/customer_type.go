package dto_response

import "myapp/model"

type CustomerTypeResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description" extensions:"x-nullable"`
	Timestamp
} // @name CustomerTypeResponse

func NewCustomerTypeResponse(customerType model.CustomerType) CustomerTypeResponse {
	r := CustomerTypeResponse{
		Id:          customerType.Id,
		Name:        customerType.Name,
		Description: customerType.Description,
		Timestamp:   Timestamp(customerType.Timestamp),
	}

	return r
}
