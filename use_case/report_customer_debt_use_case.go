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
	ReportCustomerDebtExcelSheet1Name string = "CustomerDebts"
)

type ReportCustomerDebtExcelSheet1Data struct {
	Id                           string
	CustomerDebtSource           string
	CustomerDebtSourceIdentifier string
	Status                       string
	Amount                       float64
	RemainingAmount              float64
	CustomerId                   string // G
	CustomerName                 string // H
	CreatedAt                    time.Time
}

type ReportCustomerDebtExcel struct {
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
	// 	- Bold: true
	// 	- Size: 9
	//
	// Aligment:
	// 	- Horizontal : "right"
	// 	- Vertical   : "center"
	//
	// CustomNumFmt:
	// 	- Value: "D MMM YYYY H:MM:SS"
	sheet1DataHeader3StyleId int

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

func (u *ReportCustomerDebtExcel) initSheet1(
	exportedDateTime data_type.DateTime,
	startDate data_type.Date,
	endDate data_type.Date,
	customer *model.Customer,
) (err error) {

	if err = u.initSheet1Style(); err != nil {
		return
	}

	excelFile := u.excelFile

	if err = excelFile.SetSheetName("Sheet1", ReportCustomerDebtExcelSheet1Name); err != nil {
		return
	}

	if err = SetDefaultReportExcelSheet(excelFile, ReportCustomerDebtExcelSheet1Name); err != nil {
		return
	}

	if err = FreezeRow(excelFile, ReportCustomerDebtExcelSheet1Name, 8); err != nil {
		return
	}

	if err = excelFile.SetColWidth(ReportCustomerDebtExcelSheet1Name, "A", "A", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportCustomerDebtExcelSheet1Name, "B", "B", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportCustomerDebtExcelSheet1Name, "C", "C", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportCustomerDebtExcelSheet1Name, "D", "D", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	for _, v := range []string{"E", "F", "G", "H", "I"} {
		if err = excelFile.SetColWidth(ReportCustomerDebtExcelSheet1Name, v, v, DefaultColWidthInchToExcelizeNumber(1.6)); err != nil {
			return
		}
	}

	if err = excelFile.SetCellStyle(
		ReportCustomerDebtExcelSheet1Name,
		"B1",
		"B1",
		u.sheet1Data2StyleId,
	); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(
		ReportCustomerDebtExcelSheet1Name,
		"B2",
		"B3",
		u.sheet1DataHeader3StyleId,
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportCustomerDebtExcelSheet1Name,
		"A1",
		&[]interface{}{
			"Exported Date Time",
			exportedDateTime.Time(),
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportCustomerDebtExcelSheet1Name,
		"A2",
		&[]interface{}{
			"Start Date",
			startDate.Time(),
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportCustomerDebtExcelSheet1Name,
		"A3",
		&[]interface{}{
			"End Date",
			endDate.Time(),
		},
	); err != nil {
		return
	}

	if err = SetHeaderSeparator(u.excelFile, ReportCustomerDebtExcelSheet1Name, "A6", 15); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(ReportCustomerDebtExcelSheet1Name, "A8", "I8", u.sheet1DataHeader1StyleId); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportCustomerDebtExcelSheet1Name,
		"A8",
		&[]interface{}{
			"Id",
			"Debt Source",
			"Debt Source Identifier",
			"Status",
			"Amount",
			"Remaining Amount",
			"Customer Id",
			"Customer Name",
			"Created At",
		},
	); err != nil {
		return
	}

	u.sheet1LatestDataPosY = 8

	return
}

func (u *ReportCustomerDebtExcel) initSheet1Style() (err error) {
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

	u.sheet1DataHeader3StyleId, err = u.excelFile.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
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

func (u *ReportCustomerDebtExcel) Init(
	exportedDateTime data_type.DateTime,
	startDate data_type.Date,
	endDate data_type.Date,
	customer *model.Customer,
) (err error) {
	if err = u.initSheet1(exportedDateTime, startDate, endDate, customer); err != nil {
		return
	}

	return
}

func (u *ReportCustomerDebtExcel) ToReadSeekCloserWithContentLength() (io.ReadSeekCloser, int64, error) {
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

func (u *ReportCustomerDebtExcel) AddSheet1Data(data ReportCustomerDebtExcelSheet1Data) error {
	newLatestDataPosY := u.sheet1LatestDataPosY + 1

	if err := u.excelFile.SetCellStyle(
		ReportCustomerDebtExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		fmt.Sprintf("D%d", newLatestDataPosY),
		u.sheet1Data1StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportCustomerDebtExcelSheet1Name,
		fmt.Sprintf("E%d", newLatestDataPosY),
		fmt.Sprintf("F%d", newLatestDataPosY),
		u.sheet1Data3StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportCustomerDebtExcelSheet1Name,
		fmt.Sprintf("G%d", newLatestDataPosY),
		fmt.Sprintf("H%d", newLatestDataPosY),
		u.sheet1Data1StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportCustomerDebtExcelSheet1Name,
		fmt.Sprintf("I%d", newLatestDataPosY),
		fmt.Sprintf("I%d", newLatestDataPosY),
		u.sheet1Data2StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetSheetRow(
		ReportCustomerDebtExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		&[]interface{}{
			data.Id,
			data.CustomerDebtSource,
			data.CustomerDebtSourceIdentifier,
			data.Status,
			data.Amount,
			data.RemainingAmount,
			data.CustomerId,
			data.CustomerName,
			data.CreatedAt,
		},
	); err != nil {
		return err
	}

	u.sheet1LatestDataPosY = newLatestDataPosY

	return nil
}

func (u *ReportCustomerDebtExcel) Close() error {
	if u.excelFile != nil {
		return u.excelFile.Close()
	}

	return nil
}

func NewReportCustomerDebtExcel(
	exportedDateTime data_type.DateTime,
	startDate data_type.Date,
	endDate data_type.Date,
	customer *model.Customer,
) (reportExcel *ReportCustomerDebtExcel, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("NewReportCustomerDebtExcel: %w", err)

			if reportExcel != nil {
				// make sure to close it before set to nil pointer
				reportExcel.Close()

				// set reportExcel to nil pointer before return
				reportExcel = nil
			}
		}
	}()

	reportExcel = &ReportCustomerDebtExcel{}

	reportExcel.excelFile, err = NewDefaultReportExcelFile()
	if err != nil {
		return
	}

	if err = reportExcel.Init(exportedDateTime, startDate, endDate, customer); err != nil {
		return
	}

	return
}
