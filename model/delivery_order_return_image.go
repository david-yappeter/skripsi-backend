package model

const DeliveryOrderReturnImageTableName = "delivery_order_return_images"

type DeliveryOrderReturnImage struct {
	Id                    string `db:"id"`
	DeliveryOrderReturnId string `db:"delivery_order_return_id"`
	FileId                string `db:"file_id"`

	Timestamp

	File *File `db:"-"`
}

func (m *DeliveryOrderReturnImage) TableName() string {
	return DeliveryOrderReturnImageTableName
}

func (m *DeliveryOrderReturnImage) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderReturnImage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                       m.Id,
		"delivery_order_return_id": m.DeliveryOrderReturnId,
		"file_id":                  m.FileId,
		"created_at":               m.CreatedAt,
		"updated_at":               m.UpdatedAt,
	}
}
