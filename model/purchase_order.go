package model

import "myapp/data_type"

const PurchaseOrderTableName = "purchase_orders"

type PurchaseOrder struct {
	Id                  string                        `db:"id"`
	SupplierId          string                        `db:"supplier_id"`
	UserId              string                        `db:"user_id"`
	InvoiceNumber       string                        `db:"invoice_number"`
	Date                data_type.Date                `db:"date"`
	Status              data_type.PurchaseOrderStatus `db:"status"`
	TotalEstimatedPrice float64                       `db:"total_estimated_price"`
	Timestamp

	Supplier            *Supplier            `db:"-"`
	PurchaseOrderItems  []PurchaseOrderItem  `db:"-"`
	PurchaseOrderImages []PurchaseOrderImage `db:"-"`
}

func (m *PurchaseOrder) TableName() string {
	return PurchaseOrderTableName
}

func (m *PurchaseOrder) TableIds() []string {
	return []string{"id"}
}

func (m *PurchaseOrder) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                    m.Id,
		"supplier_id":           m.SupplierId,
		"user_id":               m.UserId,
		"invoice_number":        m.InvoiceNumber,
		"date":                  m.Date,
		"status":                m.Status,
		"total_estimated_price": m.TotalEstimatedPrice,
		"created_at":            m.CreatedAt,
		"updated_at":            m.UpdatedAt,
	}
}

type PurchaseOrderQueryOption struct {
	QueryOption

	Status     *data_type.PurchaseOrderStatus
	SupplierId *string
	Phrase     *string
}

var _ PrepareOption = &PurchaseOrderQueryOption{}

func (o *PurchaseOrderQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"po.*"}
	}
}

func (o *PurchaseOrderQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
