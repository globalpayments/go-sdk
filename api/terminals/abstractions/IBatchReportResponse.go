package abstractions

import "github.com/globalpayments/go-sdk/api/entities"

type IBatchReportResponse interface {
	IDeviceResponse
	GetBatchSummary() entities.BatchSummary
	GetTransactionSummaries() []entities.TransactionSummary
}
