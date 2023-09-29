package common

import (
	"github.com/fatih/color"
	"os"
	"time"
)

var currentTime = time.Now().Format("2006-01-02 15:11")

type Logger struct {
	info  *color.Color
	warn  *color.Color
	fatal *color.Color
}

func NewLogger() *Logger {
	return &Logger{
		info:  color.New(color.FgGreen),
		warn:  color.New(color.FgYellow),
		fatal: color.New(color.FgRed),
	}
}

func (l *Logger) Info(msg string) {
	l.info.Printf("[%s][APP INFO] : %s\n", currentTime, msg)
}

func (l *Logger) Warn(msg string) {
	l.info.Printf("[%s][APP Warning] : %s\n", currentTime, msg)
}

func (l *Logger) Fatal(msg string) {
	l.info.Printf("[%s][APP Fatal] : %s\n", currentTime, msg)
	os.Exit(1)
}
