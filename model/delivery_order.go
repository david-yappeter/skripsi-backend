package model

import "myapp/data_type"

const DeliveryOrderTableName = "delivery_orders"

type DeliveryOrder struct {
	Id            string                        `db:"id"`
	CustomerId    string                        `db:"customer_id"`
	UserId        string                        `db:"user_id"`
	InvoiceNumber string                        `db:"invoice_number"`
	Date          data_type.Date                `db:"date"`
	Status        data_type.DeliveryOrderStatus `db:"status"`
	TotalPrice    float64                       `db:"total_price"`
	Timestamp

	Customer             *Customer             `db:"-"`
	DeliveryOrderItems   []DeliveryOrderItem   `db:"-"`
	DeliveryOrderImages  []DeliveryOrderImage  `db:"-"`
	DeliveryOrderDrivers []DeliveryOrderDriver `db:"-"`
}

func (m *DeliveryOrder) TableName() string {
	return DeliveryOrderTableName
}

func (m *DeliveryOrder) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrder) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             m.Id,
		"customer_id":    m.CustomerId,
		"user_id":        m.UserId,
		"invoice_number": m.InvoiceNumber,
		"date":           m.Date,
		"status":         m.Status,
		"total_price":    m.TotalPrice,
		"created_at":     m.CreatedAt,
		"updated_at":     m.UpdatedAt,
	}
}

type DeliveryOrderQueryOption struct {
	QueryOption

	ExcludeStatuses []data_type.DeliveryOrderStatus
	Status          *data_type.DeliveryOrderStatus
	CustomerId      *string
	DriverUserId    *string
	Phrase          *string
}

var _ PrepareOption = &DeliveryOrderQueryOption{}

func (o *DeliveryOrderQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"dorder.*"}
	}
}

func (o *DeliveryOrderQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
