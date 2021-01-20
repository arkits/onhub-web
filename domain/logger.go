package domain

import (
	"os"

	"github.com/op/go-logging"
)

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second
var loggingFormat = logging.MustStringFormatter(
	`%{color}%{time:2006/01/02 - 15:04:05.000} %{level:.4s} - %{shortfunc} ▶%{color:reset} %{message}`,
)

// SetupLogger sets up the application wide logger
func SetupLogger() {

	loggingBackend := logging.NewLogBackend(os.Stderr, "", 0)

	backendFormatter := logging.NewBackendFormatter(loggingBackend, loggingFormat)

	logging.SetBackend(backendFormatter)
}
