package pdf

type BaseParam struct {
	QrCodeBase64      string
	StatementDate     string
	DocumentNumber    string
	ClinicName        string
	ClinicFullAddress string
	ClinicPhoneNumber string

	CompanyName      string
	CompanyAddress   string
	CompanyTaxNumber string

	PatientId   string
	FullName    string
	CashierName string
}

type SalesInvoiceItem struct {
	Code     string
	Category string
	Name     string
	Qty      float64
	SubTotal string
}

type SalesInvoiceParam struct {
	BaseParam

	InvoiceNumber          string
	ItemsMappedByCategory  map[string][]SalesInvoiceItem
	Categories             []string
	Total                  string
	TotalSpellingIndonesia string
	TotalSpellingEnglish   string
	TotalPaidByInsurance   string
	PaidBy                 string
	TotalPaidBy            string

	RequireRenderTax bool
}
