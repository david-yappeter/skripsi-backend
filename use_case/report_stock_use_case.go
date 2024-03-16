package use_case

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"myapp/constant"
	"myapp/data_type"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"time"

	"github.com/xuri/excelize/v2"
	"golang.org/x/sync/errgroup"
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
}

type ReportDailyTransactionExcelSheet2Data struct {
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

func (u *ReportDailyTransactionExcel) initSheet2(
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

	if err = reportExcel.Init(exportedDateTime); err != nil {
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

func GenerateStockReport(
	ctx context.Context,
	repositoryManager repository.RepositoryManager,
	mainFilesystem filesystem.Client,
) (file *model.File, err error) {
	productRepository := repositoryManager.ProductRepository()
	productStockMutationRepository := repositoryManager.ProductStockMutationRepository()

	products, err := productRepository.Fetch(ctx)
	panicIfErr(err)

	baseProductUnitLoader := loader.NewBaseProductUnitLoader(repositoryManager.ProductUnitRepository())
	productStockLoader := loader.NewProductStockLoader(repositoryManager.ProductStockRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range products {
				group.Go(baseProductUnitLoader.ProductFnNotStrict(&products[i]))
				group.Go(productStockLoader.ProductFnNotStrict(&products[i]))
			}
		}),
	)

	unitLoader := loader.NewUnitLoader(repositoryManager.UnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range products {
				group.Go(unitLoader.ProductUnitFn(products[i].BaseProductUnit))
			}
		}),
	)

	// construct excel report sheets
	reportExcel, err := NewReportDailyTransactionExcel(
		util.CurrentDateTime(),
	)
	if err != nil {
		return
	}
	defer reportExcel.Close()

	for _, product := range products {
		baseUnit := "-"
		stockLeft := 0.0
		currentSellingPrice := 0.0

		if product.BaseProductUnit != nil {
			baseUnit = product.BaseProductUnit.Unit.Name
		}

		if product.Price != nil {
			currentSellingPrice = *product.Price
		}

		if product.ProductStock != nil {
			stockLeft = product.ProductStock.Qty
		}

		reportExcel.AddSheet1Data(ReportDailyTransactionExcelSheet1Data{
			ProductId:           product.Id,
			ProductName:         product.Name,
			BaseUnit:            baseUnit,
			CurrentSellingPrice: currentSellingPrice,
			IsActive:            false,
			StockLeft:           stockLeft,
		})
	}

	productStockMutations, err := productStockMutationRepository.FetchHaveQtyLeft(ctx)
	panicIfErr(err)

	productUnitLoader := loader.NewProductUnitLoader(repositoryManager.ProductUnitRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productStockMutations {
				group.Go(productUnitLoader.ProductStockMutationFn(&productStockMutations[i]))
			}
		}),
	)

	productLoader := loader.NewProductLoader(repositoryManager.ProductRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productStockMutations {
				group.Go(productLoader.ProductUnitFn(productStockMutations[i].ProductUnit))
				group.Go(unitLoader.ProductUnitFn(productStockMutations[i].ProductUnit))
			}
		}),
	)

	for _, productStockMutation := range productStockMutations {
		reportExcel.AddSheet2Data(ReportDailyTransactionExcelSheet2Data{
			ProductId:     productStockMutation.ProductUnit.ProductId,
			UnitId:        productStockMutation.ProductUnit.UnitId,
			ProductName:   productStockMutation.ProductUnit.Product.Name,
			UnitName:      productStockMutation.ProductUnit.Unit.Name,
			MutationType:  productStockMutation.Type.String(),
			Qty:           productStockMutation.Qty,
			ScaleToBase:   productStockMutation.ScaleToBase,
			BaseQty:       productStockMutation.ScaleToBase * productStockMutation.Qty,
			BaseQtyLeft:   productStockMutation.BaseQtyLeft,
			BaseQtySold:   (productStockMutation.ScaleToBase * productStockMutation.Qty) - productStockMutation.BaseQtyLeft,
			BaseCostPrice: productStockMutation.BaseCostPrice,
			MutatedAt:     productStockMutation.MutatedAt.Time(),
		})
	}

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
