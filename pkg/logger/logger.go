package logger

import (
	"os"
	
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

// InitLogger config logger
func InitLogger() {
	// Log as JSON instead of the default ASCII formatter, but wrapped with the runtime Formatter.
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "02-01-2006 15:04:05",
		},
	}
	
	// Enable line number and file logging as well
	formatter.Line = true
	formatter.File = true
	
	// Replace the default Logrus Formatter with our runtime Formatter
	logrus.SetFormatter(&formatter)
	
	return
}

// Warn prints warning message to logs
func Warn(format string, v ...interface{}) {
	logrus.Warnf(format, v...)
}

// Debug prints debug message to logs
func Debug(format string, v ...interface{}) {
	logrus.Debugf(format, v...)
}

// Info prints info message to logs
func Info(format string, v ...interface{}) {
	logrus.Printf(format, v...)
}

// Err prints error message to logs
func Err(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

// Fatal calls Err and then os.Exit(1)
func Fatal(format string, v ...interface{}) {
	Err(format, v...)
	os.Exit(1)
}
