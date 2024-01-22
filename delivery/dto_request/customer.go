package dto_request

type CustomerCreateRequest struct {
	Name     string  `json:"name" validate:"required,not_empty"`
	Email    string  `json:"email" validate:"required,not_empty,email"`
	Address  *string `json:"address" validate:"required,not_empty" extensions:"x-nullable"`
	Phone    string  `json:"phone" validate:"required,not_empty,e164"`
	IsActive bool    `json:"is_active"`
} // @name CustomerCreateRequest

type CustomerFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name email created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CustomerFetchSorts

type CustomerFetchRequest struct {
	PaginationRequest
	Sorts    CustomerFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	IsActive *bool              `json:"is_active" extensions:"x-nullable"`
	Phrase   *string            `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerFetchRequest

type CustomerGetRequest struct {
	CustomerId string `json:"-" swaggerignore:"true"`
} // @name CustomerGetRequest

type CustomerUpdateRequest struct {
	Name     string  `json:"name" validate:"required,not_empty"`
	Email    string  `json:"email" validate:"required,not_empty,email"`
	Address  *string `json:"address" validate:"required,not_empty" extensions:"x-nullable"`
	Phone    string  `json:"phone" validate:"required,not_empty,e164"`
	IsActive bool    `json:"is_active"`

	CustomerId string `json:"-" swaggerignore:"true"`
} // @name CustomerUpdateRequest

type CustomerDeleteRequest struct {
	CustomerId string `json:"-" swaggerignore:"true"`
} // @name CustomerDeleteRequest

type CustomerOptionForDeliveryOrderFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name email created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CustomerOptionForDeliveryOrderFormSorts

type CustomerOptionForDeliveryOrderFormRequest struct {
	PaginationRequest
	Sorts  CustomerOptionForDeliveryOrderFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                 `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerOptionForDeliveryOrderFormRequest
