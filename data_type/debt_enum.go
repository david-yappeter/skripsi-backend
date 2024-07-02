package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=DebtStatus,DebtSource -output=debt_enum_gen.go -swagoutput=../tool/swag/enum_gen/debt_enum_gen.go -custom
type DebtStatus int // @name DebtStatusEnum

const (
	DebtStatusUnpaid   DebtStatus = iota + 1 // UNPAID
	DebtStatusHalfPaid                       // HALF_PAID
	DebtStatusCanceled                       // CANCELED
	DebtStatusPaid                           // PAID
	DebtStatusReturned                       // RETURNED
)

type DebtSource int // @name DebtSourceEnum

const (
	DebtSourceProductReceive DebtSource = iota + 1 // PRODUCT_RECEIVE
)
