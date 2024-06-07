package model

const ProductReturnImageTableName = "product_return_images"

type ProductReturnImage struct {
	Id              string  `db:"id"`
	ProductReturnId string  `db:"product_return_id"`
	FileId          string  `db:"file_id"`
	Description     *string `db:"description"`
	Timestamp

	File *File `db:"-"`
}

func (m *ProductReturnImage) TableName() string {
	return ProductReturnImageTableName
}

func (m *ProductReturnImage) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReturnImage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"product_return_id": m.ProductReturnId,
		"file_id":           m.FileId,
		"description":       m.Description,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
