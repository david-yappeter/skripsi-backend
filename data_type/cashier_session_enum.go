package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=CashierSessionStatus -output=cashier_session_enum_gen.go -swagoutput=../tool/swag/enum_gen/cashier_session_enum_gen.go -custom
type CashierSessionStatus int // @name CashierSessionStatusEnum

const (
	CashierSessionStatusActive    CashierSessionStatus = iota + 1 // ACTIVE
	CashierSessionStatusCompleted                                 // COMPLETED
)
