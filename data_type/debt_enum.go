package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=DebtStatus,DebtDebtSource -output=debt_enum_gen.go -swagoutput=../tool/swag/enum_gen/debt_enum_gen.go -custom
type DebtStatus int // @name DebtStatusEnum

const (
	DebtStatusUnpaid   DebtStatus = iota + 1 // UNPAID
	DebtStatusCanceled                       // CANCELED
	DebtStatusPaid                           // PAID
)

type DebtDebtSource int // @name DebtDebtSourceEnum

const (
	DebtDebtSourceDeliveryOrder DebtDebtSource = iota + 1 // DELIVERY_ORDER
)
