package terminals

import (
	"errors"
	"fmt"

	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/terminals/devicecontroller"
	"github.com/globalpayments/go-sdk/api/terminals/messaging"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
	"github.com/globalpayments/go-sdk/api/terminals/terminalutilities"
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upamessageid"
	"github.com/globalpayments/go-sdk/api/terminals/upa/interfaces"
	"github.com/globalpayments/go-sdk/api/terminals/upa/responses"
	"github.com/globalpayments/go-sdk/api/terminals/upa/subgroups"
	"github.com/globalpayments/go-sdk/api/utils"
)

type UpaController struct {
	devicecontroller.DeviceController
	device        abstractions.IDeviceInterface
	onMessageSent messaging.IMessageSentInterface
}

func NewUpaController(settings *ConnectionConfig) (*UpaController, error) {
	controller := &UpaController{
		DeviceController: *devicecontroller.NewDeviceController(),
	}

	if controller.device == nil {
		controller.device = NewUpaInterface(controller)
	}

	controller.DeviceController.SetRequestIdProvider(settings.RequestIdProvider)

	switch settings.ConnectionMode {
	case connectionmodes.TCP_IP:
		controller.DeviceController.SetDeviceCommInterface(interfaces.NewUpaTcpInterface(settings))
	default:
		return nil, errors.New("unsupported connection mode")
	}

	controller.DeviceController.GetDeviceCommInterface().SetMessageSentHandler(controller.onMessageSent)
	return controller, nil
}

func (c *UpaController) SetOnMessageSent(ms messaging.IMessageSentInterface) {
	c.onMessageSent = ms
	c.DeviceController.GetDeviceCommInterface().SetMessageSentHandler(c.onMessageSent)
}

func (c *UpaController) GetOnMessageSent() messaging.IMessageSentInterface {
	return c.onMessageSent
}

func (c *UpaController) onMessageSentHandler(message string) {
	if c.onMessageSent != nil {
		c.onMessageSent.MessageSent(message)
	}
}

func (c *UpaController) ConfigureInterface() abstractions.IDeviceInterface {
	if c.device == nil {
		c.device = NewUpaInterface(c)
	}
	return c.device
}

func (c *UpaController) Send(message abstractions.IDeviceMessage) ([]byte, error) {
	return c.DeviceController.GetDeviceCommInterface().Send(message)
}

func (c *UpaController) SendWithouDisconnect(message abstractions.IDeviceMessage) ([]byte, error) {
	return c.DeviceController.GetDeviceCommInterface().SendWithouDisconnect(message)
}

func (c *UpaController) DoTransaction(messageId upamessageid.UpaMessageId, requestId int, ecrId string, paramFields *subgroups.RequestParamFields, transactionFields *subgroups.RequestTransactionFields, processingIndicators *subgroups.RequestProcessingIndicatorsFields) (*responses.UpaTransactionResponse, error) {
	body := utils.NewJsonDoc()

	if paramFields.GetElementsJson() != nil {
		body.SetJsonDoc("params", paramFields.GetElementsJson())
	}

	if transactionFields.GetElementsJson() != nil {
		body.SetJsonDoc("transaction", transactionFields.GetElementsJson())
	}

	if messageId == upamessageid.StartCardTransaction {
		body.SetJsonDoc("processingIndicators", processingIndicators.GetElementsJson())
	}

	requestIdAsString := fmt.Sprintf("%d", requestId)

	deviceMessage, _ := terminalutilities.BuildMessage(messageId, requestIdAsString, ecrId, body)

	resp, err := c.Send(deviceMessage)
	if err != nil {
		msg := err.Error()
		if msg == "Terminal did not respond in the given timeout." {
			c.device.Cancel()
		}
		return nil, err
	}

	responseObj, err := utils.ParseBytes(resp)
	if err != nil {
		return nil, err
	}

	data := responseObj.Get("data")
	if data == nil {
		return nil, exceptions.NewApiException("Terminal response was malformed.")
	}

	return responses.NewUpaTransactionResponse(data), nil
}

func (c *UpaController) ProcessTransaction(builder *builders.TerminalAuthBuilder) (terminalresponse.ITerminalResponse, error) {
	messageId, err := c.MapTransactionType(builder.GetTransactionType())
	if err != nil {
		return nil, err
	}

	requestId := c.GetRequestIdProvider().GetRequestId()

	requestParamFields := subgroups.NewRequestParamFields()
	requestParamFields.SetAuthBuilderParams(builder)

	requestTransactionFields := subgroups.NewRequestTransactionFields()
	requestTransactionFields.SetAuthBuilderParams(builder)

	processingIndicators := subgroups.NewRequestProcessingIndicatorsFields()
	return c.DoTransaction(messageId, requestId, builder.EcrId, requestParamFields, requestTransactionFields, processingIndicators)
}

func (c *UpaController) MapTransactionType(transactionType transactiontype.TransactionType) (upamessageid.UpaMessageId, error) {
	switch transactionType {
	case transactiontype.Auth:
		return upamessageid.PreAuth, nil
	case transactiontype.Sale:
		return upamessageid.Sale, nil
	case transactiontype.Void:
		return upamessageid.Void, nil
	case transactiontype.Refund:
		return upamessageid.Refund, nil
	case transactiontype.Edit:
		return upamessageid.TipAdjust, nil
	case transactiontype.Verify:
		return upamessageid.CardVerify, nil
	case transactiontype.Reversal:
		return upamessageid.Reversal, nil
	case transactiontype.Balance:
		return upamessageid.BalanceInquiry, nil
	case transactiontype.Capture:
		return upamessageid.AuthCompletion, nil
	case transactiontype.Activate:
		return upamessageid.StartCardTransaction, nil
	default:
		return "", errors.New("Selected gateway does not support this transaction type")
	}
}

func (c *UpaController) ManageTransaction(builder *builders.TerminalManageBuilder) (terminalresponse.ITerminalResponse, error) {
	messageId, err := c.MapTransactionType(builder.GetTransactionType())
	if err != nil {
		return nil, err
	}

	requestId := c.GetRequestIdProvider().GetRequestId()

	requestParamFields := subgroups.NewRequestParamFields()
	requestParamFields.SetManageBuilderParams(builder)

	requestTransactionFields := subgroups.NewRequestTransactionFields()
	requestTransactionFields.SetManageBuilderParams(builder)

	return c.DoTransaction(messageId, requestId, builder.EcrId, requestParamFields, requestTransactionFields, nil)

}
