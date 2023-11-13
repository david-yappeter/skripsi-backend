package data_type

import (
	"fmt"
)

//go:generate go run myapp/tool/stringer -linecomment -type=Permission -output=permission_enum_gen.go -swagoutput=../tool/swag/enum_gen/permission_enum_gen.go -custom
type Permission int // @name PermissionEnum

const (
	// admin appointment type
	PermissionAdminAppointmentTypeCreate Permission = iota + 1 // ADMIN_APPOINTMENT_TYPE_CREATE
	PermissionAdminAppointmentTypeList                         // ADMIN_APPOINTMENT_TYPE_LIST
	PermissionAdminAppointmentTypeShow                         // ADMIN_APPOINTMENT_TYPE_SHOW
	PermissionAdminAppointmentTypeUpdate                       // ADMIN_APPOINTMENT_TYPE_UPDATE
	PermissionAdminAppointmentTypeDelete                       // ADMIN_APPOINTMENT_TYPE_DELETE

	PermissionAdminBankCreate // ADMIN_BANK_CREATE
	PermissionAdminBankList   // ADMIN_BANK_LIST
	PermissionAdminBankShow   // ADMIN_BANK_SHOW
	PermissionAdminBankUpdate // ADMIN_BANK_UPDATE
	PermissionAdminBankDelete // ADMIN_BANK_DELETE

	// admin cash deposit
	PermissionAdminCashDepositCreate   // ADMIN_CASH_DEPOSIT_CREATE
	PermissionAdminCashDepositUpload   // ADMIN_CASH_DEPOSIT_UPLOAD
	PermissionAdminCashDepositList     // ADMIN_CASH_DEPOSIT_LIST
	PermissionAdminCashDepositShow     // ADMIN_CASH_DEPOSIT_SHOW
	PermissionAdminCashDepositDownload // ADMIN_CASH_DEPOSIT_DOWNLOAD

	// admin clinic
	PermissionAdminClinicCreate                        // ADMIN_CLINIC_CREATE
	PermissionAdminClinicList                          // ADMIN_CLINIC_LIST
	PermissionAdminClinicShow                          // ADMIN_CLINIC_SHOW
	PermissionAdminClinicUpdate                        // ADMIN_CLINIC_UPDATE
	PermissionAdminClinicOption                        // ADMIN_CLINIC_OPTION
	PermissionAdminClinicOptionForAssignUserForm       // ADMIN_CLINIC_OPTION_FOR_ASSIGN_USER_FORM
	PermissionAdminClinicOptionForDoctorFilter         // ADMIN_CLINIC_OPTION_FOR_DOCTOR_FILTER
	PermissionAdminClinicOptionForNurseFilter          // ADMIN_CLINIC_OPTION_FOR_NURSE_FILTER
	PermissionAdminClinicOptionForPharmacistFilter     // ADMIN_CLINIC_OPTION_FOR_PHARMACIST_FILTER
	PermissionAdminClinicOptionForParentClinic         // ADMIN_CLINIC_OPTION_FOR_PARENT_CLINIC
	PermissionAdminClinicOptionForUserFilter           // ADMIN_CLINIC_OPTION_FOR_USER_FILTER
	PermissionAdminClinicOptionForBankForm             // ADMIN_CLINIC_OPTION_FOR_BANK_FORM
	PermissionAdminClinicOptionForProductPromotionForm // ADMIN_CLINIC_OPTION_FOR_PRODUCT_PROMOTION_FORM
	PermissionAdminClinicOptionForProductBundlingForm  // ADMIN_CLINIC_OPTION_FOR_PRODUCT_BUNDLING_FORM

	// admin clinic type
	PermissionAdminClinicTypeOptionForClinicFilter // ADMIN_CLINIC_TYPE_OPTION_FOR_CLINIC_FILTER
	PermissionAdminClinicTypeOptionForClinicForm   // ADMIN_CLINIC_TYPE_OPTION_FOR_CLINIC_FORM

	// admin clinic user
	PermissionAdminClinicUserCreate // ADMIN_CLINIC_USER_CREATE
	PermissionAdminClinicUserDelete // ADMIN_CLINIC_USER_DELETE

	// admin clinic whitelist ip
	PermissionAdminClinicWhitelistIpCreate // ADMIN_CLINIC_WHITELIST_IP_CREATE
	PermissionAdminClinicWhitelistIpList   // ADMIN_CLINIC_WHITELIST_IP_LIST
	PermissionAdminClinicWhitelistIpDelete // ADMIN_CLINIC_WHITELIST_IP_DELETE

	// admin company
	PermissionAdminCompanyCreate                                   // ADMIN_COMPANY_CREATE
	PermissionAdminCompanyList                                     // ADMIN_COMPANY_LIST
	PermissionAdminCompanyShow                                     // ADMIN_COMPANY_SHOW
	PermissionAdminCompanyUpdate                                   // ADMIN_COMPANY_UPDATE
	PermissionAdminCompanyOption                                   // ADMIN_COMPANY_OPTION
	PermissionAdminCompanyOptionForClinicFilter                    // ADMIN_COMPANY_OPTION_FOR_CLINIC_FILTER
	PermissionAdminCompanyOptionForClinicForm                      // ADMIN_COMPANY_OPTION_FOR_CLINIC_FORM
	PermissionAdminCompanyOptionForReportNetPromoterScoreForm      // ADMIN_COMPANY_OPTION_FOR_REPORT_NET_PROMOTER_SCORE_FORM
	PermissionAdminCompanyOptionForReportServiceAndPerformanceForm // ADMIN_COMPANY_OPTION_FOR_REPORT_SERVICE_AND_PERFORMANCE_FORM
	PermissionAdminCompanyOptionForUserForm                        // ADMIN_COMPANY_OPTION_FOR_USER_FORM
	PermissionAdminCompanyOptionForBankForm                        // ADMIN_COMPANY_OPTION_FOR_BANK_FORM

	// admin doctor
	PermissionAdminDoctorCreate       // ADMIN_DOCTOR_CREATE
	PermissionAdminDoctorUploadAvatar // ADMIN_DOCTOR_UPLOAD_AVATAR
	PermissionAdminDoctorList         // ADMIN_DOCTOR_LIST
	PermissionAdminDoctorShow         // ADMIN_DOCTOR_SHOW
	PermissionAdminDoctorUpdate       // ADMIN_DOCTOR_UPDATE
	PermissionAdminDoctorOption       // ADMIN_DOCTOR_OPTION

	// admin doctor qualification
	PermissionAdminDoctorQualificationCreate // ADMIN_DOCTOR_QUALIFICATION_CREATE
	PermissionAdminDoctorQualificationUpdate // ADMIN_DOCTOR_QUALIFICATION_UPDATE
	PermissionAdminDoctorQualificationDelete // ADMIN_DOCTOR_QUALIFICATION_DELETE

	// admin examination
	PermissionAdminExaminationCreate // ADMIN_EXAMINATION_CREATE
	PermissionAdminExaminationList   // ADMIN_EXAMINATION_LIST
	PermissionAdminExaminationShow   // ADMIN_EXAMINATION_SHOW
	PermissionAdminExaminationUpdate // ADMIN_EXAMINATION_UPDATE
	PermissionAdminExaminationDelete // ADMIN_EXAMINATION_DELETE
	PermissionAdminExaminationOption // ADMIN_EXAMINATION_OPTION

	// admin examination category
	PermissionAdminExaminationCategoryCreate // ADMIN_EXAMINATION_CATEGORY_CREATE
	PermissionAdminExaminationCategoryList   // ADMIN_EXAMINATION_CATEGORY_LIST
	PermissionAdminExaminationCategoryShow   // ADMIN_EXAMINATION_CATEGORY_SHOW
	PermissionAdminExaminationCategoryUpdate // ADMIN_EXAMINATION_CATEGORY_UPDATE
	PermissionAdminExaminationCategoryDelete // ADMIN_EXAMINATION_CATEGORY_DELETE
	PermissionAdminExaminationCategoryOption // ADMIN_EXAMINATION_CATEGORY_OPTION

	// admin examination item
	PermissionAdminExaminationItemCreate // ADMIN_EXAMINATION_ITEM_CREATE
	PermissionAdminExaminationItemUpdate // ADMIN_EXAMINATION_ITEM_UPDATE
	PermissionAdminExaminationItemDelete // ADMIN_EXAMINATION_ITEM_DELETE

	// global setting
	PermissionAdminGlobalSettingList   // ADMIN_GLOBAL_SETTING_LIST
	PermissionAdminGlobalSettingUpdate // ADMIN_GLOBAL_SETTING_UPDATE

	// admin icd10
	PermissionAdminIcd10List // ADMIN_ICD_10_LIST
	PermissionAdminIcd10Tree // ADMIN_ICD_10_TREE

	// admin icd10 trending
	PermissionAdminIcd10TrendingList   // ADMIN_ICD_10_TRENDING_LIST
	PermissionAdminIcd10TrendingCreate // ADMIN_ICD_10_TRENDING_CREATE
	PermissionAdminIcd10TrendingMove   // ADMIN_ICD_10_TRENDING_MOVE
	PermissionAdminIcd10TrendingDelete // ADMIN_ICD_10_TRENDING_DELETE

	// admin injection
	PermissionAdminInjectionCreate // ADMIN_INJECTION_CREATE
	PermissionAdminInjectionList   // ADMIN_INJECTION_LIST
	PermissionAdminInjectionShow   // ADMIN_INJECTION_SHOW
	PermissionAdminInjectionUpdate // ADMIN_INJECTION_UPDATE
	PermissionAdminInjectionDelete // ADMIN_INJECTION_DELETE

	// admin injection attention
	PermissionAdminInjectionAttentionCreate // ADMIN_INJECTION_ATTENTION_CREATE
	PermissionAdminInjectionAttentionDelete // ADMIN_INJECTION_ATTENTION_DELETE

	// admin injection category
	PermissionAdminInjectionCategoryCreate // ADMIN_INJECTION_CATEGORY_CREATE
	PermissionAdminInjectionCategoryDelete // ADMIN_INJECTION_CATEGORY_DELETE

	// admin injection instruction
	PermissionAdminInjectionInstructionCreate // ADMIN_INJECTION_INSTRUCTION_CREATE
	PermissionAdminInjectionInstructionDelete // ADMIN_INJECTION_INSTRUCTION_DELETE

	// admin injection unit
	PermissionAdminInjectionUnitCreate // ADMIN_INJECTION_UNIT_CREATE
	PermissionAdminInjectionUnitUpdate // ADMIN_INJECTION_UNIT_UPDATE
	PermissionAdminInjectionUnitDelete // ADMIN_INJECTION_UNIT_DELETE

	// admin insurance
	PermissionAdminInsuranceCreate // ADMIN_INSURANCE_CREATE
	PermissionAdminInsuranceList   // ADMIN_INSURANCE_LIST
	PermissionAdminInsuranceShow   // ADMIN_INSURANCE_SHOW
	PermissionAdminInsuranceUpdate // ADMIN_INSURANCE_UPDATE
	PermissionAdminInsuranceDelete // ADMIN_INSURANCE_DELETE

	// admin insurance payor
	PermissionAdminInsurancePayorCreate // ADMIN_INSURANCE_PAYOR_CREATE
	PermissionAdminInsurancePayorList   // ADMIN_INSURANCE_PAYOR_LIST
	PermissionAdminInsurancePayorShow   // ADMIN_INSURANCE_PAYOR_SHOW
	PermissionAdminInsurancePayorUpdate // ADMIN_INSURANCE_PAYOR_UPDATE
	PermissionAdminInsurancePayorDelete // ADMIN_INSURANCE_PAYOR_DELETE

	// admin medical record template
	PermissionAdminMedicalRecordTemplateCreate // ADMIN_MEDICAL_RECORD_TEMPLATE_CREATE
	PermissionAdminMedicalRecordTemplateList   // ADMIN_MEDICAL_RECORD_TEMPLATE_LIST
	PermissionAdminMedicalRecordTemplateShow   // ADMIN_MEDICAL_RECORD_TEMPLATE_SHOW
	PermissionAdminMedicalRecordTemplateUpdate // ADMIN_MEDICAL_RECORD_TEMPLATE_UPDATE
	PermissionAdminMedicalRecordTemplateDelete // ADMIN_MEDICAL_RECORD_TEMPLATE_DELETE

	// admin medical record template diagnose
	PermissionAdminMedicalRecordTemplateDiagnoseCreate // ADMIN_MEDICAL_RECORD_TEMPLATE_DIAGNOSE_CREATE
	PermissionAdminMedicalRecordTemplateDiagnoseUpdate // ADMIN_MEDICAL_RECORD_TEMPLATE_DIAGNOSE_UPDATE
	PermissionAdminMedicalRecordTemplateDiagnoseDelete // ADMIN_MEDICAL_RECORD_TEMPLATE_DIAGNOSE_DELETE

	// admin medical record template examination
	PermissionAdminMedicalRecordTemplateExaminationCreate // ADMIN_MEDICAL_RECORD_TEMPLATE_EXAMINATION_CREATE
	PermissionAdminMedicalRecordTemplateExaminationUpdate // ADMIN_MEDICAL_RECORD_TEMPLATE_EXAMINATION_UPDATE
	PermissionAdminMedicalRecordTemplateExaminationDelete // ADMIN_MEDICAL_RECORD_TEMPLATE_EXAMINATION_DELETE

	// admin medical record template prescription
	PermissionAdminMedicalRecordTemplatePrescriptionCreate // ADMIN_MEDICAL_RECORD_TEMPLATE_PRESCRIPTION_CREATE
	PermissionAdminMedicalRecordTemplatePrescriptionUpdate // ADMIN_MEDICAL_RECORD_TEMPLATE_PRESCRIPTION_UPDATE
	PermissionAdminMedicalRecordTemplatePrescriptionDelete // ADMIN_MEDICAL_RECORD_TEMPLATE_PRESCRIPTION_DELETE

	// admin medical record template medical tool
	PermissionAdminMedicalRecordTemplateMedicalToolCreate // ADMIN_MEDICAL_RECORD_TEMPLATE_MEDICAL_TOOL_CREATE
	PermissionAdminMedicalRecordTemplateMedicalToolUpdate // ADMIN_MEDICAL_RECORD_TEMPLATE_MEDICAL_TOOL_UPDATE
	PermissionAdminMedicalRecordTemplateMedicalToolDelete // ADMIN_MEDICAL_RECORD_TEMPLATE_MEDICAL_TOOL_DELETE

	// admin medical record template treatment
	PermissionAdminMedicalRecordTemplateTreatmentCreate // ADMIN_MEDICAL_RECORD_TEMPLATE_TREATMENT_CREATE
	PermissionAdminMedicalRecordTemplateTreatmentUpdate // ADMIN_MEDICAL_RECORD_TEMPLATE_TREATMENT_UPDATE
	PermissionAdminMedicalRecordTemplateTreatmentDelete // ADMIN_MEDICAL_RECORD_TEMPLATE_TREATMENT_DELETE

	// admin medical tool
	PermissionAdminMedicalToolCreate            // ADMIN_MEDICAL_TOOL_CREATE
	PermissionAdminMedicalToolList              // ADMIN_MEDICAL_TOOL_LIST
	PermissionAdminMedicalToolShow              // ADMIN_MEDICAL_TOOL_SHOW
	PermissionAdminMedicalToolUpdate            // ADMIN_MEDICAL_TOOL_UPDATE
	PermissionAdminMedicalToolDelete            // ADMIN_MEDICAL_TOOL_DELETE
	PermissionAdminMedicalToolCreateProductUnit // ADMIN_MEDICAL_TOOL_CREATE_PRODUCT_UNIT
	PermissionAdminMedicalToolUpdateProductUnit // ADMIN_MEDICAL_TOOL_UPDATE_PRODUCT_UNIT
	PermissionAdminMedicalToolDeleteProductUnit // ADMIN_MEDICAL_TOOL_DELETE_PRODUCT_UNIT
	PermissionAdminMedicalToolOption            // ADMIN_MEDICAL_TOOL_OPTION

	// admin medicinal attention
	PermissionAdminMedicinalAttentionCreate           // ADMIN_MEDICINAL_ATTENTION_CREATE
	PermissionAdminMedicinalAttentionList             // ADMIN_MEDICINAL_ATTENTION_LIST
	PermissionAdminMedicinalAttentionUpdate           // ADMIN_MEDICINAL_ATTENTION_UPDATE
	PermissionAdminMedicinalAttentionDelete           // ADMIN_MEDICINAL_ATTENTION_DELETE
	PermissionAdminMedicinalAttentionOptionForAmpForm // ADMIN_MEDICINAL_ATTENTION_OPTION_FOR_AMP_FORM

	// admin medicinal category
	PermissionAdminMedicinalCategoryCreate           // ADMIN_MEDICINAL_CATEGORY_CREATE
	PermissionAdminMedicinalCategoryList             // ADMIN_MEDICINAL_CATEGORY_LIST
	PermissionAdminMedicinalCategoryUpdate           // ADMIN_MEDICINAL_CATEGORY_UPDATE
	PermissionAdminMedicinalCategoryDelete           // ADMIN_MEDICINAL_CATEGORY_DELETE
	PermissionAdminMedicinalCategoryOptionForAmpForm // ADMIN_MEDICINAL_CATEGORY_OPTION_FOR_AMP_FORM

	// admin medicinal classification
	PermissionAdminMedicinalClassificationCreate // ADMIN_MEDICINAL_CLASSIFICATION_CREATE
	PermissionAdminMedicinalClassificationUpdate // ADMIN_MEDICINAL_CLASSIFICATION_UPDATE
	PermissionAdminMedicinalClassificationList   // ADMIN_MEDICINAL_CLASSIFICATION_LIST
	PermissionAdminMedicinalClassificationDelete // ADMIN_MEDICINAL_CLASSIFICATION_DELETE
	PermissionAdminMedicinalClassificationOption // ADMIN_MEDICINAL_CLASSIFICATION_OPTION

	// admin medicinal controlled drug classification
	PermissionAdminMedicinalControlledDrugClassificationCreate // ADMIN_MEDICINAL_CONTROLLED_DRUG_CLASSIFICATION_CREATE
	PermissionAdminMedicinalControlledDrugClassificationList   // ADMIN_MEDICINAL_CONTROLLED_DRUG_CLASSIFICATION_LIST
	PermissionAdminMedicinalControlledDrugClassificationUpdate // ADMIN_MEDICINAL_CONTROLLED_DRUG_CLASSIFICATION_UPDATE
	PermissionAdminMedicinalControlledDrugClassificationDelete // ADMIN_MEDICINAL_CONTROLLED_DRUG_CLASSIFICATION_DELETE
	PermissionAdminMedicinalControlledDrugClassificationOption // ADMIN_MEDICINAL_CONTROLLED_DRUG_CLASSIFICATION_OPTION

	// admin medicinal form
	PermissionAdminMedicinalFormCreate // ADMIN_MEDICINAL_FORM_CREATE
	PermissionAdminMedicinalFormList   // ADMIN_MEDICINAL_FORM_LIST
	PermissionAdminMedicinalFormUpdate // ADMIN_MEDICINAL_FORM_UPDATE
	PermissionAdminMedicinalFormDelete // ADMIN_MEDICINAL_FORM_DELETE
	PermissionAdminMedicinalFormOption // ADMIN_MEDICINAL_FORM_OPTION

	// admin medicinal instruction
	PermissionAdminMedicinalInstructionCreate           // ADMIN_MEDICINAL_INSTRUCTION_CREATE
	PermissionAdminMedicinalInstructionUpdate           // ADMIN_MEDICINAL_INSTRUCTION_UPDATE
	PermissionAdminMedicinalInstructionList             // ADMIN_MEDICINAL_INSTRUCTION_LIST
	PermissionAdminMedicinalInstructionDelete           // ADMIN_MEDICINAL_INSTRUCTION_DELETE
	PermissionAdminMedicinalInstructionOptionForAmpForm // ADMIN_MEDICINAL_INSTRUCTION_OPTION_FOR_AMP_FORM

	// admin medicinal route
	PermissionAdminMedicinalRouteCreate // ADMIN_MEDICINAL_ROUTE_CREATE
	PermissionAdminMedicinalRouteList   // ADMIN_MEDICINAL_ROUTE_LIST
	PermissionAdminMedicinalRouteUpdate // ADMIN_MEDICINAL_ROUTE_UPDATE
	PermissionAdminMedicinalRouteDelete // ADMIN_MEDICINAL_ROUTE_DELETE
	PermissionAdminMedicinalRouteOption // ADMIN_MEDICINAL_ROUTE_OPTION

	// admin medicine
	PermissionAdminMedicineCreate // ADMIN_MEDICINE_CREATE
	PermissionAdminMedicineList   // ADMIN_MEDICINE_LIST
	PermissionAdminMedicineShow   // ADMIN_MEDICINE_SHOW
	PermissionAdminMedicineUpdate // ADMIN_MEDICINE_UPDATE
	PermissionAdminMedicineDelete // ADMIN_MEDICINE_DELETE

	// admin medicine attention
	PermissionAdminMedicineAttentionCreate // ADMIN_MEDICINE_ATTENTION_CREATE
	PermissionAdminMedicineAttentionDelete // ADMIN_MEDICINE_ATTENTION_DELETE

	// admin medicine category
	PermissionAdminMedicineCategoryCreate // ADMIN_MEDICINE_CATEGORY_CREATE
	PermissionAdminMedicineCategoryDelete // ADMIN_MEDICINE_CATEGORY_DELETE

	// admin medicine instruction
	PermissionAdminMedicineInstructionCreate // ADMIN_MEDICINE_INSTRUCTION_CREATE
	PermissionAdminMedicineInstructionDelete // ADMIN_MEDICINE_INSTRUCTION_DELETE

	// admin medicine unit
	PermissionAdminMedicineUnitCreate // ADMIN_MEDICINE_UNIT_CREATE
	PermissionAdminMedicineUnitUpdate // ADMIN_MEDICINE_UNIT_UPDATE
	PermissionAdminMedicineUnitDelete // ADMIN_MEDICINE_UNIT_DELETE

	// admin nurse
	PermissionAdminNurseCreate // ADMIN_NURSE_CREATE
	PermissionAdminNurseList   // ADMIN_NURSE_LIST
	PermissionAdminNurseShow   // ADMIN_NURSE_SHOW
	PermissionAdminNurseUpdate // ADMIN_NURSE_UPDATE

	// admin nurse qualification
	PermissionAdminNurseQualificationCreate // ADMIN_NURSE_QUALIFICATION_CREATE
	PermissionAdminNurseQualificationUpdate // ADMIN_NURSE_QUALIFICATION_UPDATE
	PermissionAdminNurseQualificationDelete // ADMIN_NURSE_QUALIFICATION_DELETE

	// admin pharmacist
	PermissionAdminPharmacistCreate                    // ADMIN_PHARMACIST_CREATE
	PermissionAdminPharmacistList                      // ADMIN_PHARMACIST_LIST
	PermissionAdminPharmacistShow                      // ADMIN_PHARMACIST_SHOW
	PermissionAdminPharmacistUpdate                    // ADMIN_PHARMACIST_UPDATE
	PermissionAdminPharmacistOptionForCreateClinicForm // ADMIN_PHARMACIST_OPTION_FOR_CREATE_CLINIC_FORM
	PermissionAdminPharmacistOptionForUpdateClinicForm // ADMIN_PHARMACIST_OPTION_FOR_UPDATE_CLINIC_FORM
	PermissionAdminPharmacistQualificationCreate       // ADMIN_PHARMACIST_QUALIFICATION_CREATE
	PermissionAdminPharmacistQualificationUpdate       // ADMIN_PHARMACIST_QUALIFICATION_UPDATE
	PermissionAdminPharmacistQualificationDelete       // ADMIN_PHARMACIST_QUALIFICATION_DELETE

	// admin product
	PermissionAdminProductOption                       // ADMIN_PRODUCT_OPTION
	PermissionAdminProductOptionForExaminationItem     // ADMIN_PRODUCT_OPTION_FOR_EXAMINATION_ITEM
	PermissionAdminProductOptionForRetailProduct       // ADMIN_PRODUCT_OPTION_FOR_RETAIL_PRODUCT
	PermissionAdminProductOptionForTreatmentItem       // ADMIN_PRODUCT_OPTION_FOR_TREATMENT_ITEM
	PermissionAdminProductOptionForProductPromotion    // ADMIN_PRODUCT_OPTION_FOR_PRODUCT_PROMOTION
	PermissionAdminProductOptionForProductBundlingItem // ADMIN_PRODUCT_OPTION_FOR_PRODUCT_BUNDLING_ITEM

	// admin product bundling
	PermissionAdminProductBundlingCreate        // ADMIN_PRODUCT_BUNDLING_CREATE
	PermissionAdminProductBundlingList          // ADMIN_PRODUCT_BUNDLING_LIST
	PermissionAdminProductBundlingShow          // ADMIN_PRODUCT_BUNDLING_SHOW
	PermissionAdminProductBundlingUpdate        // ADMIN_PRODUCT_BUNDLING_UPDATE
	PermissionAdminProductBundlingUpdateEndedAt // ADMIN_PRODUCT_BUNDLING_UPDATE_ENDED_AT
	PermissionAdminProductBundlingDelete        // ADMIN_PRODUCT_BUNDLING_DELETE
	PermissionAdminProductBundlingClinicCreate  // ADMIN_PRODUCT_BUNDLING_CLINIC_CREATE
	PermissionAdminProductBundlingClinicDelete  // ADMIN_PRODUCT_BUNDLING_CLINIC_DELETE
	PermissionAdminProductBundlingItemCreate    // ADMIN_PRODUCT_BUNDLING_ITEM_CREATE
	PermissionAdminProductBundlingItemDelete    // ADMIN_PRODUCT_BUNDLING_ITEM_DELETE

	// admin product promotion
	PermissionAdminProductPromotionCreate        // ADMIN_PRODUCT_PROMOTION_CREATE
	PermissionAdminProductPromotionList          // ADMIN_PRODUCT_PROMOTION_LIST
	PermissionAdminProductPromotionShow          // ADMIN_PRODUCT_PROMOTION_SHOW
	PermissionAdminProductPromotionUpdate        // ADMIN_PRODUCT_PROMOTION_UPDATE
	PermissionAdminProductPromotionUpdateEndedAt // ADMIN_PRODUCT_PROMOTION_UPDATE_ENDED_AT
	PermissionAdminProductPromotionDelete        // ADMIN_PRODUCT_PROMOTION_DELETE
	PermissionAdminProductPromotionClinicCreate  // ADMIN_PRODUCT_PROMOTION_CLINIC_CREATE
	PermissionAdminProductPromotionClinicDelete  // ADMIN_PRODUCT_PROMOTION_CLINIC_DELETE

	// admin product unit
	PermissionAdminProductUnitOption // ADMIN_PRODUCT_UNIT_OPTION

	// admin qualification
	PermissionAdminQualificationTypeCreate // ADMIN_QUALIFICATION_TYPE_CREATE
	PermissionAdminQualificationTypeList   // ADMIN_QUALIFICATION_TYPE_LIST
	PermissionAdminQualificationTypeShow   // ADMIN_QUALIFICATION_TYPE_SHOW
	PermissionAdminQualificationTypeUpdate // ADMIN_QUALIFICATION_TYPE_UPDATE
	PermissionAdminQualificationTypeDelete // ADMIN_QUALIFICATION_TYPE_DELETE
	PermissionAdminQualificationTypeOption // ADMIN_QUALIFICATION_TYPE_OPTION

	// admin retail product
	PermissionAdminRetailProductCreate            // ADMIN_RETAIL_PRODUCT_CREATE
	PermissionAdminRetailProductUpload            // ADMIN_RETAIL_PRODUCT_UPLOAD
	PermissionAdminRetailProductList              // ADMIN_RETAIL_PRODUCT_LIST
	PermissionAdminRetailProductShow              // ADMIN_RETAIL_PRODUCT_SHOW
	PermissionAdminRetailProductUpdate            // ADMIN_RETAIL_PRODUCT_UPDATE
	PermissionAdminRetailProductDelete            // ADMIN_RETAIL_PRODUCT_DELETE
	PermissionAdminRetailProductCreateProductUnit // ADMIN_RETAIL_PRODUCT_CREATE_PRODUCT_UNIT
	PermissionAdminRetailProductUpdateProductUnit // ADMIN_RETAIL_PRODUCT_UPDATE_PRODUCT_UNIT
	PermissionAdminRetailProductDeleteProductUnit // ADMIN_RETAIL_PRODUCT_DELETE_PRODUCT_UNIT

	// admin retail product category
	PermissionAdminRetailProductCategoryCreate                     // ADMIN_RETAIL_PRODUCT_CATEGORY_CREATE
	PermissionAdminRetailProductCategoryList                       // ADMIN_RETAIL_PRODUCT_CATEGORY_LIST
	PermissionAdminRetailProductCategoryShow                       // ADMIN_RETAIL_PRODUCT_CATEGORY_SHOW
	PermissionAdminRetailProductCategoryUpdate                     // ADMIN_RETAIL_PRODUCT_CATEGORY_UPDATE
	PermissionAdminRetailProductCategoryDelete                     // ADMIN_RETAIL_PRODUCT_CATEGORY_DELETE
	PermissionAdminRetailProductCategoryOptionForRetailProductForm // ADMIN_RETAIL_PRODUCT_CATEGORY_OPTION_FOR_RETAIL_PRODUCT_FORM

	// admin role
	PermissionAdminRoleList                // ADMIN_ROLE_LIST
	PermissionAdminRoleShow                // ADMIN_ROLE_SHOW
	PermissionAdminRoleOption              // ADMIN_ROLE_OPTION
	PermissionAdminRoleOptionForUserFilter // ADMIN_ROLE_OPTION_FOR_USER_FILTER
	PermissionAdminRoleOptionForUserForm   // ADMIN_ROLE_OPTION_FOR_USER_FORM

	// admin supplier
	PermissionAdminSupplierCreate                   // ADMIN_SUPPLIER_CREATE
	PermissionAdminSupplierFetch                    // ADMIN_SUPPLIER_FETCH
	PermissionAdminSupplierGet                      // ADMIN_SUPPLIER_GET
	PermissionAdminSupplierUpdate                   // ADMIN_SUPPLIER_UPDATE
	PermissionAdminSupplierDelete                   // ADMIN_SUPPLIER_DELETE
	PermissionAdminSupplierOptionForExaminationForm // ADMIN_SUPPLIER_OPTION_FOR_EXAMINATION_FORM

	// admin teleconsult type
	PermissionAdminTeleconsultTypeCreate // ADMIN_TELECONSULT_TYPE_CREATE
	PermissionAdminTeleconsultTypeList   // ADMIN_TELECONSULT_TYPE_LIST
	PermissionAdminTeleconsultTypeShow   // ADMIN_TELECONSULT_TYPE_SHOW
	PermissionAdminTeleconsultTypeUpdate // ADMIN_TELECONSULT_TYPE_UPDATE
	PermissionAdminTeleconsultTypeDelete // ADMIN_TELECONSULT_TYPE_DELETE

	// admin treatment
	PermissionAdminTreatmentCreate // ADMIN_TREATMENT_CREATE
	PermissionAdminTreatmentList   // ADMIN_TREATMENT_LIST
	PermissionAdminTreatmentShow   // ADMIN_TREATMENT_SHOW
	PermissionAdminTreatmentUpdate // ADMIN_TREATMENT_UPDATE
	PermissionAdminTreatmentDelete // ADMIN_TREATMENT_DELETE
	PermissionAdminTreatmentOption // ADMIN_TREATMENT_OPTION

	// admin treatment category
	PermissionAdminTreatmentCategoryCreate // ADMIN_TREATMENT_CATEGORY_CREATE
	PermissionAdminTreatmentCategoryList   // ADMIN_TREATMENT_CATEGORY_LIST
	PermissionAdminTreatmentCategoryShow   // ADMIN_TREATMENT_CATEGORY_SHOW
	PermissionAdminTreatmentCategoryUpdate // ADMIN_TREATMENT_CATEGORY_UPDATE
	PermissionAdminTreatmentCategoryDelete // ADMIN_TREATMENT_CATEGORY_DELETE
	PermissionAdminTreatmentCategoryOption // ADMIN_TREATMENT_CATEGORY_OPTION

	// admin treatment item
	PermissionAdminTreatmentItemCreate // ADMIN_TREATMENT_ITEM_CREATE
	PermissionAdminTreatmentItemUpdate // ADMIN_TREATMENT_ITEM_UPDATE
	PermissionAdminTreatmentItemDelete // ADMIN_TREATMENT_ITEM_DELETE

	// admin unit
	PermissionAdminUnitCreate               // ADMIN_UNIT_CREATE
	PermissionAdminUnitList                 // ADMIN_UNIT_LIST
	PermissionAdminUnitUpdate               // ADMIN_UNIT_UPDATE
	PermissionAdminUnitDelete               // ADMIN_UNIT_DELETE
	PermissionAdminUnitOptionForProductForm // ADMIN_UNIT_OPTION_FOR_PRODUCT_FORM

	// admin user
	PermissionAdminUserCreate                  // ADMIN_USER_CREATE
	PermissionAdminUserList                    // ADMIN_USER_LIST
	PermissionAdminUserShow                    // ADMIN_USER_SHOW
	PermissionAdminUserUpdate                  // ADMIN_USER_UPDATE
	PermissionAdminUserOptionForDoctorForm     // ADMIN_USER_OPTION_FOR_DOCTOR_FORM
	PermissionAdminUserOptionForNurseForm      // ADMIN_USER_OPTION_FOR_NURSE_FORM
	PermissionAdminUserOptionForPharmacistForm // ADMIN_USER_OPTION_FOR_PHARMACIST_FORM

	// admin vaccine
	PermissionAdminVaccineCreate // ADMIN_VACCINE_CREATE
	PermissionAdminVaccineList   // ADMIN_VACCINE_LIST
	PermissionAdminVaccineShow   // ADMIN_VACCINE_SHOW
	PermissionAdminVaccineUpdate // ADMIN_VACCINE_UPDATE
	PermissionAdminVaccineDelete // ADMIN_VACCINE_DELETE

	// admin vaccine attention
	PermissionAdminVaccineAttentionCreate // ADMIN_VACCINE_ATTENTION_CREATE
	PermissionAdminVaccineAttentionDelete // ADMIN_VACCINE_ATTENTION_DELETE

	// admin vaccine category
	PermissionAdminVaccineCategoryCreate // ADMIN_VACCINE_CATEGORY_CREATE
	PermissionAdminVaccineCategoryDelete // ADMIN_VACCINE_CATEGORY_DELETE

	// admin vaccine instruction
	PermissionAdminVaccineInstructionCreate // ADMIN_VACCINE_INSTRUCTION_CREATE
	PermissionAdminVaccineInstructionDelete // ADMIN_VACCINE_INSTRUCTION_DELETE

	// admin vaccine unit
	PermissionAdminVaccineUnitCreate // ADMIN_VACCINE_UNIT_CREATE
	PermissionAdminVaccineUnitUpdate // ADMIN_VACCINE_UNIT_UPDATE
	PermissionAdminVaccineUnitDelete // ADMIN_VACCINE_UNIT_DELETE

	// admin virtual medicinal product
	PermissionAdminVirtualMedicinalProductCreate // ADMIN_VIRTUAL_MEDICINAL_PRODUCT_CREATE
	PermissionAdminVirtualMedicinalProductList   // ADMIN_VIRTUAL_MEDICINAL_PRODUCT_LIST
	PermissionAdminVirtualMedicinalProductShow   // ADMIN_VIRTUAL_MEDICINAL_PRODUCT_SHOW
	PermissionAdminVirtualMedicinalProductUpdate // ADMIN_VIRTUAL_MEDICINAL_PRODUCT_UPDATE
	PermissionAdminVirtualMedicinalProductDelete // ADMIN_VIRTUAL_MEDICINAL_PRODUCT_DELETE
	PermissionAdminVirtualMedicinalProductOption // ADMIN_VIRTUAL_MEDICINAL_PRODUCT_OPTION

	// admin virtual therapeutic moiety
	PermissionAdminVirtualTherapeuticMoietyList                // ADMIN_VIRTUAL_THERAPEUTIC_MOIETY_LIST
	PermissionAdminVirtualTherapeuticMoietyShow                // ADMIN_VIRTUAL_THERAPEUTIC_MOIETY_SHOW
	PermissionAdminVirtualTherapeuticMoietyCreate              // ADMIN_VIRTUAL_THERAPEUTIC_MOIETY_CREATE
	PermissionAdminVirtualTherapeuticMoietyUpdate              // ADMIN_VIRTUAL_THERAPEUTIC_MOIETY_UPDATE
	PermissionAdminVirtualTherapeuticMoietyDelete              // ADMIN_VIRTUAL_THERAPEUTIC_MOIETY_DELETE
	PermissionAdminVirtualTherapeuticMoietyOption              // ADMIN_VIRTUAL_THERAPEUTIC_MOIETY_OPTION
	PermissionAdminVirtualTherapeuticMoietyOptionForVmpVtmForm // ADMIN_VIRTUAL_THERAPEUTIC_MOIETY_OPTION_FOR_VMP_VTM_FORM

	// admin vmp code
	PermissionAdminVmpCodeCreate // ADMIN_VMP_CODE_CREATE
	PermissionAdminVmpCodeDelete // ADMIN_VMP_CODE_DELETE

	// admin vmp vtm
	PermissionAdminVmpVtmCreate // ADMIN_VMP_VTM_CREATE
	PermissionAdminVmpVtmUpdate // ADMIN_VMP_VTM_UPDATE
	PermissionAdminVmpVtmDelete // ADMIN_VMP_VTM_DELETE

	// admin vtm code
	PermissionAdminVtmCodeCreate // ADMIN_VTM_CODE_CREATE
	PermissionAdminVtmCodeDelete // ADMIN_VTM_CODE_DELETE

	// admin voucher code
	PermissionAdminVoucherCodeCreate // ADMIN_VOUCHER_CODE_CREATE
	PermissionAdminVoucherCodeList   // ADMIN_VOUCHER_CODE_LIST
	PermissionAdminVoucherCodeShow   // ADMIN_VOUCHER_CODE_SHOW
	PermissionAdminVoucherCodeUpdate // ADMIN_VOUCHER_CODE_UPDATE
	PermissionAdminVoucherCodeDelete // ADMIN_VOUCHER_CODE_DELETE

	PermissionAdminVoucherCodeRedeemList // ADMIN_VOUCHER_CODE_REDEEM_LIST
	PermissionAdminVoucherCodeRedeemShow // ADMIN_VOUCHER_CODE_REDEEM_SHOW

	// admin voucher generate rule
	PermissionAdminVoucherGenerateRuleCreate // ADMIN_VOUCHER_GENERATE_RULE_CREATE
	PermissionAdminVoucherGenerateRuleList   // ADMIN_VOUCHER_GENERATE_RULE_LIST
	PermissionAdminVoucherGenerateRuleShow   // ADMIN_VOUCHER_GENERATE_RULE_SHOW
	PermissionAdminVoucherGenerateRuleUpdate // ADMIN_VOUCHER_GENERATE_RULE_UPDATE
	PermissionAdminVoucherGenerateRuleDelete // ADMIN_VOUCHER_GENERATE_RULE_DELETE

	// cashier sales invoice
	PermissionCashierSalesInvoiceList // CASHIER_SALES_INVOICE_LIST
	PermissionCashierSalesInvoiceShow // CASHIER_SALES_INVOICE_SHOW

	// cluster manager clinic
	PermissionClusterManagerClinicGetCashOnHand               // CLUSTER_MANAGER_CLINIC_GET_CASH_ON_HAND
	PermissionClusterManagerClinicOptionForCashierSummaryList // CLUSTER_MANAGER_CLINIC_OPTION_FOR_CASHIER_SUMMARY_LIST
	PermissionClusterManagerClinicOptionForCashDepositForm    // CLUSTER_MANAGER_CLINIC_OPTION_FOR_CASH_DEPOSIT_FORM
	PermissionClusterManagerClinicOptionForCashDepositFilter  // CLUSTER_MANAGER_CLINIC_OPTION_FOR_CASH_DEPOSIT_FILTER

	// cluster manager cashier summary
	PermissionClusterManagerCashierSummaryList             // CLUSTER_MANAGER_CASHIER_SUMMARY_LIST
	PermissionClusterManagerCashierSummaryListSummarized   // CLUSTER_MANAGER_CASHIER_SUMMARY_LIST_SUMMARIZED
	PermissionClusterManagerCashierSummaryShow             // CLUSTER_MANAGER_CASHIER_SUMMARY_SHOW
	PermissionClusterManagerCashierSummaryShowSession      // CLUSTER_MANAGER_CASHIER_SUMMARY_SHOW_SESSION
	PermissionClusterManagerCashierSummaryShowSalesInvoice // CLUSTER_MANAGER_CASHIER_SUMMARY_SHOW_SALES_INVOICE
	PermissionClusterManagerCashierSummaryDownloadCsv      // CLUSTER_MANAGER_CASHIER_SUMMARY_DOWNLOAD_CSV

	// clinic manager sales invoice
	PermissionClinicManagerSalesInvoiceList     // CLINIC_MANAGER_SALES_INVOICE_LIST
	PermissionClinicManagerSalesInvoiceDownload // CLINIC_MANAGER_SALES_INVOICE_DOWNLOAD

	// doctor patient vital check
	PermissionDoctorPatientVitalCheckCreate // DOCTOR_PATIENT_VITAL_CHECK_CREATE
	PermissionDoctorPatientVitalCheckUpdate // DOCTOR_PATIENT_VITAL_CHECK_UPDATE

	// nurse patient vital check
	PermissionNursePatientVitalCheckCreate // NURSE_PATIENT_VITAL_CHECK_CREATE
	PermissionNursePatientVitalCheckUpdate // NURSE_PATIENT_VITAL_CHECK_UPDATE

	// appointment
	PermissionAppointmentCreate                                  // APPOINTMENT_CREATE
	PermissionAppointmentList                                    // APPOINTMENT_LIST
	PermissionAppointmentListForAppointmentFormByClinicOrDoctor  // APPOINTMENT_LIST_FOR_APPOINTMENT_FORM_BY_CLINIC_OR_DOCTOR
	PermissionAppointmentListForAppointmentFormByPatient         // APPOINTMENT_LIST_FOR_APPOINTMENT_FORM_BY_PATIENT
	PermissionAppointmentListForQueueNumberGenerateFormByPatient // APPOINTMENT_LIST_FOR_QUEUE_NUMBER_GENERATE_FORM_BY_PATIENT
	PermissionAppointmentShow                                    // APPOINTMENT_SHOW
	PermissionAppointmentUpdate                                  // APPOINTMENT_UPDATE
	PermissionAppointmentUpdateReferredDoctor                    // APPOINTMENT_UPDATE_REFERRED_DOCTOR
	PermissionAppointmentApprove                                 // APPOINTMENT_APPROVE
	PermissionAppointmentCancel                                  // APPOINTMENT_CANCEL

	// appointment type
	PermissionAppointmentTypeOption // APPOINTMENT_TYPE_OPTION

	// bank
	PermissionBankOptionForCashDepositForm // BANK_OPTION_FOR_CASH_DEPOSIT_FORM

	// cash deposit
	PermissionCashDepositList     // CASH_DEPOSIT_LIST
	PermissionCashDepositShow     // CASH_DEPOSIT_SHOW
	PermissionCashDepositDownload // CASH_DEPOSIT_DOWNLOAD

	// cashier
	PermissionCashierCreate                           // CASHIER_CREATE
	PermissionCashierList                             // CASHIER_LIST
	PermissionCashierShow                             // CASHIER_SHOW
	PermissionCashierUpdate                           // CASHIER_UPDATE
	PermissionCashierDelete                           // CASHIER_DELETE
	PermissionCashierOptionForBeginCashierSessionForm // CASHIER_OPTION_FOR_START_CASHIER_SESSION_FORM

	// cashier session
	PermissionCashierSessionBegin            // CASHIER_SESSION_CREATE
	PermissionCashierSessionOpenDrawer       // CASHIER_SESSION_OPEN_DRAWER
	PermissionCashierSessionList             // CASHIER_SESSION_LIST
	PermissionCashierSessionShow             // CASHIER_SESSION_SHOW
	PermissionCashierSessionShowActive       // CASHIER_SESSION_SHOW_ACTIVE
	PermissionCashierSessionClose            // CASHIER_SESSION_CLOSE
	PermissionCashierSessionUpdateMoneyCount // CASHIER_SESSION_UPDATE_MONEY_COUNT
	PermissionCashierSessionEnd              // CASHIER_SESSION_END

	// cashier summary
	PermissionCashierSummaryCreate                      // CASHIER_SUMMARY_CREATE
	PermissionCashierSummaryList                        // CASHIER_SUMMARY_LIST
	PermissionCashierSummaryListSummarized              // CASHIER_SUMMARY_LIST_SUMMARIZED
	PermissionCashierSummaryShow                        // CASHIER_SUMMARY_SHOW
	PermissionCashierSummaryShowDailyReport             // CASHIER_SUMMARY_SHOW_DAILY_REPORT
	PermissionCashierSummaryShowSession                 // CASHIER_SUMMARY_SHOW_SESSION
	PermissionCashierSummaryShowDailyReportSession      // CASHIER_SUMMARY_SHOW_DAILY_REPORT_SESSION
	PermissionCashierSummaryShowSalesInvoice            // CASHIER_SUMMARY_SHOW_SALES_INVOICE
	PermissionCashierSummaryShowDailyReportSalesInvoice // CASHIER_SUMMARY_SHOW_DAILY_REPORT_SALES_INVOICE
	PermissionCashierSummaryDownloadCsv                 // CASHIER_SUMMARY_DOWNLOAD_CSV

	// clinic
	PermissionClinicOptionForAppointmentForm // CLINIC_OPTION_FOR_APPOINTMENT_FORM

	// clinic ottopay edc
	PermissionClinicOttopayEdcCreate               // CLINIC_OTTOPAY_EDC_CREATE
	PermissionClinicOttopayEdcList                 // CLINIC_OTTOPAY_EDC_LIST
	PermissionClinicOttopayEdcShow                 // CLINIC_OTTOPAY_EDC_SHOW
	PermissionClinicOttopayEdcUpdate               // CLINIC_OTTOPAY_EDC_UPDATE
	PermissionClinicOttopayEdcDelete               // CLINIC_OTTOPAY_EDC_DELETE
	PermissionClinicOttopayEdcOptionForCashierForm // CLINIC_OTTOPAY_EDC_OPTION_FOR_CASHIER_FORM

	// clinic printer
	PermissionClinicPrinterCreate                      // CLINIC_PRINTER_CREATE
	PermissionClinicPrinterDelete                      // CLINIC_PRINTER_DELETE
	PermissionClinicPrinterOptionForCashierForm        // CLINIC_PRINTER_OPTION_FOR_CASHIER_FORM
	PermissionClinicPrinterOptionForQueueCounter       // CLINIC_PRINTER_OPTION_FOR_QUEUE_COUNTER
	PermissionClinicPrinterOptionForPrintMedicineLabel // CLINIC_PRINTER_OPTION_FOR_PRINT_MEDICINAL_LABEL

	// clinic printer gateway
	PermissionClinicPrinterGatewayCreate // CLINIC_PRINTER_GATEWAY_CREATE
	PermissionClinicPrinterGatewayList   // CLINIC_PRINTER_GATEWAY_LIST
	PermissionClinicPrinterGatewayShow   // CLINIC_PRINTER_GATEWAY_SHOW
	PermissionClinicPrinterGatewayUpdate // CLINIC_PRINTER_GATEWAY_UPDATE
	PermissionClinicPrinterGatewayDelete // CLINIC_PRINTER_GATEWAY_DELETE

	// clinic setting
	PermissionClinicSettingList   // CLINIC_SETTING_LIST
	PermissionClinicSettingUpdate // CLINIC_SETTING_UPDATE

	// doctor
	PermissionDoctorList                     // DOCTOR_LIST
	PermissionDoctorShow                     // DOCTOR_SHOW
	PermissionDoctorOptionForAppointmentForm // DOCTOR_OPTION_FOR_APPOINTMENT_FORM
	PermissionDoctorOptionForTeleconsultForm // DOCTOR_OPTION_FOR_TELECONSULT_FORM

	// doctor memo
	PermissionDoctorMemoCreate // DOCTOR_MEMO_CREATE
	PermissionDoctorMemoUpdate // DOCTOR_MEMO_UPDATE

	// doctor referral
	PermissionDoctorReferralCreate // DOCTOR_REFERRAL_CREATE
	PermissionDoctorReferralUpdate // DOCTOR_REFERRAL_UPDATE

	// health examination result
	PermissionHealthExaminationResultCreate // HEALTH_EXAMINATION_RESULT_CREATE
	PermissionHealthExaminationResultUpdate // HEALTH_EXAMINATION_RESULT_UPDATE

	// medical summary
	PermissionMedicalSummaryCreate // MEDICAL_SUMMARY_CREATE

	// examination
	PermissionExaminationList                   // EXAMINATION_LIST
	PermissionExaminationShow                   // EXAMINATION_SHOW
	PermissionExaminationOptionForMedicalRecord // EXAMINATION_OPTION_FOR_MEDICAL_RECORD

	// icd 10
	PermissionIcd10List                                  // ICD_10_LIST
	PermissionIcd10Tree                                  // ICD_10_TREE
	PermissionIcd10OptionForPatientDiseaseHistoryForm    // ICD_10_OPTION_FOR_PATIENT_DISEASE_HISTORY_FORM
	PermissionIcd10OptionForPatientCongenitalDiseaseForm // ICD_10_OPTION_FOR_PATIENT_CONGENITAL_DISEASE_FORM

	// icd 10 trending
	PermissionIcd10TrendingList // ICD_10_TRENDING_LIST

	// injection
	PermissionInjectionShow // INJECTION_SHOW

	// insurance
	PermissionInsuranceOptionForQueueNumber // INSURANCE_OPTION_FOR_QUEUE_NUMBER

	// insurance payor
	PermissionInsurancePayorOptionForQueueNumber // INSURANCE_PAYOR_OPTION_FOR_QUEUE_NUMBER

	// cart
	PermissionCartShow // CART_SHOW

	// cart item
	PermissionCartItemCreateForActiveRetailCart                // CART_ITEM_CREATE_FOR_ACTIVE_RETAIL_CART
	PermissionCartItemCreateForActiveRetailCartByBarcodeOrCode // CART_ITEM_CREATE_FOR_ACTIVE_RETAIL_CART_BY_BARCODE_OR_CODE
	PermissionCartItemQtyUpdateInActiveCart                    // CART_ITEM_QTY_UPDATE_IN_ACTIVE_CART
	PermissionCartItemReprintDocumentFeeDelete                 // CART_ITEM_REPRINT_DOCUMENT_FEE_DELETE
	PermissionCartItemDeleteFromActiveRetailCart               // CART_ITEM_DELETE_FROM_ACTIVE_RETAIL_CART

	// medical certificate
	PermissionMedicalCertificateCreate // MEDICAL_CERTIFICATE_CREATE
	PermissionMedicalCertificateUpdate // MEDICAL_CERTIFICATE_UPDATE

	// medical certificate diagnose
	PermissionMedicalCertificateDiagnoseCreate // MEDICAL_CERTIFICATE_DIAGNOSE_CREATE
	PermissionMedicalCertificateDiagnoseDelete // MEDICAL_CERTIFICATE_DIAGNOSE_DELETE

	// medical record
	PermissionMedicalRecordBegin               // MEDICAL_RECORD_BEGIN
	PermissionMedicalRecordList                // MEDICAL_RECORD_LIST
	PermissionMedicalRecordListForStandByForm  // MEDICAL_RECORD_LIST_FOR_STAND_BY_FORM
	PermissionMedicalRecordShow                // MEDICAL_RECORD_SHOW
	PermissionMedicalRecordUpdate              // MEDICAL_RECORD_UPDATE
	PermissionMedicalRecordUpdateApplyTemplate // MEDICAL_RECORD_UPDATE_APPLY_TEMPLATE
	PermissionMedicalRecordUpdateAnnotation    // MEDICAL_RECORD_UPDATE_ANNOTATION
	PermissionMedicalRecordUpdateAmendment     // MEDICAL_RECORD_UPDATE_AMENDMENT
	PermissionMedicalRecordStandBy             // MEDICAL_RECORD_STAND_BY
	PermissionMedicalRecordResume              // MEDICAL_RECORD_RESUME
	PermissionMedicalRecordEnd                 // MEDICAL_RECORD_END

	// medical record amendment history
	PermissionMedicalRecordAmendmentHistoryList // MEDICAL_RECORD_AMENDMENT_HISTORY_LIST

	// medical record diagnose
	PermissionMedicalRecordDiagnoseCreate // MEDICAL_RECORD_DIAGNOSE_CREATE
	PermissionMedicalRecordDiagnoseUpdate // MEDICAL_RECORD_DIAGNOSE_UPDATE
	PermissionMedicalRecordDiagnoseDelete // MEDICAL_RECORD_DIAGNOSE_DELETE

	// medical record examination
	PermissionMedicalRecordExaminationCreate // MEDICAL_RECORD_EXAMINATION_CREATE
	PermissionMedicalRecordExaminationUpdate // MEDICAL_RECORD_EXAMINATION_UPDATE
	PermissionMedicalRecordExaminationDelete // MEDICAL_RECORD_EXAMINATION_DELETE

	// medical record medical tool
	PermissionMedicalRecordMedicalToolCreate // MEDICAL_RECORD_MEDICAL_TOOL_CREATE
	PermissionMedicalRecordMedicalToolUpdate // MEDICAL_RECORD_MEDICAL_TOOL_UPDATE
	PermissionMedicalRecordMedicalToolDelete // MEDICAL_RECORD_MEDICAL_TOOL_DELETE

	// medical record prescription
	PermissionMedicalRecordPrescriptionCreate // MEDICAL_RECORD_PRESCRIPTION_CREATE
	PermissionMedicalRecordPrescriptionUpdate // MEDICAL_RECORD_PRESCRIPTION_UPDATE
	PermissionMedicalRecordPrescriptionDelete // MEDICAL_RECORD_PRESCRIPTION_DELETE

	// medical record treatment
	PermissionMedicalRecordTreatmentCreate // MEDICAL_RECORD_TREATMENT_CREATE
	PermissionMedicalRecordTreatmentUpdate // MEDICAL_RECORD_TREATMENT_UPDATE
	PermissionMedicalRecordTreatmentDelete // MEDICAL_RECORD_TREATMENT_DELETE

	PermissionMedicalRecordProductBundlingCreate // MEDICAL_RECORD_PRODUCT_BUNDLING_CREATE
	PermissionMedicalRecordProductBundlingUpdate // MEDICAL_RECORD_PRODUCT_BUNDLING_UPDATE
	PermissionMedicalRecordProductBundlingDelete // MEDICAL_RECORD_PRODUCT_BUNDLING_DELETE

	// medical record template
	PermissionMedicalRecordTemplateCreate                     // MEDICAL_RECORD_TEMPLATE_CREATE
	PermissionMedicalRecordTemplateList                       // MEDICAL_RECORD_TEMPLATE_LIST
	PermissionMedicalRecordTemplateShow                       // MEDICAL_RECORD_TEMPLATE_SHOW
	PermissionMedicalRecordTemplateUpdate                     // MEDICAL_RECORD_TEMPLATE_UPDATE
	PermissionMedicalRecordTemplateDelete                     // MEDICAL_RECORD_TEMPLATE_DELETE
	PermissionMedicalRecordTemplateOptionForMedicalRecordForm // MEDICAL_RECORD_TEMPLATE_OPTION_FOR_MEDICAL_RECORD_FORM

	// medical record template diagnose
	PermissionMedicalRecordTemplateDiagnoseCreate // MEDICAL_RECORD_TEMPLATE_DIAGNOSE_CREATE
	PermissionMedicalRecordTemplateDiagnoseUpdate // MEDICAL_RECORD_TEMPLATE_DIAGNOSE_UPDATE
	PermissionMedicalRecordTemplateDiagnoseDelete // MEDICAL_RECORD_TEMPLATE_DIAGNOSE_DELETE

	// medical record template medical tool
	PermissionMedicalRecordTemplateMedicalToolCreate // MEDICAL_RECORD_TEMPLATE_MEDICAL_TOOL_CREATE
	PermissionMedicalRecordTemplateMedicalToolUpdate // MEDICAL_RECORD_TEMPLATE_MEDICAL_TOOL_UPDATE
	PermissionMedicalRecordTemplateMedicalToolDelete // MEDICAL_RECORD_TEMPLATE_MEDICAL_TOOL_DELETE

	// medical record template treatment
	PermissionMedicalRecordTemplateTreatmentCreate // MEDICAL_RECORD_TEMPLATE_TREATMENT_CREATE
	PermissionMedicalRecordTemplateTreatmentUpdate // MEDICAL_RECORD_TEMPLATE_TREATMENT_UPDATE
	PermissionMedicalRecordTemplateTreatmentDelete // MEDICAL_RECORD_TEMPLATE_TREATMENT_DELETE

	// medical record template examination
	PermissionMedicalRecordTemplateExaminationCreate // MEDICAL_RECORD_TEMPLATE_EXAMINATION_CREATE
	PermissionMedicalRecordTemplateExaminationUpdate // MEDICAL_RECORD_TEMPLATE_EXAMINATION_UPDATE
	PermissionMedicalRecordTemplateExaminationDelete // MEDICAL_RECORD_TEMPLATE_EXAMINATION_DELETE

	// medical record template prescription
	PermissionMedicalRecordTemplatePrescriptionCreate // MEDICAL_RECORD_TEMPLATE_PRESCRIPTION_CREATE
	PermissionMedicalRecordTemplatePrescriptionUpdate // MEDICAL_RECORD_TEMPLATE_PRESCRIPTION_UPDATE
	PermissionMedicalRecordTemplatePrescriptionDelete // MEDICAL_RECORD_TEMPLATE_PRESCRIPTION_DELETE

	// medical record waive
	PermissionMedicalRecordWaiveUpdate // MEDICAL_RECORD_WAIVE_UPDATE

	// medical tool
	PermissionMedicalToolOptionForMedicalRecord // MEDICAL_TOOL_OPTION_FOR_MEDICAL_RECORD

	// medicinal dispense
	PermissionMedicinalDispenseDownload               // MEDICINAL_DISPENSE_DOWNLOAD
	PermissionMedicinalDispenseCreate                 // MEDICINAL_DISPENSE_CREATE
	PermissionMedicinalDispenseUpdate                 // MEDICINAL_DISPENSE_UPDATE
	PermissionMedicinalDispenseUpload                 // MEDICINAL_DISPENSE_UPLOAD
	PermissionMedicinalDispenseList                   // MEDICINAL_DISPENSE_LIST
	PermissionMedicinalDispenseShow                   // MEDICINAL_DISPENSE_SHOW
	PermissionMedicinalDispenseAddMixedMedicineFee    // MEDICINAL_DISPENSE_ADD_MIXED_MEDICINE_FEE
	PermissionMedicinalDispenseRemoveMixedMedicineFee // MEDICINAL_DISPENSE_REMOVE_MIXED_MEDICINE_FEE
	PermissionMedicinalDispensePrintBlank             // MEDICINAL_DISPENSE_PRINT_BLANK
	PermissionMedicinalDispenseCalibrate              // MEDICINAL_DISPENSE_CALIBRATE
	PermissionMedicinalDispenseUpdateChecked          // MEDICINAL_DISPENSE_UPDATE_CHECKED
	PermissionMedicinalDispenseUpdateCompleted        // MEDICINAL_DISPENSE_UPDATE_COMPLETED

	// medicinal dispense item
	PermissionMedicinalDispenseItemPrint             // MEDICINAL_DISPENSE_ITEM_PRINT
	PermissionMedicinalDispenseItemCreateExpiredDate // MEDICINAL_DISPENSE_ITEM_CREATE_EXPIRED_DATE
	PermissionMedicinalDispenseItemUpdateExpiredDate // MEDICINAL_DISPENSE_ITEM_UPDATE_EXPIRED_DATE
	PermissionMedicinalDispenseItemDeleteExpiredDate // MEDICINAL_DISPENSE_ITEM_DELETE_EXPIRED_DATE
	PermissionMedicinalDispenseItemUpdate            // MEDICINAL_DISPENSE_ITEM_UPDATE

	// medicinal dispense prescription
	PermissionMedicinalDispensePrescriptionCreate // MEDICINAL_DISPENSE_PRESCRIPTION_CREATE
	PermissionMedicinalDispensePrescriptionUpdate // MEDICINAL_DISPENSE_PRESCRIPTION_UPDATE
	PermissionMedicinalDispensePrescriptionDelete // MEDICINAL_DISPENSE_PRESCRIPTION_DELETE

	// medicine
	PermissionMedicineShow                       // MEDICINE_SHOW
	PermissionMedicineOptionForMedicinalDispense // MEDICINE_OPTION_FOR_MEDICINAL_DISPENSE

	// net promoter score
	PermissionNetPromoterScoreStartByQueueNumber  // NET_PROMOTER_SCORE_START_BY_QUEUE_NUMBER
	PermissionNetPromoterScoreStartBySalesInvoice // NET_PROMOTER_SCORE_START_BY_SALES_INVOICE
	PermissionNetPromoterScoreGetActive           // NET_PROMOTER_SCORE_GET_ACTIVE
	PermissionNetPromoterScoreEnd                 // NET_PROMOTER_SCORE_END
	PermissionNetPromoterScoreItemUpdate          // NET_PROMOTER_SCORE_ITEM_UPDATE

	// nric card ocr
	PermissionNricCardOcrExtract // NRIC_CARD_OCR_EXTRACT
	PermissionNricCardOcrUpload  // NRIC_CARD_OCR_UPLOAD

	// nurse
	PermissionNurseList // NURSE_LIST
	PermissionNurseShow // NURSE_SHOW

	// patient
	PermissionPatientCreate                           // PATIENT_CREATE
	PermissionPatientShow                             // PATIENT_SHOW
	PermissionPatientShowMedicalSummary               // PATIENT_SHOW_MEDICAL_SUMMARY
	PermissionPatientShowMedicalNotes                 // PATIENT_SHOW_MEDICAL_NOTES
	PermissionPatientUpdate                           // PATIENT_UPDATE
	PermissionPatientUpdateMedicalSummary             // PATIENT_UPDATE_MEDICAL_SUMMARY
	PermissionPatientUpdateMedicalNotes               // PATIENT_UPDATE_MEDICAL_NOTES
	PermissionPatientOption                           // PATIENT_OPTION
	PermissionPatientOptionForAppointmentForm         // PATIENT_OPTION_FOR_APPOINTMENT_FORM
	PermissionPatientOptionForPatientDoctorForm       // PATIENT_OPTION_FOR_PATIENT_DOCTOR_FORM
	PermissionPatientOptionForPatientRelativeForm     // PATIENT_OPTION_FOR_PATIENT_RELATIVE_FORM
	PermissionPatientOptionForPatientVisitFilter      // PATIENT_OPTION_FOR_PATIENT_VISIT_FILTER
	PermissionPatientOptionForPatientVisitForm        // PATIENT_OPTION_FOR_PATIENT_VISIT_FORM
	PermissionPatientOptionForPatientVitalCheckFilter // PATIENT_OPTION_FOR_PATIENT_VITAL_CHECK_FILTER
	PermissionPatientOptionForPatientVitalCheckForm   // PATIENT_OPTION_FOR_PATIENT_VITAL_CHECK_FORM
	PermissionPatientOptionForQueueNumberGenerateForm // PATIENT_OPTION_FOR_QUEUE_NUMBER_GENERATE_FORM
	PermissionPatientOptionForTaskForm                // PATIENT_OPTION_FOR_TASK_FORM
	PermissionPatientOptionForTeleconsult             // PATIENT_OPTION_FOR_TELECONSULT
	PermissionPatientOptionForVaccination             // PATIENT_OPTION_FOR_VACCINATION

	// patient allergy
	PermissionPatientAllergyCreate // PATIENT_ALLERGY_CREATE
	PermissionPatientAllergyList   // PATIENT_ALLERGY_LIST
	PermissionPatientAllergyShow   // PATIENT_ALLERGY_SHOW
	PermissionPatientAllergyUpdate // PATIENT_ALLERGY_UPDATE
	PermissionPatientAllergyDelete // PATIENT_ALLERGY_DELETE

	// patient congenital disease
	PermissionPatientCongenitalDiseaseCreate // PATIENT_CONGENITAL_DISEASE_CREATE
	PermissionPatientCongenitalDiseaseList   // PATIENT_CONGENITAL_DISEASE_LIST
	PermissionPatientCongenitalDiseaseShow   // PATIENT_CONGENITAL_DISEASE_SHOW
	PermissionPatientCongenitalDiseaseUpdate // PATIENT_CONGENITAL_DISEASE_UPDATE
	PermissionPatientCongenitalDiseaseDelete // PATIENT_CONGENITAL_DISEASE_DELETE

	// patient disease history
	PermissionPatientDiseaseHistoryCreate // PATIENT_DISEASE_HISTORY_CREATE
	PermissionPatientDiseaseHistoryList   // PATIENT_DISEASE_HISTORY_LIST
	PermissionPatientDiseaseHistoryShow   // PATIENT_DISEASE_HISTORY_SHOW
	PermissionPatientDiseaseHistoryUpdate // PATIENT_DISEASE_HISTORY_UPDATE
	PermissionPatientDiseaseHistoryDelete // PATIENT_DISEASE_HISTORY_DELETE

	// patient doctor
	PermissionPatientDoctorCreate // PATIENT_DOCTOR_CREATE
	PermissionPatientDoctorDelete // PATIENT_DOCTOR_DELETE

	// patient document
	PermissionPatientDocumentList     // PATIENT_DOCUMENT_LIST
	PermissionPatientDocumentUpload   // PATIENT_DOCUMENT_UPLOAD
	PermissionPatientDocumentComplete // PATIENT_DOCUMENT_COMPLETE
	PermissionPatientDocumentDownload // PATIENT_DOCUMENT_DOWNLOAD
	PermissionPatientDocumentDelete   // PATIENT_DOCUMENT_DELETE

	// patient interview
	PermissionPatientInterviewCreate // PATIENT_INTERVIEW_CREATE
	PermissionPatientInterviewList   // PATIENT_INTERVIEW_LIST
	PermissionPatientInterviewShow   // PATIENT_INTERVIEW_SHOW
	PermissionPatientInterviewUpdate // PATIENT_INTERVIEW_UPDATE
	PermissionPatientInterviewDelete // PATIENT_INTERVIEW_DELETE

	// patient medicine allergy
	PermissionPatientMedicineAllergyCreate // PATIENT_MEDICINE_ALLERGY_CREATE
	PermissionPatientMedicineAllergyList   // PATIENT_MEDICINE_ALLERGY_LIST
	PermissionPatientMedicineAllergyShow   // PATIENT_MEDICINE_ALLERGY_SHOW
	PermissionPatientMedicineAllergyUpdate // PATIENT_MEDICINE_ALLERGY_UPDATE
	PermissionPatientMedicineAllergyDelete // PATIENT_MEDICINE_ALLERGY_DELETE

	// patient medicine history
	PermissionPatientMedicineHistoryCreate // PATIENT_MEDICINE_HISTORY_CREATE
	PermissionPatientMedicineHistoryList   // PATIENT_MEDICINE_HISTORY_LIST
	PermissionPatientMedicineHistoryShow   // PATIENT_MEDICINE_HISTORY_SHOW
	PermissionPatientMedicineHistoryUpdate // PATIENT_MEDICINE_HISTORY_UPDATE
	PermissionPatientMedicineHistoryDelete // PATIENT_MEDICINE_HISTORY_DELETE

	// patient relative
	PermissionPatientRelativeCreate                   // PATIENT_RELATIVE_CREATE
	PermissionPatientRelativeList                     // PATIENT_RELATIVE_LIST
	PermissionPatientRelativeUpdate                   // PATIENT_RELATIVE_UPDATE
	PermissionPatientRelativeUpdateSetDefault         // PATIENT_RELATIVE_UPDATE_SET_DEFAULT
	PermissionPatientRelativeDelete                   // PATIENT_RELATIVE_DELETE
	PermissionPatientRelativeOptionForQueueNumberForm // PATIENT_RELATIVE_OPTION_FOR_QUEUE_NUMBER_FORM

	// patient surgery
	PermissionPatientSurgeryCreate // PATIENT_SURGERY_CREATE
	PermissionPatientSurgeryList   // PATIENT_SURGERY_LIST
	PermissionPatientSurgeryShow   // PATIENT_SURGERY_SHOW
	PermissionPatientSurgeryUpdate // PATIENT_SURGERY_UPDATE
	PermissionPatientSurgeryDelete // PATIENT_SURGERY_DELETE

	// patient test result
	PermissionPatientTestResultCreate           // PATIENT_TEST_RESULT_CREATE
	PermissionPatientTestResultUpload           // PATIENT_TEST_RESULT_UPLOAD
	PermissionPatientTestResultList             // PATIENT_TEST_RESULT_LIST
	PermissionPatientTestResultShow             // PATIENT_TEST_RESULT_SHOW
	PermissionPatientTestResultUpdate           // PATIENT_TEST_RESULT_UPDATE
	PermissionPatientTestResultUpdateReadStatus // PATIENT_TEST_RESULT_UPDATE_READ_STATUS

	// patient visit
	PermissionPatientVisitList                                  // PATIENT_VISIT_LIST
	PermissionPatientVisitSummarizeHistory                      // PATIENT_VISIT_SUMMARIZE_HISTORY
	PermissionPatientVisitShow                                  // PATIENT_VISIT_SHOW
	PermissionPatientVisitInsuranceCardFileShowForQueueNumber   // PATIENT_VISIT_INSURANCE_CARD_FILE_SHOW_FOR_QUEUE_NUMBER
	PermissionPatientVisitInsuranceCardFileShowForMedicalRecord // PATIENT_VISIT_INSURANCE_CARD_FILE_SHOW_FOR_MEDICAL_RECORD
	PermissionPatientVisitUpdate                                // PATIENT_VISIT_UPDATE
	PermissionPatientVisitOptionForMedicalRecord                // PATIENT_VISIT_OPTION_FOR_MEDICAL_RECORD

	// patient vital check
	PermissionPatientVitalCheckList                 // PATIENT_VITAL_CHECK_LIST
	PermissionPatientVitalCheckListForMedicalRecord // PATIENT_VITAL_CHECK_LIST_FOR_MEDICAL_RECORD
	PermissionPatientVitalCheckShow                 // PATIENT_VITAL_CHECK_SHOW
	PermissionPatientVitalCheckShowLatest           // PATIENT_VITAL_CHECK_SHOW_LATEST
	PermissionPatientVitalCheckVisualize            // PATIENT_VITAL_CHECK_VISUALIZE
	PermissionPatientVitalCheckUpdateIsNursing      // PATIENT_VITAL_CHECK_UPDATE_IS_NURSING
	PermissionPatientVitalCheckUpdateIsPregnant     // PATIENT_VITAL_CHECK_UPDATE_IS_PREGNANT
	PermissionPatientVitalCheckUpdateIsSmoking      // PATIENT_VITAL_CHECK_UPDATE_IS_SMOKING

	// pharmacist
	PermissionPharmacistList // PHARMACIST_LIST
	PermissionPharmacistShow // PHARMACIST_SHOW

	// product
	PermissionProductListForRetailCart                      // PRODUCT_LIST_FOR_RETAIL_CART
	PermissionProductPriceTagDownload                       // PRODUCT_PRICE_TAG_DOWNLOAD
	PermissionProductOptionForMedicalRecord                 // PRODUCT_OPTION_FOR_MEDICAL_RECORD
	PermissionProductOptionForMedicinalDispensePrescription // PRODUCT_OPTION_FOR_MEDICINAL_DISPENSE_PRESCRIPTION
	PermissionProductOptionForPriceTagDownload              // PRODUCT_OPTION_FOR_PRICE_TAG_DOWNLOAD
	PermissionProductOptionForProductAdjustmentForm         // PRODUCT_OPTION_FOR_PRODUCT_ADJUSTMENT_FORM
	PermissionProductOptionForProductReceiveItem            // PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_ITEM
	PermissionProductOptionForProductTransferItemForm       // PRODUCT_OPTION_FOR_PRODUCT_TRANSFER_ITEM_FORM
	PermissionProductOptionExpiredDate                      // PRODUCT_OPTION_EXPIRED_DATE

	// product adjustment
	PermissionProductAdjustmentCreate // PRODUCT_ADJUSTMENT_CREATE
	PermissionProductAdjustmentList   // PRODUCT_ADJUSTMENT_LIST
	PermissionProductAdjustmentShow   // PRODUCT_ADJUSTMENT_SHOW
	PermissionProductAdjustmentUpdate // PRODUCT_ADJUSTMENT_UPDATE
	PermissionProductAdjustmentDelete // PRODUCT_ADJUSTMENT_DELETE

	// product adjustment item
	PermissionProductAdjustmentItemCreate // PRODUCT_ADJUSTMENT_ITEM_CREATE
	PermissionProductAdjustmentItemDelete // PRODUCT_ADJUSTMENT_ITEM_DELETE

	// product bundling
	PermissionProductBundlingShow                   // PRODUCT_BUNDLING_SHOW
	PermissionProductBundlingOptionForMedicalRecord // PRODUCT_BUNDLING_OPTION_FOR_MEDICAL_RECORD

	// product inventory
	PermissionProductInventoryList // PRODUCT_INVENTORY_LIST

	// product promotion
	PermissionProductPromotionOptionForPriceTagDownload // PRODUCT_PROMOTION_OPTION_FOR_PRICE_TAG_DOWNLOAD

	// product receive
	PermissionProductReceiveCreate // PRODUCT_RECEIVE_CREATE
	PermissionProductReceiveList   // PRODUCT_RECEIVE_LIST
	PermissionProductReceiveShow   // PRODUCT_RECEIVE_SHOW
	PermissionProductReceiveUpdate // PRODUCT_RECEIVE_UPDATE
	PermissionProductReceiveDelete // PRODUCT_RECEIVE_DELETE

	// product receive item
	PermissionProductReceiveItemCreate // PRODUCT_RECEIVE_ITEM_CREATE
	PermissionProductReceiveItemDelete // PRODUCT_RECEIVE_ITEM_DELETE

	// product unit
	PermissionProductUnitOption                             // PRODUCT_UNIT_OPTION
	PermissionProductUnitOptionForProductAdjustmentItemForm // PRODUCT_UNIT_OPTION_FOR_PRODUCT_ADJUSTMENT_ITEM_FORM
	PermissionProductUnitOptionForProductTransferItemForm   // PRODUCT_UNIT_OPTION_FOR_PRODUCT_TRANSFER_ITEM_FORM
	PermissionProductUnitOptionForPriceTagDownload          // PRODUCT_UNIT_OPTION_FOR_PRICE_TAG_DOWNLOAD

	// product transfer
	PermissionProductTransferCreate     // PRODUCT_TRANSFER_CREATE
	PermissionProductTransferList       // PRODUCT_TRANSFER_LIST
	PermissionProductTransferShow       // PRODUCT_TRANSFER_SHOW
	PermissionProductTransferUpdate     // PRODUCT_TRANSFER_UPDATE
	PermissionProductTransferDelete     // PRODUCT_TRANSFER_DELETE
	PermissionProductTransferCreateItem // PRODUCT_TRANSFER_CREATE_ITEM
	PermissionProductTransferDeleteItem // PRODUCT_TRANSFER_DELETE_ITEM

	// queue
	PermissionQueueCreate                           // QUEUE_CREATE
	PermissionQueueList                             // QUEUE_LIST
	PermissionQueueShow                             // QUEUE_SHOW
	PermissionQueueShowProcess                      // QUEUE_SHOW_PROCESS
	PermissionQueueUpdate                           // QUEUE_UPDATE
	PermissionQueueDelete                           // QUEUE_DELETE
	PermissionQueueOption                           // QUEUE_OPTION
	PermissionQueueOptionForQueueCounterCheckInForm // QUEUE_OPTION_FOR_QUEUE_COUNTER_CHECK_IN_FORM
	PermissionQueueOptionForQueueDisplay            // QUEUE_OPTION_FOR_QUEUE_DISPLAY

	// queue counter
	PermissionQueueCounterCreate                           // QUEUE_COUNTER_CREATE
	PermissionQueueCounterUpdate                           // QUEUE_COUNTER_UPDATE
	PermissionQueueCounterCheckIn                          // QUEUE_COUNTER_CHECK_IN
	PermissionQueueCounterCheckOut                         // QUEUE_COUNTER_CHECK_OUT
	PermissionQueueCounterOptionForCheckInForm             // QUEUE_COUNTER_OPTION_FOR_CHECK_IN_FORM
	PermissionQueueCounterOptionForQueueNumberCompleteForm // QUEUE_COUNTER_OPTION_FOR_QUEUE_NUMBER_COMPLETE_FORM
	PermissionQueueCounterOptionForQueueNumberGenerateForm // QUEUE_COUNTER_OPTION_FOR_QUEUE_NUMBER_GENERATE_FORM

	// queue display
	PermissionQueueDisplayCreate // QUEUE_DISPLAY_CREATE
	PermissionQueueDisplayList   // QUEUE_DISPLAY_LIST
	PermissionQueueDisplayShow   // QUEUE_DISPLAY_SHOW
	PermissionQueueDisplayUpdate // QUEUE_DISPLAY_UPDATE
	PermissionQueueDisplayDelete // QUEUE_DISPLAY_DELETE

	// queue display banner
	PermissionQueueDisplayBannerCreate // QUEUE_DISPLAY_BANNER_CREATE
	PermissionQueueDisplayBannerUpload // QUEUE_DISPLAY_BANNER_UPLOAD
	PermissionQueueDisplayBannerMove   // QUEUE_DISPLAY_BANNER_MOVE
	PermissionQueueDisplayBannerDelete // QUEUE_DISPLAY_BANNER_DELETE

	// queue display queue
	PermissionQueueDisplayQueueCreate // QUEUE_DISPLAY_QUEUE_CREATE
	PermissionQueueDisplayQueueDelete // QUEUE_DISPLAY_QUEUE_DELETE

	// queue display running text
	PermissionQueueDisplayRunningTextCreate // QUEUE_DISPLAY_RUNNING_TEXT_CREATE
	PermissionQueueDisplayRunningTextMove   // QUEUE_DISPLAY_RUNNING_TEXT_MOVE
	PermissionQueueDisplayRunningTextDelete // QUEUE_DISPLAY_RUNNING_TEXT_DELETE

	// queue number
	PermissionQueueNumberGenerate            // QUEUE_NUMBER_GENERATE
	PermissionQueueNumberUploadInsuranceCard // QUEUE_NUMBER_UPLOAD_INSURANCE_CARD
	PermissionQueueNumberGet                 // QUEUE_NUMBER_GET
	PermissionQueueNumberCurrent             // QUEUE_NUMBER_CURRENT
	PermissionQueueNumberReprint             // QUEUE_NUMBER_REPRINT
	PermissionQueueNumberUpdate              // QUEUE_NUMBER_UPDATE

	// queue number action
	PermissionQueueNumberCall       // QUEUE_NUMBER_CALL
	PermissionQueueNumberCallRepeat // QUEUE_NUMBER_CALL_REPEAT
	PermissionQueueNumberProcess    // QUEUE_NUMBER_PROCESS
	PermissionQueueNumberComplete   // QUEUE_NUMBER_COMPLETE
	PermissionQueueNumberMove       // QUEUE_NUMBER_MOVE
	PermissionQueueNumberSkip       // QUEUE_NUMBER_SKIP
	PermissionQueueNumberStandBy    // QUEUE_NUMBER_STAND_BY

	// queue session
	PermissionQueueSessionBegin                     // QUEUE_SESSION_BEGIN
	PermissionQueueSessionEnd                       // QUEUE_SESSION_END
	PermissionQueueSessionGetActive                 // QUEUE_SESSION_GET_ACTIVE
	PermissionQueueSessionCheckForActiveQueueNumber // QUEUE_SESSION_CHECK_FOR_ACTIVE_QUEUE_NUMBER

	// report daily report
	PermissionReportDailyTransactionList     // REPORT_DAILY_TRANSACTION_LIST
	PermissionReportDailyTransactionDownload // REPORT_DAILY_TRANSACTION_DOWNLOAD

	// report finance inventory
	PermissionReportFinanceInventoryCreate   // REPORT_FINANCE_INVENTORY_CREATE
	PermissionReportFinanceInventoryList     // REPORT_FINANCE_INVENTORY_LIST
	PermissionReportFinanceInventoryDownload // REPORT_FINANCE_INVENTORY_DOWNLOAD

	// report inventory stock
	PermissionReportInventoryStockCreate   // REPORT_INVENTORY_STOCK_CREATE
	PermissionReportInventoryStockList     // REPORT_INVENTORY_STOCK_LIST
	PermissionReportInventoryStockDownload // REPORT_INVENTORY_STOCK_DOWNLOAD

	// report monthly revenue
	PermissionReportMonthlyRevenueList     // REPORT_MONTHLY_REVENUE_LIST
	PermissionReportMonthlyRevenueDownload // REPORT_MONTHLY_REVENUE_DOWNLOAD

	// report month to date sales
	PermissionReportMonthToDateSalesList     // REPORT_MONTH_TO_DATE_SALES_LIST
	PermissionReportMonthToDateSalesDownload // REPORT_MONTH_TO_DATE_SALES_DOWNLOAD

	// report net promoter score
	PermissionReportNetPromoterScoreCreate   // REPORT_NET_PROMOTER_SCORE_CREATE
	PermissionReportNetPromoterScoreList     // REPORT_NET_PROMOTER_SCORE_LIST
	PermissionReportNetPromoterScoreDownload // REPORT_NET_PROMOTER_SCORE_DOWNLOAD

	// report service and performance
	PermissionReportServiceAndPerformanceCreate   // REPORT_SERVICE_AND_PERFORMANCE_CREATE
	PermissionReportServiceAndPerformanceList     // REPORT_SERVICE_AND_PERFORMANCE_LIST
	PermissionReportServiceAndPerformanceDownload // REPORT_SERVICE_AND_PERFORMANCE_DOWNLOAD

	// report waive fee
	PermissionReportWaiveFeeDownload // REPORT_WAIVE_FEE_DOWNLOAD

	// retail cart
	PermissionRetailCartGetActive                // RETAIL_CART_GET_ACTIVE
	PermissionRetailCartShow                     // RETAIL_CART_SHOW
	PermissionRetailCartHoldActive               // RETAIL_CART_HOLD_ACTIVE
	PermissionRetailCartRestore                  // RETAIL_CART_RESTORE
	PermissionRetailCartDeleteActive             // RETAIL_CART_DELETE_ACTIVE
	PermissionRetailCartDelete                   // RETAIL_CART_DELETE
	PermissionRetailCartOptionForRestoreOrDelete // RETAIL_CART_OPTION_FOR_RESTORE_OR_DELETE

	// role
	PermissionRoleOptionForQueueForm  // ROLE_OPTION_FOR_QUEUE_FORM
	PermissionRoleOptionForUserFilter // ROLE_OPTION_FOR_USER_FILTER

	// sales invoice
	PermissionSalesInvoiceGenerate          // SALES_INVOICE_GENERATE
	PermissionSalesInvoiceDownload          // SALES_INVOICE_DOWNLOAD
	PermissionSalesInvoiceVoid              // SALES_INVOICE_VOID
	PermissionSalesInvoiceCancel            // SALES_INVOICE_CANCEL
	PermissionSalesInvoiceApplyVoucherCode  // SALES_INVOICE_APPLY_VOUCHER_CODE
	PermissionSalesInvoiceRemoveVoucherCode // SALES_INVOICE_REMOVE_VOUCHER_CODE
	PermissionSalesInvoiceCheckVoucherCode  // SALES_INVOICE_CHECK_VOUCHER_CODE

	// sales payment
	PermissionSalesPaymentCreate      // SALES_PAYMENT_CREATE
	PermissionSalesPaymentShow        // SALES_PAYMENT_SHOW
	PermissionSalesPaymentPrintRetail // SALES_PAYMENT_PRINT_RETAIL
	PermissionSalesPaymentCancel      // SALES_PAYMENT_CANCEL

	// sales payment method
	PermissionSalesPaymentMethodOption // SALES_PAYMENT_METHOD_OPTION

	// supplier
	PermissionSupplierOptionForProductReceiveForm // SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM

	// task
	PermissionTaskCreate         // TASK_CREATE
	PermissionTaskList           // TASK_LIST
	PermissionTaskShow           // TASK_SHOW
	PermissionTaskUpdate         // TASK_UPDATE
	PermissionTaskUpdatePriority // TASK_UPDATE_PRIORITY
	PermissionTaskUpdateStatus   // TASK_UPDATE_STATUS
	PermissionTaskDelete         // TASK_DELETE

	// task attachment
	PermissionTaskAttachmentUpload // TASK_ATTACHMENT_UPLOAD
	PermissionTaskAttachmentShow   // TASK_ATTACHMENT_SHOW
	PermissionTaskAttachmentDelete // TASK_ATTACHMENT_DELETE

	// task comment
	PermissionTaskCommentCreate // TASK_COMMENT_CREATE
	PermissionTaskCommentList   // TASK_COMMENT_LIST
	PermissionTaskCommentUpdate // TASK_COMMENT_UPDATE
	PermissionTaskCommentDelete // TASK_COMMENT_DELETE

	// teleconsult
	PermissionTeleconsultCreate        // TELECONSULT_CREATE
	PermissionTeleconsultList          // TELECONSULT_LIST
	PermissionTeleconsultShow          // TELECONSULT_SHOW
	PermissionTeleconsultUpdate        // TELECONSULT_UPDATE
	PermissionTeleconsultCancelConsult // TELECONSULT_CANCEL_CONSULT

	// teleconsult type
	PermissionTeleconsultTypeOption // TELECONSULT_TYPE_OPTION

	// treatment
	PermissionTreatmentList                   // TREATMENT_LIST
	PermissionTreatmentShow                   // TREATMENT_SHOW
	PermissionTreatmentOptionForMedicalRecord // TREATMENT_OPTION_FOR_MEDICAL_RECORD

	// unit
	PermissionUnitOptionForProductTransferForm // UNIT_OPTION_FOR_PRODUCT_TRANSFER_FORM

	// user
	PermissionUserList                             // USER_LIST
	PermissionUserShow                             // USER_SHOW
	PermissionUserOptionForTaskForm                // USER_OPTION_FOR_TASK_FORM
	PermissionUserOptionForVoidSalesInvoiceForm    // USER_OPTION_FOR_VOID_SALES_INVOICE_FORM
	PermissionUserOptionForBeginCashierSessionForm // USER_OPTION_FOR_START_CASHIER_SESSION_FORM

	// vaccination
	PermissionVaccinationCreate // VACCINATION_CREATE
	PermissionVaccinationList   // VACCINATION_LIST
	PermissionVaccinationShow   // VACCINATION_SHOW
	PermissionVaccinationUpdate // VACCINATION_UPDATE
	PermissionVaccinationDelete // VACCINATION_DELETE

	// virtual therapeutic moiety
	PermissionVirtualTherapeuticMoietyOptionForPatientMedicineAllergy // VIRTUAL_THERAPEUTIC_MOIETY_OPTION_FOR_PATIENT_MEDICINE_ALLERGY
	PermissionVirtualTherapeuticMoietyOptionForPatientMedicineHistory // VIRTUAL_THERAPEUTIC_MOIETY_OPTION_FOR_PATIENT_MEDICINE_HISTORY
)

var permissionTypeByPermission = map[Permission]PermissionType{
	// admin appointment type
	PermissionAdminAppointmentTypeCreate: PermissionTypeAdmin,
	PermissionAdminAppointmentTypeList:   PermissionTypeAdmin,
	PermissionAdminAppointmentTypeShow:   PermissionTypeAdmin,
	PermissionAdminAppointmentTypeUpdate: PermissionTypeAdmin,
	PermissionAdminAppointmentTypeDelete: PermissionTypeAdmin,

	PermissionAdminBankCreate: PermissionTypeAdmin,
	PermissionAdminBankList:   PermissionTypeAdmin,
	PermissionAdminBankShow:   PermissionTypeAdmin,
	PermissionAdminBankUpdate: PermissionTypeAdmin,
	PermissionAdminBankDelete: PermissionTypeAdmin,

	// admin cash deposit
	PermissionAdminCashDepositCreate:   PermissionTypeAdmin,
	PermissionAdminCashDepositUpload:   PermissionTypeAdmin,
	PermissionAdminCashDepositList:     PermissionTypeAdmin,
	PermissionAdminCashDepositShow:     PermissionTypeAdmin,
	PermissionAdminCashDepositDownload: PermissionTypeAdmin,

	// admin clinic
	PermissionAdminClinicCreate:                        PermissionTypeAdmin,
	PermissionAdminClinicList:                          PermissionTypeAdmin,
	PermissionAdminClinicShow:                          PermissionTypeAdmin,
	PermissionAdminClinicUpdate:                        PermissionTypeAdmin,
	PermissionAdminClinicOption:                        PermissionTypeAdmin,
	PermissionAdminClinicOptionForAssignUserForm:       PermissionTypeAdmin,
	PermissionAdminClinicOptionForDoctorFilter:         PermissionTypeAdmin,
	PermissionAdminClinicOptionForNurseFilter:          PermissionTypeAdmin,
	PermissionAdminClinicOptionForPharmacistFilter:     PermissionTypeAdmin,
	PermissionAdminClinicOptionForParentClinic:         PermissionTypeAdmin,
	PermissionAdminClinicOptionForUserFilter:           PermissionTypeAdmin,
	PermissionAdminClinicOptionForBankForm:             PermissionTypeAdmin,
	PermissionAdminClinicOptionForProductPromotionForm: PermissionTypeAdmin,
	PermissionAdminClinicOptionForProductBundlingForm:  PermissionTypeAdmin,

	// admin clinic type
	PermissionAdminClinicTypeOptionForClinicFilter: PermissionTypeAdmin,
	PermissionAdminClinicTypeOptionForClinicForm:   PermissionTypeAdmin,

	// admin clinic user
	PermissionAdminClinicUserCreate: PermissionTypeAdmin,
	PermissionAdminClinicUserDelete: PermissionTypeAdmin,

	// admin clinic whitelist ip
	PermissionAdminClinicWhitelistIpCreate: PermissionTypeAdmin,
	PermissionAdminClinicWhitelistIpList:   PermissionTypeAdmin,
	PermissionAdminClinicWhitelistIpDelete: PermissionTypeAdmin,

	// admin company
	PermissionAdminCompanyCreate:                                   PermissionTypeAdmin,
	PermissionAdminCompanyList:                                     PermissionTypeAdmin,
	PermissionAdminCompanyShow:                                     PermissionTypeAdmin,
	PermissionAdminCompanyUpdate:                                   PermissionTypeAdmin,
	PermissionAdminCompanyOption:                                   PermissionTypeAdmin,
	PermissionAdminCompanyOptionForClinicFilter:                    PermissionTypeAdmin,
	PermissionAdminCompanyOptionForReportNetPromoterScoreForm:      PermissionTypeAdmin,
	PermissionAdminCompanyOptionForReportServiceAndPerformanceForm: PermissionTypeAdmin,
	PermissionAdminCompanyOptionForClinicForm:                      PermissionTypeAdmin,
	PermissionAdminCompanyOptionForUserForm:                        PermissionTypeAdmin,
	PermissionAdminCompanyOptionForBankForm:                        PermissionTypeAdmin,

	// admin doctor
	PermissionAdminDoctorCreate:       PermissionTypeAdmin,
	PermissionAdminDoctorUploadAvatar: PermissionTypeAdmin,
	PermissionAdminDoctorList:         PermissionTypeAdmin,
	PermissionAdminDoctorShow:         PermissionTypeAdmin,
	PermissionAdminDoctorUpdate:       PermissionTypeAdmin,
	PermissionAdminDoctorOption:       PermissionTypeAdmin,

	// admin doctor qualification
	PermissionAdminDoctorQualificationCreate: PermissionTypeAdmin,
	PermissionAdminDoctorQualificationUpdate: PermissionTypeAdmin,
	PermissionAdminDoctorQualificationDelete: PermissionTypeAdmin,

	// admin examination
	PermissionAdminExaminationCreate: PermissionTypeAdmin,
	PermissionAdminExaminationList:   PermissionTypeAdmin,
	PermissionAdminExaminationShow:   PermissionTypeAdmin,
	PermissionAdminExaminationUpdate: PermissionTypeAdmin,
	PermissionAdminExaminationDelete: PermissionTypeAdmin,
	PermissionAdminExaminationOption: PermissionTypeAdmin,

	// admin examination category
	PermissionAdminExaminationCategoryCreate: PermissionTypeAdmin,
	PermissionAdminExaminationCategoryList:   PermissionTypeAdmin,
	PermissionAdminExaminationCategoryShow:   PermissionTypeAdmin,
	PermissionAdminExaminationCategoryUpdate: PermissionTypeAdmin,
	PermissionAdminExaminationCategoryDelete: PermissionTypeAdmin,
	PermissionAdminExaminationCategoryOption: PermissionTypeAdmin,

	// admin examination item
	PermissionAdminExaminationItemCreate: PermissionTypeAdmin,
	PermissionAdminExaminationItemUpdate: PermissionTypeAdmin,
	PermissionAdminExaminationItemDelete: PermissionTypeAdmin,

	// admin global setting
	PermissionAdminGlobalSettingList:   PermissionTypeAdmin,
	PermissionAdminGlobalSettingUpdate: PermissionTypeAdmin,

	// admin icd10
	PermissionAdminIcd10List: PermissionTypeAdmin,
	PermissionAdminIcd10Tree: PermissionTypeAdmin,

	// admin icd10 trending
	PermissionAdminIcd10TrendingList:   PermissionTypeAdmin,
	PermissionAdminIcd10TrendingCreate: PermissionTypeAdmin,
	PermissionAdminIcd10TrendingMove:   PermissionTypeAdmin,
	PermissionAdminIcd10TrendingDelete: PermissionTypeAdmin,

	// admin injection
	PermissionAdminInjectionCreate: PermissionTypeAdmin,
	PermissionAdminInjectionList:   PermissionTypeAdmin,
	PermissionAdminInjectionShow:   PermissionTypeAdmin,
	PermissionAdminInjectionUpdate: PermissionTypeAdmin,
	PermissionAdminInjectionDelete: PermissionTypeAdmin,

	// admin injection attention
	PermissionAdminInjectionAttentionCreate: PermissionTypeAdmin,
	PermissionAdminInjectionAttentionDelete: PermissionTypeAdmin,

	// admin injection category
	PermissionAdminInjectionCategoryCreate: PermissionTypeAdmin,
	PermissionAdminInjectionCategoryDelete: PermissionTypeAdmin,

	// admin injection instruction
	PermissionAdminInjectionInstructionCreate: PermissionTypeAdmin,
	PermissionAdminInjectionInstructionDelete: PermissionTypeAdmin,

	// admin injection unit
	PermissionAdminInjectionUnitCreate: PermissionTypeAdmin,
	PermissionAdminInjectionUnitUpdate: PermissionTypeAdmin,
	PermissionAdminInjectionUnitDelete: PermissionTypeAdmin,

	// admin insurance
	PermissionAdminInsuranceCreate: PermissionTypeAdmin,
	PermissionAdminInsuranceList:   PermissionTypeAdmin,
	PermissionAdminInsuranceShow:   PermissionTypeAdmin,
	PermissionAdminInsuranceUpdate: PermissionTypeAdmin,
	PermissionAdminInsuranceDelete: PermissionTypeAdmin,

	// admin insurance payor
	PermissionAdminInsurancePayorCreate: PermissionTypeAdmin,
	PermissionAdminInsurancePayorList:   PermissionTypeAdmin,
	PermissionAdminInsurancePayorShow:   PermissionTypeAdmin,
	PermissionAdminInsurancePayorUpdate: PermissionTypeAdmin,
	PermissionAdminInsurancePayorDelete: PermissionTypeAdmin,

	// admin medical record template
	PermissionAdminMedicalRecordTemplateCreate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateList:   PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateShow:   PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateUpdate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateDelete: PermissionTypeAdmin,

	// admin medical record template diagnose
	PermissionAdminMedicalRecordTemplateDiagnoseCreate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateDiagnoseUpdate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateDiagnoseDelete: PermissionTypeAdmin,

	// admin medical record template examination
	PermissionAdminMedicalRecordTemplateExaminationCreate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateExaminationUpdate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateExaminationDelete: PermissionTypeAdmin,

	// admin medical record template prescription
	PermissionAdminMedicalRecordTemplatePrescriptionCreate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplatePrescriptionUpdate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplatePrescriptionDelete: PermissionTypeAdmin,

	// admin medical record template medical tool
	PermissionAdminMedicalRecordTemplateMedicalToolCreate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateMedicalToolUpdate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateMedicalToolDelete: PermissionTypeAdmin,

	// admin medical record template treatment
	PermissionAdminMedicalRecordTemplateTreatmentCreate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateTreatmentUpdate: PermissionTypeAdmin,
	PermissionAdminMedicalRecordTemplateTreatmentDelete: PermissionTypeAdmin,

	// admin medical tool
	PermissionAdminMedicalToolCreate:            PermissionTypeAdmin,
	PermissionAdminMedicalToolList:              PermissionTypeAdmin,
	PermissionAdminMedicalToolShow:              PermissionTypeAdmin,
	PermissionAdminMedicalToolUpdate:            PermissionTypeAdmin,
	PermissionAdminMedicalToolDelete:            PermissionTypeAdmin,
	PermissionAdminMedicalToolCreateProductUnit: PermissionTypeAdmin,
	PermissionAdminMedicalToolUpdateProductUnit: PermissionTypeAdmin,
	PermissionAdminMedicalToolDeleteProductUnit: PermissionTypeAdmin,
	PermissionAdminMedicalToolOption:            PermissionTypeAdmin,

	// admin medicinal attention
	PermissionAdminMedicinalAttentionCreate:           PermissionTypeAdmin,
	PermissionAdminMedicinalAttentionList:             PermissionTypeAdmin,
	PermissionAdminMedicinalAttentionUpdate:           PermissionTypeAdmin,
	PermissionAdminMedicinalAttentionDelete:           PermissionTypeAdmin,
	PermissionAdminMedicinalAttentionOptionForAmpForm: PermissionTypeAdmin,

	// admin medicinal category
	PermissionAdminMedicinalCategoryCreate:           PermissionTypeAdmin,
	PermissionAdminMedicinalCategoryList:             PermissionTypeAdmin,
	PermissionAdminMedicinalCategoryUpdate:           PermissionTypeAdmin,
	PermissionAdminMedicinalCategoryDelete:           PermissionTypeAdmin,
	PermissionAdminMedicinalCategoryOptionForAmpForm: PermissionTypeAdmin,

	// admin medicinal classification
	PermissionAdminMedicinalClassificationCreate: PermissionTypeAdmin,
	PermissionAdminMedicinalClassificationList:   PermissionTypeAdmin,
	PermissionAdminMedicinalClassificationUpdate: PermissionTypeAdmin,
	PermissionAdminMedicinalClassificationDelete: PermissionTypeAdmin,
	PermissionAdminMedicinalClassificationOption: PermissionTypeAdmin,

	// admin medicinal controlled drug classification
	PermissionAdminMedicinalControlledDrugClassificationCreate: PermissionTypeAdmin,
	PermissionAdminMedicinalControlledDrugClassificationList:   PermissionTypeAdmin,
	PermissionAdminMedicinalControlledDrugClassificationUpdate: PermissionTypeAdmin,
	PermissionAdminMedicinalControlledDrugClassificationDelete: PermissionTypeAdmin,
	PermissionAdminMedicinalControlledDrugClassificationOption: PermissionTypeAdmin,

	// admin medicinal form
	PermissionAdminMedicinalFormCreate: PermissionTypeAdmin,
	PermissionAdminMedicinalFormList:   PermissionTypeAdmin,
	PermissionAdminMedicinalFormUpdate: PermissionTypeAdmin,
	PermissionAdminMedicinalFormDelete: PermissionTypeAdmin,
	PermissionAdminMedicinalFormOption: PermissionTypeAdmin,

	// admin medicinal instruction
	PermissionAdminMedicinalInstructionCreate:           PermissionTypeAdmin,
	PermissionAdminMedicinalInstructionList:             PermissionTypeAdmin,
	PermissionAdminMedicinalInstructionUpdate:           PermissionTypeAdmin,
	PermissionAdminMedicinalInstructionDelete:           PermissionTypeAdmin,
	PermissionAdminMedicinalInstructionOptionForAmpForm: PermissionTypeAdmin,

	// admin medicinal route
	PermissionAdminMedicinalRouteCreate: PermissionTypeAdmin,
	PermissionAdminMedicinalRouteList:   PermissionTypeAdmin,
	PermissionAdminMedicinalRouteUpdate: PermissionTypeAdmin,
	PermissionAdminMedicinalRouteDelete: PermissionTypeAdmin,
	PermissionAdminMedicinalRouteOption: PermissionTypeAdmin,

	// admin medicine
	PermissionAdminMedicineCreate: PermissionTypeAdmin,
	PermissionAdminMedicineList:   PermissionTypeAdmin,
	PermissionAdminMedicineShow:   PermissionTypeAdmin,
	PermissionAdminMedicineUpdate: PermissionTypeAdmin,
	PermissionAdminMedicineDelete: PermissionTypeAdmin,

	// admin medicine attention
	PermissionAdminMedicineAttentionCreate: PermissionTypeAdmin,
	PermissionAdminMedicineAttentionDelete: PermissionTypeAdmin,

	// admin medicine category
	PermissionAdminMedicineCategoryCreate: PermissionTypeAdmin,
	PermissionAdminMedicineCategoryDelete: PermissionTypeAdmin,

	// admin medicine instruction
	PermissionAdminMedicineInstructionCreate: PermissionTypeAdmin,
	PermissionAdminMedicineInstructionDelete: PermissionTypeAdmin,

	// admin medicine unit
	PermissionAdminMedicineUnitCreate: PermissionTypeAdmin,
	PermissionAdminMedicineUnitUpdate: PermissionTypeAdmin,
	PermissionAdminMedicineUnitDelete: PermissionTypeAdmin,

	// admin nurse
	PermissionAdminNurseCreate: PermissionTypeAdmin,
	PermissionAdminNurseList:   PermissionTypeAdmin,
	PermissionAdminNurseShow:   PermissionTypeAdmin,
	PermissionAdminNurseUpdate: PermissionTypeAdmin,

	// admin doctor qualification
	PermissionAdminNurseQualificationCreate: PermissionTypeAdmin,
	PermissionAdminNurseQualificationUpdate: PermissionTypeAdmin,
	PermissionAdminNurseQualificationDelete: PermissionTypeAdmin,

	// admin pharmacist
	PermissionAdminPharmacistCreate:                    PermissionTypeAdmin,
	PermissionAdminPharmacistList:                      PermissionTypeAdmin,
	PermissionAdminPharmacistShow:                      PermissionTypeAdmin,
	PermissionAdminPharmacistUpdate:                    PermissionTypeAdmin,
	PermissionAdminPharmacistOptionForCreateClinicForm: PermissionTypeAdmin,
	PermissionAdminPharmacistOptionForUpdateClinicForm: PermissionTypeAdmin,
	PermissionAdminPharmacistQualificationCreate:       PermissionTypeAdmin,
	PermissionAdminPharmacistQualificationUpdate:       PermissionTypeAdmin,
	PermissionAdminPharmacistQualificationDelete:       PermissionTypeAdmin,

	// admin product
	PermissionAdminProductOption:                       PermissionTypeAdmin,
	PermissionAdminProductOptionForExaminationItem:     PermissionTypeAdmin,
	PermissionAdminProductOptionForRetailProduct:       PermissionTypeAdmin,
	PermissionAdminProductOptionForTreatmentItem:       PermissionTypeAdmin,
	PermissionAdminProductOptionForProductPromotion:    PermissionTypeAdmin,
	PermissionAdminProductOptionForProductBundlingItem: PermissionTypeAdmin,

	// admin product bundling
	PermissionAdminProductBundlingCreate:        PermissionTypeAdmin,
	PermissionAdminProductBundlingList:          PermissionTypeAdmin,
	PermissionAdminProductBundlingShow:          PermissionTypeAdmin,
	PermissionAdminProductBundlingUpdate:        PermissionTypeAdmin,
	PermissionAdminProductBundlingUpdateEndedAt: PermissionTypeAdmin,
	PermissionAdminProductBundlingDelete:        PermissionTypeAdmin,

	// admin product bundling clinic
	PermissionAdminProductBundlingClinicCreate: PermissionTypeAdmin,
	PermissionAdminProductBundlingClinicDelete: PermissionTypeAdmin,

	// admin product bundling item
	PermissionAdminProductBundlingItemCreate: PermissionTypeAdmin,
	PermissionAdminProductBundlingItemDelete: PermissionTypeAdmin,

	// admin product promotion
	PermissionAdminProductPromotionCreate:        PermissionTypeAdmin,
	PermissionAdminProductPromotionUpdate:        PermissionTypeAdmin,
	PermissionAdminProductPromotionList:          PermissionTypeAdmin,
	PermissionAdminProductPromotionShow:          PermissionTypeAdmin,
	PermissionAdminProductPromotionUpdateEndedAt: PermissionTypeAdmin,
	PermissionAdminProductPromotionDelete:        PermissionTypeAdmin,
	PermissionAdminProductPromotionClinicCreate:  PermissionTypeAdmin,
	PermissionAdminProductPromotionClinicDelete:  PermissionTypeAdmin,

	// admin product unit
	PermissionAdminProductUnitOption: PermissionTypeAdmin,

	// admin qualification
	PermissionAdminQualificationTypeCreate: PermissionTypeAdmin,
	PermissionAdminQualificationTypeList:   PermissionTypeAdmin,
	PermissionAdminQualificationTypeShow:   PermissionTypeAdmin,
	PermissionAdminQualificationTypeUpdate: PermissionTypeAdmin,
	PermissionAdminQualificationTypeDelete: PermissionTypeAdmin,
	PermissionAdminQualificationTypeOption: PermissionTypeAdmin,

	// admin retail product
	PermissionAdminRetailProductCreate:            PermissionTypeAdmin,
	PermissionAdminRetailProductUpload:            PermissionTypeAdmin,
	PermissionAdminRetailProductList:              PermissionTypeAdmin,
	PermissionAdminRetailProductShow:              PermissionTypeAdmin,
	PermissionAdminRetailProductUpdate:            PermissionTypeAdmin,
	PermissionAdminRetailProductDelete:            PermissionTypeAdmin,
	PermissionAdminRetailProductCreateProductUnit: PermissionTypeAdmin,
	PermissionAdminRetailProductUpdateProductUnit: PermissionTypeAdmin,
	PermissionAdminRetailProductDeleteProductUnit: PermissionTypeAdmin,

	// admin retail product category
	PermissionAdminRetailProductCategoryCreate:                     PermissionTypeAdmin,
	PermissionAdminRetailProductCategoryList:                       PermissionTypeAdmin,
	PermissionAdminRetailProductCategoryShow:                       PermissionTypeAdmin,
	PermissionAdminRetailProductCategoryUpdate:                     PermissionTypeAdmin,
	PermissionAdminRetailProductCategoryDelete:                     PermissionTypeAdmin,
	PermissionAdminRetailProductCategoryOptionForRetailProductForm: PermissionTypeAdmin,

	// admin role
	PermissionAdminRoleList:                PermissionTypeAdmin,
	PermissionAdminRoleShow:                PermissionTypeAdmin,
	PermissionAdminRoleOption:              PermissionTypeAdmin,
	PermissionAdminRoleOptionForUserFilter: PermissionTypeAdmin,
	PermissionAdminRoleOptionForUserForm:   PermissionTypeAdmin,

	// admin supplier
	PermissionAdminSupplierCreate:                   PermissionTypeAdmin,
	PermissionAdminSupplierFetch:                    PermissionTypeAdmin,
	PermissionAdminSupplierGet:                      PermissionTypeAdmin,
	PermissionAdminSupplierUpdate:                   PermissionTypeAdmin,
	PermissionAdminSupplierDelete:                   PermissionTypeAdmin,
	PermissionAdminSupplierOptionForExaminationForm: PermissionTypeAdmin,

	// admin teleconsult type
	PermissionAdminTeleconsultTypeCreate: PermissionTypeAdmin,
	PermissionAdminTeleconsultTypeList:   PermissionTypeAdmin,
	PermissionAdminTeleconsultTypeShow:   PermissionTypeAdmin,
	PermissionAdminTeleconsultTypeUpdate: PermissionTypeAdmin,
	PermissionAdminTeleconsultTypeDelete: PermissionTypeAdmin,

	// admin treatment
	PermissionAdminTreatmentCreate: PermissionTypeAdmin,
	PermissionAdminTreatmentList:   PermissionTypeAdmin,
	PermissionAdminTreatmentShow:   PermissionTypeAdmin,
	PermissionAdminTreatmentUpdate: PermissionTypeAdmin,
	PermissionAdminTreatmentDelete: PermissionTypeAdmin,
	PermissionAdminTreatmentOption: PermissionTypeAdmin,

	// admin treatment category
	PermissionAdminTreatmentCategoryCreate: PermissionTypeAdmin,
	PermissionAdminTreatmentCategoryList:   PermissionTypeAdmin,
	PermissionAdminTreatmentCategoryShow:   PermissionTypeAdmin,
	PermissionAdminTreatmentCategoryUpdate: PermissionTypeAdmin,
	PermissionAdminTreatmentCategoryDelete: PermissionTypeAdmin,
	PermissionAdminTreatmentCategoryOption: PermissionTypeAdmin,

	// admin treatment item
	PermissionAdminTreatmentItemCreate: PermissionTypeAdmin,
	PermissionAdminTreatmentItemUpdate: PermissionTypeAdmin,
	PermissionAdminTreatmentItemDelete: PermissionTypeAdmin,

	// admin unit
	PermissionAdminUnitCreate:               PermissionTypeAdmin,
	PermissionAdminUnitList:                 PermissionTypeAdmin,
	PermissionAdminUnitUpdate:               PermissionTypeAdmin,
	PermissionAdminUnitDelete:               PermissionTypeAdmin,
	PermissionAdminUnitOptionForProductForm: PermissionTypeAdmin,

	// admin user
	PermissionAdminUserCreate:                  PermissionTypeAdmin,
	PermissionAdminUserList:                    PermissionTypeAdmin,
	PermissionAdminUserShow:                    PermissionTypeAdmin,
	PermissionAdminUserUpdate:                  PermissionTypeAdmin,
	PermissionAdminUserOptionForDoctorForm:     PermissionTypeAdmin,
	PermissionAdminUserOptionForNurseForm:      PermissionTypeAdmin,
	PermissionAdminUserOptionForPharmacistForm: PermissionTypeAdmin,

	// admin vaccine
	PermissionAdminVaccineCreate: PermissionTypeAdmin,
	PermissionAdminVaccineList:   PermissionTypeAdmin,
	PermissionAdminVaccineShow:   PermissionTypeAdmin,
	PermissionAdminVaccineUpdate: PermissionTypeAdmin,
	PermissionAdminVaccineDelete: PermissionTypeAdmin,

	// admin vaccine attention
	PermissionAdminVaccineAttentionCreate: PermissionTypeAdmin,
	PermissionAdminVaccineAttentionDelete: PermissionTypeAdmin,

	// admin vaccine category
	PermissionAdminVaccineCategoryCreate: PermissionTypeAdmin,
	PermissionAdminVaccineCategoryDelete: PermissionTypeAdmin,

	// admin vaccine instruction
	PermissionAdminVaccineInstructionCreate: PermissionTypeAdmin,
	PermissionAdminVaccineInstructionDelete: PermissionTypeAdmin,

	// admin vaccine unit
	PermissionAdminVaccineUnitCreate: PermissionTypeAdmin,
	PermissionAdminVaccineUnitUpdate: PermissionTypeAdmin,
	PermissionAdminVaccineUnitDelete: PermissionTypeAdmin,

	// admin virtual medicinal product
	PermissionAdminVirtualMedicinalProductCreate: PermissionTypeAdmin,
	PermissionAdminVirtualMedicinalProductList:   PermissionTypeAdmin,
	PermissionAdminVirtualMedicinalProductShow:   PermissionTypeAdmin,
	PermissionAdminVirtualMedicinalProductUpdate: PermissionTypeAdmin,
	PermissionAdminVirtualMedicinalProductDelete: PermissionTypeAdmin,
	PermissionAdminVirtualMedicinalProductOption: PermissionTypeAdmin,

	// admin virtual therapeutic moiety
	PermissionAdminVirtualTherapeuticMoietyList:                PermissionTypeAdmin,
	PermissionAdminVirtualTherapeuticMoietyShow:                PermissionTypeAdmin,
	PermissionAdminVirtualTherapeuticMoietyCreate:              PermissionTypeAdmin,
	PermissionAdminVirtualTherapeuticMoietyUpdate:              PermissionTypeAdmin,
	PermissionAdminVirtualTherapeuticMoietyDelete:              PermissionTypeAdmin,
	PermissionAdminVirtualTherapeuticMoietyOption:              PermissionTypeAdmin,
	PermissionAdminVirtualTherapeuticMoietyOptionForVmpVtmForm: PermissionTypeAdmin,

	// admin vmp code
	PermissionAdminVmpCodeCreate: PermissionTypeAdmin,
	PermissionAdminVmpCodeDelete: PermissionTypeAdmin,

	// admin vmp vtm
	PermissionAdminVmpVtmCreate: PermissionTypeAdmin,
	PermissionAdminVmpVtmUpdate: PermissionTypeAdmin,
	PermissionAdminVmpVtmDelete: PermissionTypeAdmin,

	// admin vtm code
	PermissionAdminVtmCodeCreate: PermissionTypeAdmin,
	PermissionAdminVtmCodeDelete: PermissionTypeAdmin,

	// admin voucher code
	PermissionAdminVoucherCodeCreate: PermissionTypeAdmin,
	PermissionAdminVoucherCodeList:   PermissionTypeAdmin,
	PermissionAdminVoucherCodeShow:   PermissionTypeAdmin,
	PermissionAdminVoucherCodeUpdate: PermissionTypeAdmin,
	PermissionAdminVoucherCodeDelete: PermissionTypeAdmin,

	// admin voucher code redeem
	PermissionAdminVoucherCodeRedeemList: PermissionTypeAdmin,
	PermissionAdminVoucherCodeRedeemShow: PermissionTypeAdmin,

	// admin voucher generate rule
	PermissionAdminVoucherGenerateRuleCreate: PermissionTypeAdmin,
	PermissionAdminVoucherGenerateRuleList:   PermissionTypeAdmin,
	PermissionAdminVoucherGenerateRuleShow:   PermissionTypeAdmin,
	PermissionAdminVoucherGenerateRuleUpdate: PermissionTypeAdmin,
	PermissionAdminVoucherGenerateRuleDelete: PermissionTypeAdmin,

	// cashier sales invoice
	PermissionCashierSalesInvoiceList: PermissionTypeMustAssignAndOnSite,
	PermissionCashierSalesInvoiceShow: PermissionTypeMustAssignAndOnSite,

	// cluster manager clinic
	PermissionClusterManagerClinicGetCashOnHand:               PermissionTypeAdmin,
	PermissionClusterManagerClinicOptionForCashierSummaryList: PermissionTypeAdmin,
	PermissionClusterManagerClinicOptionForCashDepositForm:    PermissionTypeAdmin,
	PermissionClusterManagerClinicOptionForCashDepositFilter:  PermissionTypeAdmin,

	// cluster manager cashier summary
	PermissionClusterManagerCashierSummaryList:             PermissionTypeAdmin,
	PermissionClusterManagerCashierSummaryListSummarized:   PermissionTypeAdmin,
	PermissionClusterManagerCashierSummaryShow:             PermissionTypeAdmin,
	PermissionClusterManagerCashierSummaryShowSession:      PermissionTypeAdmin,
	PermissionClusterManagerCashierSummaryShowSalesInvoice: PermissionTypeAdmin,
	PermissionClusterManagerCashierSummaryDownloadCsv:      PermissionTypeAdmin,

	// clininc manager sales invoice
	PermissionClinicManagerSalesInvoiceList:     PermissionTypeMustAssignAndOnSite,
	PermissionClinicManagerSalesInvoiceDownload: PermissionTypeMustAssignAndOnSite,

	// doctor patient vital check
	PermissionDoctorPatientVitalCheckCreate: PermissionTypeMustAssignAndOnSite,
	PermissionDoctorPatientVitalCheckUpdate: PermissionTypeMustAssignAndOnSite,

	// nurse patient vital check
	PermissionNursePatientVitalCheckCreate: PermissionTypeMustAssignAndOnSite,
	PermissionNursePatientVitalCheckUpdate: PermissionTypeMustAssignAndOnSite,

	// appointment
	PermissionAppointmentCreate:                                  PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentList:                                    PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentListForAppointmentFormByClinicOrDoctor:  PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentListForAppointmentFormByPatient:         PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentListForQueueNumberGenerateFormByPatient: PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentShow:                                    PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentUpdate:                                  PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentUpdateReferredDoctor:                    PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentApprove:                                 PermissionTypeMustAssignAndOnSite,
	PermissionAppointmentCancel:                                  PermissionTypeMustAssignAndOnSite,

	// appointment type
	PermissionAppointmentTypeOption: PermissionTypeMustAssignAndOnSite,

	// bank
	PermissionBankOptionForCashDepositForm: PermissionTypeAdmin,

	// cash deposit
	PermissionCashDepositList:     PermissionTypeMustAssign,
	PermissionCashDepositShow:     PermissionTypeMustAssign,
	PermissionCashDepositDownload: PermissionTypeMustAssign,

	// cashier
	PermissionCashierCreate:                           PermissionTypeMustAssignAndOnSite,
	PermissionCashierList:                             PermissionTypeMustAssignAndOnSite,
	PermissionCashierShow:                             PermissionTypeMustAssignAndOnSite,
	PermissionCashierUpdate:                           PermissionTypeMustAssignAndOnSite,
	PermissionCashierDelete:                           PermissionTypeMustAssignAndOnSite,
	PermissionCashierOptionForBeginCashierSessionForm: PermissionTypeMustAssignAndOnSiteStrict,

	// cashier session
	PermissionCashierSessionBegin:            PermissionTypeMustAssignAndOnSiteStrict,
	PermissionCashierSessionOpenDrawer:       PermissionTypeMustAssignAndOnSiteStrict,
	PermissionCashierSessionList:             PermissionTypeMustAssignAndOnSiteStrict,
	PermissionCashierSessionShow:             PermissionTypeMustAssignAndOnSiteStrict,
	PermissionCashierSessionShowActive:       PermissionTypeMustAssignAndOnSite,
	PermissionCashierSessionClose:            PermissionTypeMustAssignAndOnSiteStrict,
	PermissionCashierSessionUpdateMoneyCount: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionCashierSessionEnd:              PermissionTypeMustAssignAndOnSiteStrict,

	// cashier summary
	PermissionCashierSummaryCreate:                      PermissionTypeMustAssignAndOnSite,
	PermissionCashierSummaryList:                        PermissionTypeMustAssign,
	PermissionCashierSummaryListSummarized:              PermissionTypeMustAssign,
	PermissionCashierSummaryDownloadCsv:                 PermissionTypeMustAssign,
	PermissionCashierSummaryShow:                        PermissionTypeMustAssign,
	PermissionCashierSummaryShowSession:                 PermissionTypeMustAssign,
	PermissionCashierSummaryShowSalesInvoice:            PermissionTypeMustAssign,
	PermissionCashierSummaryShowDailyReport:             PermissionTypeMustAssignAndOnSite,
	PermissionCashierSummaryShowDailyReportSession:      PermissionTypeMustAssignAndOnSite,
	PermissionCashierSummaryShowDailyReportSalesInvoice: PermissionTypeMustAssignAndOnSite,

	// clinic
	PermissionClinicOptionForAppointmentForm: PermissionTypeMustAssignAndOnSite,

	// clinic ottopay edc
	PermissionClinicOttopayEdcCreate:               PermissionTypeMustAssign,
	PermissionClinicOttopayEdcList:                 PermissionTypeMustAssign,
	PermissionClinicOttopayEdcShow:                 PermissionTypeMustAssign,
	PermissionClinicOttopayEdcUpdate:               PermissionTypeMustAssign,
	PermissionClinicOttopayEdcDelete:               PermissionTypeMustAssign,
	PermissionClinicOttopayEdcOptionForCashierForm: PermissionTypeMustAssign,

	// clinic printer
	PermissionClinicPrinterCreate:                      PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterDelete:                      PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterOptionForCashierForm:        PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterOptionForQueueCounter:       PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterOptionForPrintMedicineLabel: PermissionTypeMustAssignAndOnSite,

	// clinic printer gateway
	PermissionClinicPrinterGatewayCreate: PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterGatewayList:   PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterGatewayShow:   PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterGatewayUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionClinicPrinterGatewayDelete: PermissionTypeMustAssignAndOnSite,

	// clinic setting
	PermissionClinicSettingList:   PermissionTypeMustAssign,
	PermissionClinicSettingUpdate: PermissionTypeMustAssign,

	// doctor
	PermissionDoctorList:                     PermissionTypeMustAssignAndOnSite,
	PermissionDoctorShow:                     PermissionTypeMustAssignAndOnSite,
	PermissionDoctorOptionForAppointmentForm: PermissionTypeMustAssignAndOnSite,
	PermissionDoctorOptionForTeleconsultForm: PermissionTypeMustAssignAndOnSite,

	// doctor memo
	PermissionDoctorMemoCreate: PermissionTypeMustAssignAndOnSite,
	PermissionDoctorMemoUpdate: PermissionTypeMustAssignAndOnSite,

	// doctor referral
	PermissionDoctorReferralCreate: PermissionTypeMustAssignAndOnSite,
	PermissionDoctorReferralUpdate: PermissionTypeMustAssignAndOnSite,

	// health examination result
	PermissionHealthExaminationResultCreate: PermissionTypeMustAssignAndOnSite,
	PermissionHealthExaminationResultUpdate: PermissionTypeMustAssignAndOnSite,

	// medical summary
	PermissionMedicalSummaryCreate: PermissionTypeMustAssignAndOnSite,

	// examination
	PermissionExaminationList:                   PermissionTypeMustAssign,
	PermissionExaminationShow:                   PermissionTypeMustAssign,
	PermissionExaminationOptionForMedicalRecord: PermissionTypeMustAssignAndOnSite,

	// icd 10
	PermissionIcd10List: PermissionTypeMustAssignAndOnSite,
	PermissionIcd10Tree: PermissionTypeMustAssignAndOnSite,
	PermissionIcd10OptionForPatientDiseaseHistoryForm:    PermissionTypeMustAssignAndOnSite,
	PermissionIcd10OptionForPatientCongenitalDiseaseForm: PermissionTypeMustAssignAndOnSite,

	// icd 10 trending
	PermissionIcd10TrendingList: PermissionTypeMustAssignAndOnSite,

	// injection
	PermissionInjectionShow: PermissionTypeMustAssign,

	// insurance
	PermissionInsuranceOptionForQueueNumber: PermissionTypeMustAssignAndOnSiteStrict,

	PermissionInsurancePayorOptionForQueueNumber: PermissionTypeMustAssignAndOnSiteStrict,

	// cart
	PermissionCartShow: PermissionTypeMustAssignAndOnSite,

	// cart item
	PermissionCartItemCreateForActiveRetailCart:                PermissionTypeMustAssignAndOnSite,
	PermissionCartItemCreateForActiveRetailCartByBarcodeOrCode: PermissionTypeMustAssignAndOnSite,
	PermissionCartItemQtyUpdateInActiveCart:                    PermissionTypeMustAssignAndOnSite,
	PermissionCartItemReprintDocumentFeeDelete:                 PermissionTypeMustAssignAndOnSite,
	PermissionCartItemDeleteFromActiveRetailCart:               PermissionTypeMustAssignAndOnSite,

	// medical certificate
	PermissionMedicalCertificateCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalCertificateUpdate: PermissionTypeMustAssignAndOnSite,

	// medical certificate diagnose
	PermissionMedicalCertificateDiagnoseCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalCertificateDiagnoseDelete: PermissionTypeMustAssignAndOnSite,

	// medical record
	PermissionMedicalRecordBegin:               PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordList:                PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordListForStandByForm:  PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordShow:                PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordUpdate:              PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordUpdateApplyTemplate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordUpdateAnnotation:    PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordUpdateAmendment:     PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordStandBy:             PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordResume:              PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordEnd:                 PermissionTypeMustAssignAndOnSiteStrict,

	// medical record amendment history
	PermissionMedicalRecordAmendmentHistoryList: PermissionTypeMustAssignAndOnSite,

	// medical record diagnose
	PermissionMedicalRecordDiagnoseCreate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordDiagnoseUpdate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordDiagnoseDelete: PermissionTypeMustAssignAndOnSiteStrict,

	// medical record examination
	PermissionMedicalRecordExaminationCreate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordExaminationUpdate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordExaminationDelete: PermissionTypeMustAssignAndOnSiteStrict,

	// medical record medical tool
	PermissionMedicalRecordMedicalToolCreate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordMedicalToolUpdate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordMedicalToolDelete: PermissionTypeMustAssignAndOnSiteStrict,

	// medical record prescription
	PermissionMedicalRecordPrescriptionCreate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordPrescriptionUpdate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordPrescriptionDelete: PermissionTypeMustAssignAndOnSiteStrict,

	// medical record treatment
	PermissionMedicalRecordTreatmentCreate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordTreatmentUpdate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordTreatmentDelete: PermissionTypeMustAssignAndOnSiteStrict,

	// medical record product bundling
	PermissionMedicalRecordProductBundlingCreate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordProductBundlingUpdate: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionMedicalRecordProductBundlingDelete: PermissionTypeMustAssignAndOnSiteStrict,

	// medical record template
	PermissionMedicalRecordTemplateCreate:                     PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateList:                       PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateShow:                       PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateUpdate:                     PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateDelete:                     PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateOptionForMedicalRecordForm: PermissionTypeMustAssignAndOnSite,

	// medical record template diagnose
	PermissionMedicalRecordTemplateDiagnoseCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateDiagnoseUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateDiagnoseDelete: PermissionTypeMustAssignAndOnSite,

	// medical record template medical tool
	PermissionMedicalRecordTemplateMedicalToolCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateMedicalToolUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateMedicalToolDelete: PermissionTypeMustAssignAndOnSite,

	// medical record template treatment
	PermissionMedicalRecordTemplateTreatmentCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateTreatmentUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateTreatmentDelete: PermissionTypeMustAssignAndOnSite,

	// medical record template examination
	PermissionMedicalRecordTemplateExaminationCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateExaminationUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplateExaminationDelete: PermissionTypeMustAssignAndOnSite,

	// medical record template prescription
	PermissionMedicalRecordTemplatePrescriptionCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplatePrescriptionUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicalRecordTemplatePrescriptionDelete: PermissionTypeMustAssignAndOnSite,

	// medical record waives
	PermissionMedicalRecordWaiveUpdate: PermissionTypeMustAssignAndOnSiteStrict,

	// medical tool
	PermissionMedicalToolOptionForMedicalRecord: PermissionTypeMustAssignAndOnSite,

	// medicinal dispense
	PermissionMedicinalDispenseDownload:               PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseCreate:                 PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseUpdate:                 PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseUpload:                 PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseList:                   PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseShow:                   PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseAddMixedMedicineFee:    PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseRemoveMixedMedicineFee: PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispensePrintBlank:             PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseCalibrate:              PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseUpdateChecked:          PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseUpdateCompleted:        PermissionTypeMustAssignAndOnSite,

	// medicinal dispense item
	PermissionMedicinalDispenseItemPrint:             PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseItemCreateExpiredDate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseItemUpdateExpiredDate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseItemDeleteExpiredDate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispenseItemUpdate:            PermissionTypeMustAssignAndOnSite,

	// medicinal dispense prescription
	PermissionMedicinalDispensePrescriptionCreate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispensePrescriptionUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionMedicinalDispensePrescriptionDelete: PermissionTypeMustAssignAndOnSite,

	// medicine
	PermissionMedicineShow:                       PermissionTypeMustAssignAndOnSite,
	PermissionMedicineOptionForMedicinalDispense: PermissionTypeMustAssignAndOnSite,

	// net promoter score
	PermissionNetPromoterScoreStartByQueueNumber:  PermissionTypeMustAssignAndOnSite,
	PermissionNetPromoterScoreStartBySalesInvoice: PermissionTypeMustAssignAndOnSite,
	PermissionNetPromoterScoreGetActive:           PermissionTypeMustAssignAndOnSite,
	PermissionNetPromoterScoreEnd:                 PermissionTypeMustAssignAndOnSite,
	PermissionNetPromoterScoreItemUpdate:          PermissionTypeMustAssignAndOnSite,

	// nric card ocr
	PermissionNricCardOcrExtract: PermissionTypeMustAssignAndOnSite,
	PermissionNricCardOcrUpload:  PermissionTypeMustAssignAndOnSite,

	// nurse
	PermissionNurseList: PermissionTypeMustAssignAndOnSite,
	PermissionNurseShow: PermissionTypeMustAssignAndOnSite,

	// patient
	PermissionPatientCreate:                           PermissionTypeMustAssignAndOnSite,
	PermissionPatientShow:                             PermissionTypeMustAssignAndOnSite,
	PermissionPatientShowMedicalSummary:               PermissionTypeMustAssignAndOnSite,
	PermissionPatientShowMedicalNotes:                 PermissionTypeMustAssignAndOnSite,
	PermissionPatientUpdate:                           PermissionTypeMustAssignAndOnSite,
	PermissionPatientUpdateMedicalSummary:             PermissionTypeMustAssignAndOnSite,
	PermissionPatientUpdateMedicalNotes:               PermissionTypeMustAssignAndOnSite,
	PermissionPatientOption:                           PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForAppointmentForm:         PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForPatientDoctorForm:       PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForPatientRelativeForm:     PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForPatientVisitFilter:      PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForPatientVisitForm:        PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForPatientVitalCheckFilter: PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForPatientVitalCheckForm:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForQueueNumberGenerateForm: PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForTaskForm:                PermissionTypeGlobal,
	PermissionPatientOptionForTeleconsult:             PermissionTypeMustAssignAndOnSite,
	PermissionPatientOptionForVaccination:             PermissionTypeMustAssignAndOnSite,

	// patient allergy
	PermissionPatientAllergyCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientAllergyList:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientAllergyShow:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientAllergyUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientAllergyDelete: PermissionTypeMustAssignAndOnSite,

	// patient congenital disease
	PermissionPatientCongenitalDiseaseCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientCongenitalDiseaseList:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientCongenitalDiseaseShow:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientCongenitalDiseaseUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientCongenitalDiseaseDelete: PermissionTypeMustAssignAndOnSite,

	// patient disease history
	PermissionPatientDiseaseHistoryCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientDiseaseHistoryList:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientDiseaseHistoryShow:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientDiseaseHistoryUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientDiseaseHistoryDelete: PermissionTypeMustAssignAndOnSite,

	// patient doctor
	PermissionPatientDoctorCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientDoctorDelete: PermissionTypeMustAssignAndOnSite,

	// patient document
	PermissionPatientDocumentList:     PermissionTypeMustAssignAndOnSite,
	PermissionPatientDocumentUpload:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientDocumentComplete: PermissionTypeMustAssignAndOnSite,
	PermissionPatientDocumentDownload: PermissionTypeMustAssignAndOnSite,
	PermissionPatientDocumentDelete:   PermissionTypeMustAssignAndOnSite,

	// patient interview
	PermissionPatientInterviewCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientInterviewList:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientInterviewShow:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientInterviewUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientInterviewDelete: PermissionTypeMustAssignAndOnSite,

	// patient medicine allergy
	PermissionPatientMedicineAllergyCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineAllergyList:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineAllergyShow:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineAllergyUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineAllergyDelete: PermissionTypeMustAssignAndOnSite,

	// patient medicine history
	PermissionPatientMedicineHistoryCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineHistoryList:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineHistoryShow:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineHistoryUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientMedicineHistoryDelete: PermissionTypeMustAssignAndOnSite,

	// patient relative
	PermissionPatientRelativeCreate:                   PermissionTypeMustAssignAndOnSite,
	PermissionPatientRelativeList:                     PermissionTypeMustAssignAndOnSite,
	PermissionPatientRelativeUpdate:                   PermissionTypeMustAssignAndOnSite,
	PermissionPatientRelativeUpdateSetDefault:         PermissionTypeMustAssignAndOnSite,
	PermissionPatientRelativeDelete:                   PermissionTypeMustAssignAndOnSite,
	PermissionPatientRelativeOptionForQueueNumberForm: PermissionTypeMustAssignAndOnSite,

	// patient surgery
	PermissionPatientSurgeryCreate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientSurgeryList:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientSurgeryShow:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientSurgeryUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionPatientSurgeryDelete: PermissionTypeMustAssignAndOnSite,

	// patient test result
	PermissionPatientTestResultCreate:           PermissionTypeMustAssignAndOnSite,
	PermissionPatientTestResultUpload:           PermissionTypeMustAssignAndOnSite,
	PermissionPatientTestResultList:             PermissionTypeMustAssignAndOnSite,
	PermissionPatientTestResultShow:             PermissionTypeMustAssignAndOnSite,
	PermissionPatientTestResultUpdate:           PermissionTypeMustAssignAndOnSite,
	PermissionPatientTestResultUpdateReadStatus: PermissionTypeMustAssignAndOnSite,

	// patient visit
	PermissionPatientVisitList:                                  PermissionTypeMustAssignAndOnSite,
	PermissionPatientVisitSummarizeHistory:                      PermissionTypeMustAssignAndOnSite,
	PermissionPatientVisitShow:                                  PermissionTypeMustAssignAndOnSite,
	PermissionPatientVisitInsuranceCardFileShowForMedicalRecord: PermissionTypeMustAssignAndOnSite,
	PermissionPatientVisitInsuranceCardFileShowForQueueNumber:   PermissionTypeMustAssignAndOnSite,
	PermissionPatientVisitUpdate:                                PermissionTypeMustAssignAndOnSite,
	PermissionPatientVisitOptionForMedicalRecord:                PermissionTypeMustAssignAndOnSite,

	// patient vital check
	PermissionPatientVitalCheckList:                 PermissionTypeMustAssignAndOnSite,
	PermissionPatientVitalCheckListForMedicalRecord: PermissionTypeMustAssignAndOnSite,
	PermissionPatientVitalCheckShow:                 PermissionTypeMustAssignAndOnSite,
	PermissionPatientVitalCheckShowLatest:           PermissionTypeMustAssignAndOnSite,
	PermissionPatientVitalCheckVisualize:            PermissionTypeMustAssignAndOnSite,
	PermissionPatientVitalCheckUpdateIsNursing:      PermissionTypeMustAssignAndOnSite,
	PermissionPatientVitalCheckUpdateIsSmoking:      PermissionTypeMustAssignAndOnSite,
	PermissionPatientVitalCheckUpdateIsPregnant:     PermissionTypeMustAssignAndOnSite,

	// pharmacist
	PermissionPharmacistList: PermissionTypeMustAssignAndOnSite,
	PermissionPharmacistShow: PermissionTypeMustAssignAndOnSite,

	// product
	PermissionProductListForRetailCart:                      PermissionTypeMustAssignAndOnSite,
	PermissionProductPriceTagDownload:                       PermissionTypeMustAssignAndOnSite,
	PermissionProductOptionForMedicalRecord:                 PermissionTypeMustAssignAndOnSite,
	PermissionProductOptionForMedicinalDispensePrescription: PermissionTypeMustAssignAndOnSite,
	PermissionProductOptionForPriceTagDownload:              PermissionTypeMustAssignAndOnSite,
	PermissionProductOptionForProductAdjustmentForm:         PermissionTypeMustAssignAndOnSite,
	PermissionProductOptionForProductReceiveItem:            PermissionTypeMustAssignAndOnSite,
	PermissionProductOptionForProductTransferItemForm:       PermissionTypeMustAssignAndOnSite,
	PermissionProductOptionExpiredDate:                      PermissionTypeMustAssignAndOnSite,

	// product adjustment
	PermissionProductAdjustmentCreate: PermissionTypeMustAssignAndOnSite,
	PermissionProductAdjustmentList:   PermissionTypeMustAssignAndOnSite,
	PermissionProductAdjustmentShow:   PermissionTypeMustAssignAndOnSite,
	PermissionProductAdjustmentUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionProductAdjustmentDelete: PermissionTypeMustAssignAndOnSite,

	// product adjustment item
	PermissionProductAdjustmentItemCreate: PermissionTypeMustAssignAndOnSite,
	PermissionProductAdjustmentItemDelete: PermissionTypeMustAssignAndOnSite,

	// product bundling
	PermissionProductBundlingShow:                   PermissionTypeMustAssignAndOnSite,
	PermissionProductBundlingOptionForMedicalRecord: PermissionTypeMustAssignAndOnSite,

	// product inventory
	PermissionProductInventoryList: PermissionTypeMustAssign,

	// product promotion
	PermissionProductPromotionOptionForPriceTagDownload: PermissionTypeMustAssignAndOnSite,

	// product receive
	PermissionProductReceiveCreate: PermissionTypeMustAssign,
	PermissionProductReceiveList:   PermissionTypeMustAssign,
	PermissionProductReceiveShow:   PermissionTypeMustAssign,
	PermissionProductReceiveUpdate: PermissionTypeMustAssign,
	PermissionProductReceiveDelete: PermissionTypeMustAssign,

	// product receive item
	PermissionProductReceiveItemCreate: PermissionTypeMustAssign,
	PermissionProductReceiveItemDelete: PermissionTypeMustAssign,

	// product unit
	PermissionProductUnitOption:                             PermissionTypeMustAssign,
	PermissionProductUnitOptionForProductAdjustmentItemForm: PermissionTypeMustAssignAndOnSite,
	PermissionProductUnitOptionForProductTransferItemForm:   PermissionTypeMustAssignAndOnSite,
	PermissionProductUnitOptionForPriceTagDownload:          PermissionTypeMustAssignAndOnSite,

	// product transfer
	PermissionProductTransferCreate:     PermissionTypeMustAssignAndOnSite,
	PermissionProductTransferList:       PermissionTypeMustAssignAndOnSite,
	PermissionProductTransferShow:       PermissionTypeMustAssignAndOnSite,
	PermissionProductTransferUpdate:     PermissionTypeMustAssignAndOnSite,
	PermissionProductTransferDelete:     PermissionTypeMustAssignAndOnSite,
	PermissionProductTransferCreateItem: PermissionTypeMustAssignAndOnSite,
	PermissionProductTransferDeleteItem: PermissionTypeMustAssignAndOnSite,

	// queue
	PermissionQueueCreate:                           PermissionTypeMustAssignAndOnSite,
	PermissionQueueList:                             PermissionTypeMustAssignAndOnSite,
	PermissionQueueShow:                             PermissionTypeMustAssignAndOnSite,
	PermissionQueueShowProcess:                      PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueUpdate:                           PermissionTypeMustAssignAndOnSite,
	PermissionQueueDelete:                           PermissionTypeMustAssignAndOnSite,
	PermissionQueueOption:                           PermissionTypeMustAssignAndOnSite,
	PermissionQueueOptionForQueueCounterCheckInForm: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueOptionForQueueDisplay:            PermissionTypeMustAssignAndOnSite,

	// queue counter
	PermissionQueueCounterCreate:                           PermissionTypeMustAssignAndOnSite,
	PermissionQueueCounterUpdate:                           PermissionTypeMustAssignAndOnSite,
	PermissionQueueCounterCheckIn:                          PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueCounterCheckOut:                         PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueCounterOptionForCheckInForm:             PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueCounterOptionForQueueNumberCompleteForm: PermissionTypeMustAssignAndOnSite,
	PermissionQueueCounterOptionForQueueNumberGenerateForm: PermissionTypeMustAssignAndOnSite,

	// queue display
	PermissionQueueDisplayCreate: PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayList:   PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayShow:   PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayDelete: PermissionTypeMustAssignAndOnSite,

	// queue display banner
	PermissionQueueDisplayBannerCreate: PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayBannerUpload: PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayBannerMove:   PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayBannerDelete: PermissionTypeMustAssignAndOnSite,

	// queue display queue
	PermissionQueueDisplayQueueCreate: PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayQueueDelete: PermissionTypeMustAssignAndOnSite,

	// queue display running text
	PermissionQueueDisplayRunningTextCreate: PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayRunningTextMove:   PermissionTypeMustAssignAndOnSite,
	PermissionQueueDisplayRunningTextDelete: PermissionTypeMustAssignAndOnSite,

	// queue number
	PermissionQueueNumberGenerate:            PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberUploadInsuranceCard: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberGet:                 PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberCurrent:             PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberReprint:             PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberUpdate:              PermissionTypeMustAssignAndOnSiteStrict,

	// queue number action
	PermissionQueueNumberCall:       PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberCallRepeat: PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberProcess:    PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberComplete:   PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberMove:       PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberSkip:       PermissionTypeMustAssignAndOnSiteStrict,
	PermissionQueueNumberStandBy:    PermissionTypeMustAssignAndOnSiteStrict,

	// queue session
	PermissionQueueSessionBegin:                     PermissionTypeMustAssignAndOnSite,
	PermissionQueueSessionEnd:                       PermissionTypeMustAssignAndOnSite,
	PermissionQueueSessionGetActive:                 PermissionTypeMustAssignAndOnSite,
	PermissionQueueSessionCheckForActiveQueueNumber: PermissionTypeMustAssignAndOnSite,

	// report daily transaction
	PermissionReportDailyTransactionList:     PermissionTypeMustAssign,
	PermissionReportDailyTransactionDownload: PermissionTypeMustAssign,

	// report finance inventory
	PermissionReportFinanceInventoryCreate:   PermissionTypeMustAssign,
	PermissionReportFinanceInventoryList:     PermissionTypeMustAssign,
	PermissionReportFinanceInventoryDownload: PermissionTypeMustAssign,

	// report inventory stock
	PermissionReportInventoryStockCreate:   PermissionTypeMustAssign,
	PermissionReportInventoryStockList:     PermissionTypeMustAssign,
	PermissionReportInventoryStockDownload: PermissionTypeMustAssign,

	// report monthly revenue
	PermissionReportMonthlyRevenueList:     PermissionTypeAdmin,
	PermissionReportMonthlyRevenueDownload: PermissionTypeAdmin,

	// report month to date sales
	PermissionReportMonthToDateSalesList:     PermissionTypeMustAssign,
	PermissionReportMonthToDateSalesDownload: PermissionTypeMustAssign,

	// report net promoter score
	PermissionReportNetPromoterScoreCreate:   PermissionTypeAdmin,
	PermissionReportNetPromoterScoreList:     PermissionTypeAdmin,
	PermissionReportNetPromoterScoreDownload: PermissionTypeAdmin,

	// report service and performance
	PermissionReportServiceAndPerformanceCreate:   PermissionTypeAdmin,
	PermissionReportServiceAndPerformanceList:     PermissionTypeAdmin,
	PermissionReportServiceAndPerformanceDownload: PermissionTypeAdmin,

	// report waive fee
	PermissionReportWaiveFeeDownload: PermissionTypeMustAssign,

	// retail cart
	PermissionRetailCartGetActive:                PermissionTypeMustAssignAndOnSite,
	PermissionRetailCartShow:                     PermissionTypeMustAssignAndOnSite,
	PermissionRetailCartHoldActive:               PermissionTypeMustAssignAndOnSite,
	PermissionRetailCartRestore:                  PermissionTypeMustAssignAndOnSite,
	PermissionRetailCartDelete:                   PermissionTypeMustAssignAndOnSite,
	PermissionRetailCartDeleteActive:             PermissionTypeMustAssignAndOnSite,
	PermissionRetailCartOptionForRestoreOrDelete: PermissionTypeMustAssignAndOnSite,

	// role
	PermissionRoleOptionForQueueForm:  PermissionTypeMustAssign,
	PermissionRoleOptionForUserFilter: PermissionTypeMustAssign,

	// sales invoice
	PermissionSalesInvoiceGenerate:          PermissionTypeMustAssignAndOnSite,
	PermissionSalesInvoiceDownload:          PermissionTypeMustAssignAndOnSite,
	PermissionSalesInvoiceVoid:              PermissionTypeMustAssignAndOnSite,
	PermissionSalesInvoiceCancel:            PermissionTypeMustAssignAndOnSite,
	PermissionSalesInvoiceApplyVoucherCode:  PermissionTypeMustAssignAndOnSite,
	PermissionSalesInvoiceRemoveVoucherCode: PermissionTypeMustAssignAndOnSite,
	PermissionSalesInvoiceCheckVoucherCode:  PermissionTypeMustAssignAndOnSite,

	// sales payment
	PermissionSalesPaymentCreate:      PermissionTypeMustAssignAndOnSite,
	PermissionSalesPaymentShow:        PermissionTypeMustAssignAndOnSite,
	PermissionSalesPaymentPrintRetail: PermissionTypeMustAssignAndOnSite,
	PermissionSalesPaymentCancel:      PermissionTypeMustAssignAndOnSite,

	// sales payment method
	PermissionSalesPaymentMethodOption: PermissionTypeMustAssignAndOnSiteStrict,

	// supplier
	PermissionSupplierOptionForProductReceiveForm: PermissionTypeMustAssignAndOnSiteStrict,

	// task
	PermissionTaskCreate:         PermissionTypeGlobal,
	PermissionTaskList:           PermissionTypeGlobal,
	PermissionTaskShow:           PermissionTypeGlobal,
	PermissionTaskUpdate:         PermissionTypeGlobal,
	PermissionTaskUpdatePriority: PermissionTypeGlobal,
	PermissionTaskUpdateStatus:   PermissionTypeGlobal,
	PermissionTaskDelete:         PermissionTypeGlobal,

	// task attachment
	PermissionTaskAttachmentUpload: PermissionTypeGlobal,
	PermissionTaskAttachmentShow:   PermissionTypeGlobal,
	PermissionTaskAttachmentDelete: PermissionTypeGlobal,

	// task comment
	PermissionTaskCommentCreate: PermissionTypeGlobal,
	PermissionTaskCommentList:   PermissionTypeGlobal,
	PermissionTaskCommentUpdate: PermissionTypeGlobal,
	PermissionTaskCommentDelete: PermissionTypeGlobal,

	// teleconsult
	PermissionTeleconsultCreate:        PermissionTypeMustAssignAndOnSite,
	PermissionTeleconsultList:          PermissionTypeMustAssignAndOnSite,
	PermissionTeleconsultShow:          PermissionTypeMustAssignAndOnSite,
	PermissionTeleconsultUpdate:        PermissionTypeMustAssignAndOnSite,
	PermissionTeleconsultCancelConsult: PermissionTypeMustAssignAndOnSite,

	// teleconsult type
	PermissionTeleconsultTypeOption: PermissionTypeMustAssignAndOnSite,

	// treatment
	PermissionTreatmentList:                   PermissionTypeMustAssign,
	PermissionTreatmentShow:                   PermissionTypeMustAssign,
	PermissionTreatmentOptionForMedicalRecord: PermissionTypeMustAssignAndOnSite,

	// unit
	PermissionUnitOptionForProductTransferForm: PermissionTypeMustAssignAndOnSite,

	// user
	PermissionUserList:                             PermissionTypeMustAssignAndOnSite,
	PermissionUserShow:                             PermissionTypeMustAssignAndOnSite,
	PermissionUserOptionForTaskForm:                PermissionTypeGlobal,
	PermissionUserOptionForVoidSalesInvoiceForm:    PermissionTypeMustAssignAndOnSiteStrict,
	PermissionUserOptionForBeginCashierSessionForm: PermissionTypeMustAssignAndOnSiteStrict,

	// vaccination
	PermissionVaccinationCreate: PermissionTypeMustAssignAndOnSite,
	PermissionVaccinationList:   PermissionTypeMustAssignAndOnSite,
	PermissionVaccinationShow:   PermissionTypeMustAssignAndOnSite,
	PermissionVaccinationUpdate: PermissionTypeMustAssignAndOnSite,
	PermissionVaccinationDelete: PermissionTypeMustAssignAndOnSite,

	// virtual therapeutic moiety
	PermissionVirtualTherapeuticMoietyOptionForPatientMedicineAllergy: PermissionTypeMustAssignAndOnSite,
	PermissionVirtualTherapeuticMoietyOptionForPatientMedicineHistory: PermissionTypeMustAssignAndOnSite,
}

// Don't use this for use case
func (p Permission) PermissionType() PermissionType {
	permissionType, exist := permissionTypeByPermission[p]
	if !exist {
		panic(fmt.Errorf("Permission %s is not registered in permissionTypeByPermission", p))
	}

	return permissionType
}
