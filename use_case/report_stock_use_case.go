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
	ReportStockExcelSheet1Name string = "Product Stock"
	ReportStockExcelSheet2Name string = "Stock Mutation"
)

type ReportStockExcelSheet1Data struct {
	ProductId           string
	ProductName         string
	BaseUnit            string
	CurrentSellingPrice float64
	IsActive            bool
	StockLeft           float64
}

type ReportStockExcelSheet2Data struct {
	ProductId     string
	UnitId        string
	ProductName   string
	UnitName      string
	MutationType  string
	Qty           float64
	ScaleToBase   float64
	BaseQty       float64
	BaseQtyLeft   float64
	BaseQtySold   float64
	BaseCostPrice float64
	MutatedAt     time.Time
}

type ReportStockExcel struct {
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

func (u *ReportStockExcel) initSheet1(
	dateTime data_type.DateTime,
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

	if err = FreezeRow(excelFile, ReportStockExcelSheet1Name, 4); err != nil {
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
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "D", "D", DefaultColWidthInchToExcelizeNumber(1.7)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "E", "E", DefaultColWidthInchToExcelizeNumber(1.27)); err != nil {
		return
	}
	if err = excelFile.SetColWidth(ReportStockExcelSheet1Name, "F", "F", DefaultColWidthInchToExcelizeNumber(1.21)); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(
		ReportStockExcelSheet1Name,
		"B1",
		"B1",
		u.sheet1Data2StyleId,
	); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportStockExcelSheet1Name,
		"A1",
		&[]interface{}{
			"Time",
			dateTime.Time(),
		},
	); err != nil {
		return
	}

	if err = SetHeaderSeparator(u.excelFile, ReportStockExcelSheet1Name, "A2", 13); err != nil {
		return
	}

	if err = excelFile.SetCellStyle(ReportStockExcelSheet1Name, "A4", "F4", u.sheet1DataHeader1StyleId); err != nil {
		return
	}

	if err = excelFile.SetSheetRow(
		ReportStockExcelSheet1Name,
		"A4",
		&[]interface{}{
			"Product Id",
			"Product Name",
			"Base Unit",
			"Current Selling Price",
			"Is Active",
			"Stock Left",
		},
	); err != nil {
		return
	}

	u.sheet1LatestDataPosY = 4

	return
}

func (u *ReportStockExcel) initSheet2(
	dateTime data_type.DateTime,
) (err error) {
	if u.excelFile == nil {
		return
	}

	if err = u.initSheet2Style(); err != nil {
		return
	}

	if _, err = u.excelFile.NewSheet(ReportStockExcelSheet2Name); err != nil {
		return
	}

	if err = SetDefaultReportExcelSheet(u.excelFile, ReportStockExcelSheet2Name); err != nil {
		return
	}

	if err = FreezeRow(u.excelFile, ReportStockExcelSheet2Name, 4); err != nil {
		return
	}

	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "A", "A", DefaultColWidthInchToExcelizeNumber(1.7)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "B", "B", DefaultColWidthInchToExcelizeNumber(1.7)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "C", "C", DefaultColWidthInchToExcelizeNumber(1.5)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "D", "D", DefaultColWidthInchToExcelizeNumber(1.5)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "E", "E", DefaultColWidthInchToExcelizeNumber(2)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "F", "F", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "G", "G", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "H", "H", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "I", "I", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "J", "J", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "K", "K", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
		return
	}
	if err = u.excelFile.SetColWidth(ReportStockExcelSheet2Name, "L", "L", DefaultColWidthInchToExcelizeNumber(1.44)); err != nil {
		return
	}

	if err = u.excelFile.SetCellStyle(
		ReportStockExcelSheet2Name,
		"B1",
		"B1",
		u.sheet1Data2StyleId,
	); err != nil {
		return
	}

	if err = u.excelFile.SetSheetRow(
		ReportStockExcelSheet2Name,
		"A1",
		&[]interface{}{
			"Time",
			dateTime.Time(),
		},
	); err != nil {
		return
	}

	if err = SetHeaderSeparator(u.excelFile, ReportStockExcelSheet2Name, "A2", 13); err != nil {
		return
	}

	if err = u.excelFile.SetCellStyle(ReportStockExcelSheet2Name, "A4", "L4", u.sheet2DataHeader1StyleId); err != nil {
		return
	}

	if err = u.excelFile.SetSheetRow(
		ReportStockExcelSheet2Name,
		"A4",
		&[]interface{}{
			"Product Id",
			"Unit Id",
			"Product Name",
			"Unit Name",
			"Mutation Type",
			"Quantity",
			"Scale To Base",
			"Base Qty",
			"Base Qty Left",
			"Base Qty Sold",
			"Base Cost Price",
			"Mutated At",
		},
	); err != nil {
		return
	}

	u.sheet2LatestDataPosY = 4

	return
}

func (u *ReportStockExcel) initSheet1Style() (err error) {
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

func (u *ReportStockExcel) initSheet2Style() (err error) {
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

func (u *ReportStockExcel) Init(
	dateTime data_type.DateTime,
) (err error) {
	if err = u.initSheet1(dateTime); err != nil {
		return
	}

	if err = u.initSheet2(dateTime); err != nil {
		return
	}

	return
}

func (u *ReportStockExcel) ToReadSeekCloserWithContentLength() (io.ReadSeekCloser, int64, error) {
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

func (u *ReportStockExcel) AddSheet1Data(data ReportStockExcelSheet1Data) error {
	newLatestDataPosY := u.sheet1LatestDataPosY + 1

	if err := u.excelFile.SetCellStyle(
		ReportStockExcelSheet1Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		fmt.Sprintf("F%d", newLatestDataPosY),
		u.sheet1Data1StyleId,
	); err != nil {
		return err
	}

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

func (u *ReportStockExcel) AddSheet2Data(data ReportStockExcelSheet2Data) error {
	newLatestDataPosY := u.sheet2LatestDataPosY + 1

	if err := u.excelFile.SetCellStyle(
		ReportStockExcelSheet2Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		fmt.Sprintf("K%d", newLatestDataPosY),
		u.sheet2Data1StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetCellStyle(
		ReportStockExcelSheet2Name,
		fmt.Sprintf("L%d", newLatestDataPosY),
		fmt.Sprintf("L%d", newLatestDataPosY),
		u.sheet1Data2StyleId,
	); err != nil {
		return err
	}

	if err := u.excelFile.SetSheetRow(
		ReportStockExcelSheet2Name,
		fmt.Sprintf("A%d", newLatestDataPosY),
		&[]interface{}{
			data.ProductId,
			data.UnitId,
			data.ProductName,
			data.UnitName,
			data.MutationType,
			data.Qty,
			data.ScaleToBase,
			data.BaseQty,
			data.BaseQtyLeft,
			data.BaseQtySold,
			data.BaseCostPrice,
			data.MutatedAt,
		},
	); err != nil {
		return err
	}

	u.sheet2LatestDataPosY = newLatestDataPosY

	return nil
}

func (u *ReportStockExcel) Close() error {
	if u.excelFile != nil {
		return u.excelFile.Close()
	}

	return nil
}

func NewReportStockExcel(
	exportedDateTime data_type.DateTime,
) (reportExcel *ReportStockExcel, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("NewReportStockExcel: %w", err)

			if reportExcel != nil {
				// make sure to close it before set to nil pointer
				reportExcel.Close()

				// set reportExcel to nil pointer before return
				reportExcel = nil
			}
		}
	}()

	reportExcel = &ReportStockExcel{}

	reportExcel.excelFile, err = NewDefaultReportExcelFile()
	if err != nil {
		return
	}

	if err = reportExcel.Init(exportedDateTime); err != nil {
		return
	}

	return
}
