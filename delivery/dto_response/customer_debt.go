package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type CustomerDebtResponse struct {
	Id              string                           `json:"id"`
	CustomerId      string                           `json:"customer_id"`
	DebtSource      data_type.CustomerDebtDebtSource `json:"debt_source"`
	DebtSourceId    string                           `json:"debt_source_id"`
	DueDate         data_type.NullDate               `json:"due_date"`
	Status          data_type.CustomerDebtStatus     `json:"status"`
	Amount          float64                          `json:"amount"`
	RemainingAmount float64                          `json:"remaining_amount"`

	Timestamp
} // @name CustomerDebtResponse

func NewCustomerDebtResponse(unit model.CustomerDebt) CustomerDebtResponse {
	r := CustomerDebtResponse{
		Id:              unit.Id,
		CustomerId:      unit.CustomerId,
		DebtSource:      unit.DebtSource,
		DebtSourceId:    unit.DebtSourceId,
		DueDate:         unit.DueDate,
		Status:          unit.Status,
		Amount:          unit.Amount,
		RemainingAmount: unit.RemainingAmount,
		Timestamp:       Timestamp(unit.Timestamp),
	}

	return r
}
