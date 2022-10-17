package logger

import (
	"fmt"
	"log"
	"os"
)

type LevelFalg string

const (
	DebugFalg = "debug"
	InfoFalg  = "info"
	WarnFalg  = "warn"
	ErrorFalg = "error"
	PaincFalg = "painc"
)

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PaincLevel
)

func NewLeve(flag LevelFalg) Level {
	switch flag {
	case DebugFalg:
		return DebugLevel
	case InfoFalg:
		return InfoLevel
	case WarnFalg:
		return WarnLevel
	case ErrorFalg:
		return ErrorLevel
	case PaincFalg:
		return PaincLevel
	default:
		return DebugLevel
	}
}

func (l Level) Enable(level Level) bool {
	return level >= l
}

type Logger struct {
	logger *log.Logger
	level  Level
}

func NewLogger(prefix string, level Level) *Logger {
	return &Logger{
		logger: log.New(os.Stdout, fmt.Sprintf("[%s] ", prefix), log.LUTC|log.Ldate|log.Llongfile),
		level:  level,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if DebugLevel.Enable(l.level) {
		l.logger.Println("[Debug] ", v)
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if DebugLevel.Enable(l.level) {
		l.logger.Printf("[Debug] %s\n", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if InfoLevel.Enable(l.level) {
		l.logger.Println("[Info] ", v)
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if InfoLevel.Enable(l.level) {
		l.logger.Printf("[Info] %s\n", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if WarnLevel.Enable(l.level) {
		l.logger.Println("[Warn] ", v)
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if WarnLevel.Enable(l.level) {
		l.logger.Printf("[Warn] %s\n", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if ErrorLevel.Enable(l.level) {
		l.logger.Println("[Error] ", v)
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if ErrorLevel.Enable(l.level) {
		l.logger.Printf("[Error] %s\n", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Painc(v ...interface{}) {
	if PaincLevel.Enable(l.level) {
		l.logger.Panicln("[Painc] ", v)
	}
}

func (l *Logger) Paincf(format string, v ...interface{}) {
	if PaincLevel.Enable(l.level) {
		l.logger.Panicf("[Painc] %s\n", fmt.Sprintf(format, v...))
	}
}
