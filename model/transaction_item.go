package model

const TransactionItemTableName = "transaction_items"

type TransactionItem struct {
	Id            string  `db:"id"`
	TransactionId string  `db:"transaction_id"`
	ProductUnitId string  `db:"product_unit_id"`
	Qty           float64 `db:"qty"`
	Timestamp

	ProductUnit *ProductUnit `db:"-"`
}

func (m *TransactionItem) TableName() string {
	return TransactionItemTableName
}

func (m *TransactionItem) TableIds() []string {
	return []string{"id"}
}

func (m *TransactionItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              m.Id,
		"transaction_id":  m.TransactionId,
		"product_unit_id": m.ProductUnitId,
		"qty":             m.Qty,
		"created_at":      m.CreatedAt,
		"updated_at":      m.UpdatedAt,
	}
}

type TransactionItemQueryOption struct {
	QueryOption

	CartId *string
	Phrase *string
}

var _ PrepareOption = &TransactionItemQueryOption{}

func (o *TransactionItemQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"u.*"}
	}
}

func (o *TransactionItemQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
