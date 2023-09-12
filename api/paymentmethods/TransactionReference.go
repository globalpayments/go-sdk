package paymentmethods

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"reflect"
)

type TransactionReference struct {
	AlternativePaymentType          string
	AcquiringInstitutionID          string
	AuthCode                        string
	BatchNumber                     *int
	ClientTransactionID             string
	MessageTypeIndicator            string
	OrderID                         string
	OriginalAmount                  float64
	OriginalApprovedAmount          float64
	OriginalPaymentMethod           abstractions.IPaymentMethod
	OriginalProcessingCode          string
	OriginalTransactionTime         string
	PartialApproval                 bool
	PaymentMethodType               paymentmethodtype.PaymentMethodType
	PosDataCode                     string
	SequenceNumber                  *int
	SystemTraceAuditNumber          string
	TransactionID                   string
	OriginalTransactionDate         string
	ResponseCode                    string
	UseAuthorizedAmount             bool
	TransactionIdentifier           string
	OriginalInvoiceNumber           string
	OriginalTransactionInfo         string
	OriginalPOSEntryMode            string
	OriginalTransactionType         transactiontype.TransactionType
	ApprovalCode                    string
	OriginalMessageCode             string
	MastercardBanknetRefNo          string
	MastercardBanknetSettlementDate string
	DebitAuthorizer                 string
	VisaTransactionID               string
	DiscoverNetworkRefID            string
}

func (t *TransactionReference) GetOriginalApprovedAmount() float64 {
	if t.OriginalApprovedAmount != 0 {
		return t.OriginalApprovedAmount
	}
	return t.OriginalAmount
}

func (t *TransactionReference) GetPaymentMethodType() paymentmethodtype.PaymentMethodType {
	if t.OriginalPaymentMethod != nil {
		return t.OriginalPaymentMethod.GetPaymentMethodType()
	}
	return t.PaymentMethodType
}

func NewTransactionReference() *TransactionReference {
	return &TransactionReference{}
}

func (tr *TransactionReference) SetTransactionId(tid string) {
	tr.TransactionID = tid
}

func IsItATransactionReference(i interface{}) bool {
	t := reflect.TypeOf(i)
	typeString := t.String()
	return typeString == "paymentmethods.TransactionReference"
}
