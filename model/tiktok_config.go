package model

const TiktokConfigTableName = "tiktok_configs"

type TiktokConfig struct {
	AppKey       string  `db:"app_key"`
	AppSecret    string  `db:"app_secret"`
	WarehouseId  string  `db:"warehouse_id"`
	ShopId       string  `db:"shop_id"`
	ShopCipher   string  `db:"shop_cipher"`
	AccessToken  *string `db:"access_token"`
	RefreshToken *string `db:"refresh_token"`
	Timestamp
}

func (m *TiktokConfig) TableName() string {
	return TiktokConfigTableName
}

func (m *TiktokConfig) TableIds() []string {
	return []string{"app_key"}
}

func (m *TiktokConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"app_key":       m.AppKey,
		"app_secret":    m.AppSecret,
		"warehouse_id":  m.WarehouseId,
		"shop_id":       m.ShopId,
		"shop_cipher":   m.ShopCipher,
		"access_token":  m.AccessToken,
		"refresh_token": m.RefreshToken,
		"created_at":    m.CreatedAt,
		"updated_at":    m.UpdatedAt,
	}
}

type TiktokConfigQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &TiktokConfigQueryOption{}

func (o *TiktokConfigQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"tc.*"}
	}
}

func (o *TiktokConfigQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
