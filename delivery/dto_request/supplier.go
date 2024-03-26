package dto_request

type SupplierCreateRequest struct {
	SupplierTypeId string  `json:"supplier_type_id" validate:"required,not_empty,uuid"`
	Code           string  `json:"code" validate:"required,not_empty,alphanumdashdotslash" example:"2024/slash/alphabet-dash"`
	Name           string  `json:"name" validate:"required,not_empty"`
	IsActive       bool    `json:"is_active"`
	Address        string  `json:"address" validate:"required,not_empty"`
	Phone          string  `json:"phone" validate:"required,not_empty,e164"`
	Email          *string `json:"email" validate:"omitempty,not_empty,email" extensions:"x-nullable"`
	Description    *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierCreateRequest

type SupplierFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name SupplierFetchSorts

type SupplierFetchRequest struct {
	PaginationRequest
	IsActive       *bool              `json:"is_active"`
	SupplierTypeId *string            `json:"supplier_type_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	Sorts          SupplierFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase         *string            `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierFetchRequest

type SupplierGetRequest struct {
	SupplierId string `json:"-" swaggerignore:"true"`
} // @name SupplierGetRequest

type SupplierUpdateRequest struct {
	SupplierTypeId string  `json:"supplier_type_id" validate:"required,not_empty,uuid"`
	Name           string  `json:"name" validate:"required,not_empty"`
	IsActive       bool    `json:"is_active"`
	Address        string  `json:"address" validate:"required,not_empty"`
	Phone          string  `json:"phone" validate:"required,not_empty,e164"`
	Email          *string `json:"email" validate:"omitempty,not_empty,email" extensions:"x-nullable"`
	Description    *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`

	SupplierId string `json:"-" swaggerignore:"true"`
} // @name SupplierUpdateRequest

type SupplierDeleteRequest struct {
	SupplierId string `json:"-" swaggerignore:"true"`
} // @name SupplierDeleteRequest

type SupplierOptionForProductReceiveFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=code name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name SupplierOptionForProductReceiveFormSorts

type SupplierOptionForProductReceiveFormRequest struct {
	PaginationRequest
	Sorts            SupplierOptionForProductReceiveFormSorts `json:"sorts" validate:"unique=Field,dive"`
	ProductReceiveId string                                   `json:"product_receive_id" validate:"required,not_empty,uuid" extensions:"x-nullable"`
	Phrase           *string                                  `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierOptionForProductReceiveFormRequest

type SupplierOptionForProductReceiveFilterRequest struct {
	PaginationRequest
	Phrase *string `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierOptionForProductReceiveFilterRequest
