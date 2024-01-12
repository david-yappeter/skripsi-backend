package dto_response

import "myapp/model"

type BalanceResponse struct {
	Id            string  `json:"id"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	BankName      string  `json:"bank_name"`
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`

	Timestamp
} // @name BalanceResponse

func NewBalanceResponse(balance model.Balance) BalanceResponse {
	r := BalanceResponse{
		Id:            balance.Id,
		AccountNumber: balance.AccountNumber,
		AccountName:   balance.AccountName,
		BankName:      balance.BankName,
		Name:          balance.Name,
		Amount:        balance.Amount,
		Timestamp:     Timestamp(balance.Timestamp),
	}

	return r
}
