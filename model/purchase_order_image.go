package model

const PurchaseOrderImageTableName = "purchase_order_images"

type PurchaseOrderImage struct {
	Id              string  `db:"id"`
	PurchaseOrderId string  `db:"purchase_order_id"`
	FileId          string  `db:"file_id"`
	Description     *string `db:"description"`
	Timestamp

	File *File `db:"-"`
}

func (m *PurchaseOrderImage) TableName() string {
	return PurchaseOrderImageTableName
}

func (m *PurchaseOrderImage) TableIds() []string {
	return []string{"id"}
}

func (m *PurchaseOrderImage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"purchase_order_id": m.PurchaseOrderId,
		"file_id":           m.FileId,
		"description":       m.Description,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
