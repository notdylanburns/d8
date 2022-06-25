package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type logLevel int

const (
	fatal logLevel = iota - 4 // fatal    : -4
	error                         // error    : -3
	warn                          // warn     : -2
	log                           // log      : -1
	_                             // SKIP     : 0
	info                          // info     : 1
	debug                         // debug    : 2
)

var logLabels = map[logLevel]string{
	fatal: "FATAL",
	error: "ERROR",
	warn:  "WARNING",
	log:   "LOG",
	info:  "INFO",
	debug: "DEBUG",
}

var logColors = map[logLevel]string{
	fatal: "\x1b[31m",
	error: "\x1b[31m",
	warn:  "\x1b[33m",
	info:  "\x1b[34m",
	log:   "\x1b[30m",
	debug: "\x1b[32m",
}

func (l logLevel) String() string {
	str, ok := logLabels[l]
	if !ok {
		str = "????"
	}
	return str
}

var colorReset = "\x1b[0m"
var loggerLevel logLevel

func init() {
	_level, ok := os.LookupEnv("SCIF_MESSAGELEVEL")
	if !ok {
		loggerLevel = debug
	} else {
		_levelint, err := strconv.Atoi(_level)
		if err != nil {
			loggerLevel = debug
		} else {
			loggerLevel = logLevel(_levelint)
		}
	}
}

func prefix(level logLevel) string {

	messageColor, ok := logColors[level]
	if !ok {
		messageColor = "\x1b[0m"
	}

	if loggerLevel < debug {
		return fmt.Sprintf("%s%-8s%s ", messageColor, level.String()+":", colorReset)
	}

	pc, _, _, ok := runtime.Caller(3)
	details := runtime.FuncForPC(pc)

	var funcName string
	if ok && details == nil {
		funcName = "????()"
	} else {
		funcNameSplit := strings.Split(details.Name(), ".")
		funcName = funcNameSplit[len(funcNameSplit)-1] + "()"
	}

	return fmt.Sprintf("%s%-8s%-19s%-30s", messageColor, level, colorReset, funcName)
}

func writef(level logLevel, format string, a ...interface{}) {
	if loggerLevel < level {
		return
	}

	message := fmt.Sprintf(format, a...)
	message = strings.TrimSuffix(message, "\n")

	fmt.Fprintf(os.Stderr, "%s%s\n", prefix(level), message)
}

func Debugf(format string, a ...interface{}) {
	writef(debug, format, a...)
}

func Errorf(format string, a ...interface{}) {
	writef(error, format, a...)
}

func Fatalf(format string, a ...interface{}) {
	writef(fatal, format, a...)
	os.Exit(255)
}

func Infof(format string, a ...interface{}) {
	writef(info, format, a...)
}

func Warningf(format string, a ...interface{}) {
	writef(warn, format, a...)
}

func SetLevel(l int) {
	loggerLevel = logLevel(l)
}

func DisableColor() {
	logColors = map[logLevel]string{
		fatal: "",
		error: "",
		warn:  "",
		info:  "",
		log:   "",
		debug: "",
	}
	colorReset = ""
}

func Writer() io.Writer {
	if loggerLevel <= -1 {
		return ioutil.Discard
	}
	return os.Stderr
}