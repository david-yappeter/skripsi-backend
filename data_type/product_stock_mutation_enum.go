package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=ProductStockMutationType -output=product_stock_mutation_enum_gen.go -swagoutput=../tool/swag/enum_gen/product_stock_mutation_enum_gen.go -custom
type ProductStockMutationType int // @name ProductStockMutationTypeEnum

const (
	ProductStockMutationTypeProductReceiveItem          ProductStockMutationType = iota + 1 // PRODUCT_RECEIVE_ITEM
	ProductStockMutationTypeDeliveryOrderItemCostCancel                                     // DELIVERY_ORDER_ITEM_COST_CANCEL
)
