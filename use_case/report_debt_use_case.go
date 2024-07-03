package use_case

import (
	"bytes"
	"fmt"
	"io"
	"myapp/data_type"
	"myapp/util"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	ReportDebtExcelSheet1Name string = "Debts"
)

type ReportDebtExcelSheet1Data struct {
	Id                   string
	DebtSource           string
	DebtSourceIdentifier string
	Status               string
	Amount               float64
	RemainingAmount      float64
	SupplierId           string // G
	SupplierCode         string // H
	SupplierName         string
	CreatedAt            time.Time
}

type ReportDebtExcel struct {
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

func (u *ReportDebtExcel) initSheet1(
	exportedDateTime data_type.DateTime,
	startDate data_type.Date,
	endDate data_type.Date,
) (err error) {

	if err = u.initSheet1Style(); err != nil {
		return
	}

	excelFile := u.excelFile

	if err = excelFile.SetSheetName("Sheet1", ReportDebtExcelSheet1Name); err != nil {
		return
	}

	if err = SetDefaultReportExcelSheet(excelFile, ReportDebtExcelSheet1Name); err != nil {
		return
	}

	if err = FreezeRow(excelFile, ReportDebtExcelSheet1Name, 8); err != nil {
		return
	}

	if err = excelFile.SetColWidth(ReportDebtExcelSheet1Name, "A", "A", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportDebtExcelSheet1Name, "B", "B", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportDebtExcelSheet1Name, "C", "C", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportDebtExcelSheet1Name, "D", "D", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	for _, v := range []string{"E", "F", "G", "H", "I", "J"} {
		if err = excelFile.SetColWidth(ReportDebtExcelSheet1Name, v, v, DefaultColWidthInchToExcelizeNumber(1.6)); err != nil {
			return
		}
	}

	if err = excelFile.SetCellStyle(
		ReportDebtExcelSheet1Name,
		"B1",
		"B1",
		u.sheet1Data2StyleId,
	); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(
		ReportDebtExcelSheet1Name,
		"B3",
		"B4",
		u.sheet1Data2StyleId,
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportDebtExcelSheet1Name,
		"A1",
		&[]interface{}{
			"Exported Date Time",
			exportedDateTime.Time(),
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportDebtExcelSheet1Name,
		"A2",
		&[]interface{}{
			"Start Date",
			startDate.Time(),
		},
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportDebtExcelSheet1Name,
		"A3",
		&[]interface{}{
			"End Date",
			endDate.Time(),
		},
	); err != nil {
		return
	}

	if err = SetHeaderSeparator(u.excelFile, ReportDebtExcelSheet1Name, "A6", 15); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(ReportDebtExcelSheet1Name, "A8", "J8", u.sheet1DataHeader1StyleId); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportDebtExcelSheet1Name,
		"A8",
		&[]interface{}{
			"Debt Id",
			"Debt Source",
			"Debt Source Identifier",
			"Status",
			"Amount",
			"Remaining Amount",
			"Supplier Id",
			"Supplier Code",
			"Supplier Name",
			"Created At",
		},
	); err != nil {
		return
	}

	u.sheet1LatestDataPosY = 8

	return
}

func (u *ReportDebtExcel) initSheet1Style() (err error) {
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

func (u *ReportDebtExcel) Init(
	exportedDateTime data_type.DateTime,
	startDate data_type.Date,
	endDate data_type.Date,
) (err error) {
	if err = u.initSheet1(exportedDateTime, startDate, endDate); err != nil {
		return
	}

	return
}

func (u *ReportDebtExcel) ToReadSeekCloserWithContentLength() (io.ReadSeekCloser, int64, error) {
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

func (u *ReportDebtExcel) AddSheet1Data(data ReportDebtExcelSheet1Data) error {
	newLatestDataPosY := u.sheet1LatestDataPosY + 1

	if err := u.excelFile.SetCellStyle(
		ReportDebtExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		fmt.Sprintf("D%d", newLatestDataPosY),
		u.sheet1Data1StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportDebtExcelSheet1Name,
		fmt.Sprintf("E%d", newLatestDataPosY),
		fmt.Sprintf("F%d", newLatestDataPosY),
		u.sheet1Data3StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportDebtExcelSheet1Name,
		fmt.Sprintf("G%d", newLatestDataPosY),
		fmt.Sprintf("I%d", newLatestDataPosY),
		u.sheet1Data1StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportDebtExcelSheet1Name,
		fmt.Sprintf("J%d", newLatestDataPosY),
		fmt.Sprintf("J%d", newLatestDataPosY),
		u.sheet1Data3StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetSheetRow(
		ReportDebtExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		&[]interface{}{
			data.Id,
			data.DebtSource,
			data.DebtSourceIdentifier,
			data.Status,
			data.Amount,
			data.RemainingAmount,
			data.SupplierId,
			data.SupplierCode,
			data.SupplierName,
			data.CreatedAt,
		},
	); err != nil {
		return err
	}

	u.sheet1LatestDataPosY = newLatestDataPosY

	return nil
}

func (u *ReportDebtExcel) Close() error {
	if u.excelFile != nil {
		return u.excelFile.Close()
	}

	return nil
}

func NewReportDebtExcel(
	exportedDateTime data_type.DateTime,
	startDate data_type.Date,
	endDate data_type.Date,
) (reportExcel *ReportDebtExcel, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("NewReportDebtExcel: %w", err)

			if reportExcel != nil {
				// make sure to close it before set to nil pointer
				reportExcel.Close()

				// set reportExcel to nil pointer before return
				reportExcel = nil
			}
		}
	}()

	reportExcel = &ReportDebtExcel{}

	reportExcel.excelFile, err = NewDefaultReportExcelFile()
	if err != nil {
		return
	}

	if err = reportExcel.Init(exportedDateTime, startDate, endDate); err != nil {
		return
	}

	return
}
