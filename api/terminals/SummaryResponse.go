package terminals

import (
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/enums/summarytype"
	"github.com/shopspring/decimal"
)

type SummaryResponse struct {
	Amount           *decimal.Decimal
	AmountDue        *decimal.Decimal
	AuthorizedAmount *decimal.Decimal
	Count            *int
	SummaryType      summarytype.SummaryType
	TotalAmount      *decimal.Decimal
	Transactions     []entities.TransactionSummary
}

func NewSummaryResponse() *SummaryResponse {
	return &SummaryResponse{
		Transactions: make([]entities.TransactionSummary, 0),
	}
}

func (s *SummaryResponse) GetAmount() *decimal.Decimal {
	return s.Amount
}

func (s *SummaryResponse) GetAmountDue() *decimal.Decimal {
	return s.AmountDue
}

func (s *SummaryResponse) GetAuthorizedAmount() *decimal.Decimal {
	return s.AuthorizedAmount
}

func (s *SummaryResponse) GetCount() *int {
	return s.Count
}

func (s *SummaryResponse) GetSummaryType() summarytype.SummaryType {
	return s.SummaryType
}

func (s *SummaryResponse) GetTotalAmount() *decimal.Decimal {
	return s.TotalAmount
}

func (s *SummaryResponse) GetTransactions() []entities.TransactionSummary {
	return s.Transactions
}
