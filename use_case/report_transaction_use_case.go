package use_case

import (
	"bytes"
	"fmt"
	"io"
	"myapp/data_type"
	"myapp/model"
	"myapp/util"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	ReportTransactionExcelSheet1Name string = "Transactions"
)

type ReportTransactionExcelSheet1Data struct {
	Id              string
	Status          string
	PaymentMethod   string
	ProductId       string
	UnitId          string
	ProductName     string
	UnitName        string  // G
	Qty             float64 // H
	PricePerUnit    float64
	DiscountPerUnit float64
	Total           float64
	CostPerUnit     float64
	Revenue         float64
	PaymentAt       time.Time // N
}

type ReportTransactionExcel struct {
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

func (u *ReportTransactionExcel) initSheet1(
	dateTime data_type.DateTime,
	cashierSession model.CashierSession,
) (err error) {

	if err = u.initSheet1Style(); err != nil {
		return
	}

	excelFile := u.excelFile

	if err = excelFile.SetSheetName("Sheet1", ReportTransactionExcelSheet1Name); err != nil {
		return
	}

	if err = SetDefaultReportExcelSheet(excelFile, ReportTransactionExcelSheet1Name); err != nil {
		return
	}

	if err = FreezeRow(excelFile, ReportTransactionExcelSheet1Name, 8); err != nil {
		return
	}

	if err = excelFile.SetColWidth(ReportTransactionExcelSheet1Name, "A", "A", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportTransactionExcelSheet1Name, "B", "B", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportTransactionExcelSheet1Name, "C", "C", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportTransactionExcelSheet1Name, "D", "D", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	for _, v := range []string{"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"} {
		if err = excelFile.SetColWidth(ReportTransactionExcelSheet1Name, v, v, DefaultColWidthInchToExcelizeNumber(1.6)); err != nil {
			return
		}
	}

	if err = excelFile.SetCellStyle(
		ReportTransactionExcelSheet1Name,
		"B1",
		"B1",
		u.sheet1Data2StyleId,
	); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(
		ReportTransactionExcelSheet1Name,
		"B3",
		"B4",
		u.sheet1Data2StyleId,
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"A1",
		&[]interface{}{
			"Time",
			dateTime.Time(),
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"A2",
		&[]interface{}{
			"Cashier Session Id",
			cashierSession.Id,
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"A3",
		&[]interface{}{
			"Started At",
			cashierSession.StartedAt.Time(),
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"A4",
		&[]interface{}{
			"Ended At",
			cashierSession.StartedAt.Time(),
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"A5",
		&[]interface{}{
			"Starting Cash",
			cashierSession.StartingCash,
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"A6",
		&[]interface{}{
			"Ending Cash",
			*cashierSession.EndingCash,
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"C5",
		&[]interface{}{
			"User Id",
			cashierSession.User.Id,
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"C6",
		&[]interface{}{
			"User Name",
			cashierSession.User.Name,
		},
	); err != nil {
		return
	}

	if err = SetHeaderSeparator(u.excelFile, ReportTransactionExcelSheet1Name, "A6", 15); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(ReportTransactionExcelSheet1Name, "A8", "N8", u.sheet1DataHeader1StyleId); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		"A8",
		&[]interface{}{
			"Transaction Id",
			"Status",
			"Payment Method",
			"Product Id",
			"Unit Id",
			"Product Name",
			"Unit Name",
			"Qty",
			"Price Per Unit",
			"Discount Per Unit",
			"Total",
			"Cost Per Unit",
			"Revenue",
			"Payment At",
		},
	); err != nil {
		return
	}

	u.sheet1LatestDataPosY = 8

	return
}

func (u *ReportTransactionExcel) initSheet1Style() (err error) {
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

func (u *ReportTransactionExcel) Init(
	dateTime data_type.DateTime,
	cashierSession model.CashierSession,
) (err error) {
	if err = u.initSheet1(dateTime, cashierSession); err != nil {
		return
	}

	return
}

func (u *ReportTransactionExcel) ToReadSeekCloserWithContentLength() (io.ReadSeekCloser, int64, error) {
	reader := bytes.NewReader(nil)
	if u.excelFile != nil {
		bytesBuffer, err := u.excelFile.WriteToBuffer()
		if err != nil {
			return nil, 0, err
		}

		reader = bytes.NewReader(bytesBuffer.Bytes())
	}

	readSeekCloser := util.ReadSeekNopCloser(reader)

	seeker, ok := readSeekCloser.(io.Seeker)
	if !ok {
		panic("does not support seeking")
	}

	contentLength, err := seeker.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, 0, err
	}
	_, err = seeker.Seek(0, io.SeekStart)
	if err != nil {
		return nil, 0, err
	}

	return readSeekCloser, contentLength, nil
}

func (u *ReportTransactionExcel) AddSheet1Data(data ReportTransactionExcelSheet1Data) error {
	newLatestDataPosY := u.sheet1LatestDataPosY + 1

	if err := u.excelFile.SetCellStyle(
		ReportTransactionExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		fmt.Sprintf("M%d", newLatestDataPosY),
		u.sheet1Data1StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportTransactionExcelSheet1Name,
		fmt.Sprintf("N%d", newLatestDataPosY),
		fmt.Sprintf("N%d", newLatestDataPosY),
		u.sheet1Data2StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetSheetRow(
		ReportTransactionExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		&[]interface{}{
			data.Id,
			data.Status,
			data.PaymentMethod,
			data.ProductId,
			data.UnitId,
			data.ProductName,
			data.UnitName,
			data.Qty,
			data.PricePerUnit,
			data.DiscountPerUnit,
			data.Total,
			data.CostPerUnit,
			data.Revenue,
			data.PaymentAt,
		},
	); err != nil {
		return err
	}

	u.sheet1LatestDataPosY = newLatestDataPosY

	return nil
}

func (u *ReportTransactionExcel) Close() error {
	if u.excelFile != nil {
		return u.excelFile.Close()
	}

	return nil
}

func NewReportTransactionExcel(
	exportedDateTime data_type.DateTime,
	cashierSession model.CashierSession,
) (reportExcel *ReportTransactionExcel, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("NewReportTransactionExcel: %w", err)

			if reportExcel != nil {
				// make sure to close it before set to nil pointer
				reportExcel.Close()

				// set reportExcel to nil pointer before return
				reportExcel = nil
			}
		}
	}()

	reportExcel = &ReportTransactionExcel{}

	reportExcel.excelFile, err = NewDefaultReportExcelFile()
	if err != nil {
		return
	}

	if err = reportExcel.Init(exportedDateTime, cashierSession); err != nil {
		return
	}

	return
}
