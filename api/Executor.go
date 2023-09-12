package api

import (
	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
)

type ExecutableTerminal interface {
	ExecuteWithName(name string, device builders.ITerminalBuilderDevice) (terminalresponse.ITerminalResponse, error)
}

func ExecuteTerminal(terminal ExecutableTerminal) (terminalresponse.ITerminalResponse, error) {
	return ExecuteTerminalWithName("default", terminal)
}

func ExecuteTerminalWithName(name string, terminal ExecutableTerminal) (terminalresponse.ITerminalResponse, error) {
	device, err := GetServiceContainerInstance().GetDeviceController(name)
	if err != nil {
		return nil, err
	}
	return terminal.ExecuteWithName(name, device)
}
