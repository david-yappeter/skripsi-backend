package dto_request

import "myapp/data_type"

type DashboardSummarizeTransactionRequest struct {
	StartDate data_type.Date `json:"start_date"`
	EndDate   data_type.Date `json:"end_date"`
} // @name DashboardSummarizeTransactionRequest
