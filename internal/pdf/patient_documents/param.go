package pdf

type BaseParam struct {
	QrCodeBase64       string
	InchargeDoctorName string
	StatementDate      string
	DocumentNumber     string

	ClinicName        string
	ClinicFullAddress string
	ClinicPhoneNumber string

	CompanyName    string
	CompanyAddress string

	PatientId      string
	IdentityNumber string
	FullName       string
	DateOfBirth    string
	Gender         string
	HomeAddress    string
	PhoneNumber    string

	DoctorPraticeFullAddress    string
	DoctorPracticeLicenseNumber string
	IdHeader                    string
	EnHeader                    string
	NotRequirePatientInfoTitle  bool

	RenderPatientGender      bool
	RenderPatientPhoneNumber bool
	RenderPatientHomeAddress bool
}

type DoctorMemoParam struct {
	BaseParam

	Note string
}

type DoctorReferralParam struct {
	BaseParam

	Examination     string
	FollowingAction string
}

type MedicalCertificateParam struct {
	BaseParam

	Diagnoses        []string
	RestStartedAt    string
	RestEndedAt      string
	Duration         string
	DurationSpelling string
}

type HealthExaminationResultParam struct {
	BaseParam

	BloodPressureDiastolicMmhg string
	BloodPressureSystolicMmhg  string
	PulseRateBpm               string
	BloodGlucoseLevelMgdl      string
	OxygenSaturation           string
	WeightKg                   string
	HeightCm                   string
	BMI                        string
	BodyTemperatureCelcius     string

	ExaminationSummaryDate                  string
	ExamiantionSummaryNote                  string
	ExaminationSummaryDiagnose              string
	ExaminationSummaryTherapyAndInstruction string
}

type MedicalSummaryParam struct {
	BaseParam

	GuardianFullName    string
	GuardianDateOfBirth string
	GuardianGender      string
	GuardianHomeAddress string
	GuardianPhoneNumber string

	Allergies           string
	CongenitalDiseases  string
	SurgeryHistories    string
	MedicationHistories string
	Vaccinations        []MedicalSummaryVaccination
	Examinations        []MedicalSummaryExamination
}

type MedicalSummaryVaccination struct {
	Date        string
	VaccineName string
	Note        string
}

type MedicalSummaryExamination struct {
	Date                   string
	Note                   string
	Diagnoses              string
	TherapyAndInstructions string
}
