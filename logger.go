package gosteam

type Logging interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
}

var logger Logging = nil

func SetLogger(customLogger Logging) {
	logger = customLogger
}
