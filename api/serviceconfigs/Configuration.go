package serviceconfigs

import (
	"net/url"

	"github.com/globalpayments/go-sdk/api/entities/enums/environment"
	"github.com/globalpayments/go-sdk/api/entities/enums/host"
	"github.com/globalpayments/go-sdk/api/entities/enums/hosterror"
	"github.com/globalpayments/go-sdk/api/logging"
)

type Configuration struct {
	EnableLogging       bool
	RequestLogger       logging.IRequestLogger
	Environment         environment.Environment
	ServiceURL          string
	WebProxy            *url.URL
	SimulatedHostErrors map[host.Host][]hosterror.HostError
	Timeout             int
	Validated           bool
	DynamicHeaders      map[string]string
}

func (c *Configuration) GetRequestLogger() logging.IRequestLogger {
	return c.RequestLogger
}

func (c *Configuration) SetRequestLogger(r logging.IRequestLogger) {
	c.RequestLogger = r
}

func (c *Configuration) GetTimeout() int {
	return c.Timeout
}

func (c *Configuration) SetTimeout(timeout int) {
	c.Timeout = timeout
}

func (c *Configuration) Validate() error {
	c.Validated = true
	return nil
}
