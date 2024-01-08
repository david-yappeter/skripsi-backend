package dto_request

type AdminSupplierCreateRequest struct {
	SupplierTypeId string  `json:"supplier_type_id" validate:"required,not_empty,uuid"`
	Code           string  `json:"code" validate:"required,not_empty,alphanumdashdotslash" example:"2024/slash/alphabet-dash"`
	Name           string  `json:"name" validate:"required,not_empty"`
	IsActive       bool    `json:"is_active"`
	Address        string  `json:"address" validate:"required,not_empty"`
	Phone          string  `json:"phone" validate:"required,not_empty,e164"`
	Email          *string `json:"email" validate:"omitempty,not_empty,email" extensions:"x-nullable"`
	Description    *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminSupplierCreateRequest

type AdminSupplierFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name AdminSupplierFetchSorts

type AdminSupplierFetchRequest struct {
	PaginationRequest
	IsActive       *bool                   `json:"is_active"`
	SupplierTypeId *string                 `json:"supplier_type_id" validate:"omitempty,not_empty,uuid" extensions:"x-nullable"`
	Sorts          AdminSupplierFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase         *string                 `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminSupplierFetchRequest

type AdminSupplierGetRequest struct {
	SupplierId string `json:"-" swaggerignore:"true"`
} // @name AdminSupplierGetRequest

type AdminSupplierUpdateRequest struct {
	SupplierTypeId string  `json:"supplier_type_id" validate:"required,not_empty,uuid"`
	Name           string  `json:"name" validate:"required,not_empty"`
	IsActive       bool    `json:"is_active"`
	Address        string  `json:"address" validate:"required,not_empty"`
	Phone          string  `json:"phone" validate:"required,not_empty,e164"`
	Email          *string `json:"email" validate:"omitempty,not_empty,email" extensions:"x-nullable"`
	Description    *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`

	SupplierId string `json:"-" swaggerignore:"true"`
} // @name AdminSupplierUpdateRequest

type AdminSupplierDeleteRequest struct {
	SupplierId string `json:"-" swaggerignore:"true"`
} // @name AdminSupplierDeleteRequest
