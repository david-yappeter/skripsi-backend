package dto_request

type AdminSupplierTypeCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminSupplierTypeCreateRequest

type AdminSupplierTypeFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name AdminSupplierTypeFetchSorts

type AdminSupplierTypeFetchRequest struct {
	PaginationRequest
	Sorts  AdminSupplierTypeFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                     `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminSupplierTypeFetchRequest

type AdminSupplierTypeGetRequest struct {
	SupplierTypeId string `json:"-" swaggerignore:"true"`
} // @name AdminSupplierTypeGetRequest

type AdminSupplierTypeUpdateRequest struct {
	SupplierTypeId string `json:"-" swaggerignore:"true"`

	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name AdminSupplierTypeUpdateRequest

type AdminSupplierTypeDeleteRequest struct {
	SupplierTypeId string `json:"-" swaggerignore:"true"`
} // @name AdminSupplierTypeDeleteRequest
