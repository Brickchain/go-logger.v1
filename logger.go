/*
Logger for Brickchain software
*/
package logger

import (
	"os"
	"github.com/Sirupsen/logrus"
	"io"
	"sync"
)

var (
	ctxlogger *logrus.Entry
	mu *sync.Mutex
)

func init() {
	mu = &sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	logrus.SetOutput(os.Stdout)
	hostname, _ := os.Hostname()
	ctxlogger = logrus.WithField("pid", os.Getpid()).WithField("hostname", hostname)

}

func GetLogger() *logrus.Entry {
	return ctxlogger
}

func AddContext(key string, value interface{}) {
	mu.Lock()
	defer mu.Unlock()
	ctxlogger = ctxlogger.WithField(key, value)
}

func SetFormatter(formatter string) {
	var _formatter logrus.Formatter
	switch formatter {
	case "json":
		_formatter = &logrus.JSONFormatter{}
	default:
		_formatter = &logrus.TextFormatter{}
	}
	mu.Lock()
	defer mu.Unlock()
	data := ctxlogger.Data
	logrus.SetFormatter(_formatter)
	ctxlogger = logrus.WithFields(data)
}

func SetOutput(out io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	data := ctxlogger.Data
	logrus.SetOutput(out)
	ctxlogger = logrus.WithFields(data)
}

func SetLevel(level string) {
	_level, err := logrus.ParseLevel(level)
	if err != nil {
		_level = logrus.InfoLevel
	}
	mu.Lock()
	defer mu.Unlock()
	data := ctxlogger.Data
	logrus.SetLevel(_level)
	ctxlogger = logrus.WithFields(data)
}

func GetLoglevel() string {
	return logrus.GetLevel().String()
}

func WithField(key string, value interface{}) *logrus.Entry {
	return ctxlogger.WithField(key, value)
}

// Wrapper for Logrus Debug()
func Debug(args ...interface{}) {
	ctxlogger.Debug(args...)
}

// Wrapper for Logrus Info()
func Info(args ...interface{}) {
	ctxlogger.Info(args...)
}

// Wrapper for Logrus Warn()
func Warn(args ...interface{}) {
	ctxlogger.Warn(args...)
}

// Wrapper for Logrus Error()
func Error(args ...interface{}) {
	ctxlogger.Error(args...)
}

// Wrapper for Logrus Fatal()
func Fatal(args ...interface{}) {
	ctxlogger.Fatal(args...)
}

func Errorf(format string, args ...interface{}) {
	ctxlogger.Errorf(format, args...)
}

func Infof(format string, args ...interface{}) {
	ctxlogger.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	ctxlogger.Warningf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	ctxlogger.Debugf(format, args...)
}
