package dto_request

type BalanceCreateRequest struct {
	AccountNumber string `json:"account_number" validate:"required,not_empty"`
	AccountName   string `json:"account_name" validate:"required,not_empty"`
	BankName      string `json:"bank_name" validate:"required,not_empty"`
	Name          string `json:"name" validate:"required,not_empty"`
} // @name BalanceCreateRequest

type BalanceFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=account_number account_name bank_name name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name BalanceFetchSorts

type BalanceFetchRequest struct {
	PaginationRequest
	Sorts  BalanceFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string           `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name BalanceFetchRequest

type BalanceGetRequest struct {
	BalanceId string `json:"-" swaggerignore:"true"`
} // @name BalanceGetRequest

type BalanceUpdateRequest struct {
	AccountNumber string `json:"account_number" validate:"required,not_empty"`
	AccountName   string `json:"account_name" validate:"required,not_empty"`
	BankName      string `json:"bank_name" validate:"required,not_empty"`
	Name          string `json:"name" validate:"required,not_empty"`

	BalanceId string `json:"-" swaggerignore:"true"`
} // @name BalanceUpdateRequest

type BalanceDeleteRequest struct {
	BalanceId string `json:"-" swaggerignore:"true"`
} // @name BalanceDeleteRequest
