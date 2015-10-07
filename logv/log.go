package logv

import (
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
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
	level  int32
	mu     sync.RWMutex
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
		logger: log.New(os.Stdout, "", 0),
		level:  0,
	}
}

func (l *logger) SetOutput(out io.Writer) {
	l.mu.Lock()
	l.logger = log.New(out, "", 0)
	l.mu.Unlock()
}

func (l *logger) SetLevel(level int) {
	atomic.StoreInt32(&l.level, int32(level))
}

func (l *logger) V(level int) bool {
	return int32(level) <= atomic.LoadInt32(&l.level)
}

func (l *logger) getLogger() *log.Logger {
	l.mu.RLock()
	lgr := l.logger
	l.mu.RUnlock()
	return lgr
}

func (l *logger) Print(v ...interface{}) {
	lgr := l.getLogger()
	if lgr != nil {
		lgr.Print(v)
	}
}

func (l *logger) Println(v ...interface{}) {
	lgr := l.getLogger()
	if lgr != nil {
		lgr.Println(v)
	}
}

func (l *logger) Printf(format string, v ...interface{}) {
	lgr := l.getLogger()
	if lgr != nil {
		lgr.Printf(format, v...)
	}
}
