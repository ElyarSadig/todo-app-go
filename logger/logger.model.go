package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	reset        = "\033[0m"
	red          = "\033[31m"
	yellow       = "\033[33m"
	blue         = "\033[34m"
	brightYellow = "\033[93m"
)

type Log struct {
	logger *log.Logger
}

func New() Logger {
	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return &Log{logger: l}
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown file:0"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func (l *Log) Info(msg string) {
	l.logger.Printf("%s %sINFO: %s%s", getCallerInfo(), blue, msg, reset)
}

func (l *Log) Warn(msg string) {
	l.logger.Printf("%s %sWARN: %s%s", getCallerInfo(), yellow, msg, reset)
}

func (l *Log) Error(msg string) {
	l.logger.Printf("%s %sERROR: %s%s", getCallerInfo(), red, msg, reset)
}

func (l *Log) Debug(msg string) {
	l.logger.Printf("%s %sDEBUG: %s%s", getCallerInfo(), brightYellow, msg, reset)
}

func (l *Log) Fatal(msg string) {
	l.logger.Printf("%s %sFATAL: %s%s", getCallerInfo(), red, msg, reset)
	os.Exit(1)
}
