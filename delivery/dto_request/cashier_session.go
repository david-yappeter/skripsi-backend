package dto_request

type CashierSessionStartRequest struct {
	StartingCash float64 `db:"starting_cash"`
} // @name CashierSessionCreateRequest

type CashierSessionFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CashierSessionFetchSorts

type CashierSessionFetchRequest struct {
	PaginationRequest
	Sorts  CashierSessionFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                  `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CashierSessionFetchRequest

type CashierSessionGetRequest struct {
	CashierSessionId string `json:"-" swaggerignore:"true"`
} // @name CashierSessionGetRequest

type CashierSessionMarkCompleteRequest struct {
	EndingCash       float64 `json:"ending_cash"`
	CashierSessionId string  `json:"-" swaggerignore:"true"`
} // @name CashierSessionMarkCompleteRequest
