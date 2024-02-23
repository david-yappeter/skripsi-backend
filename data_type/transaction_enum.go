package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=TransactionStatus -output=transaction_enum_gen.go -swagoutput=../tool/swag/enum_gen/transaction_enum_gen.go -custom
type TransactionStatus int // @name TransactionStatusEnum

const (
	TransactionStatusUnpaid TransactionStatus = iota + 1 // UNPAID
	TransactionStatusPaid                                // PAID
)
