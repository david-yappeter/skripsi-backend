package dto_request

import (
	"mime/multipart"
	"myapp/data_type"
)

type DebtFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name DebtFetchSorts

type DebtUploadImageRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
} // @name DebtUploadImageRequest

type DebtFetchRequest struct {
	PaginationRequest
	Sorts  DebtFetchSorts        `json:"sorts" validate:"unique=Field,dive"`
	Status *data_type.DebtStatus `json:"status" validate:"omitempty,data_type_enum"`
	Phrase *string               `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name DebtFetchRequest

type DebtGetRequest struct {
	DebtId string `json:"-" swaggerignore:"true"`
} // @name DebtGetRequest

type DebtPaymentRequest struct {
	ImageFilePath string  `json:"image_file_path" validate:"required,not_empty"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
	Description   *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`

	DebtId string `json:"-" swaggerignore:"true"`
} // @name DebtPaymentRequest
