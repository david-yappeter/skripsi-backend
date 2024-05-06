package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type DebtResponse struct {
	Id              string               `json:"id"`
	DebtSource      data_type.DebtSource `json:"debt_source"`
	DebtSourceId    string               `json:"debt_source_id"`
	DueDate         data_type.NullDate   `json:"due_date"`
	Status          data_type.DebtStatus `json:"status"`
	Amount          float64              `json:"amount"`
	RemainingAmount float64              `json:"remaining_amount"`

	Timestamp
	Payments []DebtPaymentResponse `json:"payments"`
} // @name DebtResponse

func NewDebtResponse(debt model.Debt) DebtResponse {
	r := DebtResponse{
		Id:              debt.Id,
		DebtSource:      debt.DebtSource,
		DebtSourceId:    debt.DebtSourceId,
		DueDate:         debt.DueDate,
		Status:          debt.Status,
		Amount:          debt.Amount,
		RemainingAmount: debt.RemainingAmount,
		Timestamp:       Timestamp(debt.Timestamp),
	}

	for _, debtPayment := range debt.DebtPayments {
		r.Payments = append(r.Payments, NewDebtPaymentResponse(debtPayment))
	}

	return r
}
