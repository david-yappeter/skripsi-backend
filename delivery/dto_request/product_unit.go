package dto_request

type ProductUnitCreateRequest struct {
	ToUnitId  *string `json:"to_unit_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	UnitId    string  `json:"unit_id" validate:"required,not_empty,uuid"`
	ProductId string  `json:"product_id" validate:"required,not_empty,uuid"`
	Scale     float64 `json:"scale" validate:"required,not_empty,gte=1"`
} // @name ProductUnitCreateRequest

type ProductUnitGetRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name ProductUnitGetRequest

type ProductUnitUpdateRequest struct {
	UnitId    string `json:"unit_id" validate:"required,not_empty,uuid"`
	ProductId string `json:"product_id" validate:"required,not_empty,uuid"`

	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name ProductUnitUpdateRequest

type ProductUnitDeleteRequest struct {
	ProductUnitId string `json:"-" swaggerignore:"true"`
} // @name ProductUnitDeleteRequest

type ProductUnitOptionForProductReceiveFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductUnitOptionForProductReceiveFormSorts

type ProductUnitOptionForProductReceiveFormRequest struct {
	PaginationRequest
	ProductReceiveId string                                      `json:"product_receive_id" validate:"required,not_empty,uuid"`
	Sorts            ProductUnitOptionForProductReceiveFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase           *string                                     `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductUnitOptionForProductReceiveFormRequest

type ProductUnitOptionForDeliveryOrderFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name ProductUnitOptionForDeliveryOrderFormSorts

type ProductUnitOptionForDeliveryOrderFormRequest struct {
	PaginationRequest
	DeliveryOrderId string                                     `json:"delivery_order_id" validate:"required,not_empty,uuid"`
	Sorts           ProductUnitOptionForDeliveryOrderFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase          *string                                    `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name ProductUnitOptionForDeliveryOrderFormRequest
