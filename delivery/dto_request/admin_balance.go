package dto_request

type AdminBalanceCreateRequest struct {
	AccountNumber string `json:"account_number" validate:"required,not_empty"`
	AccountName   string `json:"account_name" validate:"required,not_empty"`
	BankName      string `json:"bank_name" validate:"required,not_empty"`
	Name          string `json:"name" validate:"required,not_empty"`
} // @name AdminBalanceCreateRequest

type AdminBalanceFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=account_number account_name bank_name name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name AdminBalanceFetchSorts

type AdminBalanceFetchRequest struct {
	PaginationRequest
	Sorts  AdminBalanceFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminBalanceFetchRequest

type AdminBalanceGetRequest struct {
	BalanceId string `json:"-" swaggerignore:"true"`
} // @name AdminBalanceGetRequest

type AdminBalanceUpdateRequest struct {
	AccountNumber string `json:"account_number" validate:"required,not_empty"`
	AccountName   string `json:"account_name" validate:"required,not_empty"`
	BankName      string `json:"bank_name" validate:"required,not_empty"`
	Name          string `json:"name" validate:"required,not_empty"`

	BalanceId string `json:"-" swaggerignore:"true"`
} // @name AdminBalanceUpdateRequest

type AdminBalanceDeleteRequest struct {
	BalanceId string `json:"-" swaggerignore:"true"`
} // @name AdminBalanceDeleteRequest
