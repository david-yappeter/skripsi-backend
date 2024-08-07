package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type CustomerDebtResponse struct {
	Id              string                           `json:"id"`
	CustomerId      string                           `json:"customer_id"`
	DebtSource      data_type.CustomerDebtDebtSource `json:"debt_source"`
	DueDate         data_type.NullDate               `json:"due_date"`
	Status          data_type.CustomerDebtStatus     `json:"status"`
	Amount          float64                          `json:"amount"`
	RemainingAmount float64                          `json:"remaining_amount"`

	Timestamp

	Customer      *CustomerResponse         `json:"customer" extensions:"x-nullable"`
	Payments      []CustomerPaymentResponse `json:"payments" extensions:"x-nullable"`
	DeliveryOrder *DeliveryOrderResponse    `json:"delivery_order" extensions:"x-nullable"`
} // @name CustomerDebtResponse

func NewCustomerDebtResponse(customerDebt model.CustomerDebt) CustomerDebtResponse {
	r := CustomerDebtResponse{
		Id:              customerDebt.Id,
		CustomerId:      customerDebt.CustomerId,
		DebtSource:      customerDebt.DebtSource,
		DueDate:         customerDebt.DueDate,
		Status:          customerDebt.Status,
		Amount:          customerDebt.Amount,
		RemainingAmount: customerDebt.RemainingAmount,
		Timestamp:       Timestamp(customerDebt.Timestamp),
	}

	if customerDebt.Customer != nil {
		r.Customer = NewCustomerResponseP(*customerDebt.Customer)
	}

	if customerDebt.DeliveryOrder != nil {
		r.DeliveryOrder = NewDeliveryOrderResponseP(*customerDebt.DeliveryOrder)
	}

	for _, customerPayment := range customerDebt.CustomerPayments {
		r.Payments = append(r.Payments, NewCustomerPaymentResponse(customerPayment))
	}

	return r
}
