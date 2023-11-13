package model

import (
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/util"
	"strings"

	"github.com/Masterminds/squirrel"
)

type BaseModel interface {
	TableName() string
	TableIds() []string
	ToMap() map[string]interface{}
	GetCreatedAt() data_type.DateTime
	GetUpdatedAt() data_type.DateTime
	SetCreatedAt(dateTime data_type.DateTime)
	SetUpdatedAt(dateTime data_type.DateTime)
}

type Timestamp struct {
	CreatedAt data_type.DateTime `db:"created_at"`
	UpdatedAt data_type.DateTime `db:"updated_at"`
}

func (m Timestamp) GetCreatedAt() data_type.DateTime {
	return m.CreatedAt
}

func (m Timestamp) GetUpdatedAt() data_type.DateTime {
	return m.UpdatedAt
}

func (m *Timestamp) SetCreatedAt(dateTime data_type.DateTime) {
	m.CreatedAt = dateTime
}

func (m *Timestamp) SetUpdatedAt(dateTime data_type.DateTime) {
	m.UpdatedAt = dateTime
}

type Sorts []struct {
	Field     string
	Direction string
}

type PrepareOption interface {
	SetDefaultFields()
	SetDefaultSorts()

	GetFields() []string
	GetPage() *int
	GetLimit() *int
	GetSorts() Sorts
	GetDisableSorts() bool
	GetIsCount() bool

	TranslateSorts()
	FixInconsistentSort()
}

type QueryOption struct {
	Fields       []string
	Page         *int
	Limit        *int
	Sorts        Sorts
	DisableSorts bool
	IsCount      bool
}

var _ PrepareOption = &QueryOption{}

func NewQueryOptionWithPagination(
	page *int,
	limit *int,
	sorts Sorts,
) QueryOption {
	if page == nil {
		page = util.IntP(constant.PaginationDefaultPage)
	}

	if limit == nil {
		limit = util.IntP(constant.PaginationDefaultLimit)
	}

	return QueryOption{
		Page:  page,
		Limit: limit,
		Sorts: sorts,
	}
}

func (o *QueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}
}

func (o *QueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "created_at", Direction: "desc"},
		}
	}
}

func (o *QueryOption) GetDisableSorts() bool {
	return o.DisableSorts
}

func (o QueryOption) GetFields() []string {
	return o.Fields
}

func (o QueryOption) GetPage() *int {
	return o.Page
}

func (o QueryOption) GetLimit() *int {
	return o.Limit
}

func (o QueryOption) GetSorts() Sorts {
	return o.Sorts
}

func (o QueryOption) GetIsCount() bool {
	return o.IsCount
}

// This func is used to meet PrepareOption interface
func (o *QueryOption) TranslateSorts() {}

func (o *QueryOption) FixInconsistentSort() {
	var (
		consistentField    string = "created_at"
		hasConsistentField bool   = false
	)

	for _, sort := range o.Sorts {
		if sort.Field == consistentField {
			hasConsistentField = true
			break
		}
	}

	if !hasConsistentField {
		o.Sorts = append(
			o.Sorts,
			Sorts{
				{Field: consistentField, Direction: "asc"},
			}...,
		)
	}
}

func defaultTranslateSort(field string, direction string) Sorts {
	return Sorts{
		{Field: field, Direction: direction},
	}
}

func TranslateSorts(sorts Sorts, translateFn func(field string, direction string) Sorts) Sorts {
	newSorts := Sorts{}
	for _, sort := range sorts {
		newSorts = append(
			newSorts,
			translateFn(sort.Field, sort.Direction)...,
		)
	}

	return newSorts
}

func Prepare(stmt squirrel.SelectBuilder, option PrepareOption) squirrel.SelectBuilder {
	if option.GetIsCount() {
		stmt = stmt.Column("COUNT(*) row_count")

		return stmt
	}

	option.SetDefaultFields()

	disableSorts := option.GetDisableSorts()
	if !disableSorts {
		option.SetDefaultSorts()
		option.TranslateSorts()
		option.FixInconsistentSort()
	}

	var (
		fields []string = option.GetFields()
		sorts  Sorts    = Sorts{}
		page   *int     = option.GetPage()
		limit  *int     = option.GetLimit()
	)

	if !disableSorts {
		sorts = option.GetSorts()
	}

	stmt = stmt.Columns(fields...)

	if len(sorts) > 0 {
		for _, sort := range sorts {
			if util.StringInSlice(strings.ToUpper(sort.Direction), []string{"ASC", "DESC"}) {
				stmt = stmt.OrderBy(fmt.Sprintf("%s %s", sort.Field, strings.ToUpper(sort.Direction)))
			}
		}
	}

	if page != nil && limit != nil && *page >= 1 && *limit >= 1 {
		offset := (*page - 1) * *limit
		stmt = stmt.Offset(uint64(offset))
		stmt = stmt.Limit(uint64(*limit))
	}

	return stmt
}
