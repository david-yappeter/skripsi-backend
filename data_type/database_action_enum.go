package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=DatabaseAction -output=database_action_enum_gen.go -swagoutput=../tool/swag/enum_gen/database_action_enum_gen.go -custom
type DatabaseAction int // @name DocumentNumberTypeEnum

const (
	DatabaseActionInsert DatabaseAction = iota + 1 // INSERT
	DatabaseActionUpdate                           // UPDATE
	DatabaseActionDelete                           // DELETE
)
