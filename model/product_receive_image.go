package model

const ProductReceiveImageTableName = "product_receive_images"

type ProductReceiveImage struct {
	Id               string  `db:"id"`
	ProductReceiveId string  `db:"product_receive_id"`
	FileId           string  `db:"file_id"`
	Description      *string `db:"description"`
	Timestamp

	File *File `db:"-"`
}

func (m *ProductReceiveImage) TableName() string {
	return ProductReceiveImageTableName
}

func (m *ProductReceiveImage) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReceiveImage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                 m.Id,
		"product_receive_id": m.ProductReceiveId,
		"file_id":            m.FileId,
		"description":        m.Description,
		"created_at":         m.CreatedAt,
		"updated_at":         m.UpdatedAt,
	}
}
