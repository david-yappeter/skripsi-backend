package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ArrayString struct {
	values []interface{}
}

func (dt ArrayString) arrStringToArrInterface(arr []string) (a []interface{}) {
	for _, val := range arr {
		a = append(a, val)
	}
	return
}

func (dt ArrayString) IsUnique() bool {
	isUnique := true
	uniqueMap := map[string]bool{}
	values := dt.Values()
	for idx := 0; idx < len(values); idx++ {
		if uniqueMap[values[idx]] {
			isUnique = false
			break
		}
		uniqueMap[values[idx]] = true
	}

	return isUnique
}

func (dt ArrayString) Values() []string {
	arr := []string{}
	for _, val := range dt.values {
		arr = append(arr, val.(string))
	}
	return arr
}

func (dt ArrayString) Contain(v string) bool {
	mapValues := map[string]bool{}
	for _, val := range dt.values {
		mapValues[val.(string)] = true
	}
	return mapValues[v]
}

func (dt ArrayString) Contains(arr []string, oneOf bool) bool {
	mapValues := map[string]bool{}
	match := 0
	for _, val := range dt.values {
		mapValues[val.(string)] = true
	}
	for _, val := range arr {
		if mapValues[val] {
			match++
		}
	}

	if oneOf {
		return match > 0
	} else {
		return match == len(arr)
	}
}

func (dt ArrayString) String() string {
	return fmt.Sprintf("%+v", dt.values)
}

func (dt ArrayString) isValid() bool {
	if len(dt.values) == 0 {
		return true
	}
	for _, val := range dt.values {
		switch val.(type) {
		case string:
		default:
			return false
		}
	}
	return true
}

func (dt ArrayString) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.values)
}

func (dt *ArrayString) UnmarshalJSON(b []byte) error {
	var arr []interface{}
	if err := json.Unmarshal(b, &arr); err != nil {
		return err
	}
	dt.values = arr
	return nil
}

func (dt *ArrayString) UnmarshalText(b []byte) error {
	s := string(b)
	dt.values = []interface{}{}
	if s == "" {
		return nil
	}

	dt.values = dt.arrStringToArrInterface(strings.Split(s, ","))
	return nil
}

func (dt *ArrayString) Scan(value interface{}) error {
	if value == nil {
		dt.values = []interface{}{}
		return nil
	}

	switch v := value.(type) {
	case string:
		dt.values = dt.arrStringToArrInterface(strings.Split(v, ","))
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dt)
	}

	return nil
}

func (dt ArrayString) Value() (driver.Value, error) {
	if len(dt.values) == 0 {
		return nil, nil
	}

	return strings.Join(dt.Values(), ","), nil
}

func NewArrayString(v []string) ArrayString {
	dt := ArrayString{}
	dt.values = []interface{}{}
	for _, val := range v {
		dt.values = append(dt.values, val)
	}
	return dt
}

func ArrayStringValidationFn(sl validator.StructLevel) {
	arrayString := sl.Current().Interface().(ArrayString)
	if !arrayString.isValid() {
		sl.ReportError(arrayString, "", "", "data_type_array_string", "")
	}
}
