package nlog

import (
	"fmt"
	"io"
	"log"
)

const (
	No = iota
	Critical
	Error
	Warn
	Info
	Debug
	Verbose
)

const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LstdFlags     = log.LstdFlags
)

var levelName = map[int]string{
	Critical: "CRITICAL",
	Error:    "ERROR",
	Warn:     "WARN",
	Info:     "INFO",
	Debug:    "DEBUG",
	Verbose:  "VERBOSE",
}

type Logger interface {
	SetLevel(level int)
	Verbosef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Criticalf(format string, v ...interface{})
}

type Config struct {
	Out io.Writer
	Level int
	Prefix string
	Flag int
}

type myLogger struct {
	level  int
	logger *log.Logger
}

var _ Logger = (*myLogger)(nil)

func NewLogger(config *Config) Logger {
	return &myLogger{
		level:  config.Level,
		logger: log.New(config.Out, config.Prefix, config.Flag),
	}
}

func LevelToName(level int) string {
	name, ok := levelName[level]
	if ok {
		return name
	}
	return ""
}

func (l *myLogger) SetLevel(level int) {
	l.level = level
}

func (l *myLogger) Verbosef(format string, v ...interface{}) {
	l.Printf(Verbose, format, v...)
}

func (l *myLogger) Debugf(format string, v ...interface{}) {
	l.Printf(Debug, format, v...)
}

func (l *myLogger) Infof(format string, v ...interface{}) {
	l.Printf(Info, format, v...)
}

func (l *myLogger) Warnf(format string, v ...interface{}) {
	l.Printf(Warn, format, v...)
}

func (l *myLogger) Errorf(format string, v ...interface{}) {
	l.Printf(Error, format, v...)
}

func (l *myLogger) Criticalf(format string, v ...interface{}) {
	l.Printf(Critical, format, v...)
}

func (l *myLogger) Printf(level int, format string, v ...interface{}) {
	if level > l.level {
		return
	}
	name := LevelToName(level)
	if name != "" {
		format = fmt.Sprintf("[%s] %s", name, format)
		l.logger.Printf(format, v...)
	}
}
