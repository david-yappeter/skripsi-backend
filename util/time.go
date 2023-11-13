package util

import (
	"myapp/data_type"
	"myapp/global"
	"time"
)

func CurrentDateTime() data_type.DateTime {
	return data_type.NewDateTime(time.Now())
}

func CurrentNullDateTime() data_type.NullDateTime {
	return CurrentDateTime().NullDateTime()
}

func ParseDateTime(s string) data_type.DateTime {
	dt := data_type.DateTime{}
	if err := dt.UnmarshalText([]byte(s)); err != nil {
		return data_type.DateTime{}
	}

	return dt
}

func ParseNullDateTime(s string) data_type.NullDateTime {
	ndt := data_type.NullDateTime{}
	if err := ndt.UnmarshalText([]byte(s)); err != nil {
		return data_type.NullDateTime{}
	}

	return ndt
}

func CurrentDate() data_type.Date {
	return data_type.NewDate(time.Now().In(global.GetTimeLocation()))
}

func CurrentNullDate() data_type.NullDate {
	return CurrentDate().NullDate()
}

func ParseDate(s string) data_type.Date {
	d := data_type.Date{}
	if err := d.UnmarshalText([]byte(s)); err != nil {
		return data_type.Date{}
	}

	return d
}

func ParseNullDate(s string) data_type.NullDate {
	nd := data_type.NullDate{}
	if err := nd.UnmarshalText([]byte(s)); err != nil {
		return data_type.NullDate{}
	}

	return nd
}
