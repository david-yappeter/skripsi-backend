package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=TiktokProductStatus,TiktokProductDimensionUnit,TiktokProductPackageWeight -output=tiktok_product_enum_gen.go -swagoutput=../tool/swag/enum_gen/tiktok_product_enum_gen.go -custom
type TiktokProductStatus int // @name TiktokProductStatusEnum

const (
	TiktokProductStatusActive   TiktokProductStatus = iota + 1 // ACTIVE
	TiktokProductStatusInActive                                // IN_ACTIVE
)

type TiktokProductDimensionUnit int // @name TiktokProductDimensionUnitEnum

const (
	TiktokProductDimensionUnitCentimeter TiktokProductDimensionUnit = iota + 1 // CENTIMETER
)

type TiktokProductPackageWeight int // @name TiktokProductPackageWeightEnum

const (
	TiktokProductPackageWeightKilogram TiktokProductPackageWeight = iota + 1 // KILOGRAM
)
