package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type CustomerPaymentResponse struct {
	Id             string             `json:"id"`
	UserId         string             `json:"user_id"`
	ImageFileId    string             `json:"image_file_id"`
	CustomerDebtId string             `json:"customer_debt_id"`
	Amount         float64            `json:"amount"`
	Description    *string            `json:"description"`
	PaidAt         data_type.DateTime `json:"paid_at"`

	Timestamp

	ImageFile *FileResponse `json:"image_file" extensions:"x-nullable"`
} // @name CustomerPaymentResponse

func NewCustomerPaymentResponse(customerPayment model.CustomerPayment) CustomerPaymentResponse {
	r := CustomerPaymentResponse{
		Id:             customerPayment.Id,
		UserId:         customerPayment.UserId,
		ImageFileId:    customerPayment.ImageFileId,
		CustomerDebtId: customerPayment.CustomerDebtId,
		Amount:         customerPayment.Amount,
		Description:    customerPayment.Description,
		PaidAt:         customerPayment.PaidAt,
		Timestamp:      Timestamp(customerPayment.Timestamp),
	}

	if customerPayment.ImageFile != nil {
		r.ImageFile = NewFileResponseP(*customerPayment.ImageFile)
	}

	return r
}
