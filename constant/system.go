package constant

import "time"

	const NilAsString = `<nil>`

const SystemTreeSeparator = "|" // example: model.Clinic (IdTree), model.Icd10 (CodeTree, IdTree)

const HeaderApiVersionKey = "X-API-Version"
const HeaderRequestIdKey = "X-Request-ID"

const PatientDoctorExpiredDuration = 7 * 24 * time.Hour
const PatientVisitIntervalForAllInDay = 7
const PatientVitalCheckIntervalForMedicalRecordInDay = 365
const PatientTestResultTaskDueDateDuration = 24 * time.Hour

const OneMonthInDays = 30

const DefaultUnitName = "Kali"
const DefaultUnitNameForBundling = "Paket"

const ProductCodePrefixForRetail = "1"
const ProductCodePrefixForNonRetail = "5"
