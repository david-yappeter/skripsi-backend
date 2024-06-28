package model

const DeliveryOrderReviewTableName = "delivery_order_reviews"

type DeliveryOrderReview struct {
	Id              string  `db:"id"`
	DeliveryOrderId string  `db:"delivery_order_id"`
	StarRating      int     `db:"star_rating"`
	Description     *string `db:"description"`

	Timestamp

	DeliveryOrder *DeliveryOrder `db:"-"`
}

func (m *DeliveryOrderReview) TableName() string {
	return DeliveryOrderReviewTableName
}

func (m *DeliveryOrderReview) TableIds() []string {
	return []string{"id"}
}

func (m *DeliveryOrderReview) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"delivery_order_id": m.DeliveryOrderId,
		"star_rating":       m.StarRating,
		"description":       m.Description,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}

type DeliveryOrderReviewQueryOption struct {
	QueryOption

	StarRating *int
}

var _ PrepareOption = &DeliveryOrderReviewQueryOption{}

func (o *DeliveryOrderReviewQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"dor.*"}
	}
}

func (o *DeliveryOrderReviewQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}

func (o *DeliveryOrderReviewQueryOption) TranslateSorts() {
	translateFn := func(field string, direction string) Sorts {
		switch field {

		default:
			return defaultTranslateSort(field, direction)
		}
	}

	o.Sorts = TranslateSorts(o.Sorts, translateFn)
}
