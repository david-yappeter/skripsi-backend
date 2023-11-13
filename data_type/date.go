package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"myapp/global"
	"time"

	"github.com/go-playground/validator/v10"
)

type Date struct {
	origin   time.Time
	parseErr error
}

func (dt Date) layout() string {
	return "2006-01-02"
}

func (dt Date) IsoLayout() string {
	return "YYYY-MM-DD"
}

func (dt *Date) parse(s string) {
	dt.origin, dt.parseErr = time.Time{}, nil
	if s == "" {
		return
	}

	t, err := time.Parse(dt.layout(), s)
	dt.parseErr = err
	if err == nil {
		year, month, day := t.Date()
		*dt = newDate(year, month, day)
	}
}

func (dt Date) HasParseErr() bool {
	hasParseError := dt.parseErr != nil
	return hasParseError
}

func (dt Date) IsZero() bool {
	return dt.Time().IsZero()
}

func (dt Date) IsValid() bool {
	return !(dt.IsZero() || dt.HasParseErr())
}

func (dt Date) Format(layout string) string {
	return dt.Time().Format(layout)
}

func (dt Date) DateStartOfMonth() Date {
	year, month, _ := dt.Time().Date()

	return newDate(year, month, 1)
}

func (dt Date) DateEndOfMonth() Date {
	var (
		year, month, _ = dt.Time().Date()

		nextMonth1stDay = newDate(year, month+1, 1)
	)

	return nextMonth1stDay.AddDay(-1)
}

func (dt Date) DateTimeEndOfDay() DateTime {
	year, month, day := dt.Time().Date()
	return NewDateTime(newDate(year, month, day+1).Time().Add(-1 * time.Microsecond))
}

func (dt Date) DateTimeStartOfDay() DateTime {
	year, month, day := dt.Time().Date()
	return NewDateTime(newDate(year, month, day).Time())
}

func (dt Date) NullDate() NullDate {
	return NewNullDate(&dt)
}

func (dt Date) Time() time.Time {
	return dt.origin
}

func (dt Date) AddDay(d int) Date {
	year, month, day := dt.Time().Date()

	return newDate(year, month, day+d)
}

func (dt Date) SubInDay(u Date) int {
	duration := dt.Time().Sub(u.Time())

	return int(duration.Hours()) / 24
}

func (dt Date) SubInYear(u Date) int {
	return dt.SubInDay(u) / 365
}

func (dt Date) IsEqual(u Date) bool {
	return dt.Time().Equal(u.Time())
}

func (dt Date) IsGreaterThan(u Date) bool {
	return dt.Time().After(u.Time())
}

func (dt Date) IsGreaterThanOrEqual(u Date) bool {
	return dt.IsGreaterThan(u) || dt.IsEqual(u)
}

func (dt Date) IsLessThan(u Date) bool {
	return dt.Time().Before(u.Time())
}

func (dt Date) IsLessThanOrEqual(u Date) bool {
	return dt.IsLessThan(u) || dt.IsEqual(u)
}

func (dt Date) String() string {
	var (
		l string = dt.layout()
		s string = "0000-00-00"
	)

	if !dt.IsZero() {
		s = dt.Time().Format(l)
	}

	return s
}

func (dt Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.String())
}

func (dt *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	dt.parse(s)

	return nil
}

func (dt *Date) UnmarshalText(b []byte) error {
	dt.parse(string(b))
	return nil
}

func (dt *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		year, month, day := v.Date()
		*dt = newDate(year, month, day)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dt)
	}

	return nil
}

func (dt Date) Value() (driver.Value, error) {
	return dt.String(), nil
}

// what matters in t is date
func NewDate(t time.Time) Date {
	year, month, day := t.Date()
	return newDate(year, month, day)
}

func newDate(year int, month time.Month, day int) Date {
	return Date{origin: time.Date(year, month, day, 0, 0, 0, 0, global.GetTimeLocation())}
}

func DateValidationFn(sl validator.StructLevel) {
	date := sl.Current().Interface().(Date)
	if !date.IsValid() {
		sl.ReportError(date, "", "", "data_type_date", "")
	}
}
