package dto_request

type SupplierTypeCreateRequest struct {
	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierTypeCreateRequest

type SupplierTypeFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name SupplierTypeFetchSorts

type SupplierTypeFetchRequest struct {
	PaginationRequest
	Sorts  SupplierTypeFetchSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierTypeFetchRequest

type SupplierTypeGetRequest struct {
	SupplierTypeId string `json:"-" swaggerignore:"true"`
} // @name SupplierTypeGetRequest

type SupplierTypeUpdateRequest struct {
	SupplierTypeId string `json:"-" swaggerignore:"true"`

	Name        string  `json:"name" validate:"required,not_empty"`
	Description *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierTypeUpdateRequest

type SupplierTypeDeleteRequest struct {
	SupplierTypeId string `json:"-" swaggerignore:"true"`
} // @name SupplierTypeDeleteRequest

type SupplierTypeOptionForSupplierFormSorts []struct {
	Field     string `json:"field" validate:"required,oneof=name created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name SupplierTypeOptionForSupplierFormSorts

type SupplierTypeOptionForSupplierFormRequest struct {
	PaginationRequest
	Sorts  SupplierTypeOptionForSupplierFormSorts `json:"sorts" validate:"unique=Field,dive"`
	Phrase *string                                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name SupplierTypeOptionForSupplierFormRequest
