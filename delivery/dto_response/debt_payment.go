package dto_response

import (
	"myapp/data_type"
	"myapp/model"
)

type DebtPaymentResponse struct {
	Id          string             `json:"id"`
	UserId      string             `json:"user_id"`
	ImageFileId string             `json:"image_file_id"`
	DebtId      string             `json:"debt_id"`
	Amount      float64            `json:"amount"`
	Description *string            `json:"description"`
	PaidAt      data_type.DateTime `json:"paid_at"`

	Timestamp

	ImageFile *FileResponse `json:"image_file" extensions:"x-nullable"`
} // @name DebtPaymentResponse

func NewDebtPaymentResponse(debtPayment model.DebtPayment) DebtPaymentResponse {
	r := DebtPaymentResponse{
		Id:          debtPayment.Id,
		UserId:      debtPayment.UserId,
		ImageFileId: debtPayment.ImageFileId,
		DebtId:      debtPayment.DebtId,
		Amount:      debtPayment.Amount,
		Description: debtPayment.Description,
		PaidAt:      debtPayment.PaidAt,
		Timestamp:   Timestamp(debtPayment.Timestamp),
	}

	if debtPayment.ImageFile != nil {
		r.ImageFile = NewFileResponseP(*debtPayment.ImageFile)
	}

	return r
}
