package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=TiktokProductStatus,TiktokProductDimensionUnit,TiktokProductPackageWeigth -output=tiktok_product_enum_gen.go -swagoutput=../tool/swag/enum_gen/tiktok_product_enum_gen.go -custom
type TiktokProductStatus int // @name TiktokProductStatus

const (
	TiktokProductStatusActive   TiktokProductStatus = iota + 1 // ACTIVE
	TiktokProductStatusInActive                                // IN_ACTIVE
)

type TiktokProductDimensionUnit int // @name TiktokProductDimensionUnit

const (
	TiktokProductDimensionUnitCentimeter TiktokProductDimensionUnit = iota + 1 // CENTIMETER
)

type TiktokProductPackageWeigth int // @name TiktokProductPackageWeigth

const (
	TiktokProductPackageWeigthKilogram TiktokProductPackageWeigth = iota + 1 // KILOGRAM
)
