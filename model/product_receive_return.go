package model

const ProductReceiveReturnTableName = "product_receive_returns"

type ProductReceiveReturn struct {
	Id               string  `db:"id"`
	ProductReceiveId string  `db:"product_receive_id"`
	UserId           string  `db:"user_id"`
	Description      *string `db:"description"`

	Timestamp
}

func (m *ProductReceiveReturn) TableName() string {
	return ProductReceiveReturnTableName
}

func (m *ProductReceiveReturn) TableIds() []string {
	return []string{"id"}
}

func (m *ProductReceiveReturn) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                 m.Id,
		"product_receive_id": m.ProductReceiveId,
		"user_id":            m.UserId,
		"description":        m.Description,
		"created_at":         m.CreatedAt,
		"updated_at":         m.UpdatedAt,
	}
}
