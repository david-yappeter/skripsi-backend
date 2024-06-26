package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type CashierSessionResponse struct {
	Id           string                         `json:"id"`
	UserId       string                         `json:"user_id"`
	Status       data_type.CashierSessionStatus `json:"status"`
	StartingCash float64                        `json:"starting_cash"`
	EndingCash   *float64                       `json:"ending_cash" extensions:"x-nullable"`
	StartedAt    data_type.DateTime             `json:"started_at"`
	EndedAt      data_type.NullDateTime         `json:"ended_at" extensions:"x-nullable"`
	Timestamp

	User *UserResponse `json:"user" extensions:"x-nullable"`
} // @name CashierSessionResponse

func NewCashierSessionResponse(cashierSession model.CashierSession) CashierSessionResponse {
	r := CashierSessionResponse{
		Id:           cashierSession.Id,
		UserId:       cashierSession.UserId,
		Status:       cashierSession.Status,
		StartingCash: cashierSession.StartingCash,
		EndingCash:   cashierSession.EndingCash,
		StartedAt:    cashierSession.StartedAt,
		EndedAt:      cashierSession.EndedAt,
		Timestamp:    Timestamp(cashierSession.Timestamp),
	}

	if cashierSession.User != nil {
		r.User = NewUserResponseP(*cashierSession.User)
	}

	return r
}

func NewCashierSessionResponseP(cashierSession model.CashierSession) *CashierSessionResponse {
	r := NewCashierSessionResponse(cashierSession)
	return &r
}
