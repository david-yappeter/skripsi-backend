package dto_request

type CustomerTypeCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeCreateRequest

type CustomerTypeFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CustomerTypeFetchSorts

type CustomerTypeFetchRequest struct {
	PaginationRequest
	Sorts  CustomerTypeFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeFetchRequest

type CustomerTypeGetRequest struct {
	CustomerTypeId string `json:"-" swaggerignore:"true"`
} // @name CustomerTypeGetRequest

type CustomerTypeUpdateRequest struct {
	CustomerTypeId string `json:"-" swaggerignore:"true"`

	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeUpdateRequest

type CustomerTypeDeleteRequest struct {
	CustomerTypeId string `json:"-" swaggerignore:"true"`
} // @name CustomerTypeDeleteRequest

type CustomerTypeOptionForCustomerFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CustomerTypeOptionForCustomerFormSorts

type CustomerTypeAddDiscountRequest struct {
	ProductId          string   `json:"product_id" validate:"required,not_empty,uuid"`
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage" validate:"omitempty,gt=0,lte=100"`
	DiscountAmount     *float64 `json:"discount_amount" validate:"required_without=DiscountPercentage,omitempty,excluded_with=DiscountPercentage,gt=0"`

	CustomerTypeId string `json:"-" swaggerignore:"true"`
} // @name CustomerTypeAddDiscountRequest

type CustomerTypeUpdateDiscountRequest struct {
	IsActive           bool     `json:"is_active"`
	DiscountPercentage *float64 `json:"discount_percentage" validate:"omitempty,gt=0,lte=100"`
	DiscountAmount     *float64 `json:"discount_amount" validate:"required_without=DiscountPercentage,omitempty,excluded_with=DiscountPercentage,gt=0"`

	CustomerTypeDiscountId string `json:"-" swaggerignore:"true"`
	CustomerTypeId         string `json:"-" swaggerignore:"true"`
} // @name CustomerTypeUpdateDiscountRequest

type CustomerTypeDeleteDiscountRequest struct {
	CustomerTypeDiscountId string `json:"-" swaggerignore:"true"`
	CustomerTypeId         string `json:"-" swaggerignore:"true"`
} // @name CustomerTypeDeleteDiscountRequest

type CustomerTypeOptionForCustomerFormRequest struct {
	PaginationRequest
	Sorts  CustomerTypeOptionForCustomerFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeOptionForCustomerFormRequest

type CustomerTypeOptionForWhatsappProductPriceChangeBroadcastFormRequest struct {
	PaginationRequest
	Sorts  CustomerTypeOptionForCustomerFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeOptionForWhatsappProductPriceChangeBroadcastFormRequest
