package model

const ProductReceiveReturnImageTableName = "product_receive_return_images"

type ProductReceiveReturnImage struct {
	Id                     string `db:"id"`
	ProductReceiveReturnId string `db:"product_receive_return_id"`
	FileId                 string `db:"file_id"`

	Timestamp
}

func (m *ProductReceiveReturnImage) TableName() string {
	return ProductReceiveReturnImageTableName
}

func (m *ProductReceiveReturnImage) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReceiveReturnImage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                        m.Id,
		"product_receive_return_id": m.ProductReceiveReturnId,
		"file_id":                   m.FileId,
		"created_at":                m.CreatedAt,
		"updated_at":                m.UpdatedAt,
	}
}
