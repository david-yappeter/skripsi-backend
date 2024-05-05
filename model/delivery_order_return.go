package model

const DeliveryOrderReturnTableName = "delivery_order_returns"

type DeliveryOrderReturn struct {
	Id              string  `db:"id"`
	DeliveryOrderId string  `db:"delivery_order_id"`
	UserId          string  `db:"user_id"`
	Description     *string `db:"description"`

	Timestamp

	User                      *User                      `db:"-"`
	DeliveryOrderReturnImages []DeliveryOrderReturnImage `db:"-"`
}

func (m *DeliveryOrderReturn) TableName() string {
	return DeliveryOrderReturnTableName
}

func (m *DeliveryOrderReturn) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderReturn) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"delivery_order_id": m.DeliveryOrderId,
		"user_id":           m.UserId,
		"description":       m.Description,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
