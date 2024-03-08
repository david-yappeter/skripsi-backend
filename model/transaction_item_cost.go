package model

const TransactionItemCostTableName = "transaction_item_cost"

type TransactionItemCost struct {
	Id                string  `db:"id"`
	TransactionItemId string  `db:"transaction_item_id"`
	Qty               float64 `db:"qty"`
	BaseCostPrice     float64 `db:"base_cost_price"`
	TotalCostPrice    float64 `db:"total_cost_price"`

	Timestamp
}

func (m *TransactionItemCost) TableName() string {
	return TransactionItemCostTableName
}

func (m *TransactionItemCost) TableIds() []string {
	return []string{"id"}
}

func (m *TransactionItemCost) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                  m.Id,
		"transaction_item_id": m.TransactionItemId,
		"qty":                 m.Qty,
		"base_cost_price":     m.BaseCostPrice,
		"total_cost_price":    m.TotalCostPrice,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
	}
}

type TransactionItemCostQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &TransactionItemCostQueryOption{}

func (o *TransactionItemCostQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"tic.*"}
	}
}

func (o *TransactionItemCostQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
