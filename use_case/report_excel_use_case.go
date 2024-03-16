package use_case

import (
	"errors"
	"fmt"
	"log"
	"math"
	"myapp/util"

	"github.com/xuri/excelize/v2"
)

const (
	defaultGwsMedicaLogoWidthInInch  float64 = 2.15
	defaultGwsMedicaLogoHeightInInch float64 = 2.67 // Consider the column height is 0.22"

	minColumnPrefix string = "A"
	maxColumnPrefix string = "XFD"

	defaultColWidthInInch  float64 = 1.3
	defaultRowHeightInInch float64 = 0.22

	defaultColWidthScale  float64 = 11.808787277
	defaultRowHeightScale float64 = 71.428571429

	offsetXYScale float64 = 96.296296296
)

var (
	ErrDefaultRowHeightIsNotEqual = fmt.Errorf("default row height is not equal to %f", defaultRowHeightInInch)
	ErrDefaultColWidthIsNilOrZero = errors.New("default column width is a nil pointer or zero")
	ErrPositionXIsOutOfCellRange  = errors.New("position x is out of cell")
	ErrPositionYIsOutOfCellRange  = errors.New("position y is out of cell")
)

func DefaultColWidthInchToExcelizeNumber(inch float64) float64 {
	return inch * defaultColWidthScale
}

func DefaultColWidthExcelizeNumberToInch(excelizeNumber float64) float64 {
	return excelizeNumber / defaultColWidthScale
}

func DefaultRowHeightInchToExcelizeNumber(inch float64) float64 {
	return inch * defaultRowHeightScale
}

func DefaultRowHeightExcelizeNumberToInch(excelizeNumber float64) float64 {
	return excelizeNumber / defaultRowHeightScale
}

func OffsetXYInchToExcelizeNumber(inch float64) float64 {
	return math.Ceil(inch * offsetXYScale)
}

// FreezeRow must be on top of everything
func FreezeRow(file *excelize.File, sheet string, firstNRows int) error {
	var (
		xSplit int = 0
		ySplit int = firstNRows
	)

	// To get selected cell coordinate just add 1 to both xSplit and ySplit
	topLeftCell, err := excelize.CoordinatesToCellName(xSplit+1, ySplit+1)
	if err != nil {
		return err
	}

	// Notes:
	// 1. Auto set Panes.ActivePane to "topRight"
	// 2. Panes.Options doesn't work for Google Spreadsheet
	// 3. Panes.Options for Microsoft Office Excel will have prompt "some file is corrupted"
	if err := file.SetPanes(
		sheet,
		&excelize.Panes{
			Freeze:      true,
			Split:       false,
			XSplit:      xSplit,
			YSplit:      ySplit,
			TopLeftCell: topLeftCell,
			ActivePane:  "topRight",
		},
	); err != nil {
		return err
	}

	return nil
}

func SetDefaultReportExcelSheet(file *excelize.File, sheet string) error {
	var (
		defaultColWidth  float64 = DefaultColWidthInchToExcelizeNumber(defaultColWidthInInch)
		defaultRowHeight float64 = DefaultRowHeightInchToExcelizeNumber(defaultRowHeightInInch)
	)

	if err := file.SetSheetProps(
		sheet,
		&excelize.SheetPropsOptions{
			DefaultColWidth:  &defaultColWidth,
			DefaultRowHeight: &defaultRowHeight,
		},
	); err != nil {
		return err
	}

	if err := file.SetSheetView(
		sheet,
		0,
		&excelize.ViewOptions{
			ShowGridLines: util.BoolP(true),
		},
	); err != nil {
		return err
	}

	return nil
}

// Set border bottom from startCell to endCell
func SetHeaderSeparator(file *excelize.File, sheet string, startCell string, colCount int) error {
	cellCoordinateX, cellCoordinateY, err := excelize.CellNameToCoordinates(startCell)
	if err != nil {
		return err
	}

	endCell, err := excelize.CoordinatesToCellName(cellCoordinateX+colCount-1, cellCoordinateY)
	if err != nil {
		return err
	}

	headerSeparatorStyleId, err := file.NewStyle(
		&excelize.Style{
			Border: []excelize.Border{
				{Type: "left", Style: 0},
				{Type: "right", Style: 0},
				{Type: "top", Style: 0},
				{Type: "bottom", Color: "000000", Style: 2},
			},
		},
	)
	if err != nil {
		return err
	}

	if err := file.SetCellStyle(sheet, startCell, endCell, headerSeparatorStyleId); err != nil {
		return err
	}

	return nil
}

func NewDefaultReportExcelFile() (file *excelize.File, err error) {
	defer func() {
		if err != nil {
			if closeErr := file.Close(); closeErr != nil {
				log.Printf("close excel file err: %v\n", closeErr)
			}

			// if error exist, need to set file to nil pointer after close
			file = nil
		}
	}()

	file = excelize.NewFile()

	if err = file.SetDefaultFont("Arial"); err != nil {
		return
	}

	return
}

func NewDefaultHeaderTitleStyle(file *excelize.File) (int, error) {
	headerTitleStyleId, err := file.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 15,
			},
		},
	)

	return headerTitleStyleId, err
}

func NewDefaultHeaderSubTitleStyle(file *excelize.File) (int, error) {
	headerSubTitleStyleId, err := file.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 10,
			},
		},
	)

	return headerSubTitleStyleId, err
}

func NewDefaultTableHeaderStyle(file *excelize.File) (int, error) {
	tableHeaderStyleId, err := file.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Vertical: "center",
			},
		},
	)

	return tableHeaderStyleId, err
}

func NewDefaultTableHeaderNumberStyle(file *excelize.File) (int, error) {
	tableHeaderNumberStyleId, err := file.NewStyle(
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

	return tableHeaderNumberStyleId, err
}

func NewDefaultTableBodyStyle(file *excelize.File) (int, error) {
	tableBodyStyleId, err := file.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
		},
	)

	return tableBodyStyleId, err
}

func NewDefaultTableBodyDateStyle(file *excelize.File) (int, error) {
	tableBodyDateStyleId, err := file.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "left",
			},
			CustomNumFmt: util.StringP("D MMM YYYY H:MM:SS"),
		},
	)

	return tableBodyDateStyleId, err
}

func NewDefaultTableBodyCurrencyStyle(file *excelize.File) (int, error) {
	tableBodyCurrencyStyleId, err := file.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Size: 9,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "right",
			},
			DecimalPlaces: 2,
			CustomNumFmt:  util.StringP("_-[$Rp-421]* #,##0.00_-;_-[$Rp-421]* -#,##0.00_-;_-[$Rp-421]* -??_-;_-@"),
		},
	)

	return tableBodyCurrencyStyleId, err
}
