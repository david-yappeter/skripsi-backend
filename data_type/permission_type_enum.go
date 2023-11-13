package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=PermissionType -output=permission_type_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_type_enum_gen.go -custom
type PermissionType int // @name PermissionTypeEnum

const (
	PermissionTypeAdmin                     PermissionType = iota + 1 // ADMIN
	PermissionTypeGlobal                                              // GLOBAL
	PermissionTypeMustAssign                                          // MUST_ASSIGN
	PermissionTypeMustOnSite                                          // MUST_ON_SITE
	PermissionTypeMustAssignAndOnSite                                 // MUST_ASSIGN_AND_ON_SITE
	PermissionTypeMustAssignAndOnSiteStrict                           // MUST_ASSIGN_AND_ON_SITE_STRICT
)
