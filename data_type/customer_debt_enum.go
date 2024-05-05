package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=CustomerDebtStatus,CustomerDebtDebtSource -output=customer_debt_enum_gen.go -swagoutput=../tool/swag/enum_gen/customer_debt_enum_gen.go -custom
type CustomerDebtStatus int // @name CustomerDebtStatusEnum

const (
	CustomerDebtStatusUnpaid   CustomerDebtStatus = iota + 1 // UNPAID
	CustomerDebtStatusCanceled                               // CANCELED
	CustomerDebtStatusPaid                                   // PAID
	CustomerDebtStatusReturned                               // RETURNED
)

type CustomerDebtDebtSource int // @name CustomerDebtDebtSourceEnum

const (
	CustomerDebtDebtSourceDeliveryOrder CustomerDebtDebtSource = iota + 1 // DELIVERY_ORDER
)
