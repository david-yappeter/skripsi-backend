package use_case

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"myapp/constant"
	"myapp/data_type"
	"myapp/internal/filesystem"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"github.com/xuri/excelize/v2"
)

const (
	ReportStockExcelSheet1Name string = "Product Stock"
	ReportStockExcelSheet2Name string = "Stock Mutation"
)

type ReportDailyTransactionExcelSheet1Data struct {
	ProductId           string
	ProductName         string
	BaseUnit            string
	CurrentSellingPrice float64
	IsActive            bool
	StockLeft           float64

	// SalesType       string
	// InvoiceId       string
	// InvoiceDateTime time.Time
	// SalesAmount     float64
	// Voucher         float64
	// InsuranceName   *string
	// InsurancePayor  *string
	// CardNumber      *string
	// PaidByInsurance float64
	// PaidByPatient   float64
	// TransactionId   *string
	// PaymentMethod   *string
	// Mid             *string
	// ReferenceNumber *string
}

type ReportDailyTransactionExcelSheet2Data struct {
	PaymentMethod      string
	TotalSalesAmount   float64
	TotalVoucher       float64
	TotalPaymentAmount float64
}

type ReportDailyTransactionExcel struct {
	excelFile *excelize.File

	// Font:
	// 	- Bold : true
	// 	- Size : 15
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet1Header1StyleId int
	// Font:
	// 	- Bold : true
	// 	- Size : 10
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet1Header2StyleId int

	// Font:
	// 	- Bold : true
	// 	- Size : 9
	//
	// Alignment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet1DataHeader1StyleId int
	// Font:
	// 	- Bold: true
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "right"
	// 	- Vertical   : "center"
	sheet1DataHeader2StyleId int

	// Font:
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet1Data1StyleId int
	// Font:
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	//
	// CustomNumFmt:
	// 	- Value: "D MMM YYYY H:MM:SS"
	sheet1Data2StyleId int
	// Font:
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	//
	// CustomNumFmt: _(\Rp* #,##0.00_);_(\Rp* (#,##0.00);_(\Rp* "-"??_);_(@_)
	sheet1Data3StyleId int

	sheet1LatestDataPosY int

	// Font:
	// 	- Bold : true
	// 	- Size : 15
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet2Header1StyleId int
	// Font:
	// 	- Bold : true
	// 	- Size : 10
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet2Header2StyleId int

	// Font:
	// 	- Bold : true
	// 	- Size : 9
	//
	// Alignment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet2DataHeader1StyleId int
	// Font:
	// 	- Bold: true
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "right"
	// 	- Vertical   : "center"
	sheet2DataHeader2StyleId int

	// Font:
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	sheet2Data1StyleId int
	// Font:
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "left"
	// 	- Vertical   : "center"
	//
	// CustomNumFmt:
	// 	- Value: `_-* #,##0_-;-* #,##0_-;_-* "-"_-;_-@`
	sheet2Data2StyleId int

	sheet2LatestDataPosY int
}

func (u *ReportDailyTransactionExcel) initSheet1(
	date data_type.Date,
) (err error) {

	if err = u.initSheet1Style(); err != nil {
		return
	}

	excelFile := u.excelFile

	if err = excelFile.SetSheetName("Sheet1", ReportStockExcelSheet1Name); err != nil {
		return
	}

	if err = SetDefaultReportExcelSheet(excelFile, ReportStockExcelSheet1Name); err != nil {
		return
	}

	if err = FreezeRow(excelFile, ReportStockExcelSheet1Name, 8); err != nil {
		return
	}

	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "A", "A", DefaultColWidthInchToExcelizeNumber(0.86)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "B", "B", DefaultColWidthInchToExcelizeNumber(1.21)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "C", "C", DefaultColWidthInchToExcelizeNumber(1.15)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "D", "D", DefaultColWidthInchToExcelizeNumber(1.21)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "E", "E", DefaultColWidthInchToExcelizeNumber(1.27)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "F", "F", DefaultColWidthInchToExcelizeNumber(1.21)); err != nil {
		return
	}

	// if err = excelFile.SetCellStyle(
	// 	ReportStockExcelSheet1Name,
	// 	"B2",
	// 	"B2",
	// 	u.sheet1Header1StyleId,
	// ); err != nil {
	// 	return
	// }
	// if err = excelFile.SetCellStyle(
	// 	ReportStockExcelSheet1Name,
	// 	"B3",
	// 	"B6",
	// 	u.sheet1Header2StyleId,
	// ); err != nil {
	// 	return
	// }

	// if err = excelFile.SetSheetCol(
	// 	ReportStockExcelSheet1Name,
	// 	"B2",
	// 	&[]interface{}{
	// 		ReportStockExcelSheet1Name,
	// 		companyName,
	// 		fmt.Sprintf("Clinic: %s", clinicName),
	// 		fmt.Sprintf("Date: %s", date.Format("02 January 2006")),
	// 	},
	// ); err != nil {
	// 	return
	// }

	// if err = SetHeaderSeparator(excelFile, ReportStockExcelSheet1Name, "A6", 14); err != nil {
	// 	return
	// }

	if err = excelFile.SetCellStyle(ReportStockExcelSheet1Name, "A8", "F8", u.sheet1DataHeader1StyleId); err != nil {
		return
	}
	// if err = excelFile.SetCellStyle(ReportStockExcelSheet1Name, "F8", "H8", u.sheet1DataHeader2StyleId); err != nil {
	// 	return
	// }
	// if err = excelFile.SetCellStyle(ReportStockExcelSheet1Name, "I8", "L8", u.sheet1DataHeader1StyleId); err != nil {
	// 	return
	// }

	if err = excelFile.SetSheetRow(
		ReportStockExcelSheet1Name,
		"A8",
		&[]interface{}{
			"Product Id",
			"Product Name",
			"Base Unit",
			"Current Selling Price",
			"Is Active",
			"Stock Left",

			// "Sales Type",
			// "Invoice ID",
			// "Invoice Datetime",
			// "Sales Amount",
			// "Voucher",
			// "Insurance Name",
			// "Insurance Payor",
			// "Card Number",
			// "Paid By Insurance",
			// "Paid By Patient",
			// "Transaction ID",
			// "Payment Method",
			// "MID",
			// "Reference Number",
		},
	); err != nil {
		return
	}

	u.sheet1LatestDataPosY = 8

	return
}

// func (u *ReportDailyTransactionExcel) initSheet2(
// 	companyName string,
// 	clinicName string,
// 	date data_type.Date,
// ) (err error) {
// 	if u.excelFile == nil {
// 		return
// 	}

// 	if err = u.initSheet2Style(); err != nil {
// 		return
// 	}

// 	if _, err = u.excelFile.NewSheet(ReportStockExcelSheet2Name); err != nil {
// 		return
// 	}

// 	if err = SetDefaultReportExcelSheet(u.excelFile, ReportStockExcelSheet2Name); err != nil {
// 		return
// 	}

// 	if err = FreezeRow(u.excelFile, ReportStockExcelSheet2Name, 8); err != nil {
// 		return
// 	}

// 	// if err = AddGwsMedikaLogoToCell(
// 	// 	u.excelFile,
// 	// 	ReportStockExcelSheet2Name,
// 	// 	"A2",
// 	// 	0.27,
// 	// 	0.31,
// 	// 	0.70,
// 	// 	0.82,
// 	// ); err != nil {
// 	// 	return
// 	// }

// 	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "A", "A", DefaultColWidthInchToExcelizeNumber(1.32)); err != nil {
// 		return
// 	}
// 	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "B", "B", DefaultColWidthInchToExcelizeNumber(1.34)); err != nil {
// 		return
// 	}
// 	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "C", "C", DefaultColWidthInchToExcelizeNumber(1.28)); err != nil {
// 		return
// 	}
// 	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "D", "D", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
// 		return
// 	}

// 	if err = u.excelFile.SetCellStyle(
// 		ReportStockExcelSheet2Name,
// 		"B2",
// 		"B2",
// 		u.sheet2Header1StyleId,
// 	); err != nil {
// 		return
// 	}
// 	if err = u.excelFile.SetCellStyle(
// 		ReportStockExcelSheet2Name,
// 		"B3",
// 		"B6",
// 		u.sheet2Header2StyleId,
// 	); err != nil {
// 		return
// 	}

// 	if err = u.excelFile.SetSheetCol(
// 		ReportStockExcelSheet2Name,
// 		"B2",
// 		&[]interface{}{
// 			ReportStockExcelSheet2Name,
// 			companyName,
// 			fmt.Sprintf("Clinic: %s", clinicName),
// 			fmt.Sprintf("Date: %s", date.Format("02 January 2006")),
// 		},
// 	); err != nil {
// 		return
// 	}

// 	if err = SetHeaderSeparator(u.excelFile, ReportStockExcelSheet2Name, "A6", 13); err != nil {
// 		return
// 	}

// 	if err = u.excelFile.SetCellStyle(ReportStockExcelSheet2Name, "A8", "A8", u.sheet2DataHeader1StyleId); err != nil {
// 		return
// 	}
// 	if err = u.excelFile.SetCellStyle(ReportStockExcelSheet2Name, "B8", "D8", u.sheet2DataHeader2StyleId); err != nil {
// 		return
// 	}

// 	if err = u.excelFile.SetSheetRow(
// 		ReportStockExcelSheet2Name,
// 		"A8",
// 		&[]interface{}{
// 			"Payment Method",
// 			"Total Sales Amount",
// 			"Total Voucher",
// 			"Total Payment Amount",
// 		},
// 	); err != nil {
// 		return
// 	}

// 	u.sheet2LatestDataPosY = 8

// 	return
// }

func (u *ReportDailyTransactionExcel) initSheet1Style() (err error) {
	if u.excelFile == nil {
		return
	}

	u.sheet1Header1StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 15,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet1Header2StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 10,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet1DataHeader1StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet1DataHeader2StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet1Data1StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet1Data2StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			CustomNumFmt: util.StringP("D MMM YYYY H:MM:SS"),
		},
	)
	if err != nil {
		return
	}

	u.sheet1Data3StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			CustomNumFmt: util.StringP(`_(\Rp* #,##0.00_);_(\Rp* (#,##0.00);_(\Rp* "-"??_);_(@_)`),
		},
	)
	if err != nil {
		return
	}

	return
}

func (u *ReportDailyTransactionExcel) initSheet2Style() (err error) {
	if u.excelFile == nil {
		return
	}

	u.sheet2Header1StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 15,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet2Header2StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 10,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet2DataHeader1StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet2DataHeader2StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet2Data1StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
		},
	)
	if err != nil {
		return
	}

	u.sheet2Data2StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			CustomNumFmt: util.StringP(`_-* #,##0_-;-* #,##0_-;_-* "-"_-;_-@`),
		},
	)
	if err != nil {
		return
	}

	return
}

func (u *ReportDailyTransactionExcel) Init(
	date data_type.Date,
) (err error) {
	if err = u.initSheet1(date); err != nil {
		return
	}

	// if err = u.initSheet2(companyName, clinicName, date); err != nil {
	// 	return
	// }

	return
}

func (u *ReportDailyTransactionExcel) ToReadSeekCloser() (io.ReadSeekCloser, error) {
	reader := bytes.NewReader(nil)
	if u.excelFile != nil {
		bytesBuffer, err := u.excelFile.WriteToBuffer()
		if err != nil {
			return nil, err
		}

		reader = bytes.NewReader(bytesBuffer.Bytes())
	}

	return util.ReadSeekNopCloser(reader), nil
}

func (u *ReportDailyTransactionExcel) AddSheet1Data(data ReportDailyTransactionExcelSheet1Data) error {
	newLatestDataPosY := u.sheet1LatestDataPosY + 1

	if err := u.excelFile.SetCellStyle(
		ReportStockExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		fmt.Sprintf("F%d", newLatestDataPosY),
		u.sheet1Data1StyleId,
	); err != nil {
		return err
	}
	// if err := u.excelFile.SetCellStyle(
	// 	ReportStockExcelSheet1Name,
	// 	fmt.Sprintf("C%d", newLatestDataPosY),
	// 	fmt.Sprintf("C%d", newLatestDataPosY),
	// 	u.sheet1Data2StyleId,
	// ); err != nil {
	// 	return err
	// }
	// if err := u.excelFile.SetCellStyle(
	// 	ReportStockExcelSheet1Name,
	// 	fmt.Sprintf("D%d", newLatestDataPosY),
	// 	fmt.Sprintf("E%d", newLatestDataPosY),
	// 	u.sheet1Data3StyleId,
	// ); err != nil {
	// 	return err
	// }
	// if err := u.excelFile.SetCellStyle(
	// 	ReportStockExcelSheet1Name,
	// 	fmt.Sprintf("F%d", newLatestDataPosY),
	// 	fmt.Sprintf("H%d", newLatestDataPosY),
	// 	u.sheet1Data1StyleId,
	// ); err != nil {
	// 	return err
	// }
	// if err := u.excelFile.SetCellStyle(
	// 	ReportStockExcelSheet1Name,
	// 	fmt.Sprintf("I%d", newLatestDataPosY),
	// 	fmt.Sprintf("J%d", newLatestDataPosY),
	// 	u.sheet1Data3StyleId,
	// ); err != nil {
	// 	return err
	// }
	// if err := u.excelFile.SetCellStyle(
	// 	ReportStockExcelSheet1Name,
	// 	fmt.Sprintf("K%d", newLatestDataPosY),
	// 	fmt.Sprintf("N%d", newLatestDataPosY),
	// 	u.sheet1Data1StyleId,
	// ); err != nil {
	// 	return err
	// }

	// var (
	// 	insuranceName   interface{} = "-"
	// 	insurancePayor  interface{} = "-"
	// 	cardNumber      interface{} = "-"
	// 	transactionId   interface{} = "-"
	// 	paymentMethod   interface{} = "-"
	// 	mid             interface{} = "-"
	// 	referenceNumber interface{} = "-"
	// )

	// if data.InsuranceName != nil {
	// 	insuranceName = *data.InsuranceName
	// }

	// if data.InsurancePayor != nil {
	// 	insurancePayor = *data.InsurancePayor
	// }

	// if data.CardNumber != nil {
	// 	cardNumber = *data.CardNumber
	// }

	// if data.TransactionId != nil {
	// 	transactionId = *data.TransactionId
	// }

	// if data.PaymentMethod != nil {
	// 	paymentMethod = *data.PaymentMethod
	// }

	// if data.Mid != nil {
	// 	mid = *data.Mid
	// }

	// if data.ReferenceNumber != nil {
	// 	referenceNumber = *data.ReferenceNumber
	// }

	if err := u.excelFile.SetSheetRow(
		ReportStockExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		&[]interface{}{
			data.ProductId,
			data.ProductName,
			data.BaseUnit,
			data.CurrentSellingPrice,
			data.IsActive,
			data.StockLeft,
		},
	); err != nil {
		return err
	}

	u.sheet1LatestDataPosY = newLatestDataPosY

	return nil
}

func (u *ReportDailyTransactionExcel) AddSheet2Data(data ReportDailyTransactionExcelSheet2Data) error {
	newLatestDataPosY := u.sheet2LatestDataPosY + 1

	if err := u.excelFile.SetCellStyle(
		ReportStockExcelSheet2Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		fmt.Sprintf("A%d", newLatestDataPosY),
		u.sheet2Data1StyleId,
	); err != nil {
		return err
	}
	if err := u.excelFile.SetCellStyle(
		ReportStockExcelSheet2Name,
		fmt.Sprintf("B%d", newLatestDataPosY),
		fmt.Sprintf("D%d", newLatestDataPosY),
		u.sheet2Data2StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetSheetRow(
		ReportStockExcelSheet2Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		&[]interface{}{
			data.PaymentMethod,
			data.TotalSalesAmount,
			data.TotalVoucher,
			data.TotalPaymentAmount,
		},
	); err != nil {
		return err
	}

	u.sheet2LatestDataPosY = newLatestDataPosY

	return nil
}

func (u *ReportDailyTransactionExcel) Close() error {
	if u.excelFile != nil {
		return u.excelFile.Close()
	}

	return nil
}

func NewReportDailyTransactionExcel(
	exportedDateTime data_type.DateTime,
) (reportExcel *ReportDailyTransactionExcel, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("NewReportDailyTransactionExcel: %w", err)

			if reportExcel != nil {
				// make sure to close it before set to nil pointer
				reportExcel.Close()

				// set reportExcel to nil pointer before return
				reportExcel = nil
			}
		}
	}()

	reportExcel = &ReportDailyTransactionExcel{}

	reportExcel.excelFile, err = NewDefaultReportExcelFile()
	if err != nil {
		return
	}

	if err = reportExcel.Init(exportedDateTime.Date()); err != nil {
		return
	}

	return
}

type ReportDailyTransactionUseCase interface {
	// create
	// Generate(ctx context.Context, clinicId string, date data_type.Date) (*model.Report, error)

	// read
	// Fetch(ctx context.Context, request dto_request.ReportDailyTransactionFetchRequest) ([]model.Report, int)
	// Download(ctx context.Context, request dto_request.ReportDailyTransactionDownloadRequest) (
	// 	ioReadCloser io.ReadCloser,
	// 	contentLength int64,
	// 	contentType string,
	// 	filename string,
	// )
}

type reportDailyTransactionUseCase struct {
	repositoryManager repository.RepositoryManager

	mainFilesystem filesystem.Client
}

// func NewReportDailyTransactionUseCase(
// 	repositoryManager repository.RepositoryManager,
// 	queueManager queue.QueueManager,

// 	mainFilesystem filesystem.Client,
// ) ReportDailyTransactionUseCase {
// 	return &reportDailyTransactionUseCase{
// 		repositoryManager: repositoryManager,
// 		queueManager:      queueManager,

// 		mainFilesystem: mainFilesystem,
// 	}
// }

// func (u *reportDailyTransactionUseCase) Generate(
// 	ctx context.Context,
// 	clinicId string,
// 	date data_type.Date,
// ) (*model.Report, error) {
// 	var (
// 		authorizedUser, _ = model.GetUserCtx(ctx)

// 		requestedUserId *string = nil
// 		isScheduled     bool    = false

// 		clinicRepository = u.repositoryManager.ClinicRepository()
// 		reportRepository = u.repositoryManager.ReportRepository()

// 		reportQueue = u.queueManager.ReportQueue()
// 	)

// 	if authorizedUser != nil {
// 		requestedUserId = &authorizedUser.Id
// 		isScheduled = true
// 	}

// 	if _, err := clinicRepository.GetById(ctx, clinicId); err != nil {
// 		if err == constant.ErrNoData {
// 			return nil, fmt.Errorf("clinic [id: %s]: %w", clinicId, err)
// 		}

// 		return nil, err
// 	}

// 	report := model.Report{
// 		Id:              util.NewUuid(),
// 		ClinicId:        &clinicId,
// 		CompanyId:       nil,
// 		RequestedUserId: requestedUserId,
// 		FileId:          nil,
// 		Type:            data_type.ReportTypeDailyTransaction,
// 		ReportStartTime: date.DateTimeStartOfDay().NullDateTime(),
// 		ReportEndTime:   date.DateTimeEndOfDay().NullDateTime(),
// 		Status:          data_type.ReportStatusPending,
// 		IsScheduled:     isScheduled,
// 		RequestedAt:     util.CurrentNullDateTime(),
// 		ProcessedAt:     data_type.NewNullDateTime(nil),
// 		CompletedAt:     data_type.NewNullDateTime(nil),
// 	}

// 	if err := reportRepository.Insert(ctx, &report); err != nil {
// 		return nil, fmt.Errorf("insert report: %w", err)
// 	}

// 	if err := reportQueue.Publish(
// 		ctx,
// 		model_queue.Report{
// 			Id: report.Id,
// 		},
// 	); err != nil {
// 		if err := UpdateReportWhenPublishQueueFailed(ctx, reportRepository, &report, err); err != nil {
// 			return nil, fmt.Errorf("update report: %w", err)
// 		}

// 		return nil, fmt.Errorf("publish report: %w", err)
// 	}

// 	return &report, nil
// }

// func (u *reportDailyTransactionUseCase) Fetch(ctx context.Context, request dto_request.ReportDailyTransactionFetchRequest) ([]model.Report, int) {
// 	var (
// 		activeClinicId = model.MustGetActiveClinicIdCtx(ctx)

// 		reportRepository  = u.repositoryManager.ReportRepository()
// 		clinicRepository  = u.repositoryManager.ClinicRepository()
// 		companyRepository = u.repositoryManager.CompanyRepository()
// 		userRepository    = u.repositoryManager.UserRepository()
// 		fileRepository    = u.repositoryManager.FileRepository()

// 		queryOption = model.ReportQueryOption{
// 			QueryOption: model.QueryOption{
// 				Page:  request.Page,
// 				Limit: request.Limit,
// 				Sorts: model.Sorts(request.Sorts),
// 			},
// 			ClinicId: &activeClinicId,
// 			Type:     data_type.ReportTypeP(data_type.ReportTypeDailyTransaction),
// 		}
// 	)

// 	reports, err := reportRepository.Fetch(ctx, queryOption)
// 	panicIfErr(err)

// 	count, err := reportRepository.Count(ctx, queryOption)
// 	panicIfErr(err)

// 	mustLoadReportsData(
// 		ctx,
// 		clinicRepository,
// 		companyRepository,
// 		userRepository,
// 		fileRepository,
// 		u.mainFilesystem,
// 		reports,
// 	)

// 	return reports, count
// }

// func (u *reportDailyTransactionUseCase) Download(ctx context.Context, request dto_request.ReportDailyTransactionDownloadRequest) (io.ReadCloser, int64, string, string) {
// 	var (
// 		activeClinicId = model.MustGetActiveClinicIdCtx(ctx)

// 		fileRepository   = u.repositoryManager.FileRepository()
// 		reportRepository = u.repositoryManager.ReportRepository()

// 		report = mustGetReportPWithinClinic(ctx, reportRepository, request.Id, data_type.ReportTypeDailyTransaction, activeClinicId, false)
// 	)

// 	mustValidateAllowDownload(ctx, report)

// 	file := mustGetFileP(ctx, fileRepository, *report.FileId, true)

// 	reader, err := u.mainFilesystem.Stream(ctx, file.Path)
// 	if err != nil {
// 		if err == filesystem.ErrFileNotExist {
// 			panic(dto_response.NewBadRequestErrorResponse("Report file not found"))
// 		}

// 		panic(err)
// 	}

// 	return reader, reader.ContentLength(), reader.ContentType(), file.Name
// }

func GenerateDailyTransactionReport(
	ctx context.Context,
	// report *model.Report,
	repositoryManager repository.RepositoryManager,
	mainFilesystem filesystem.Client,
) (file *model.File, err error) {
	// var (
	// 	clinicRepository              repository.ClinicRepository              = repositoryManager.ClinicRepository()
	// 	companyRepository             repository.CompanyRepository             = repositoryManager.CompanyRepository()
	// 	salesInvoiceRepository        repository.SalesInvoiceRepository        = repositoryManager.SalesInvoiceRepository()
	// 	salesPaymentRepository        repository.SalesPaymentRepository        = repositoryManager.SalesPaymentRepository()
	// 	salesPaymentMethodRepository  repository.SalesPaymentMethodRepository  = repositoryManager.SalesPaymentMethodRepository()
	// 	salesPaymentOttopayRepository repository.SalesPaymentOttopayRepository = repositoryManager.SalesPaymentOttopayRepository()
	// 	salesPaymentMandiriRepository repository.SalesPaymentMandiriRepository = repositoryManager.SalesPaymentMandiriRepository()
	// 	salesPaymentManualRepository  repository.SalesPaymentManualRepository  = repositoryManager.SalesPaymentManualRepository()
	// 	insuranceRepository           repository.InsuranceRepository           = repositoryManager.InsuranceRepository()
	// 	insurancePayorRepository      repository.InsurancePayorRepository      = repositoryManager.InsurancePayorRepository()
	// )

	// defer func() {
	// 	if err != nil {
	// 		reportErrs := ReportErrs{
	// 			message: "generateDailyTransactionReport",
	// 			errs: []error{
	// 				err,
	// 			},
	// 		}

	// 		if file != nil {
	// 			if deleteErr := mainFilesystem.Delete(file.Path); deleteErr != nil {
	// 				reportErrs.errs = append(reportErrs.errs, deleteErr)
	// 			}

	// 			file = nil
	// 		}

	// 		err = reportErrs
	// 	}
	// }()

	// if report.Type != data_type.ReportTypeDailyTransaction {
	// 	err = errors.New("type is not daily transaction")
	// 	return
	// }

	// if report.ClinicId == nil {
	// 	err = errors.New("clinic id is empty")
	// 	return
	// }

	// if report.ReportStartTime.DateTimeP() == nil {
	// 	err = errors.New("report start time is empty")
	// 	return
	// }

	// if report.ReportEndTime.DateTimeP() == nil {
	// 	err = errors.New("report end time is empty")
	// 	return
	// }

	// clinic, err := clinicRepository.GetById(ctx, *report.ClinicId)
	// if err != nil {
	// 	if err == constant.ErrNoData {
	// 		err = errors.New("clinic data not found")
	// 	}

	// 	return
	// }

	// company, err := companyRepository.GetById(ctx, clinic.CompanyId)
	// if err != nil {
	// 	if err == constant.ErrNoData {
	// 		err = errors.New("company data not found")
	// 	}

	// 	return
	// }

	// reportStartTime := report.ReportStartTime
	// reportEndTime := report.ReportEndTime

	// // fetch sales paid sales invoices by date
	// queryOption := model.SalesInvoiceQueryOption{
	// 	QueryOption: model.QueryOption{
	// 		Sorts: model.Sorts{
	// 			{Field: "invoice_number", Direction: "asc"},
	// 		},
	// 	},
	// 	ClinicIds: []string{clinic.Id},
	// 	Statuses: []data_type.SalesInvoiceStatus{
	// 		data_type.SalesInvoiceStatusPaid,
	// 	},
	// 	GeneratedAtStart: reportStartTime,
	// 	GeneratedAtEnd:   reportEndTime,
	// }
	// salesInvoices, err := salesInvoiceRepository.Fetch(ctx, queryOption)
	// if err != nil {
	// 	return
	// }

	// // load sales invoices data
	// var (
	// 	salesPaymentLoader        = loader.NewSalesPaymentLoader(salesPaymentRepository)
	// 	salesPaymentMethodLoader  = loader.NewSalesPaymentMethodLoader(salesPaymentMethodRepository)
	// 	salesPaymentOttopayLoader = loader.NewSalesPaymentOttopayLoader(salesPaymentOttopayRepository)
	// 	salesPaymentMandiriLoader = loader.NewSalesPaymentMandiriLoader(salesPaymentMandiriRepository)
	// 	salesPaymentManualLoader  = loader.NewSalesPaymentManualLoader(salesPaymentManualRepository)
	// 	insuranceLoader           = loader.NewInsuranceLoader(insuranceRepository)
	// 	insurancePayorLoader      = loader.NewInsurancePayorLoader(insurancePayorRepository)
	// )

	// if err = util.Await(func(group *errgroup.Group) {
	// 	for i := range salesInvoices {
	// 		group.Go(salesPaymentLoader.SalesInvoiceNotStrictFn(&salesInvoices[i]))
	// 		group.Go(insuranceLoader.SalesInvoiceFn(&salesInvoices[i]))
	// 		group.Go(insurancePayorLoader.SalesInvoiceFn(&salesInvoices[i]))
	// 	}
	// }); err != nil {
	// 	return
	// }

	// if err = util.Await(func(group *errgroup.Group) {
	// 	for i := range salesInvoices {
	// 		if salesInvoices[i].SalesPayment != nil {
	// 			group.Go(salesPaymentMethodLoader.SalesPaymentFn(salesInvoices[i].SalesPayment))
	// 		}
	// 	}
	// }); err != nil {
	// 	return
	// }

	// if err = util.Await(func(group *errgroup.Group) {
	// 	for i := range salesInvoices {
	// 		if salesInvoices[i].SalesPayment != nil {
	// 			switch salesInvoices[i].SalesPayment.SalesPaymentMethod.Name {
	// 			case data_type.SalesPaymentMethodOttopayQr,
	// 				data_type.SalesPaymentMethodOttopayDebitCard,
	// 				data_type.SalesPaymentMethodOttopayCreditCard:
	// 				group.Go(salesPaymentOttopayLoader.SalesPaymentFn(salesInvoices[i].SalesPayment))

	// 			case data_type.SalesPaymentMethodMandiriCreditCard,
	// 				data_type.SalesPaymentMethodMandiriDebitCard:
	// 				group.Go(salesPaymentMandiriLoader.SalesPaymentFn(salesInvoices[i].SalesPayment))

	// 			case data_type.SalesPaymentMethodManualMandiriQr,
	// 				data_type.SalesPaymentMethodManualMandiriCreditCard,
	// 				data_type.SalesPaymentMethodManualMandiriDebitCard,
	// 				data_type.SalesPaymentMethodManualOttopayQr,
	// 				data_type.SalesPaymentMethodManualOttopayDebitCard,
	// 				data_type.SalesPaymentMethodManualOttopayCreditCard,
	// 				data_type.SalesPaymentMethodManualSinarmasCreditCard,
	// 				data_type.SalesPaymentMethodManualSinarmasDebitCard,
	// 				data_type.SalesPaymentMethodManualHalodoc,
	// 				data_type.SalesPaymentMethodManualGoodDoctor,
	// 				data_type.SalesPaymentMethodManualKlikDokter,
	// 				data_type.SalesPaymentMethodManualKredivo,
	// 				data_type.SalesPaymentMethodManualXendit:
	// 				group.Go(salesPaymentManualLoader.SalesPaymentFn(salesInvoices[i].SalesPayment))
	// 			}
	// 		}
	// 	}
	// }); err != nil {
	// 	return
	// }

	// construct excel report sheets
	reportExcel, err := NewReportDailyTransactionExcel(
		// company.Name,
		// clinic.Name,
		// reportEndTime.DateTime(),
		util.CurrentDateTime(),
	)
	if err != nil {
		return
	}
	defer reportExcel.Close()

	// for _, salesInvoice := range salesInvoices {
	// 	// calculate remaining parameter sheet 1
	// 	var (
	// 		insuranceName  *string
	// 		insurancePayor *string

	// 		paidByPatient   float64
	// 		transactionId   *string
	// 		paymentMethod   *string
	// 		mid             *string
	// 		referenceNumber *string
	// 	)

	// 	if salesInvoice.Insurance != nil {
	// 		insuranceName = &salesInvoice.Insurance.Name
	// 	}

	// 	if salesInvoice.InsurancePayor != nil {
	// 		insurancePayor = &salesInvoice.InsurancePayor.Name
	// 	}

	// 	if salesPayment := salesInvoice.SalesPayment; salesPayment != nil {
	// 		paidByPatient = salesPayment.Total
	// 		transactionId = &salesPayment.PaymentNumber

	// 		if salesPayment.SalesPaymentMethod != nil {
	// 			paymentMethod = &salesPayment.SalesPaymentMethod.DisplayName
	// 		}

	// 		if salesPayment.SalesPaymentOttopay != nil {
	// 			mid = &salesPayment.SalesPaymentOttopay.Mid
	// 		}

	// 		if salesPayment.SalesPaymentMandiri != nil {
	// 			mid = &salesPayment.SalesPaymentMandiri.MerchantId
	// 		}

	// 		if salesPayment.SalesPaymentManual != nil {
	// 			referenceNumber = &salesPayment.SalesPaymentManual.ReferenceNumber
	// 		}
	// 	}

	// 	sheetParameters := ReportDailyTransactionExcelSheet1Data{
	// 		SalesType:       salesInvoice.Type.String(),
	// 		InvoiceId:       salesInvoice.InvoiceNumber,
	// 		InvoiceDateTime: salesInvoice.CreatedAt.DateTime().Time(),
	// 		SalesAmount:     salesInvoice.Subtotal,
	// 		Voucher:         salesInvoice.DiscountAmount,
	// 		InsuranceName:   insuranceName,
	// 		InsurancePayor:  insurancePayor,
	// 		CardNumber:      salesInvoice.InsuranceCardNumber,
	// 		PaidByInsurance: salesInvoice.InsuranceAmount,
	// 		PaidByPatient:   paidByPatient,
	// 		TransactionId:   transactionId,
	// 		PaymentMethod:   paymentMethod,
	// 		Mid:             mid,
	// 		ReferenceNumber: referenceNumber,
	// 	}

	// 	reportExcel.AddSheet1Data(sheetParameters)
	// }

	// calculate and group payment method assign parameter
	// var sheet2DataByPaymentMethodDisplayName = map[string]*ReportDailyTransactionExcelSheet2Data{}
	// for _, salesInvoice := range salesInvoices {
	// 	salesPayment := salesInvoice.SalesPayment
	// 	if salesPayment == nil {
	// 		continue
	// 	}

	// 	paymentMethod := salesPayment.SalesPaymentMethod.DisplayName
	// 	if sheet2DataByPaymentMethodDisplayName[paymentMethod] == nil {
	// 		sheet2DataByPaymentMethodDisplayName[paymentMethod] = &ReportDailyTransactionExcelSheet2Data{
	// 			PaymentMethod:      paymentMethod,
	// 			TotalSalesAmount:   0,
	// 			TotalVoucher:       0,
	// 			TotalPaymentAmount: 0,
	// 		}
	// 	}

	// 	currentPaymentMethod := sheet2DataByPaymentMethodDisplayName[paymentMethod]

	// 	currentPaymentMethod.TotalSalesAmount += salesInvoice.Subtotal
	// 	currentPaymentMethod.TotalVoucher += salesInvoice.DiscountAmount
	// 	currentPaymentMethod.TotalPaymentAmount += salesPayment.Total
	// }

	// salesPaymentMethodDisplayNames := []string{}
	// for salesPaymentMethodDisplayName := range sheet2DataByPaymentMethodDisplayName {
	// 	salesPaymentMethodDisplayNames = append(salesPaymentMethodDisplayNames, salesPaymentMethodDisplayName)
	// }

	// // sort asc
	// sort.Slice(salesPaymentMethodDisplayNames, func(i, j int) bool {
	// 	return salesPaymentMethodDisplayNames[i] < salesPaymentMethodDisplayNames[j]
	// })

	// for _, salesPaymentMethodDisplayName := range salesPaymentMethodDisplayNames {
	// 	reportExcel.AddSheet2Data(*sheet2DataByPaymentMethodDisplayName[salesPaymentMethodDisplayName])
	// }

	readSeekCloser, err := reportExcel.ToReadSeekCloser()
	if err != nil {
		return
	}

	// filename := fmt.Sprintf("%s.xlsx", report.Id)
	filename := fmt.Sprintf("%s.xlsx", "test_excel")

	file = &model.File{
		Id:   util.NewUuid(),
		Name: filename,
		Path: fmt.Sprintf("%s/%s", constant.ReportProductStockPath, filename),
		Type: data_type.FileTypeCustomerPaymentImage,
	}

	if err = mainFilesystem.Write(ctx, readSeekCloser, file.Path); err != nil {
		return
	}

	return
}
