package model

const ShopeeConfigTableName = "shopee_configs"

type ShopeeConfig struct {
	PartnerId    string  `db:"partner_id"`
	PartnerKey   string  `db:"partner_key"`
	AccessToken  *string `db:"access_token"`
	RefreshToken *string `db:"refresh_token"`
	Timestamp
}

func (m *ShopeeConfig) TableName() string {
	return ShopeeConfigTableName
}

func (m *ShopeeConfig) TableIds() []string {
	return []string{"partner_id"}
}

func (m *ShopeeConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"partner_id":    m.PartnerId,
		"partner_key":   m.PartnerKey,
		"access_token":  m.AccessToken,
		"refresh_token": m.RefreshToken,
		"created_at":    m.CreatedAt,
		"updated_at":    m.UpdatedAt,
	}
}

type ShopeeConfigQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &ShopeeConfigQueryOption{}

func (o *ShopeeConfigQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"sc.*"}
	}
}

func (o *ShopeeConfigQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
