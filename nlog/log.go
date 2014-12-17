package nlog

import (
	"log"
	"bufio"
	"sync"
	"io"
	"fmt"
	"strings"
)

const (
	LogLevelNo = iota
	LogLevelCritical
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
)

const (
	Ldate = log.Ldate
	Ltime = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile = log.Llongfile
	Lshortfile = log.Lshortfile
	LstdFlags = log.LstdFlags
)

var logLevelName = map[int]string{
	LogLevelCritical: "CRITICAL",
	LogLevelError:    "ERROR",
	LogLevelWarn:     "WARN",
	LogLevelInfo:     "INFO",
	LogLevelDebug:    "DEBUG",
}

var logNameLevel = map[string]int{
	"CRITICAL": LogLevelCritical,
	"ERROR":    LogLevelError,
	"WARN":     LogLevelWarn,
	"INFO":     LogLevelInfo,
	"DEBUG":    LogLevelDebug,
}

type Logger interface {
	SetLevel(level int)
	Printf(level int, format string, v ...interface{})
	Flush()
}

type bufLogger struct {
	level  int
	logger *log.Logger
	buf    *bufio.Writer
	mu     sync.Mutex
}

var _ Logger = (*bufLogger)(nil)

func NewLogger(out io.Writer, level int, prefix string, flg int) Logger {
	buf := bufio.NewWriter(out)
	return &bufLogger{
		level:  level,
		buf:    buf,
		logger: log.New(buf, prefix, flg),
	}
}

func GetLogLevelByName(name string) int {
	level, ok := logNameLevel[strings.ToUpper(name)]
	if ok {
		return level
	}
	return LogLevelNo
}

func (l *bufLogger) SetLevel(level int) {
	l.level = level
}

func (l *bufLogger) Printf(level int, format string, v ...interface{}) {
	if level > l.level {
		return
	}
	name, ok := logLevelName[level]
	if ok {
		format = fmt.Sprintf("[%s] %s", name, format)
		l.mu.Lock()
		defer l.mu.Unlock()
		l.logger.Printf(format, v...)
	}
}

func (l *bufLogger) Flush() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf.Flush()
}
