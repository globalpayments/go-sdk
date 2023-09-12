package builders

import "github.com/globalpayments/go-sdk/api/terminals/terminalresponse"

type ITerminalBuilderDevice interface {
	ManageTransaction(tmb *TerminalManageBuilder) (terminalresponse.ITerminalResponse, error)
	ProcessTransaction(tab *TerminalAuthBuilder) (terminalresponse.ITerminalResponse, error)
}
