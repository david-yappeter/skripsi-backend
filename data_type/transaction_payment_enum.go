package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=TransactionPaymentType -output=transaction_payment_enum_gen.go -swagoutput=../tool/swag/enum_gen/transaction_payment_enum_gen.go -custom
type TransactionPaymentType int // @name TransactionPaymentTypeEnum

const (
	TransactionPaymentTypeCash        TransactionPaymentType = iota + 1 // CASH
	TransactionPaymentTypeBcaTransfer                                   // BCA_TRANSFER
)
