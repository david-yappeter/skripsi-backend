package model

const BalanceTableName = "balances"

type Balance struct {
	Id            string  `db:"id"`
	AccountNumber string  `db:"account_number"`
	AccountName   string  `db:"account_name"`
	BankName      string  `db:"bank_name"`
	Name          string  `db:"name"`
	Amount        float64 `db:"amount"`

	Timestamp
}

func (m *Balance) TableName() string {
	return BalanceTableName
}

func (m *Balance) TableIds() []string {
	return []string{"id"}
}

func (m *Balance) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             m.Id,
		"account_number": m.AccountNumber,
		"account_name":   m.AccountName,
		"bank_name":      m.BankName,
		"name":           m.Name,
		"amount":         m.Amount,
		"created_at":     m.CreatedAt,
		"updated_at":     m.UpdatedAt,
	}
}

type BalanceQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &BalanceQueryOption{}

func (o *BalanceQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"b.*"}
	}
}

func (o *BalanceQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
