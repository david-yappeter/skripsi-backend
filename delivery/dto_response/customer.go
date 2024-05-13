package dto_response

import (
	"myapp/model"
	"myapp/util"
)

type CustomerResponse struct {
	Id             string  `json:"id"`
	CustomerTypeId *string `json:"customer_type_id" extensions:"x-nullable"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Address        string  `json:"address"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Phone          string  `json:"phone"`
	IsActive       bool    `json:"is_active"`

	Timestamp

	CustomerType *CustomerTypeResponse `json:"customer_type" extensions:"x-nullable"`
} // @name CustomerResponse

func NewCustomerResponse(customer model.Customer) CustomerResponse {
	r := CustomerResponse{
		Id:             customer.Id,
		CustomerTypeId: customer.CustomerTypeId,
		Name:           customer.Name,
		Email:          customer.Email,
		Address:        customer.Address,
		Latitude:       customer.Latitude,
		Longitude:      customer.Longitude,
		Phone:          customer.Phone,
		IsActive:       customer.IsActive,
		Timestamp:      Timestamp(customer.Timestamp),
	}

	if customer.CustomerType != nil {
		r.CustomerType = util.Pointer(NewCustomerTypeResponse(*customer.CustomerType))
	}

	return r
}

func NewCustomerResponseP(customer model.Customer) *CustomerResponse {
	r := NewCustomerResponse(customer)

	return &r
}

type CustomerDebtSummaryResponse struct {
	CustomerId   string  `json:"customer_id"`
	CustomerName string  `json:"customer_name"`
	TotalDebt    float64 `json:"total_debt"`
}

func NewCustomerDebtSummaryResponse(customerDebtSummary model.CustomerDebtSummary) CustomerDebtSummaryResponse {
	r := CustomerDebtSummaryResponse{
		CustomerId:   customerDebtSummary.CustomerId,
		CustomerName: customerDebtSummary.CustomerName,
		TotalDebt:    customerDebtSummary.TotalDebt,
	}
	return r
}
