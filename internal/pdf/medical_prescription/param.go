package pdf

type BaseParam struct {
	QrCodeBase64   string
	StatementDate  string
	DocumentNumber string

	ClinicName        string
	ClinicFullAddress string
	ClinicPhoneNumber string

	CompanyName    string
	CompanyAddress string

	PatientId          string
	PatientFullName    string
	PatientDateOfBirth string

	PharmacyName string

	MainPharmacistName          string
	MainPharmacistLicenseNumber string

	CheckerPharmacistName string
}

type MedicalPrescriptionItem struct {
	Name       string
	Qty        float64
	Unit       string
	Categories string

	DoseIndonesia string
	DoseEnglish   string

	TimeInstructionIndonesia string
	TimeInstructionEnglish   string

	MealInstructionIndonesia string
	MealInstructionEnglish   string
}

type MedicalPrescriptionParam struct {
	BaseParam

	Items  []MedicalPrescriptionItem
	Advice string
}
