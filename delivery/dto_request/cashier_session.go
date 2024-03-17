package dto_request

import "myapp/data_type"

type CashierSessionStartRequest struct {
	StartingCash float64 `db:"starting_cash"`
} // @name CashierSessionStartRequest

type CashierSessionFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CashierSessionFetchSorts

type CashierSessionFetchRequest struct {
	PaginationRequest

	UserId    *string                         `json:"user_id" validate:"omitempty,not_empty,uuid"`
	Status    *data_type.CashierSessionStatus `json:"status" validate:"omitempty,data_type_enum"`
	Sorts     CashierSessionFetchSorts        `json:"sorts" validate:"unique=Field,dive"`
	Phrase    *string                         `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
	StartedAt data_type.NullDateTime          `json:"started_at"`
	EndedAt   data_type.NullDateTime          `json:"ended_at"`
} // @name CashierSessionFetchRequest

type CashierSessionGetRequest struct {
	CashierSessionId string `json:"-" swaggerignore:"true"`
} // @name CashierSessionGetRequest

type CashierSessionEndRequest struct {
	EndingCash       float64 `json:"ending_cash"`
	CashierSessionId string  `json:"-" swaggerignore:"true"`
} // @name CashierSessionEndRequest
