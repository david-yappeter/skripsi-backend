package model

import "myapp/data_type"

const ProductReceiveTableName = "product_receives"

type ProductReceive struct {
	Id            string                         `db:"id"`
	SupplierId    string                         `db:"supplier_id"`
	UserId        string                         `db:"user_id"`
	InvoiceNumber string                         `db:"invoice_number"`
	Date          data_type.Date                 `db:"date"`
	Status        data_type.ProductReceiveStatus `db:"status"`
	TotalPrice    float64                        `db:"total_price"`
	Timestamp

	Supplier             *Supplier             `db:"-"`
	ProductReceiveItems  []ProductReceiveItem  `db:"-"`
	ProductReceiveImages []ProductReceiveImage `db:"-"`
}

func (m *ProductReceive) TableName() string {
	return ProductReceiveTableName
}

func (m *ProductReceive) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReceive) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             m.Id,
		"supplier_id":    m.SupplierId,
		"user_id":        m.UserId,
		"invoice_number": m.InvoiceNumber,
		"date":           m.Date,
		"status":         m.Status,
		"total_price":    m.TotalPrice,
		"created_at":     m.CreatedAt,
		"updated_at":     m.UpdatedAt,
	}
}

type ProductReceiveQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &ProductReceiveQueryOption{}

func (o *ProductReceiveQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"pr.*"}
	}
}

func (o *ProductReceiveQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
