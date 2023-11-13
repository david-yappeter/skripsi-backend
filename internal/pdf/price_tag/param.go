package pdf

type PriceTagPdfParam struct {
	Items         []PriceTagPdfItem
	IsYellowPaper bool
}

type PriceTagPdfItem struct {
	Code         string
	Barcode      string
	BarcodeImage string
	Name         string
	Price        string
	PrintDate    string

	UnitName string

	NormalPrice          string
	StartPromotionPeriod string
	EndPromotionPeriod   string
	HasPromotion         bool
}
