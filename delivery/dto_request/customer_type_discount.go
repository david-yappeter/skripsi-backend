package dto_request

type CustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastFormRequest struct {
	PaginationRequest
	Sorts          CustomerTypeOptionForCustomerFormSorts `json:"sorts" validate:"unique=Field,dive"`
	CustomerTypeId string                                 `json:"customer_type_id" validate:"required,not_empty,uuid"`
	Phrase         *string                                `json:"phrase" validate:"omitempty,not_empty" extensions:"x-nullable"`
} // @name CustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastFormRequest
