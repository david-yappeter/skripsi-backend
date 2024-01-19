// Code generated by "go run myapp/tool/stringer -linecomment -type=RoleType -output=role_type_enum_gen.go -swagoutput=../tool/swag/enum_gen/role_type_enum_gen.go -custom"; DO NOT EDIT.

package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RoleTypeSuperAdmin-1]
	_ = x[RoleTypeGlobal-2]
	_ = x[RoleTypeAdmin-3]
}

const _RoleType_nameReadable = "SUPER_ADMIN, GLOBAL, ADMIN"

const _RoleType_name = "SUPER_ADMINGLOBALADMIN"

var _RoleType_index = [...]uint8{0, 11, 17, 22}

func (i *RoleType) determine(s string) {
	switch s {
	case "SUPER_ADMIN":
		*i = RoleTypeSuperAdmin
	case "GLOBAL":
		*i = RoleTypeGlobal
	case "ADMIN":
		*i = RoleTypeAdmin
	default:
		*i = 0
	}
}

func (i RoleType) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i RoleType) GetValidValuesString() string {
	return _RoleType_nameReadable
}

func (i RoleType) String() string {
	i -= 1
	if i < 0 || i >= RoleType(len(_RoleType_index)-1) {
		return "RoleType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _RoleType_name[_RoleType_index[i]:_RoleType_index[i+1]]
}

func (i RoleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *RoleType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.determine(s)

	return nil
}

func (i *RoleType) UnmarshalText(b []byte) error {
	i.determine(string(b))

	return nil
}

func (i *RoleType) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i RoleType) Value() (driver.Value, error) {
	return i.String(), nil
}

func RoleTypeP(v RoleType) *RoleType {
	return &v
}

func ListRoleType() []RoleType {
	return []RoleType{
		RoleTypeSuperAdmin,
		RoleTypeGlobal,
		RoleTypeAdmin,
	}
}

func ListRoleTypeString() []string {
	return []string{
		RoleTypeSuperAdmin.String(),
		RoleTypeGlobal.String(),
		RoleTypeAdmin.String(),
	}
}
