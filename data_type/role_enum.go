package data_type

//go:generate go run myapp/tool/stringer -linecomment -type=Role -output=role_enum_gen.go -swagoutput=../tool/swag/enum_gen/role_enum_gen.go -custom
type Role int // @name RoleEnum

const (
	RoleSuperAdmin             Role = iota + 1 // Super Admin
	RoleClusterManager                         // Cluster Manager
	RoleClinicManager                          // Clinic Manager
	RoleDoctor                                 // Doctor
	RoleNurse                                  // Nurse
	RoleCashier                                // Cashier
	RolePharmacist                             // Pharmacist
	RoleOperationManager                       // Operation Manager
	RolePharmacyManager                        // Pharmacy Manager
	RoleFinanceManager                         // Finance Manager
	RoleAssistantClinicManager                 // Assistant Clinic Manager
)

var roleTypeByRole = map[Role]RoleType{
	RoleSuperAdmin:             RoleTypeSuperAdmin,
	RoleClusterManager:         RoleTypeGlobal,
	RoleClinicManager:          RoleTypeGlobal,
	RoleDoctor:                 RoleTypeDoctor,
	RoleNurse:                  RoleTypeNurse,
	RoleCashier:                RoleTypeGlobal,
	RolePharmacist:             RoleTypePharmacist,
	RoleOperationManager:       RoleTypeAdmin,
	RolePharmacyManager:        RoleTypeAdmin,
	RoleFinanceManager:         RoleTypeAdmin,
	RoleAssistantClinicManager: RoleTypeGlobal,
}

func (r Role) RoleType() RoleType {
	return roleTypeByRole[r]
}

func (r Role) Permissions() []Permission {
	switch r {
	case RoleSuperAdmin:
		return GetRoleSuperAdminPermissions()
	case RoleClusterManager:
		return GetClusterManagerPermissions()
	case RoleClinicManager:
		return GetRoleClinicManagerPermissions()
	case RoleDoctor:
		return GetRoleDoctorPermissions()
	case RoleNurse:
		return GetRoleNursePermissions()
	case RoleCashier:
		return GetRoleCashierPermissions()
	case RolePharmacist:
		return GetRolePharmacistPermissions()
	case RoleOperationManager:
		return GetRoleOperationManagerPermissions()
	case RolePharmacyManager:
		return GetRolePharmacyManagerPermissions()
	case RoleFinanceManager:
		return GetRoleFinanceManagerPermissions()
	case RoleAssistantClinicManager:
		return GetRoleAssistantClinicManagerPermissions()
	}

	return []Permission{}
}

func GetRoleSuperAdminPermissions() []Permission {
	return []Permission{
		// admin appointment type
		PermissionAdminAppointmentTypeCreate,
		PermissionAdminAppointmentTypeList,
		PermissionAdminAppointmentTypeShow,
		PermissionAdminAppointmentTypeUpdate,
		PermissionAdminAppointmentTypeDelete,

		// admin bank
		PermissionAdminBankCreate,
		PermissionAdminBankList,
		PermissionAdminBankShow,
		PermissionAdminBankUpdate,
		PermissionAdminBankDelete,

		// admin clinic
		PermissionAdminClinicCreate,
		PermissionAdminClinicList,
		PermissionAdminClinicShow,
		PermissionAdminClinicUpdate,
		PermissionAdminClinicOption,
		PermissionAdminClinicOptionForAssignUserForm,
		PermissionAdminClinicOptionForDoctorFilter,
		PermissionAdminClinicOptionForNurseFilter,
		PermissionAdminClinicOptionForPharmacistFilter,
		PermissionAdminClinicOptionForParentClinic,
		PermissionAdminClinicOptionForUserFilter,
		PermissionAdminClinicOptionForBankForm,
		PermissionAdminClinicOptionForProductPromotionForm,
		PermissionAdminClinicOptionForProductBundlingForm,

		// admin clinic type
		PermissionAdminClinicTypeOptionForClinicFilter,
		PermissionAdminClinicTypeOptionForClinicForm,

		// admin clinic user
		PermissionAdminClinicUserCreate,
		PermissionAdminClinicUserDelete,

		// admin clinic whitelist ip
		PermissionAdminClinicWhitelistIpCreate,
		PermissionAdminClinicWhitelistIpList,
		PermissionAdminClinicWhitelistIpDelete,

		// admin company
		PermissionAdminCompanyCreate,
		PermissionAdminCompanyList,
		PermissionAdminCompanyShow,
		PermissionAdminCompanyUpdate,
		PermissionAdminCompanyOption,
		PermissionAdminCompanyOptionForClinicFilter,
		PermissionAdminCompanyOptionForClinicForm,
		PermissionAdminCompanyOptionForReportNetPromoterScoreForm,
		PermissionAdminCompanyOptionForReportServiceAndPerformanceForm,
		PermissionAdminCompanyOptionForUserForm,
		PermissionAdminCompanyOptionForBankForm,

		// admin doctor
		PermissionAdminDoctorCreate,
		PermissionAdminDoctorUploadAvatar,
		PermissionAdminDoctorList,
		PermissionAdminDoctorShow,
		PermissionAdminDoctorUpdate,
		PermissionAdminDoctorOption,

		// admin doctor qualification
		PermissionAdminDoctorQualificationCreate,
		PermissionAdminDoctorQualificationUpdate,
		PermissionAdminDoctorQualificationDelete,

		// admin examination
		PermissionAdminExaminationCreate,
		PermissionAdminExaminationList,
		PermissionAdminExaminationShow,
		PermissionAdminExaminationUpdate,
		PermissionAdminExaminationDelete,
		PermissionAdminExaminationOption,

		// admin examination category
		PermissionAdminExaminationCategoryCreate,
		PermissionAdminExaminationCategoryList,
		PermissionAdminExaminationCategoryShow,
		PermissionAdminExaminationCategoryUpdate,
		PermissionAdminExaminationCategoryDelete,
		PermissionAdminExaminationCategoryOption,

		// admin examination item
		PermissionAdminExaminationItemCreate,
		PermissionAdminExaminationItemUpdate,
		PermissionAdminExaminationItemDelete,

		// admin global setting
		PermissionAdminGlobalSettingList,
		PermissionAdminGlobalSettingUpdate,

		// admin icd 10
		PermissionAdminIcd10List,
		PermissionAdminIcd10Tree,

		// admin icd 10 trending
		PermissionAdminIcd10TrendingList,
		PermissionAdminIcd10TrendingCreate,
		PermissionAdminIcd10TrendingMove,
		PermissionAdminIcd10TrendingDelete,

		// admin injection
		PermissionAdminInjectionCreate,
		PermissionAdminInjectionList,
		PermissionAdminInjectionShow,
		PermissionAdminInjectionUpdate,
		PermissionAdminInjectionDelete,

		// admin injection attention
		PermissionAdminInjectionAttentionCreate,
		PermissionAdminInjectionAttentionDelete,

		// admin injection category
		PermissionAdminInjectionCategoryCreate,
		PermissionAdminInjectionCategoryDelete,

		// admin injection instruction
		PermissionAdminInjectionInstructionCreate,
		PermissionAdminInjectionInstructionDelete,

		// admin injection unit
		PermissionAdminInjectionUnitCreate,
		PermissionAdminInjectionUnitUpdate,
		PermissionAdminInjectionUnitDelete,

		// admin insurance
		PermissionAdminInsuranceCreate,
		PermissionAdminInsuranceList,
		PermissionAdminInsuranceShow,
		PermissionAdminInsuranceUpdate,
		PermissionAdminInsuranceDelete,

		// admin insurance payor
		PermissionAdminInsurancePayorCreate,
		PermissionAdminInsurancePayorList,
		PermissionAdminInsurancePayorShow,
		PermissionAdminInsurancePayorUpdate,
		PermissionAdminInsurancePayorDelete,

		// admin medical record template
		PermissionAdminMedicalRecordTemplateCreate,
		PermissionAdminMedicalRecordTemplateList,
		PermissionAdminMedicalRecordTemplateShow,
		PermissionAdminMedicalRecordTemplateUpdate,
		PermissionAdminMedicalRecordTemplateDelete,

		// admin medical record template diagnose
		PermissionAdminMedicalRecordTemplateDiagnoseCreate,
		PermissionAdminMedicalRecordTemplateDiagnoseUpdate,
		PermissionAdminMedicalRecordTemplateDiagnoseDelete,

		// admin medical record template examination
		PermissionAdminMedicalRecordTemplateExaminationCreate,
		PermissionAdminMedicalRecordTemplateExaminationUpdate,
		PermissionAdminMedicalRecordTemplateExaminationDelete,

		// admin medical record template prescription
		PermissionAdminMedicalRecordTemplatePrescriptionCreate,
		PermissionAdminMedicalRecordTemplatePrescriptionUpdate,
		PermissionAdminMedicalRecordTemplatePrescriptionDelete,

		// admin medical record template medical tool
		PermissionAdminMedicalRecordTemplateMedicalToolCreate,
		PermissionAdminMedicalRecordTemplateMedicalToolUpdate,
		PermissionAdminMedicalRecordTemplateMedicalToolDelete,

		// admin medical record template treatment
		PermissionAdminMedicalRecordTemplateTreatmentCreate,
		PermissionAdminMedicalRecordTemplateTreatmentUpdate,
		PermissionAdminMedicalRecordTemplateTreatmentDelete,

		// admin medical tool
		PermissionAdminMedicalToolCreate,
		PermissionAdminMedicalToolList,
		PermissionAdminMedicalToolShow,
		PermissionAdminMedicalToolUpdate,
		PermissionAdminMedicalToolDelete,
		PermissionAdminMedicalToolCreateProductUnit,
		PermissionAdminMedicalToolUpdateProductUnit,
		PermissionAdminMedicalToolDeleteProductUnit,
		PermissionAdminMedicalToolOption,

		// admin medicinal attention
		PermissionAdminMedicinalAttentionCreate,
		PermissionAdminMedicinalAttentionList,
		PermissionAdminMedicinalAttentionUpdate,
		PermissionAdminMedicinalAttentionDelete,
		PermissionAdminMedicinalAttentionOptionForAmpForm,

		// admin medicinal category
		PermissionAdminMedicinalCategoryCreate,
		PermissionAdminMedicinalCategoryList,
		PermissionAdminMedicinalCategoryUpdate,
		PermissionAdminMedicinalCategoryDelete,
		PermissionAdminMedicinalCategoryOptionForAmpForm,

		// admin medicinal classification
		PermissionAdminMedicinalClassificationCreate,
		PermissionAdminMedicinalClassificationList,
		PermissionAdminMedicinalClassificationUpdate,
		PermissionAdminMedicinalClassificationDelete,
		PermissionAdminMedicinalClassificationOption,

		// admin medicinal controlled drug classification
		PermissionAdminMedicinalControlledDrugClassificationCreate,
		PermissionAdminMedicinalControlledDrugClassificationList,
		PermissionAdminMedicinalControlledDrugClassificationUpdate,
		PermissionAdminMedicinalControlledDrugClassificationDelete,
		PermissionAdminMedicinalControlledDrugClassificationOption,

		// admin medicinal instruction
		PermissionAdminMedicinalInstructionCreate,
		PermissionAdminMedicinalInstructionList,
		PermissionAdminMedicinalInstructionUpdate,
		PermissionAdminMedicinalInstructionDelete,
		PermissionAdminMedicinalInstructionOptionForAmpForm,

		// admin medicinal form
		PermissionAdminMedicinalFormCreate,
		PermissionAdminMedicinalFormList,
		PermissionAdminMedicinalFormUpdate,
		PermissionAdminMedicinalFormDelete,
		PermissionAdminMedicinalFormOption,

		// admin medicinal route
		PermissionAdminMedicinalRouteCreate,
		PermissionAdminMedicinalRouteList,
		PermissionAdminMedicinalRouteUpdate,
		PermissionAdminMedicinalRouteDelete,
		PermissionAdminMedicinalRouteOption,

		// admin medicine
		PermissionAdminMedicineCreate,
		PermissionAdminMedicineList,
		PermissionAdminMedicineShow,
		PermissionAdminMedicineUpdate,
		PermissionAdminMedicineDelete,

		// admin medicine attention
		PermissionAdminMedicineAttentionCreate,
		PermissionAdminMedicineAttentionDelete,

		// admin medicine category
		PermissionAdminMedicineCategoryCreate,
		PermissionAdminMedicineCategoryDelete,

		// admin medicine instruction
		PermissionAdminMedicineInstructionCreate,
		PermissionAdminMedicineInstructionDelete,

		// admin medicine unit
		PermissionAdminMedicineUnitCreate,
		PermissionAdminMedicineUnitUpdate,
		PermissionAdminMedicineUnitDelete,

		// admin nurse
		PermissionAdminNurseCreate,
		PermissionAdminNurseList,
		PermissionAdminNurseShow,
		PermissionAdminNurseUpdate,

		// admin nurse qualification
		PermissionAdminNurseQualificationCreate,
		PermissionAdminNurseQualificationUpdate,
		PermissionAdminNurseQualificationDelete,

		// admin pharmacist
		PermissionAdminPharmacistCreate,
		PermissionAdminPharmacistList,
		PermissionAdminPharmacistShow,
		PermissionAdminPharmacistUpdate,
		PermissionAdminPharmacistOptionForCreateClinicForm,
		PermissionAdminPharmacistOptionForUpdateClinicForm,
		PermissionAdminPharmacistQualificationCreate,
		PermissionAdminPharmacistQualificationUpdate,
		PermissionAdminPharmacistQualificationDelete,

		// admin product
		PermissionAdminProductOption,
		PermissionAdminProductOptionForExaminationItem,
		PermissionAdminProductOptionForTreatmentItem,
		PermissionAdminProductOptionForProductPromotion,
		PermissionAdminProductOptionForRetailProduct,
		PermissionAdminProductOptionForProductBundlingItem,

		// admin product promotion
		PermissionAdminProductPromotionCreate,
		PermissionAdminProductPromotionList,
		PermissionAdminProductPromotionShow,
		PermissionAdminProductPromotionUpdate,
		PermissionAdminProductPromotionUpdateEndedAt,
		PermissionAdminProductPromotionDelete,
		PermissionAdminProductPromotionClinicCreate,
		PermissionAdminProductPromotionClinicDelete,

		// admin product bundling
		PermissionAdminProductBundlingCreate,
		PermissionAdminProductBundlingList,
		PermissionAdminProductBundlingShow,
		PermissionAdminProductBundlingUpdate,
		PermissionAdminProductBundlingUpdateEndedAt,
		PermissionAdminProductBundlingDelete,

		// admin product bundling clinic
		PermissionAdminProductBundlingClinicCreate,
		PermissionAdminProductBundlingClinicDelete,

		// admin product bundling item
		PermissionAdminProductBundlingItemCreate,
		PermissionAdminProductBundlingItemDelete,

		// admin product unit
		PermissionAdminProductUnitOption,

		// admin qualification
		PermissionAdminQualificationTypeCreate,
		PermissionAdminQualificationTypeList,
		PermissionAdminQualificationTypeShow,
		PermissionAdminQualificationTypeUpdate,
		PermissionAdminQualificationTypeDelete,
		PermissionAdminQualificationTypeOption,

		// admin retail product
		PermissionAdminRetailProductCreate,
		PermissionAdminRetailProductUpload,
		PermissionAdminRetailProductList,
		PermissionAdminRetailProductShow,
		PermissionAdminRetailProductUpdate,
		PermissionAdminRetailProductDelete,
		PermissionAdminRetailProductCreateProductUnit,
		PermissionAdminRetailProductUpdateProductUnit,
		PermissionAdminRetailProductDeleteProductUnit,

		// admin retail product category
		PermissionAdminRetailProductCategoryCreate,
		PermissionAdminRetailProductCategoryList,
		PermissionAdminRetailProductCategoryShow,
		PermissionAdminRetailProductCategoryUpdate,
		PermissionAdminRetailProductCategoryDelete,
		PermissionAdminRetailProductCategoryOptionForRetailProductForm,

		// admin role
		PermissionAdminRoleList,
		PermissionAdminRoleShow,
		PermissionAdminRoleOption,
		PermissionAdminRoleOptionForUserFilter,
		PermissionAdminRoleOptionForUserForm,

		// admin supplier
		PermissionAdminSupplierCreate,
		PermissionAdminSupplierFetch,
		PermissionAdminSupplierGet,
		PermissionAdminSupplierUpdate,
		PermissionAdminSupplierDelete,
		PermissionAdminSupplierOptionForExaminationForm,

		// admin teleconsult type
		PermissionAdminTeleconsultTypeCreate,
		PermissionAdminTeleconsultTypeList,
		PermissionAdminTeleconsultTypeShow,
		PermissionAdminTeleconsultTypeUpdate,
		PermissionAdminTeleconsultTypeDelete,

		// admin treatment
		PermissionAdminTreatmentCreate,
		PermissionAdminTreatmentList,
		PermissionAdminTreatmentShow,
		PermissionAdminTreatmentUpdate,
		PermissionAdminTreatmentDelete,
		PermissionAdminTreatmentOption,

		// admin treatment category
		PermissionAdminTreatmentCategoryCreate,
		PermissionAdminTreatmentCategoryList,
		PermissionAdminTreatmentCategoryShow,
		PermissionAdminTreatmentCategoryUpdate,
		PermissionAdminTreatmentCategoryDelete,
		PermissionAdminTreatmentCategoryOption,

		// admin treatment item
		PermissionAdminTreatmentItemCreate,
		PermissionAdminTreatmentItemUpdate,
		PermissionAdminTreatmentItemDelete,

		// admin unit
		PermissionAdminUnitCreate,
		PermissionAdminUnitList,
		PermissionAdminUnitUpdate,
		PermissionAdminUnitDelete,
		PermissionAdminUnitOptionForProductForm,

		// admin user
		PermissionAdminUserCreate,
		PermissionAdminUserList,
		PermissionAdminUserShow,
		PermissionAdminUserUpdate,
		PermissionAdminUserOptionForDoctorForm,
		PermissionAdminUserOptionForNurseForm,
		PermissionAdminUserOptionForPharmacistForm,

		// admin vaccine
		PermissionAdminVaccineCreate,
		PermissionAdminVaccineList,
		PermissionAdminVaccineShow,
		PermissionAdminVaccineUpdate,
		PermissionAdminVaccineDelete,

		// admin vaccine attention
		PermissionAdminVaccineAttentionCreate,
		PermissionAdminVaccineAttentionDelete,

		// admin vaccine category
		PermissionAdminVaccineCategoryCreate,
		PermissionAdminVaccineCategoryDelete,

		// admin vaccine instruction
		PermissionAdminVaccineInstructionCreate,
		PermissionAdminVaccineInstructionDelete,

		// admin vaccine unit
		PermissionAdminVaccineUnitCreate,
		PermissionAdminVaccineUnitUpdate,
		PermissionAdminVaccineUnitDelete,

		// admin virtual medicinal product
		PermissionAdminVirtualMedicinalProductCreate,
		PermissionAdminVirtualMedicinalProductList,
		PermissionAdminVirtualMedicinalProductShow,
		PermissionAdminVirtualMedicinalProductUpdate,
		PermissionAdminVirtualMedicinalProductDelete,
		PermissionAdminVirtualMedicinalProductOption,

		// admin virtual therapeutic moiety
		PermissionAdminVirtualTherapeuticMoietyList,
		PermissionAdminVirtualTherapeuticMoietyShow,
		PermissionAdminVirtualTherapeuticMoietyCreate,
		PermissionAdminVirtualTherapeuticMoietyUpdate,
		PermissionAdminVirtualTherapeuticMoietyDelete,
		PermissionAdminVirtualTherapeuticMoietyOption,
		PermissionAdminVirtualTherapeuticMoietyOptionForVmpVtmForm,

		// admin vmp code
		PermissionAdminVmpCodeCreate,
		PermissionAdminVmpCodeDelete,

		// admin vmp vtm
		PermissionAdminVmpVtmCreate,
		PermissionAdminVmpVtmUpdate,
		PermissionAdminVmpVtmDelete,

		// admin vtm code
		PermissionAdminVtmCodeCreate,
		PermissionAdminVtmCodeDelete,

		// admin voucher code
		PermissionAdminVoucherCodeCreate,
		PermissionAdminVoucherCodeList,
		PermissionAdminVoucherCodeShow,
		PermissionAdminVoucherCodeUpdate,
		PermissionAdminVoucherCodeDelete,

		// admin voucher code redeem
		PermissionAdminVoucherCodeRedeemList,
		PermissionAdminVoucherCodeRedeemShow,

		// admin voucher generate rule
		PermissionAdminVoucherGenerateRuleCreate,
		PermissionAdminVoucherGenerateRuleList,
		PermissionAdminVoucherGenerateRuleShow,
		PermissionAdminVoucherGenerateRuleUpdate,
		PermissionAdminVoucherGenerateRuleDelete,

		// patient
		PermissionPatientOptionForTaskForm,

		// report monthly revenue
		PermissionReportMonthlyRevenueList,
		PermissionReportMonthlyRevenueDownload,

		// report net promoter score
		PermissionReportNetPromoterScoreCreate,
		PermissionReportNetPromoterScoreList,
		PermissionReportNetPromoterScoreDownload,

		// report service and performance
		PermissionReportServiceAndPerformanceCreate,
		PermissionReportServiceAndPerformanceList,
		PermissionReportServiceAndPerformanceDownload,

		// task
		PermissionTaskCreate,
		PermissionTaskList,
		PermissionTaskShow,
		PermissionTaskUpdate,
		PermissionTaskUpdatePriority,
		PermissionTaskUpdateStatus,
		PermissionTaskDelete,

		// task attachment
		PermissionTaskAttachmentUpload,
		PermissionTaskAttachmentShow,
		PermissionTaskAttachmentDelete,

		// task comment
		PermissionTaskCommentCreate,
		PermissionTaskCommentList,
		PermissionTaskCommentUpdate,
		PermissionTaskCommentDelete,

		// user
		PermissionUserOptionForTaskForm,
	}
}

func GetClusterManagerPermissions() []Permission {
	return []Permission{
		// bank
		PermissionBankOptionForCashDepositForm,

		// cash deposit
		PermissionAdminCashDepositCreate,
		PermissionAdminCashDepositUpload,
		PermissionAdminCashDepositList,
		PermissionAdminCashDepositShow,
		PermissionAdminCashDepositDownload,

		// cashier summary
		PermissionClusterManagerCashierSummaryList,
		PermissionClusterManagerCashierSummaryListSummarized,
		PermissionClusterManagerCashierSummaryShow,
		PermissionClusterManagerCashierSummaryShowSession,
		PermissionClusterManagerCashierSummaryShowSalesInvoice,
		PermissionClusterManagerCashierSummaryDownloadCsv,

		// clinic
		PermissionClusterManagerClinicGetCashOnHand,
		PermissionClusterManagerClinicOptionForCashierSummaryList,
		PermissionClusterManagerClinicOptionForCashDepositForm,
		PermissionClusterManagerClinicOptionForCashDepositFilter,
	}
}

func GetRoleClinicManagerPermissions() []Permission {
	return []Permission{
		// appointment
		PermissionAppointmentCreate,
		PermissionAppointmentList,
		PermissionAppointmentListForAppointmentFormByClinicOrDoctor,
		PermissionAppointmentListForAppointmentFormByPatient,
		PermissionAppointmentListForQueueNumberGenerateFormByPatient,
		PermissionAppointmentShow,
		PermissionAppointmentUpdate,
		PermissionAppointmentUpdateReferredDoctor,
		PermissionAppointmentCancel,

		// appointment type
		PermissionAppointmentTypeOption,

		// cash deposit
		PermissionCashDepositList,
		PermissionCashDepositShow,
		PermissionCashDepositDownload,

		// cashier
		PermissionCashierCreate,
		PermissionCashierList,
		PermissionCashierShow,
		PermissionCashierUpdate,
		PermissionCashierDelete,

		// cashier session
		PermissionCashierSessionList,
		PermissionCashierSessionShow,
		PermissionCashierSessionUpdateMoneyCount,
		PermissionCashierSessionEnd,

		// cashier summary
		PermissionCashierSummaryCreate,
		PermissionCashierSummaryList,
		PermissionCashierSummaryListSummarized,
		PermissionCashierSummaryDownloadCsv,
		PermissionCashierSummaryShow,
		PermissionCashierSummaryShowSession,
		PermissionCashierSummaryShowSalesInvoice,
		PermissionCashierSummaryShowDailyReport,
		PermissionCashierSummaryShowDailyReportSession,
		PermissionCashierSummaryShowDailyReportSalesInvoice,

		// clinic
		PermissionClinicOptionForAppointmentForm,

		// clinic manager sales invoice
		PermissionClinicManagerSalesInvoiceList,
		PermissionClinicManagerSalesInvoiceDownload,

		// clinic ottopay edc
		PermissionClinicOttopayEdcCreate,
		PermissionClinicOttopayEdcList,
		PermissionClinicOttopayEdcShow,
		PermissionClinicOttopayEdcUpdate,
		PermissionClinicOttopayEdcDelete,
		PermissionClinicOttopayEdcOptionForCashierForm,

		// clinic printer
		PermissionClinicPrinterCreate,
		PermissionClinicPrinterDelete,
		PermissionClinicPrinterOptionForCashierForm,
		PermissionClinicPrinterOptionForQueueCounter,

		// clinic printer gateway
		PermissionClinicPrinterGatewayCreate,
		PermissionClinicPrinterGatewayList,
		PermissionClinicPrinterGatewayShow,
		PermissionClinicPrinterGatewayUpdate,
		PermissionClinicPrinterGatewayDelete,

		// clinic setting
		PermissionClinicSettingList,

		// doctor
		PermissionDoctorList,
		PermissionDoctorShow,
		PermissionDoctorOptionForAppointmentForm,

		// examination
		PermissionExaminationList,
		PermissionExaminationShow,

		// nurse
		PermissionNurseList,
		PermissionNurseShow,

		// patient
		PermissionPatientOptionForAppointmentForm,
		PermissionPatientOptionForPatientDoctorForm,
		PermissionPatientOptionForQueueNumberGenerateForm,
		PermissionPatientOptionForTaskForm,

		// patient doctor
		PermissionPatientDoctorCreate,
		PermissionPatientDoctorDelete,

		// patient visit
		PermissionPatientVisitInsuranceCardFileShowForQueueNumber,

		// pharmacist
		PermissionPharmacistList,
		PermissionPharmacistShow,

		// queue
		PermissionQueueCreate,
		PermissionQueueList,
		PermissionQueueShow,
		PermissionQueueUpdate,
		PermissionQueueDelete,
		PermissionQueueOption,
		PermissionQueueOptionForQueueDisplay,

		// queue number
		PermissionQueueNumberGet,
		PermissionQueueNumberUpdate,

		// queue number action
		PermissionQueueNumberProcess,
		PermissionQueueNumberMove,

		// queue counter
		PermissionQueueCounterCreate,
		PermissionQueueCounterUpdate,
		PermissionQueueCounterOptionForQueueNumberGenerateForm,

		// queue display
		PermissionQueueDisplayCreate,
		PermissionQueueDisplayList,
		PermissionQueueDisplayShow,
		PermissionQueueDisplayUpdate,
		PermissionQueueDisplayDelete,

		// queue display banner
		PermissionQueueDisplayBannerCreate,
		PermissionQueueDisplayBannerUpload,
		PermissionQueueDisplayBannerMove,
		PermissionQueueDisplayBannerDelete,

		// queue display queue
		PermissionQueueDisplayQueueCreate,
		PermissionQueueDisplayQueueDelete,

		// queue display running text
		PermissionQueueDisplayRunningTextCreate,
		PermissionQueueDisplayRunningTextMove,
		PermissionQueueDisplayRunningTextDelete,

		// queue session
		PermissionQueueSessionBegin,
		PermissionQueueSessionEnd,
		PermissionQueueSessionGetActive,
		PermissionQueueSessionCheckForActiveQueueNumber,

		// report daily transaction
		PermissionReportDailyTransactionList,
		PermissionReportDailyTransactionDownload,

		// report inventory stock
		PermissionReportInventoryStockCreate,
		PermissionReportInventoryStockList,
		PermissionReportInventoryStockDownload,

		// report waive fee
		PermissionReportWaiveFeeDownload,

		// role
		PermissionRoleOptionForQueueForm,
		PermissionRoleOptionForUserFilter,

		// task
		PermissionTaskCreate,
		PermissionTaskList,
		PermissionTaskShow,
		PermissionTaskUpdate,
		PermissionTaskUpdatePriority,
		PermissionTaskUpdateStatus,
		PermissionTaskDelete,

		// task attachment
		PermissionTaskAttachmentUpload,
		PermissionTaskAttachmentShow,
		PermissionTaskAttachmentDelete,

		// task comment
		PermissionTaskCommentCreate,
		PermissionTaskCommentList,
		PermissionTaskCommentUpdate,
		PermissionTaskCommentDelete,

		// treatment
		PermissionTreatmentList,
		PermissionTreatmentShow,

		// user
		PermissionUserList,
		PermissionUserShow,
		PermissionUserOptionForTaskForm,
	}
}

func GetRoleDoctorPermissions() []Permission {
	return []Permission{
		// cart
		PermissionCartItemReprintDocumentFeeDelete,

		// doctor patient vital check
		PermissionDoctorPatientVitalCheckCreate,
		PermissionDoctorPatientVitalCheckUpdate,

		// appointment
		PermissionAppointmentCreate,
		PermissionAppointmentList,
		PermissionAppointmentListForAppointmentFormByClinicOrDoctor,
		PermissionAppointmentListForAppointmentFormByPatient,
		PermissionAppointmentShow,
		PermissionAppointmentUpdate,
		PermissionAppointmentUpdateReferredDoctor,
		PermissionAppointmentApprove,
		PermissionAppointmentCancel,

		// appointment type
		PermissionAppointmentTypeOption,

		// clinic
		PermissionClinicOptionForAppointmentForm,

		// doctor
		PermissionDoctorOptionForAppointmentForm,
		PermissionDoctorOptionForTeleconsultForm,

		// doctor memo
		PermissionDoctorMemoCreate,
		PermissionDoctorMemoUpdate,

		// doctor referral
		PermissionDoctorReferralCreate,
		PermissionDoctorReferralUpdate,

		// examination
		PermissionExaminationList,
		PermissionExaminationShow,
		PermissionExaminationOptionForMedicalRecord,

		// icd 10
		PermissionIcd10List,
		PermissionIcd10Tree,
		PermissionIcd10OptionForPatientDiseaseHistoryForm,
		PermissionIcd10OptionForPatientCongenitalDiseaseForm,

		// icd 10 trending
		PermissionIcd10TrendingList,

		// injection
		PermissionInjectionShow,

		// health examination result
		PermissionHealthExaminationResultCreate,
		PermissionHealthExaminationResultUpdate,

		// medical certificate
		PermissionMedicalCertificateCreate,
		PermissionMedicalCertificateUpdate,

		// medical certificate diagnose
		PermissionMedicalCertificateDiagnoseCreate,
		PermissionMedicalCertificateDiagnoseDelete,

		// medical record
		PermissionMedicalRecordBegin,
		PermissionMedicalRecordList,
		PermissionMedicalRecordListForStandByForm,
		PermissionMedicalRecordShow,
		PermissionMedicalRecordUpdate,
		PermissionMedicalRecordUpdateApplyTemplate,
		PermissionMedicalRecordUpdateAnnotation,
		PermissionMedicalRecordUpdateAmendment,
		PermissionMedicalRecordStandBy,
		PermissionMedicalRecordResume,
		PermissionMedicalRecordEnd,

		// medical record amendment history
		PermissionMedicalRecordAmendmentHistoryList,

		// medical record diagnose
		PermissionMedicalRecordDiagnoseCreate,
		PermissionMedicalRecordDiagnoseUpdate,
		PermissionMedicalRecordDiagnoseDelete,

		// medical record examination
		PermissionMedicalRecordExaminationCreate,
		PermissionMedicalRecordExaminationUpdate,
		PermissionMedicalRecordExaminationDelete,

		// medical record medical tool
		PermissionMedicalRecordMedicalToolCreate,
		PermissionMedicalRecordMedicalToolUpdate,
		PermissionMedicalRecordMedicalToolDelete,

		// medical record prescription
		PermissionMedicalRecordPrescriptionCreate,
		PermissionMedicalRecordPrescriptionUpdate,
		PermissionMedicalRecordPrescriptionDelete,

		// medical record treatment
		PermissionMedicalRecordTreatmentCreate,
		PermissionMedicalRecordTreatmentUpdate,
		PermissionMedicalRecordTreatmentDelete,

		// medical record product bundling
		PermissionMedicalRecordProductBundlingCreate,
		PermissionMedicalRecordProductBundlingUpdate,
		PermissionMedicalRecordProductBundlingDelete,

		// medical record template
		PermissionMedicalRecordTemplateCreate,
		PermissionMedicalRecordTemplateList,
		PermissionMedicalRecordTemplateShow,
		PermissionMedicalRecordTemplateUpdate,
		PermissionMedicalRecordTemplateDelete,
		PermissionMedicalRecordTemplateOptionForMedicalRecordForm,

		// medical record template diagnose
		PermissionMedicalRecordTemplateDiagnoseCreate,
		PermissionMedicalRecordTemplateDiagnoseUpdate,
		PermissionMedicalRecordTemplateDiagnoseDelete,

		// medical record template medical tool
		PermissionMedicalRecordTemplateMedicalToolCreate,
		PermissionMedicalRecordTemplateMedicalToolUpdate,
		PermissionMedicalRecordTemplateMedicalToolDelete,

		// medical record template treatment
		PermissionMedicalRecordTemplateTreatmentCreate,
		PermissionMedicalRecordTemplateTreatmentUpdate,
		PermissionMedicalRecordTemplateTreatmentDelete,

		// medical record template examination
		PermissionMedicalRecordTemplateExaminationCreate,
		PermissionMedicalRecordTemplateExaminationUpdate,
		PermissionMedicalRecordTemplateExaminationDelete,

		// medical record template prescription
		PermissionMedicalRecordTemplatePrescriptionCreate,
		PermissionMedicalRecordTemplatePrescriptionUpdate,
		PermissionMedicalRecordTemplatePrescriptionDelete,

		// medical record waive
		PermissionMedicalRecordWaiveUpdate,

		// medical summary
		PermissionMedicalSummaryCreate,

		// medical tool
		PermissionMedicalToolOptionForMedicalRecord,

		// medicine
		PermissionMedicineShow,

		// patient
		PermissionPatientShow,
		PermissionPatientShowMedicalSummary,
		PermissionPatientShowMedicalNotes,
		PermissionPatientUpdate,
		PermissionPatientUpdateMedicalSummary,
		PermissionPatientUpdateMedicalNotes,
		PermissionPatientOption,
		PermissionPatientOptionForAppointmentForm,
		PermissionPatientOptionForPatientVitalCheckForm,
		PermissionPatientOptionForPatientVitalCheckFilter,
		PermissionPatientOptionForTeleconsult,
		PermissionPatientOptionForVaccination,
		PermissionPatientOptionForTaskForm,

		// patient allergy
		PermissionPatientAllergyCreate,
		PermissionPatientAllergyList,
		PermissionPatientAllergyShow,
		PermissionPatientAllergyUpdate,
		PermissionPatientAllergyDelete,

		// patient congenital disease
		PermissionPatientCongenitalDiseaseCreate,
		PermissionPatientCongenitalDiseaseList,
		PermissionPatientCongenitalDiseaseShow,
		PermissionPatientCongenitalDiseaseUpdate,
		PermissionPatientCongenitalDiseaseDelete,

		// patient disease history
		PermissionPatientDiseaseHistoryCreate,
		PermissionPatientDiseaseHistoryShow,
		PermissionPatientDiseaseHistoryList,
		PermissionPatientDiseaseHistoryUpdate,
		PermissionPatientDiseaseHistoryDelete,

		// patient document
		PermissionPatientDocumentList,
		PermissionPatientDocumentUpload,
		PermissionPatientDocumentComplete,
		PermissionPatientDocumentDownload,
		PermissionPatientDocumentDelete,

		// patient interview
		PermissionPatientInterviewCreate,
		PermissionPatientInterviewList,
		PermissionPatientInterviewShow,
		PermissionPatientInterviewUpdate,
		PermissionPatientInterviewDelete,

		// patient medicine allergy
		PermissionPatientMedicineAllergyCreate,
		PermissionPatientMedicineAllergyList,
		PermissionPatientMedicineAllergyShow,
		PermissionPatientMedicineAllergyUpdate,
		PermissionPatientMedicineAllergyDelete,

		// patient medicine history
		PermissionPatientMedicineHistoryCreate,
		PermissionPatientMedicineHistoryList,
		PermissionPatientMedicineHistoryShow,
		PermissionPatientMedicineHistoryUpdate,
		PermissionPatientMedicineHistoryDelete,

		// patient surgery
		PermissionPatientSurgeryCreate,
		PermissionPatientSurgeryList,
		PermissionPatientSurgeryShow,
		PermissionPatientSurgeryUpdate,
		PermissionPatientSurgeryDelete,

		// patient test result
		PermissionPatientTestResultCreate,
		PermissionPatientTestResultUpload,
		PermissionPatientTestResultList,
		PermissionPatientTestResultShow,
		PermissionPatientTestResultUpdate,
		PermissionPatientTestResultUpdateReadStatus,

		// patient vital check
		PermissionPatientVitalCheckListForMedicalRecord,
		PermissionPatientVitalCheckShow,
		PermissionPatientVitalCheckShowLatest,
		PermissionPatientVitalCheckVisualize,
		PermissionPatientVitalCheckUpdateIsNursing,
		PermissionPatientVitalCheckUpdateIsPregnant,
		PermissionPatientVitalCheckUpdateIsSmoking,

		// patient visit
		PermissionPatientVisitOptionForMedicalRecord,
		PermissionPatientVisitSummarizeHistory,
		PermissionPatientVisitInsuranceCardFileShowForMedicalRecord,

		// product
		PermissionProductOptionForMedicalRecord,

		// product bundling
		PermissionProductBundlingShow,
		PermissionProductBundlingOptionForMedicalRecord,

		// queue
		PermissionQueueShowProcess,
		PermissionQueueOption,
		PermissionQueueOptionForQueueCounterCheckInForm,

		// queue counter
		PermissionQueueCounterCheckIn,
		PermissionQueueCounterCheckOut,
		PermissionQueueCounterOptionForCheckInForm,
		PermissionQueueCounterOptionForQueueNumberCompleteForm,

		// queue number
		PermissionQueueNumberCurrent,

		// queue number action
		PermissionQueueNumberCall,
		PermissionQueueNumberCallRepeat,
		PermissionQueueNumberProcess,
		PermissionQueueNumberComplete,
		PermissionQueueNumberMove,
		PermissionQueueNumberSkip,
		PermissionQueueNumberStandBy,

		// task
		PermissionTaskCreate,
		PermissionTaskList,
		PermissionTaskShow,
		PermissionTaskUpdate,
		PermissionTaskUpdatePriority,
		PermissionTaskUpdateStatus,
		PermissionTaskDelete,

		// task attachment
		PermissionTaskAttachmentUpload,
		PermissionTaskAttachmentShow,
		PermissionTaskAttachmentDelete,

		// task comment
		PermissionTaskCommentCreate,
		PermissionTaskCommentList,
		PermissionTaskCommentUpdate,
		PermissionTaskCommentDelete,

		// teleconsult
		PermissionTeleconsultCreate,
		PermissionTeleconsultList,
		PermissionTeleconsultShow,
		PermissionTeleconsultUpdate,
		PermissionTeleconsultCancelConsult,

		// teleconsult type
		PermissionTeleconsultTypeOption,

		// treatment
		PermissionTreatmentList,
		PermissionTreatmentShow,
		PermissionTreatmentOptionForMedicalRecord,

		// user
		PermissionUserOptionForTaskForm,

		// vaccination
		PermissionVaccinationCreate,
		PermissionVaccinationList,
		PermissionVaccinationShow,
		PermissionVaccinationUpdate,
		PermissionVaccinationDelete,

		// virtual therapeutic moiety
		PermissionVirtualTherapeuticMoietyOptionForPatientMedicineAllergy,
		PermissionVirtualTherapeuticMoietyOptionForPatientMedicineHistory,
	}
}

func GetRoleNursePermissions() []Permission {
	return []Permission{
		// nurse patient vital check
		PermissionNursePatientVitalCheckCreate,
		PermissionNursePatientVitalCheckUpdate,

		// appointment
		PermissionAppointmentCreate,
		PermissionAppointmentList,
		PermissionAppointmentListForAppointmentFormByClinicOrDoctor,
		PermissionAppointmentListForAppointmentFormByPatient,
		PermissionAppointmentListForQueueNumberGenerateFormByPatient,
		PermissionAppointmentShow,
		PermissionAppointmentUpdate,
		PermissionAppointmentUpdateReferredDoctor,
		PermissionAppointmentCancel,

		// appointment type
		PermissionAppointmentTypeOption,

		// clinic
		PermissionClinicOptionForAppointmentForm,

		// doctor
		PermissionDoctorOptionForAppointmentForm,

		// insurance
		PermissionInsuranceOptionForQueueNumber,

		// insurance payor
		PermissionInsurancePayorOptionForQueueNumber,

		// nric card ocr,
		PermissionNricCardOcrExtract,
		PermissionNricCardOcrUpload,

		// queue
		PermissionQueueList,
		PermissionQueueShow,
		PermissionQueueShowProcess,
		PermissionQueueOption,
		PermissionQueueOptionForQueueCounterCheckInForm,

		// queue counter
		PermissionQueueCounterCheckIn,
		PermissionQueueCounterCheckOut,
		PermissionQueueCounterOptionForCheckInForm,
		PermissionQueueCounterOptionForQueueNumberCompleteForm,
		PermissionQueueCounterOptionForQueueNumberGenerateForm,

		// queue number
		PermissionQueueNumberGenerate,
		PermissionQueueNumberUploadInsuranceCard,
		PermissionQueueNumberGet,
		PermissionQueueNumberCurrent,
		PermissionQueueNumberUpdate,
		PermissionQueueNumberReprint,

		// queue number action
		PermissionQueueNumberCall,
		PermissionQueueNumberCallRepeat,
		PermissionQueueNumberProcess,
		PermissionQueueNumberComplete,
		PermissionQueueNumberMove,
		PermissionQueueNumberSkip,
		PermissionQueueNumberStandBy,

		// patient
		PermissionPatientCreate,
		PermissionPatientShow,
		PermissionPatientUpdate,
		PermissionPatientOption,
		PermissionPatientOptionForAppointmentForm,
		PermissionPatientOptionForPatientVisitForm,
		PermissionPatientOptionForPatientVisitFilter,
		PermissionPatientOptionForPatientVitalCheckForm,
		PermissionPatientOptionForPatientVitalCheckFilter,
		PermissionPatientOptionForQueueNumberGenerateForm,
		PermissionPatientOptionForVaccination,
		PermissionPatientOptionForPatientRelativeForm,
		PermissionPatientOptionForTaskForm,

		// patient relative
		PermissionPatientRelativeCreate,
		PermissionPatientRelativeList,
		PermissionPatientRelativeUpdate,
		PermissionPatientRelativeUpdateSetDefault,
		PermissionPatientRelativeDelete,
		PermissionPatientRelativeOptionForQueueNumberForm,

		// patient test result
		PermissionPatientTestResultCreate,
		PermissionPatientTestResultList,
		PermissionPatientTestResultShow,
		PermissionPatientTestResultUpdate,
		PermissionPatientTestResultUpdateReadStatus,

		// patient visit
		PermissionPatientVisitList,
		PermissionPatientVisitShow,
		PermissionPatientVisitInsuranceCardFileShowForQueueNumber,
		PermissionPatientVisitUpdate,

		// patient vital check
		PermissionPatientVitalCheckList,
		PermissionPatientVitalCheckShow,

		// task
		PermissionTaskCreate,
		PermissionTaskList,
		PermissionTaskShow,
		PermissionTaskUpdate,
		PermissionTaskUpdatePriority,
		PermissionTaskUpdateStatus,
		PermissionTaskDelete,

		// task attachment
		PermissionTaskAttachmentUpload,
		PermissionTaskAttachmentShow,
		PermissionTaskAttachmentDelete,

		// task comment
		PermissionTaskCommentCreate,
		PermissionTaskCommentList,
		PermissionTaskCommentUpdate,
		PermissionTaskCommentDelete,

		// user
		PermissionUserOptionForTaskForm,

		// vaccination
		PermissionVaccinationCreate,
		PermissionVaccinationList,
		PermissionVaccinationShow,
		PermissionVaccinationUpdate,
		PermissionVaccinationDelete,
	}
}

func GetRoleCashierPermissions() []Permission {
	return []Permission{
		// cashier sales invoice
		PermissionCashierSalesInvoiceList,
		PermissionCashierSalesInvoiceShow,

		PermissionCashierOptionForBeginCashierSessionForm,

		// cashier session
		PermissionCashierSessionBegin,
		PermissionCashierSessionShowActive,
		PermissionCashierSessionOpenDrawer,
		PermissionCashierSessionClose,

		// cart
		PermissionCartShow,

		// cart item
		PermissionCartItemCreateForActiveRetailCart,
		PermissionCartItemCreateForActiveRetailCartByBarcodeOrCode,
		PermissionCartItemQtyUpdateInActiveCart,
		PermissionCartItemDeleteFromActiveRetailCart,

		// insurance
		PermissionInsuranceOptionForQueueNumber,

		// insurance payor
		PermissionInsurancePayorOptionForQueueNumber,

		// net promoter score
		PermissionNetPromoterScoreStartByQueueNumber,
		PermissionNetPromoterScoreStartBySalesInvoice,
		PermissionNetPromoterScoreGetActive,
		PermissionNetPromoterScoreItemUpdate,
		PermissionNetPromoterScoreEnd,

		// patient
		PermissionPatientOptionForTaskForm,

		// patient visit
		PermissionPatientVisitInsuranceCardFileShowForQueueNumber,

		// product
		PermissionProductListForRetailCart,
		PermissionProductOptionForPriceTagDownload,

		// product bundling
		PermissionProductBundlingShow,

		// product promotion
		PermissionProductPromotionOptionForPriceTagDownload,

		// product unit
		PermissionProductUnitOptionForPriceTagDownload,

		// price tag
		PermissionProductPriceTagDownload,

		// medicinal dispense
		PermissionMedicinalDispenseUpload,
		PermissionMedicinalDispenseDownload,
		PermissionMedicinalDispenseCreate,
		PermissionMedicinalDispenseUpdate,

		// queue
		PermissionQueueShowProcess,
		PermissionQueueOption,
		PermissionQueueOptionForQueueCounterCheckInForm,

		// queue counter
		PermissionQueueCounterCheckIn,
		PermissionQueueCounterCheckOut,
		PermissionQueueCounterOptionForCheckInForm,
		PermissionQueueCounterOptionForQueueNumberCompleteForm,
		PermissionQueueCounterOptionForQueueNumberGenerateForm,

		// queue number
		PermissionQueueNumberGenerate,
		PermissionQueueNumberUploadInsuranceCard,
		PermissionQueueNumberGet,
		PermissionQueueNumberCurrent,
		PermissionQueueNumberUpdate,
		PermissionQueueNumberReprint,

		// queue number action
		PermissionQueueNumberCall,
		PermissionQueueNumberCallRepeat,
		PermissionQueueNumberProcess,
		PermissionQueueNumberComplete,
		PermissionQueueNumberMove,
		PermissionQueueNumberSkip,
		PermissionQueueNumberStandBy,

		// retail cart
		PermissionRetailCartGetActive,
		PermissionRetailCartShow,
		PermissionRetailCartHoldActive,
		PermissionRetailCartRestore,
		PermissionRetailCartDeleteActive,
		PermissionRetailCartDelete,
		PermissionRetailCartOptionForRestoreOrDelete,

		// sales invoice
		PermissionSalesInvoiceGenerate,
		PermissionSalesInvoiceDownload,
		PermissionSalesInvoiceVoid,
		PermissionSalesInvoiceCancel,
		PermissionSalesInvoiceApplyVoucherCode,
		PermissionSalesInvoiceRemoveVoucherCode,
		PermissionSalesInvoiceCheckVoucherCode,

		// sales payment
		PermissionSalesPaymentCreate,
		PermissionSalesPaymentShow,
		PermissionSalesPaymentPrintRetail,
		PermissionSalesPaymentCancel,

		// sales payment method
		PermissionSalesPaymentMethodOption,

		// task
		PermissionTaskCreate,
		PermissionTaskList,
		PermissionTaskShow,
		PermissionTaskUpdate,
		PermissionTaskUpdatePriority,
		PermissionTaskUpdateStatus,
		PermissionTaskDelete,

		// task attachment
		PermissionTaskAttachmentUpload,
		PermissionTaskAttachmentShow,
		PermissionTaskAttachmentDelete,

		// task comment
		PermissionTaskCommentCreate,
		PermissionTaskCommentList,
		PermissionTaskCommentUpdate,
		PermissionTaskCommentDelete,

		// user
		PermissionUserOptionForVoidSalesInvoiceForm,
		PermissionUserOptionForBeginCashierSessionForm,
		PermissionUserOptionForTaskForm,
	}
}

func GetRolePharmacistPermissions() []Permission {
	return []Permission{
		// cart item
		PermissionCartItemReprintDocumentFeeDelete,

		// clinic printer gateway
		PermissionClinicPrinterOptionForPrintMedicineLabel,

		// insurance
		PermissionInsuranceOptionForQueueNumber,

		// insurance payor
		PermissionInsurancePayorOptionForQueueNumber,

		// medicinal dispense
		PermissionMedicinalDispenseDownload,
		PermissionMedicinalDispenseCreate,
		PermissionMedicinalDispenseUpdate,
		PermissionMedicinalDispenseUpload,
		PermissionMedicinalDispenseList,
		PermissionMedicinalDispenseShow,
		PermissionMedicinalDispenseAddMixedMedicineFee,
		PermissionMedicinalDispenseRemoveMixedMedicineFee,
		PermissionMedicinalDispensePrintBlank,
		PermissionMedicinalDispenseCalibrate,
		PermissionMedicinalDispenseUpdateChecked,
		PermissionMedicinalDispenseUpdateCompleted,

		// medicinal dispense item
		PermissionMedicinalDispenseItemPrint,
		PermissionMedicinalDispenseItemCreateExpiredDate,
		PermissionMedicinalDispenseItemUpdateExpiredDate,
		PermissionMedicinalDispenseItemDeleteExpiredDate,
		PermissionMedicinalDispenseItemUpdate,

		// medicinal dispense prescription
		PermissionMedicinalDispensePrescriptionCreate,
		PermissionMedicinalDispensePrescriptionUpdate,
		PermissionMedicinalDispensePrescriptionDelete,

		// medicine
		PermissionMedicineShow,
		PermissionMedicineOptionForMedicinalDispense,

		// patient
		PermissionPatientOptionForQueueNumberGenerateForm,
		PermissionPatientOptionForTaskForm,

		// patient visit
		PermissionPatientVisitInsuranceCardFileShowForQueueNumber,

		// product
		PermissionProductOptionForMedicinalDispensePrescription,
		PermissionProductPriceTagDownload,
		PermissionProductOptionForProductReceiveItem,
		PermissionProductOptionForPriceTagDownload,
		PermissionProductOptionForProductTransferItemForm,
		PermissionProductOptionForProductAdjustmentForm,
		PermissionProductOptionExpiredDate,

		// product adjustment
		PermissionProductAdjustmentCreate,
		PermissionProductAdjustmentList,
		PermissionProductAdjustmentShow,
		PermissionProductAdjustmentUpdate,
		PermissionProductAdjustmentDelete,

		// product adjustment item
		PermissionProductAdjustmentItemCreate,
		PermissionProductAdjustmentItemDelete,

		// product inventory
		PermissionProductInventoryList,

		// product promotion
		PermissionProductPromotionOptionForPriceTagDownload,

		// product receive
		PermissionProductReceiveCreate,
		PermissionProductReceiveList,
		PermissionProductReceiveShow,
		PermissionProductReceiveUpdate,
		PermissionProductReceiveDelete,

		// product receive item
		PermissionProductReceiveItemCreate,
		PermissionProductReceiveItemDelete,

		// product transfer
		PermissionProductTransferCreate,
		PermissionProductTransferList,
		PermissionProductTransferShow,
		PermissionProductTransferUpdate,
		PermissionProductTransferDelete,
		PermissionProductTransferCreateItem,
		PermissionProductTransferDeleteItem,

		// product unit
		PermissionProductUnitOption,
		PermissionProductUnitOptionForProductAdjustmentItemForm,
		PermissionProductUnitOptionForProductTransferItemForm,
		PermissionProductUnitOptionForPriceTagDownload,

		// queue
		PermissionQueueShowProcess,
		PermissionQueueOption,
		PermissionQueueOptionForQueueCounterCheckInForm,

		// queue counter
		PermissionQueueCounterCheckIn,
		PermissionQueueCounterCheckOut,
		PermissionQueueCounterOptionForCheckInForm,
		PermissionQueueCounterOptionForQueueNumberCompleteForm,
		PermissionQueueCounterOptionForQueueNumberGenerateForm,

		// queue number
		PermissionQueueNumberGenerate,
		PermissionQueueNumberUploadInsuranceCard,
		PermissionQueueNumberGet,
		PermissionQueueNumberCurrent,
		PermissionQueueNumberUpdate,
		PermissionQueueNumberReprint,

		// queue number action
		PermissionQueueNumberCall,
		PermissionQueueNumberCallRepeat,
		PermissionQueueNumberProcess,
		PermissionQueueNumberComplete,
		PermissionQueueNumberMove,
		PermissionQueueNumberSkip,
		PermissionQueueNumberStandBy,

		// supplier
		PermissionSupplierOptionForProductReceiveForm,

		// task
		PermissionTaskCreate,
		PermissionTaskList,
		PermissionTaskShow,
		PermissionTaskUpdate,
		PermissionTaskUpdatePriority,
		PermissionTaskUpdateStatus,
		PermissionTaskDelete,

		// task attachment
		PermissionTaskAttachmentUpload,
		PermissionTaskAttachmentShow,
		PermissionTaskAttachmentDelete,

		// task comment
		PermissionTaskCommentCreate,
		PermissionTaskCommentList,
		PermissionTaskCommentUpdate,
		PermissionTaskCommentDelete,

		// unit
		PermissionUnitOptionForProductTransferForm,

		// user
		PermissionUserOptionForTaskForm,
	}
}

func GetRoleOperationManagerPermissions() []Permission {
	return []Permission{
		// admin appointment type
		PermissionAdminAppointmentTypeCreate,
		PermissionAdminAppointmentTypeList,
		PermissionAdminAppointmentTypeShow,
		PermissionAdminAppointmentTypeUpdate,
		PermissionAdminAppointmentTypeDelete,

		// admin clinic
		PermissionAdminClinicCreate,
		PermissionAdminClinicList,
		PermissionAdminClinicShow,
		PermissionAdminClinicUpdate,
		PermissionAdminClinicOption,
		PermissionAdminClinicOptionForAssignUserForm,
		PermissionAdminClinicOptionForDoctorFilter,
		PermissionAdminClinicOptionForNurseFilter,
		PermissionAdminClinicOptionForPharmacistFilter,
		PermissionAdminClinicOptionForParentClinic,
		PermissionAdminClinicOptionForUserFilter,
		PermissionAdminClinicOptionForProductPromotionForm,
		PermissionAdminClinicOptionForProductBundlingForm,

		// admin clinic type
		PermissionAdminClinicTypeOptionForClinicFilter,
		PermissionAdminClinicTypeOptionForClinicForm,

		// admin clinic user
		PermissionAdminClinicUserCreate,
		PermissionAdminClinicUserDelete,

		// admin clinic whitelist ip
		PermissionAdminClinicWhitelistIpCreate,
		PermissionAdminClinicWhitelistIpList,
		PermissionAdminClinicWhitelistIpDelete,

		// admin company
		PermissionAdminCompanyCreate,
		PermissionAdminCompanyList,
		PermissionAdminCompanyShow,
		PermissionAdminCompanyUpdate,
		PermissionAdminCompanyOption,
		PermissionAdminCompanyOptionForClinicFilter,
		PermissionAdminCompanyOptionForClinicForm,
		PermissionAdminCompanyOptionForReportNetPromoterScoreForm,
		PermissionAdminCompanyOptionForReportServiceAndPerformanceForm,
		PermissionAdminCompanyOptionForUserForm,

		// admin doctor
		PermissionAdminDoctorCreate,
		PermissionAdminDoctorUploadAvatar,
		PermissionAdminDoctorList,
		PermissionAdminDoctorShow,
		PermissionAdminDoctorUpdate,
		PermissionAdminDoctorOption,

		// admin doctor qualification
		PermissionAdminDoctorQualificationCreate,
		PermissionAdminDoctorQualificationUpdate,
		PermissionAdminDoctorQualificationDelete,

		// admin examination
		PermissionAdminExaminationCreate,
		PermissionAdminExaminationList,
		PermissionAdminExaminationShow,
		PermissionAdminExaminationUpdate,
		PermissionAdminExaminationDelete,
		PermissionAdminExaminationOption,

		// admin examination category
		PermissionAdminExaminationCategoryCreate,
		PermissionAdminExaminationCategoryList,
		PermissionAdminExaminationCategoryShow,
		PermissionAdminExaminationCategoryUpdate,
		PermissionAdminExaminationCategoryDelete,
		PermissionAdminExaminationCategoryOption,

		// admin examination item
		PermissionAdminExaminationItemCreate,
		PermissionAdminExaminationItemUpdate,
		PermissionAdminExaminationItemDelete,

		// admin global setting
		PermissionAdminGlobalSettingList,
		PermissionAdminGlobalSettingUpdate,

		// admin icd 10
		PermissionAdminIcd10List,
		PermissionAdminIcd10Tree,

		// admin icd 10 trending
		PermissionAdminIcd10TrendingList,
		PermissionAdminIcd10TrendingCreate,
		PermissionAdminIcd10TrendingMove,
		PermissionAdminIcd10TrendingDelete,

		// admin injection
		PermissionAdminInjectionList,
		PermissionAdminInjectionShow,

		// admin vaccine
		PermissionAdminVaccineList,
		PermissionAdminVaccineShow,

		// admin medical record template
		PermissionAdminMedicalRecordTemplateCreate,
		PermissionAdminMedicalRecordTemplateList,
		PermissionAdminMedicalRecordTemplateShow,
		PermissionAdminMedicalRecordTemplateUpdate,
		PermissionAdminMedicalRecordTemplateDelete,

		// admin medical record template diagnose
		PermissionAdminMedicalRecordTemplateDiagnoseCreate,
		PermissionAdminMedicalRecordTemplateDiagnoseUpdate,
		PermissionAdminMedicalRecordTemplateDiagnoseDelete,

		// admin medical record template examination
		PermissionAdminMedicalRecordTemplateExaminationCreate,
		PermissionAdminMedicalRecordTemplateExaminationUpdate,
		PermissionAdminMedicalRecordTemplateExaminationDelete,

		// admin medical record template prescription
		PermissionAdminMedicalRecordTemplatePrescriptionCreate,
		PermissionAdminMedicalRecordTemplatePrescriptionUpdate,
		PermissionAdminMedicalRecordTemplatePrescriptionDelete,

		// admin medical record template medical tool
		PermissionAdminMedicalRecordTemplateMedicalToolCreate,
		PermissionAdminMedicalRecordTemplateMedicalToolUpdate,
		PermissionAdminMedicalRecordTemplateMedicalToolDelete,

		// admin medical record template treatment
		PermissionAdminMedicalRecordTemplateTreatmentCreate,
		PermissionAdminMedicalRecordTemplateTreatmentUpdate,
		PermissionAdminMedicalRecordTemplateTreatmentDelete,

		// admin medical tool
		PermissionAdminMedicalToolList,
		PermissionAdminMedicalToolShow,
		PermissionAdminMedicalToolOption,

		// admin medicine
		PermissionAdminMedicineList,
		PermissionAdminMedicineShow,

		// admin nurse
		PermissionAdminNurseCreate,
		PermissionAdminNurseList,
		PermissionAdminNurseShow,
		PermissionAdminNurseUpdate,

		// admin nurse qualification
		PermissionAdminNurseQualificationCreate,
		PermissionAdminNurseQualificationUpdate,
		PermissionAdminNurseQualificationDelete,

		// admin pharmacist
		PermissionAdminPharmacistCreate,
		PermissionAdminPharmacistList,
		PermissionAdminPharmacistShow,
		PermissionAdminPharmacistUpdate,
		PermissionAdminPharmacistOptionForCreateClinicForm,
		PermissionAdminPharmacistOptionForUpdateClinicForm,
		PermissionAdminPharmacistQualificationCreate,
		PermissionAdminPharmacistQualificationUpdate,
		PermissionAdminPharmacistQualificationDelete,

		// admin product
		PermissionAdminProductOption,
		PermissionAdminProductOptionForExaminationItem,
		PermissionAdminProductOptionForTreatmentItem,
		PermissionAdminProductOptionForProductPromotion,
		PermissionAdminProductOptionForProductBundlingItem,

		// admin product promotion
		PermissionAdminProductPromotionCreate,
		PermissionAdminProductPromotionList,
		PermissionAdminProductPromotionShow,
		PermissionAdminProductPromotionUpdate,
		PermissionAdminProductPromotionUpdateEndedAt,
		PermissionAdminProductPromotionDelete,
		PermissionAdminProductPromotionClinicCreate,
		PermissionAdminProductPromotionClinicDelete,

		// admin product bundling
		PermissionAdminProductBundlingCreate,
		PermissionAdminProductBundlingList,
		PermissionAdminProductBundlingShow,
		PermissionAdminProductBundlingUpdate,
		PermissionAdminProductBundlingUpdateEndedAt,
		PermissionAdminProductBundlingDelete,

		// admin product bundling clinic
		PermissionAdminProductBundlingClinicCreate,
		PermissionAdminProductBundlingClinicDelete,

		// admin product bundling item
		PermissionAdminProductBundlingItemCreate,
		PermissionAdminProductBundlingItemDelete,

		// admin product unit
		PermissionAdminProductUnitOption,

		// admin qualification type
		PermissionAdminQualificationTypeCreate,
		PermissionAdminQualificationTypeList,
		PermissionAdminQualificationTypeShow,
		PermissionAdminQualificationTypeUpdate,
		PermissionAdminQualificationTypeDelete,
		PermissionAdminQualificationTypeOption,

		// admin retail product
		PermissionAdminRetailProductList,
		PermissionAdminRetailProductShow,

		// admin role
		PermissionAdminRoleOptionForUserFilter,
		PermissionAdminRoleOptionForUserForm,

		// admin supplier
		PermissionAdminSupplierCreate,
		PermissionAdminSupplierFetch,
		PermissionAdminSupplierGet,
		PermissionAdminSupplierUpdate,
		PermissionAdminSupplierDelete,
		PermissionAdminSupplierOptionForExaminationForm,

		// admin treatment
		PermissionAdminTreatmentCreate,
		PermissionAdminTreatmentList,
		PermissionAdminTreatmentShow,
		PermissionAdminTreatmentUpdate,
		PermissionAdminTreatmentDelete,
		PermissionAdminTreatmentOption,

		// admin treatment category
		PermissionAdminTreatmentCategoryCreate,
		PermissionAdminTreatmentCategoryList,
		PermissionAdminTreatmentCategoryShow,
		PermissionAdminTreatmentCategoryUpdate,
		PermissionAdminTreatmentCategoryDelete,
		PermissionAdminTreatmentCategoryOption,

		// admin treatment item
		PermissionAdminTreatmentItemCreate,
		PermissionAdminTreatmentItemUpdate,
		PermissionAdminTreatmentItemDelete,

		// admin user
		PermissionAdminUserCreate,
		PermissionAdminUserList,
		PermissionAdminUserShow,
		PermissionAdminUserUpdate,
		PermissionAdminUserOptionForDoctorForm,
		PermissionAdminUserOptionForNurseForm,
		PermissionAdminUserOptionForPharmacistForm,

		// admin voucher code
		PermissionAdminVoucherCodeList,
		PermissionAdminVoucherCodeShow,

		// admin voucher code redeem
		PermissionAdminVoucherCodeRedeemList,
		PermissionAdminVoucherCodeRedeemShow,

		// admin voucher generate rule
		PermissionAdminVoucherGenerateRuleList,
		PermissionAdminVoucherGenerateRuleShow,

		// appointment
		PermissionAppointmentList,
		PermissionAppointmentShow,

		// bank
		PermissionBankOptionForCashDepositForm,

		// cash deposit
		PermissionAdminCashDepositCreate,
		PermissionAdminCashDepositUpload,
		PermissionAdminCashDepositList,
		PermissionAdminCashDepositShow,
		PermissionAdminCashDepositDownload,

		// cashier summary
		PermissionClusterManagerCashierSummaryList,
		PermissionClusterManagerCashierSummaryListSummarized,
		PermissionClusterManagerCashierSummaryShow,
		PermissionClusterManagerCashierSummaryShowSession,
		PermissionClusterManagerCashierSummaryShowSalesInvoice,
		PermissionClusterManagerCashierSummaryDownloadCsv,

		// clinic
		PermissionClusterManagerClinicGetCashOnHand,
		PermissionClusterManagerClinicOptionForCashierSummaryList,
		PermissionClusterManagerClinicOptionForCashDepositForm,
		PermissionClusterManagerClinicOptionForCashDepositFilter,

		// clinic setting
		PermissionClinicSettingList,
		PermissionClinicSettingUpdate,

		// product unit
		PermissionProductUnitOption,

		// report daily transaction
		PermissionReportDailyTransactionList,
		PermissionReportDailyTransactionDownload,

		// report inventory stock
		PermissionReportInventoryStockCreate,
		PermissionReportInventoryStockList,
		PermissionReportInventoryStockDownload,

		// report net promoter score
		PermissionReportNetPromoterScoreCreate,
		PermissionReportNetPromoterScoreList,
		PermissionReportNetPromoterScoreDownload,

		// report service and performance
		PermissionReportServiceAndPerformanceCreate,
		PermissionReportServiceAndPerformanceList,
		PermissionReportServiceAndPerformanceDownload,

		// report month to date sales
		PermissionReportMonthToDateSalesList,
		PermissionReportMonthToDateSalesDownload,

		// report waive fee
		PermissionReportWaiveFeeDownload,
	}
}

func GetRolePharmacyManagerPermissions() []Permission {
	return []Permission{
		// admin clinic
		PermissionAdminClinicOptionForProductPromotionForm,
		PermissionAdminClinicOptionForProductBundlingForm,

		// admin injection
		PermissionAdminInjectionCreate,
		PermissionAdminInjectionList,
		PermissionAdminInjectionShow,
		PermissionAdminInjectionUpdate,
		PermissionAdminInjectionDelete,

		// admin injection attention
		PermissionAdminInjectionAttentionCreate,
		PermissionAdminInjectionAttentionDelete,

		// admin injection category
		PermissionAdminInjectionCategoryCreate,
		PermissionAdminInjectionCategoryDelete,

		// admin injection instruction
		PermissionAdminInjectionInstructionCreate,
		PermissionAdminInjectionInstructionDelete,

		// admin injection unit
		PermissionAdminInjectionUnitCreate,
		PermissionAdminInjectionUnitUpdate,
		PermissionAdminInjectionUnitDelete,

		// admin medical tool
		PermissionAdminMedicalToolCreate,
		PermissionAdminMedicalToolList,
		PermissionAdminMedicalToolShow,
		PermissionAdminMedicalToolUpdate,
		PermissionAdminMedicalToolDelete,
		PermissionAdminMedicalToolCreateProductUnit,
		PermissionAdminMedicalToolUpdateProductUnit,
		PermissionAdminMedicalToolDeleteProductUnit,

		// admin medicinal attention
		PermissionAdminMedicinalAttentionCreate,
		PermissionAdminMedicinalAttentionList,
		PermissionAdminMedicinalAttentionUpdate,
		PermissionAdminMedicinalAttentionDelete,
		PermissionAdminMedicinalAttentionOptionForAmpForm,

		// admin medicinal category
		PermissionAdminMedicinalCategoryCreate,
		PermissionAdminMedicinalCategoryList,
		PermissionAdminMedicinalCategoryUpdate,
		PermissionAdminMedicinalCategoryDelete,
		PermissionAdminMedicinalCategoryOptionForAmpForm,

		// admin medicinal classification
		PermissionAdminMedicinalClassificationCreate,
		PermissionAdminMedicinalClassificationList,
		PermissionAdminMedicinalClassificationUpdate,
		PermissionAdminMedicinalClassificationDelete,
		PermissionAdminMedicinalClassificationOption,

		// admin medicinal controlled drug classification
		PermissionAdminMedicinalControlledDrugClassificationCreate,
		PermissionAdminMedicinalControlledDrugClassificationList,
		PermissionAdminMedicinalControlledDrugClassificationUpdate,
		PermissionAdminMedicinalControlledDrugClassificationDelete,
		PermissionAdminMedicinalControlledDrugClassificationOption,

		// admin medicinal instruction
		PermissionAdminMedicinalInstructionCreate,
		PermissionAdminMedicinalInstructionList,
		PermissionAdminMedicinalInstructionUpdate,
		PermissionAdminMedicinalInstructionDelete,
		PermissionAdminMedicinalInstructionOptionForAmpForm,

		// admin medicinal form
		PermissionAdminMedicinalFormCreate,
		PermissionAdminMedicinalFormList,
		PermissionAdminMedicinalFormUpdate,
		PermissionAdminMedicinalFormDelete,
		PermissionAdminMedicinalFormOption,

		// admin medicinal route
		PermissionAdminMedicinalRouteCreate,
		PermissionAdminMedicinalRouteList,
		PermissionAdminMedicinalRouteUpdate,
		PermissionAdminMedicinalRouteDelete,
		PermissionAdminMedicinalRouteOption,

		// admin medicine
		PermissionAdminMedicineCreate,
		PermissionAdminMedicineList,
		PermissionAdminMedicineShow,
		PermissionAdminMedicineUpdate,
		PermissionAdminMedicineDelete,

		// admin medicine attention
		PermissionAdminMedicineAttentionCreate,
		PermissionAdminMedicineAttentionDelete,

		// admin medicine category
		PermissionAdminMedicineCategoryCreate,
		PermissionAdminMedicineCategoryDelete,

		// admin medicine instruction
		PermissionAdminMedicineInstructionCreate,
		PermissionAdminMedicineInstructionDelete,

		// admin medicine unit
		PermissionAdminMedicineUnitCreate,
		PermissionAdminMedicineUnitUpdate,
		PermissionAdminMedicineUnitDelete,

		// admin product
		PermissionAdminProductOptionForProductPromotion,
		PermissionAdminProductOptionForProductBundlingItem,

		// admin product promotion
		PermissionAdminProductPromotionCreate,
		PermissionAdminProductPromotionList,
		PermissionAdminProductPromotionShow,
		PermissionAdminProductPromotionUpdate,
		PermissionAdminProductPromotionUpdateEndedAt,
		PermissionAdminProductPromotionDelete,
		PermissionAdminProductPromotionClinicCreate,
		PermissionAdminProductPromotionClinicDelete,

		// admin product bundling
		PermissionAdminProductBundlingCreate,
		PermissionAdminProductBundlingList,
		PermissionAdminProductBundlingShow,
		PermissionAdminProductBundlingUpdate,
		PermissionAdminProductBundlingUpdateEndedAt,
		PermissionAdminProductBundlingDelete,

		// admin product bundling clinic
		PermissionAdminProductBundlingClinicCreate,
		PermissionAdminProductBundlingClinicDelete,

		// admin product bundling item
		PermissionAdminProductBundlingItemCreate,
		PermissionAdminProductBundlingItemDelete,

		// admin product unit
		PermissionAdminProductUnitOption,

		// admin retail product
		PermissionAdminRetailProductCreate,
		PermissionAdminRetailProductUpload,
		PermissionAdminRetailProductList,
		PermissionAdminRetailProductShow,
		PermissionAdminRetailProductUpdate,
		PermissionAdminRetailProductDelete,
		PermissionAdminRetailProductCreateProductUnit,
		PermissionAdminRetailProductUpdateProductUnit,
		PermissionAdminRetailProductDeleteProductUnit,

		// admin retail product category
		PermissionAdminRetailProductCategoryCreate,
		PermissionAdminRetailProductCategoryList,
		PermissionAdminRetailProductCategoryShow,
		PermissionAdminRetailProductCategoryUpdate,
		PermissionAdminRetailProductCategoryDelete,
		PermissionAdminRetailProductCategoryOptionForRetailProductForm,

		// admin supplier
		PermissionAdminSupplierCreate,
		PermissionAdminSupplierFetch,
		PermissionAdminSupplierGet,
		PermissionAdminSupplierUpdate,
		PermissionAdminSupplierDelete,

		// admin unit
		PermissionAdminUnitCreate,
		PermissionAdminUnitList,
		PermissionAdminUnitUpdate,
		PermissionAdminUnitDelete,
		PermissionAdminUnitOptionForProductForm,

		// admin vaccine
		PermissionAdminVaccineCreate,
		PermissionAdminVaccineList,
		PermissionAdminVaccineShow,
		PermissionAdminVaccineUpdate,
		PermissionAdminVaccineDelete,

		// admin vaccine attention
		PermissionAdminVaccineAttentionCreate,
		PermissionAdminVaccineAttentionDelete,

		// admin vaccine category
		PermissionAdminVaccineCategoryCreate,
		PermissionAdminVaccineCategoryDelete,

		// admin vaccine instruction
		PermissionAdminVaccineInstructionCreate,
		PermissionAdminVaccineInstructionDelete,

		// admin vaccine unit
		PermissionAdminVaccineUnitCreate,
		PermissionAdminVaccineUnitUpdate,
		PermissionAdminVaccineUnitDelete,

		// admin virtual medicinal product
		PermissionAdminVirtualMedicinalProductCreate,
		PermissionAdminVirtualMedicinalProductList,
		PermissionAdminVirtualMedicinalProductShow,
		PermissionAdminVirtualMedicinalProductUpdate,
		PermissionAdminVirtualMedicinalProductDelete,
		PermissionAdminVirtualMedicinalProductOption,

		// admin virtual therapeutic moiety
		PermissionAdminVirtualTherapeuticMoietyList,
		PermissionAdminVirtualTherapeuticMoietyShow,
		PermissionAdminVirtualTherapeuticMoietyCreate,
		PermissionAdminVirtualTherapeuticMoietyUpdate,
		PermissionAdminVirtualTherapeuticMoietyDelete,
		PermissionAdminVirtualTherapeuticMoietyOption,
		PermissionAdminVirtualTherapeuticMoietyOptionForVmpVtmForm,

		// admin vmp code
		PermissionAdminVmpCodeCreate,
		PermissionAdminVmpCodeDelete,

		// admin vmp vtm
		PermissionAdminVmpVtmCreate,
		PermissionAdminVmpVtmUpdate,
		PermissionAdminVmpVtmDelete,

		// admin vtm code
		PermissionAdminVtmCodeCreate,
		PermissionAdminVtmCodeDelete,

		// cashier summary
		PermissionClusterManagerCashierSummaryList,
		PermissionClusterManagerCashierSummaryListSummarized,
		PermissionClusterManagerCashierSummaryShow,
		PermissionClusterManagerCashierSummaryShowSession,
		PermissionClusterManagerCashierSummaryShowSalesInvoice,
		PermissionClusterManagerCashierSummaryDownloadCsv,

		// clinic
		PermissionClusterManagerClinicOptionForCashierSummaryList,

		// price tag
		PermissionProductPriceTagDownload,

		// product
		PermissionProductOptionForProductReceiveItem,
		PermissionProductOptionForPriceTagDownload,
		PermissionProductOptionForProductTransferItemForm,
		PermissionProductOptionForProductAdjustmentForm,
		PermissionProductOptionExpiredDate,

		// product adjustment
		PermissionProductAdjustmentCreate,
		PermissionProductAdjustmentList,
		PermissionProductAdjustmentShow,
		PermissionProductAdjustmentUpdate,
		PermissionProductAdjustmentDelete,

		// product adjustment item
		PermissionProductAdjustmentItemCreate,
		PermissionProductAdjustmentItemDelete,

		// product inventory
		PermissionProductInventoryList,

		// product promotion
		PermissionProductPromotionOptionForPriceTagDownload,

		// product unit
		PermissionProductUnitOption,
		PermissionProductUnitOptionForProductAdjustmentItemForm,
		PermissionProductUnitOptionForProductTransferItemForm,

		// product receive
		PermissionProductReceiveCreate,
		PermissionProductReceiveList,
		PermissionProductReceiveShow,
		PermissionProductReceiveUpdate,
		PermissionProductReceiveDelete,

		// product receive item
		PermissionProductReceiveItemCreate,
		PermissionProductReceiveItemDelete,

		// product transfer
		PermissionProductTransferCreate,
		PermissionProductTransferList,
		PermissionProductTransferShow,
		PermissionProductTransferUpdate,
		PermissionProductTransferDelete,

		// product transfer item
		PermissionProductTransferCreateItem,
		PermissionProductTransferDeleteItem,

		// supplier
		PermissionSupplierOptionForProductReceiveForm,

		// report inventory stock
		PermissionReportInventoryStockCreate,
		PermissionReportInventoryStockList,
		PermissionReportInventoryStockDownload,

		// report daily transaction
		PermissionReportDailyTransactionList,
		PermissionReportDailyTransactionDownload,

		// report month to date sales
		PermissionReportMonthToDateSalesList,
		PermissionReportMonthToDateSalesDownload,
	}
}

func GetRoleFinanceManagerPermissions() []Permission {
	return []Permission{
		// admin bank
		PermissionAdminBankCreate,
		PermissionAdminBankList,
		PermissionAdminBankShow,
		PermissionAdminBankUpdate,
		PermissionAdminBankDelete,

		// admin clinic
		PermissionAdminClinicOptionForBankForm,

		// admin company
		PermissionAdminCompanyOptionForBankForm,

		// admin voucher code
		PermissionAdminVoucherCodeCreate,
		PermissionAdminVoucherCodeList,
		PermissionAdminVoucherCodeShow,
		PermissionAdminVoucherCodeUpdate,
		PermissionAdminVoucherCodeDelete,

		// admin voucher code redeem
		PermissionAdminVoucherCodeRedeemList,
		PermissionAdminVoucherCodeRedeemShow,

		// admin voucher generate rule
		PermissionAdminVoucherGenerateRuleCreate,
		PermissionAdminVoucherGenerateRuleList,
		PermissionAdminVoucherGenerateRuleShow,
		PermissionAdminVoucherGenerateRuleUpdate,
		PermissionAdminVoucherGenerateRuleDelete,

		// bank
		PermissionBankOptionForCashDepositForm,

		// cash deposit
		PermissionAdminCashDepositCreate,
		PermissionAdminCashDepositUpload,
		PermissionAdminCashDepositList,
		PermissionAdminCashDepositShow,
		PermissionAdminCashDepositDownload,

		// cashier summary
		PermissionClusterManagerCashierSummaryList,
		PermissionClusterManagerCashierSummaryListSummarized,
		PermissionClusterManagerCashierSummaryShow,
		PermissionClusterManagerCashierSummaryShowSession,
		PermissionClusterManagerCashierSummaryShowSalesInvoice,
		PermissionClusterManagerCashierSummaryDownloadCsv,

		// clinic
		PermissionClusterManagerClinicGetCashOnHand,
		PermissionClusterManagerClinicOptionForCashierSummaryList,
		PermissionClusterManagerClinicOptionForCashDepositForm,
		PermissionClusterManagerClinicOptionForCashDepositFilter,

		// report daily transaction
		PermissionReportDailyTransactionList,
		PermissionReportDailyTransactionDownload,

		// report finance inventory
		PermissionReportFinanceInventoryCreate,
		PermissionReportFinanceInventoryList,
		PermissionReportFinanceInventoryDownload,

		// report monthly revenue
		PermissionReportMonthlyRevenueList,
		PermissionReportMonthlyRevenueDownload,

		// report month to date sales
		PermissionReportMonthToDateSalesList,
		PermissionReportMonthToDateSalesDownload,
	}
}

func GetRoleAssistantClinicManagerPermissions() []Permission {
	return []Permission{
		// appointment
		PermissionAppointmentListForQueueNumberGenerateFormByPatient,

		// cashier session
		PermissionCashierSessionList,
		PermissionCashierSessionShow,
		PermissionCashierSessionUpdateMoneyCount,
		PermissionCashierSessionEnd,

		// cashier summary
		PermissionCashierSummaryCreate,
		PermissionCashierSummaryShow,
		PermissionCashierSummaryShowSession,
		PermissionCashierSummaryShowSalesInvoice,
		PermissionCashierSummaryShowDailyReport,
		PermissionCashierSummaryShowDailyReportSession,
		PermissionCashierSummaryShowDailyReportSalesInvoice,

		// clinic manager
		PermissionClinicManagerSalesInvoiceList,
		PermissionClinicManagerSalesInvoiceDownload,

		// clinic printer
		PermissionClinicPrinterOptionForQueueCounter,

		// patient
		PermissionPatientOptionForQueueNumberGenerateForm,

		// patient relative
		PermissionPatientRelativeOptionForQueueNumberForm,

		// patient visit
		PermissionPatientVisitInsuranceCardFileShowForQueueNumber,

		// queue
		PermissionQueueCreate,
		PermissionQueueList,
		PermissionQueueShow,
		PermissionQueueUpdate,
		PermissionQueueDelete,
		PermissionQueueOption,
		PermissionQueueOptionForQueueDisplay,

		// queue number
		PermissionQueueNumberGet,
		PermissionQueueNumberUpdate,

		// queue number action
		PermissionQueueNumberMove,

		// queue counter
		PermissionQueueCounterCreate,
		PermissionQueueCounterUpdate,
		PermissionQueueCounterOptionForQueueNumberGenerateForm,

		// queue display
		PermissionQueueDisplayCreate,
		PermissionQueueDisplayList,
		PermissionQueueDisplayShow,
		PermissionQueueDisplayUpdate,
		PermissionQueueDisplayDelete,

		// queue display banner
		PermissionQueueDisplayBannerCreate,
		PermissionQueueDisplayBannerUpload,
		PermissionQueueDisplayBannerMove,
		PermissionQueueDisplayBannerDelete,

		// queue display queue
		PermissionQueueDisplayQueueCreate,
		PermissionQueueDisplayQueueDelete,

		// queue display running text
		PermissionQueueDisplayRunningTextCreate,
		PermissionQueueDisplayRunningTextMove,
		PermissionQueueDisplayRunningTextDelete,

		// queue session
		PermissionQueueSessionBegin,
		PermissionQueueSessionEnd,
		PermissionQueueSessionGetActive,
		PermissionQueueSessionCheckForActiveQueueNumber,
	}
}

func ListRoleTypeDoctor() []Role {
	roles := []Role{}
	for _, role := range ListRole() {
		if role.RoleType() == RoleTypeDoctor {
			roles = append(roles, role)
		}
	}

	return roles
}

func ListRoleTypeNurse() []Role {
	roles := []Role{}
	for _, role := range ListRole() {
		if role.RoleType() == RoleTypeNurse {
			roles = append(roles, role)
		}
	}

	return roles
}

func ListRoleTypePharmacist() []Role {
	roles := []Role{}
	for _, role := range ListRole() {
		if role.RoleType() == RoleTypePharmacist {
			roles = append(roles, role)
		}
	}

	return roles
}
