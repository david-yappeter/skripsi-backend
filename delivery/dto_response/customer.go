package dto_response

import "myapp/model"

type CustomerResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`

	Timestamp
} // @name CustomerResponse

func NewCustomerResponse(customer model.Customer) CustomerResponse {
	r := CustomerResponse{
		Id:        customer.Id,
		Name:      customer.Name,
		Email:     customer.Email,
		Address:   customer.Address,
		Phone:     customer.Phone,
		IsActive:  customer.IsActive,
		Timestamp: Timestamp(customer.Timestamp),
	}

	return r
}

func NewCustomerResponseP(customer model.Customer) *CustomerResponse {
	r := NewCustomerResponse(customer)

	return &r
}
