package dto_request

import "myapp/data_type"

type CashierSessionStartRequest struct {
	StartingCash float64 `json:"starting_cash" validate:"gte=0"`
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

type CashierSessionFetchTransactionSorts []struct {
	Field     string `json:"field" validate:"required,oneof=payment_at updated_at" example:"payment_at"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"desc"`
} // @name CashierSessionFetchTransactionSorts

type CashierSessionFetchTransactionRequest struct {
	PaginationRequest
	Sorts CashierSessionFetchTransactionSorts `json:"sorts" validate:"unique=Field,dive"`

	CashierSessionId string                       `json:"cashier_session_id" validate:"required,not_empty,uuid"`
	Status           *data_type.TransactionStatus `json:"status" validate:"omitempty,data_type_enum"`
} // @name CashierSessionFetchTransactionRequest

type CashierSessionDownloadReportRequest struct {
	CashierSessionId string `json:"-" swaggerignore:"true"`
} // @name CashierSessionDownloadReportRequest

type CashierSessionEndRequest struct {
} // @name CashierSessionEndRequest
