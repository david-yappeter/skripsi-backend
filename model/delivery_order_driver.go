package model

const DeliveryOrderDriverTableName = "delivery_order_drivers"

type DeliveryOrderDriver struct {
	Id              string `db:"id"`
	DeliveryOrderId string `db:"delivery_order_id"`
	DriverUserId    string `db:"driver_user_id"`

	Timestamp
}

func (m *DeliveryOrderDriver) TableName() string {
	return DeliveryOrderDriverTableName
}

func (m *DeliveryOrderDriver) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderDriver) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"delivery_order_id": m.DeliveryOrderId,
		"driver_user_id":    m.DriverUserId,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
