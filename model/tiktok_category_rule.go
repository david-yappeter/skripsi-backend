package model

type TiktokCategoryRule struct {
	PackageDimensionIsRequired bool `db:"-"`
	SizeChartIsSupported       bool `db:"-"`
	SizeChartIsRequired        bool `db:"-"`
}
