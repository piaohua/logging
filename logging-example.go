package main

import (
	"fmt"
	"io"
	"os"
	"time"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{pid}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {
	//var buf bytes.Buffer
	//fmt.Fprintf(&buf, *fmt, Args...)
	fmt.Fprintf(os.Stderr, "log: exiting because of error: %s\n", "ss")
	os.Stderr.Write([]byte("ERROR: logging before flag.Parse: \n"))
	//
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	now := time.Now()
	fmt.Printf("new -> %v\n", now)
	file, _, err := create("INFO", now)
	if err != nil {
		panic(err)
	}
	fmt.Printf("name -> %v\n", file.Name)
	fmt.Printf("name -> %v\n", file.Name)
	file.Write([]byte(boldcolors[CRITICAL] + "[CRITICAL]\n"))
	file.Write([]byte(boldcolors[ERROR] + "[ERROR]\n"))
	file.Write([]byte(boldcolors[WARNING] + "[WARNING]\n"))
	file.Write([]byte(boldcolors[NOTICE] + "[NOTICE]\n"))
	file.Write([]byte(boldcolors[INFO] + "[INFO]\n"))
	file.Write([]byte(boldcolors[DEBUG] + "[DEBUG]\n"))
}

func example2() {
	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)
	//
	log.Debugf("debug %s", Password("secret"))
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}

func example3() {
	os.Stderr.Write([]byte(boldcolors[CRITICAL] + "[CRITICAL] sssss\n"))
	os.Stderr.Write([]byte(boldcolors[ERROR] + "[ERROR]\n"))
	os.Stderr.Write([]byte(boldcolors[WARNING] + "[WARNING]\n"))
	os.Stderr.Write([]byte(boldcolors[NOTICE] + "[NOTICE]\n"))
	os.Stderr.Write([]byte(boldcolors[INFO] + "[INFO]\n"))
	os.Stderr.Write([]byte(boldcolors[DEBUG] + "[DEBUG]\n"))
}

func example4() {
	doFmtVerbLevelColor("bold", CRITICAL, os.Stderr)
	os.Stderr.Write([]byte("[CRITICAL] "))
	doFmtVerbLevelColor("reset", CRITICAL, os.Stderr)
	os.Stderr.Write([]byte("[messages] \n"))
	doFmtVerbLevelColor("", CRITICAL, os.Stderr)
	os.Stderr.Write([]byte("[CRITICAL] "))
	doFmtVerbLevelColor("reset", CRITICAL, os.Stderr)
	os.Stderr.Write([]byte("[messages] \n"))
	//
	doFmtVerbLevelColor("bold", ERROR, os.Stderr)
	os.Stderr.Write([]byte("[ERROR] "))
	doFmtVerbLevelColor("reset", ERROR, os.Stderr)
	os.Stderr.Write([]byte("[messages] \n"))
	doFmtVerbLevelColor("", ERROR, os.Stderr)
	os.Stderr.Write([]byte("[ERROR] "))
	doFmtVerbLevelColor("reset", ERROR, os.Stderr)
	os.Stderr.Write([]byte("[messages] \n"))
}

// Level defines all available log levels for log messages.
type Level int

// Log levels.
const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

type color int

const (
	ColorBlack = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

var (
	colors = []string{
		CRITICAL: ColorSeq(ColorMagenta),
		ERROR:    ColorSeq(ColorRed),
		WARNING:  ColorSeq(ColorYellow),
		NOTICE:   ColorSeq(ColorGreen),
		DEBUG:    ColorSeq(ColorCyan),
	}
	boldcolors = []string{
		CRITICAL: ColorSeqBold(ColorMagenta),
		ERROR:    ColorSeqBold(ColorRed),
		WARNING:  ColorSeqBold(ColorYellow),
		NOTICE:   ColorSeqBold(ColorGreen),
		DEBUG:    ColorSeqBold(ColorCyan),
	}
)

func ColorSeq(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

func ColorSeqBold(color color) string {
	return fmt.Sprintf("\033[%d;1m", int(color))
}

func doFmtVerbLevelColor(layout string, level Level, output io.Writer) {
	if layout == "bold" {
		output.Write([]byte(boldcolors[level]))
	} else if layout == "reset" {
		output.Write([]byte("\033[0m"))
	} else {
		output.Write([]byte(colors[level]))
	}
}
