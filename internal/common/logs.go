package common

import (
	"log"
	"os"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

type Logger struct {
	info  *log.Logger
	warn  *log.Logger
	fatal *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "[APP INFO] ", log.Ldate|log.Lshortfile),
		warn:  log.New(os.Stdout, "[APP WARNING] ", log.Ldate|log.Lshortfile),
		fatal: log.New(os.Stdout, "[APP FATAL] ", log.Ldate|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.info.Println(colorGreen + msg + colorReset)
}

func (l *Logger) Warn(msg string) {
	l.warn.Println(colorYellow + msg + colorReset)
}

func (l *Logger) Fatal(msg string) {
	l.fatal.Println(colorRed + msg + colorReset)
	os.Exit(1)
}

func LogInfo(logs *Logger, msg string) {
	logs.Info(msg)
}

func LogWarning(logs *Logger, err error) {
	if err != nil {
		logs.Warn(err.Error())
	}
}

func LogError(logs *Logger, err error) {
	if err != nil {
		logs.Fatal(err.Error())
	}
}
