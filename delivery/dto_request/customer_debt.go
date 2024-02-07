package dto_request

import (
	"mime/multipart"
	"myapp/data_type"
)

type CustomerDebtFetchSorts []struct {
	Field     string `json:"field" validate:"required,oneof=created_at updated_at" example:"name"`
	Direction string `json:"direction" validate:"required,oneof=asc desc" example:"asc"`
} // @name CustomerDebtFetchSorts

type CustomerDebtUploadImageRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
} // @name CustomerDebtUploadImageRequest

type CustomerDebtFetchRequest struct {
	PaginationRequest
	Sorts      CustomerDebtFetchSorts        `json:"sorts" validate:"unique=Field,dive"`
	CustomerId *string                       `json:"customer_id" validate:"omitempty,not_empty,uuid"`
	Status     *data_type.CustomerDebtStatus `json:"status" validate:"omitempty,data_type_enum"`
	Phrase     *string                       `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerDebtFetchRequest

type CustomerDebtGetRequest struct {
	CustomerDebtId string `json:"-" swaggerignore:"true"`
} // @name CustomerDebtGetRequest

type CustomerDebtPaymentRequest struct {
	ImageFilePath string  `json:"image_file_path" validate:"required,not_empty"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
	Description   *string `json:"description" validate:"omitempty,not_empty" extensions:"x-nullable"`

	CustomerDebtId string `json:"-" swaggerignore:"true"`
} // @name CustomerDebtPaymentRequest
