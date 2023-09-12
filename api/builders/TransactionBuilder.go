package builders

import (
	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/host"
	"github.com/globalpayments/go-sdk/api/entities/enums/hosterror"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	networkentities "github.com/globalpayments/go-sdk/api/network/entities"
	"github.com/globalpayments/go-sdk/api/network/enums/cardissuerentrytag"
	paymentmethods "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
)

type TransactionBuilder struct {
	*BaseBuilder
	transactionType         transactiontype.TransactionType
	transactionModifier     transactionmodifier.TransactionModifier
	paymentMethod           paymentmethods.IPaymentMethod
	simulatedHostErrors     map[host.Host][]hosterror.HostError
	batchNumber             *int
	companyId               string
	description             string
	issuerData              map[cardissuerentrytag.CardIssuerEntryTag]string
	followOnStan            *int
	priorMessageInformation *networkentities.PriorMessageInformation
	sequenceNumber          *int
	systemTraceAuditNumber  *int
	uniqueDeviceId          string
	transactionMatchingData *networkentities.TransactionMatchingData
	terminalError           bool
	taxAmount               *decimal.Decimal
	tipAmount               *decimal.Decimal
	surchargeAmount         *decimal.Decimal
	cashBackAmount          *decimal.Decimal
	invoiceNumber           string
	cvn                     string
	amount                  *decimal.Decimal
	tagData                 string
	emvMaxPinEntry          string
	serviceCode             string
	cardSequenceNumber      string
	ecommerceAuthIndicator  string
	ecommerceData1          string
	ecommerceData2          string
	transactionDate         string
	transactionTime         string
	zipCode                 string
}

func NewTransactionBuilder(transactionType transactiontype.TransactionType, paymentMethod paymentmethods.IPaymentMethod) *TransactionBuilder {
	return &TransactionBuilder{
		BaseBuilder:         NewBaseBuilder(),
		transactionType:     transactionType,
		paymentMethod:       paymentMethod,
		simulatedHostErrors: make(map[host.Host][]hosterror.HostError),
	}
}

func (builder *TransactionBuilder) WithServiceCode(serviceCode string) *TransactionBuilder {
	builder.serviceCode = serviceCode
	return builder
}

func (builder *TransactionBuilder) WithEcommerceAuthIndicator(ecommerceAuthIndicator string) *TransactionBuilder {
	builder.ecommerceAuthIndicator = ecommerceAuthIndicator
	return builder
}

func (builder *TransactionBuilder) WithEcommerceData1(ecommerceData1 string) *TransactionBuilder {
	builder.ecommerceData1 = ecommerceData1
	return builder
}

func (builder *TransactionBuilder) WithEcommerceData2(ecommerceData2 string) *TransactionBuilder {
	builder.ecommerceData2 = ecommerceData2
	return builder
}

func (builder *TransactionBuilder) WithZipCode(zipCode string) *TransactionBuilder {
	builder.zipCode = zipCode
	return builder
}

func (builder *TransactionBuilder) GetTransactionType() transactiontype.TransactionType {
	return builder.transactionType
}

func (builder *TransactionBuilder) SetTransactionType(transactionType transactiontype.TransactionType) {
	builder.transactionType = transactionType
}

func (builder *TransactionBuilder) GetTransactionModifier() transactionmodifier.TransactionModifier {
	return builder.transactionModifier
}

func (builder *TransactionBuilder) SetTransactionModifier(transactionModifier transactionmodifier.TransactionModifier) {
	builder.transactionModifier = transactionModifier
}

func (builder *TransactionBuilder) GetPaymentMethod() paymentmethods.IPaymentMethod {
	return builder.paymentMethod
}

func (builder *TransactionBuilder) SetPaymentMethod(paymentMethod paymentmethods.IPaymentMethod) {
	builder.paymentMethod = paymentMethod
}

func (builder *TransactionBuilder) GetSimulatedHostErrors() map[host.Host][]hosterror.HostError {
	return builder.simulatedHostErrors
}

func (builder *TransactionBuilder) GetBatchNumber() *int {
	return builder.batchNumber
}

func (builder *TransactionBuilder) GetCompanyId() string {
	return builder.companyId
}

func (builder *TransactionBuilder) GetIssuerData() map[cardissuerentrytag.CardIssuerEntryTag]string {
	return builder.issuerData
}

func (builder *TransactionBuilder) GetFollowOnStan() *int {
	return builder.followOnStan
}

func (builder *TransactionBuilder) GetPriorMessageInformation() *networkentities.PriorMessageInformation {
	return builder.priorMessageInformation
}

func (builder *TransactionBuilder) SetPriorMessageInformation(priorMessageInformation *networkentities.PriorMessageInformation) {
	builder.priorMessageInformation = priorMessageInformation
}

func (builder *TransactionBuilder) GetSequenceNumber() *int {
	return builder.sequenceNumber
}

func (builder *TransactionBuilder) GetSystemTraceAuditNumber() *int {
	return builder.systemTraceAuditNumber
}

func (builder *TransactionBuilder) GetUniqueDeviceId() string {
	return builder.uniqueDeviceId
}

func (builder *TransactionBuilder) GetTransactionMatchingData() *networkentities.TransactionMatchingData {
	return builder.transactionMatchingData
}

func (builder *TransactionBuilder) IsTerminalError() bool {
	return builder.terminalError
}

func (builder *TransactionBuilder) WithDescription(description string) *TransactionBuilder {
	builder.description = description
	return builder
}

func (builder *TransactionBuilder) WithIssuerData(issuerData map[cardissuerentrytag.CardIssuerEntryTag]string) *TransactionBuilder {
	builder.issuerData = issuerData
	return builder
}

func (builder *TransactionBuilder) WithFollowOnStan(followOnStan *int) *TransactionBuilder {
	builder.followOnStan = followOnStan
	return builder
}

func (builder *TransactionBuilder) WithSequenceNumber(sequenceNumber *int) *TransactionBuilder {
	builder.sequenceNumber = sequenceNumber
	return builder
}

func (builder *TransactionBuilder) WithSystemTraceAuditNumber(systemTraceAuditNumber *int) *TransactionBuilder {
	builder.systemTraceAuditNumber = systemTraceAuditNumber
	return builder
}

func (builder *TransactionBuilder) WithUniqueDeviceId(uniqueDeviceId string) *TransactionBuilder {
	builder.uniqueDeviceId = uniqueDeviceId
	return builder
}

func (builder *TransactionBuilder) WithTransactionMatchingData(transactionMatchingData *networkentities.TransactionMatchingData) *TransactionBuilder {
	builder.transactionMatchingData = transactionMatchingData
	return builder
}

func (builder *TransactionBuilder) SetTerminalError(terminalError bool) {
	builder.terminalError = terminalError
}

func (builder *TransactionBuilder) SetTaxAmount(taxAmount *decimal.Decimal) {
	builder.taxAmount = taxAmount
}

func (builder *TransactionBuilder) SetTipAmount(tipAmount *decimal.Decimal) {
	builder.tipAmount = tipAmount
}

func (builder *TransactionBuilder) SetSurchargeAmount(surchargeAmount *decimal.Decimal) {
	builder.surchargeAmount = surchargeAmount
}

func (builder *TransactionBuilder) SetCashBackAmount(cashBackAmount *decimal.Decimal) {
	builder.cashBackAmount = cashBackAmount
}

func (builder *TransactionBuilder) SetInvoiceNumber(invoiceNumber string) {
	builder.invoiceNumber = invoiceNumber
}

func (builder *TransactionBuilder) SetCvn(cvn string) {
	builder.cvn = cvn
}

func (builder *TransactionBuilder) SetAmount(amount *decimal.Decimal) {
	builder.amount = amount
}

func (builder *TransactionBuilder) SetTagData(tagData string) {
	builder.tagData = tagData
}

func (builder *TransactionBuilder) SetEmvMaxPinEntry(emvMaxPinEntry string) {
	builder.emvMaxPinEntry = emvMaxPinEntry
}

func (builder *TransactionBuilder) SetCardSequenceNumber(cardSequenceNumber string) {
	builder.cardSequenceNumber = cardSequenceNumber
}

func (builder *TransactionBuilder) SetEcommerceAuthIndicator(ecommerceAuthIndicator string) {
	builder.ecommerceAuthIndicator = ecommerceAuthIndicator
}

func (builder *TransactionBuilder) SetEcommerceData1(ecommerceData1 string) {
	builder.ecommerceData1 = ecommerceData1
}

func (builder *TransactionBuilder) SetEcommerceData2(ecommerceData2 string) {
	builder.ecommerceData2 = ecommerceData2
}

func (builder *TransactionBuilder) SetTransactionDate(transactionDate string) {
	builder.transactionDate = transactionDate
}

func (builder *TransactionBuilder) SetTransactionTime(transactionTime string) {
	builder.transactionTime = transactionTime
}

func (builder *TransactionBuilder) SetZipCode(zipCode string) {
	builder.zipCode = zipCode
}
