package logv

import (
	"io"
	"log"
	"os"
)

const (
	No = iota
	Fatal
	Err
	Warn
	Info
	Debug
	Trace
)

type Logger interface {
	SetOutput(out io.Writer)
	SetLevel(level int)
	V(level int) bool
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	logger *log.Logger
	level  int
}

var (
	_             Logger = (*logger)(nil)
	defaultLogger Logger
)

func init() {
	defaultLogger = NewLogger()
}

func SetOutput(out io.Writer) {
	defaultLogger.SetOutput(out)
}

func SetLevel(level int) {
	defaultLogger.SetLevel(level)
}

func V(level int) bool {
	return defaultLogger.V(level)
}

func Print(v ...interface{}) {
	defaultLogger.Print(v...)
}

func Println(v ...interface{}) {
	defaultLogger.Println(v...)
}

func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

func NewLogger() Logger {
	return &logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		level:  Warn,
	}
}

func (l *logger) SetOutput(out io.Writer) {
	l.logger = log.New(out, "", log.LstdFlags)
}

func (l *logger) SetLevel(level int) {
	l.level = level
}

func (l *logger) V(level int) bool {
	return level <= l.level && level > No
}

func (l *logger) Print(v ...interface{}) {
	if l.logger != nil {
		l.logger.Print(v)
	}
}

func (l *logger) Println(v ...interface{}) {
	if l.logger != nil {
		l.logger.Println(v)
	}
}

func (l *logger) Printf(format string, v ...interface{}) {
	if l.logger != nil {
		l.logger.Printf(format, v...)
	}
}
