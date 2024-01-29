package model

const DeliveryOrderImageTableName = "delivery_order_images"

type DeliveryOrderImage struct {
	Id              string  `db:"id"`
	DeliveryOrderId string  `db:"delivery_order_id"`
	FileId          string  `db:"file_id"`
	Description     *string `db:"description"`
	Timestamp

	File *File `db:"-"`
}

func (m *DeliveryOrderImage) TableName() string {
	return DeliveryOrderImageTableName
}

func (m *DeliveryOrderImage) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderImage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"delivery_order_id": m.DeliveryOrderId,
		"file_id":           m.FileId,
		"description":       m.Description,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
