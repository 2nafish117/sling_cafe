package log

// @TODO: improve logging to have colors
// https://github.com/op/go-logging

import (
	"os"

	"github.com/op/go-logging"
)

// Log is a logger instance
var Log *logging.Logger

func init() {
	Log = logging.MustGetLogger("example")
	// Example format string. Everything except the message has a custom color
	// which is dependent on the Log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	var format = logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{level:.4s} %{shortfunc} ▶ %{id:03x}%{color:reset} %{message}`,
	)
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used Log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)

	// Log.Debugf("debug %s", Password("secret"))
	// Log.Info("info")
	// Log.Notice("notice")
	// Log.Warning("warning")
	// Log.Error("err")
	// Log.Critical("crit")
}

// // Password is just an example type implementing the Redactor interface. Any
// // time this is logged, the Redacted() function will be called.
// type Password string

// func (p Password) Redacted() interface{} {
// 	return logging.Redact(string(p))
// }
