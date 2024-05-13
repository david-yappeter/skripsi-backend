package use_case

import (
	"context"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"sort"

	"golang.org/x/sync/errgroup"
)

type dashboardLoaderParams struct {
}

type DashboardUseCase interface {
	// read
	SummarizeDebt(ctx context.Context) ([]model.CustomerDebtSummary, []model.SupplierDebtSummary)
	SummarizeTransaction(ctx context.Context, request dto_request.DashboardSummarizeTransactionRequest) []model.TransactionSummary
}

type dashboardUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewDashboardUseCase(
	repositoryManager repository.RepositoryManager,
) DashboardUseCase {
	return &dashboardUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *dashboardUseCase) SummarizeDebt(ctx context.Context) ([]model.CustomerDebtSummary, []model.SupplierDebtSummary) {
	customerDebtSummaries, err := u.repositoryManager.CustomerRepository().FetchTopNCustomerDebtSummaryOrderByTotalDebt(ctx, 5)
	panicIfErr(err)

	supplierDebtSummaries, err := u.repositoryManager.SupplierRepository().FetchTopNSupplierDebtSummaryOrderByTotalDebt(ctx, 5)
	panicIfErr(err)

	return customerDebtSummaries, supplierDebtSummaries
}

func (u *dashboardUseCase) SummarizeTransaction(ctx context.Context, request dto_request.DashboardSummarizeTransactionRequest) []model.TransactionSummary {
	transactions, err := u.repositoryManager.TransactionRepository().Fetch(ctx, model.TransactionQueryOption{
		QueryOption: model.QueryOption{
			Page:  nil,
			Limit: nil,
		},
		CashierSessionId: nil,
		Status:           data_type.TransactionStatusP(data_type.TransactionStatusPaid),
		PaymentStartedAt: request.StartDate.DateTimeStartOfDay().NullDateTime(),
		PaymentEndedAt:   request.EndDate.DateTimeEndOfDay().NullDateTime(),
	})
	panicIfErr(err)

	transactionItemsLoader := loader.NewTransactionItemsLoader(u.repositoryManager.TransactionItemRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			group.Go(transactionItemsLoader.TransactionFn(&transactions[i]))
		}
	}))

	transactionItemCostsLoader := loader.NewTransactionItemCostsLoader(u.repositoryManager.TransactionItemCostRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			for j := range transactions[i].TransactionItems {
				group.Go(transactionItemCostsLoader.TransactionItemFn(&transactions[i].TransactionItems[j]))
			}
		}
	}))

	transactionSummaryMapByDate := map[string]*model.TransactionSummary{}

	for _, transaction := range transactions {
		if transactionSummaryMapByDate[transaction.PaymentAt.Date().IsoLayout()] != nil {
			transactionSummaryMapByDate[transaction.PaymentAt.Date().IsoLayout()].TotalGrossSales += transaction.Total
			transactionSummaryMapByDate[transaction.PaymentAt.Date().IsoLayout()].TotalNetSales += transaction.NetTotal()
			transactionSummaryMapByDate[transaction.PaymentAt.Date().IsoLayout()].Transactions = append(transactionSummaryMapByDate[transaction.PaymentAt.Date().IsoLayout()].Transactions, transaction)
		} else {
			transactionSummaryMapByDate[transaction.PaymentAt.Date().IsoLayout()] = &model.TransactionSummary{
				Date:            transaction.PaymentAt.Date(),
				TotalGrossSales: transaction.Total,
				TotalNetSales:   transaction.NetTotal(),
				Transactions:    []model.Transaction{transaction},
			}
		}
	}

	transactionSummaries := []model.TransactionSummary{}
	for _, transactionSummary := range transactionSummaryMapByDate {
		transactionSummaries = append(transactionSummaries, *transactionSummary)
	}

	sort.Slice(transactionSummaries, func(i, j int) bool {
		return transactionSummaries[i].Date.IsLessThanOrEqual(transactionSummaries[j].Date)
	})

	return transactionSummaries
}
