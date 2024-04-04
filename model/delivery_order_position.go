package model

const DeliveryOrderPositionTableName = "delivery_order_positions"

type DeliveryOrderPosition struct {
	Id              string  `db:"id"`
	DeliveryOrderId string  `db:"delivery_order_id"`
	DriverUserId    string  `db:"driver_user_id"`
	Latitude        float64 `db:"latitude"`
	Longitude       float64 `db:"longitude"`
	Bearing         float64 `db:"bearing"`

	Timestamp
}

func (m *DeliveryOrderPosition) TableName() string {
	return DeliveryOrderPositionTableName
}

func (m *DeliveryOrderPosition) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderPosition) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"delivery_order_id": m.DeliveryOrderId,
		"driver_user_id":    m.DriverUserId,
		"latitude":          m.Latitude,
		"longitude":         m.Longitude,
		"bearing":           m.Bearing,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}

type DeliveryOrderPositionQueryOption struct {
	QueryOption

	Phrase *string
}

var _ PrepareOption = &DeliveryOrderPositionQueryOption{}

func (o *DeliveryOrderPositionQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"dop.*"}
	}
}

func (o *DeliveryOrderPositionQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}
