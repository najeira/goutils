package nlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	No = iota
	Fatal
	Error
	Warn
	Info
	Debug
	Trace
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
	Fatal: "FATAL",
	Error: "ERROR",
	Warn:  "WARN",
	Info:  "INFO",
	Debug: "DEBUG",
	Trace: "TRACE",
}

type Logger interface {
	V(level int) bool
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Config struct {
	Out    io.Writer
	Level  int
	Prefix string
	Flag   int
}

type myLogger struct {
	level  int
	logger *log.Logger
}

var _ Logger = (*myLogger)(nil)

func NewLogger(config *Config) Logger {
	var out io.Writer = nil
	var level int = Info
	var prefix string = ""
	var flag int = 0
	if config != nil {
		out = config.Out
		level = config.Level
		prefix = config.Prefix
		flag = config.Flag
	}
	if out == nil {
		out = os.Stdout
	}
	return &myLogger{
		level:  level,
		logger: log.New(out, prefix, flag),
	}
}

func LevelToName(level int) string {
	if name, ok := levelName[level]; ok {
		return name
	}
	return ""
}

func NameToLevel(name string) int {
	name = strings.ToUpper(name)
	for level, value := range levelName {
		if name == value {
			return level
		}
	}
	return No
}

func (l *myLogger) V(level int) bool {
	return level <= l.level && level != No
}

func (l *myLogger) Tracef(format string, v ...interface{}) {
	l.Printf(Trace, format, v...)
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

func (l *myLogger) Fatalf(format string, v ...interface{}) {
	l.Printf(Fatal, format, v...)
	os.Exit(1)
}

func (l *myLogger) Printf(level int, format string, v ...interface{}) {
	if level > l.level {
		return
	}
	if name := LevelToName(level); name != "" {
		format = fmt.Sprintf("[%s] %s", name, format)
		l.logger.Printf(format, v...)
	}
}
