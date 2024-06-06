package model

import "myapp/data_type"

const ProductReturnTableName = "product_returns"

type ProductReturn struct {
	Id            string                        `db:"id"`
	SupplierId    string                        `db:"supplier_id"`
	UserId        string                        `db:"user_id"`
	InvoiceNumber string                        `db:"invoice_number"`
	Date          data_type.Date                `db:"date"`
	Status        data_type.ProductReturnStatus `db:"status"`
	Timestamp

	Supplier            *Supplier            `db:"-"`
	ProductReturnItems  []ProductReturnItem  `db:"-"`
	ProductReturnImages []ProductReturnImage `db:"-"`
}

func (m *ProductReturn) TableName() string {
	return ProductReturnTableName
}

func (m *ProductReturn) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReturn) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             m.Id,
		"supplier_id":    m.SupplierId,
		"user_id":        m.UserId,
		"invoice_number": m.InvoiceNumber,
		"date":           m.Date,
		"status":         m.Status,
		"created_at":     m.CreatedAt,
		"updated_at":     m.UpdatedAt,
	}
}

type ProductReturnQueryOption struct {
	QueryOption

	Status     *data_type.ProductReturnStatus
	SupplierId *string
	Phrase     *string
}

var _ PrepareOption = &ProductReturnQueryOption{}

func (o *ProductReturnQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"pr.*"}
	}
}

func (o *ProductReturnQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
