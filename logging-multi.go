package main

import logging "github.com/op/go-logging"

var log = logging.MustGetLogger("log_multi")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {

	log1 := logging.NewMemoryBackend(8)
	log2 := logging.NewMemoryBackend(8)

	leveled1 := logging.AddModuleLevel(log1)
	leveled2 := logging.AddModuleLevel(log2)

	multi := logging.MultiLogger(leveled1, leveled2)
	multi.SetLevel(logging.DEBUG, "log_multi")
	logging.SetBackend(multi)

	log.Debugf("debug %s", Password("secret"))
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")

}
