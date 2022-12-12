package taifn

import (
	"net/http"
)

type TaifnLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type TaifnContext interface {
	Logger() TaifnLogger
}

type TaifnHandlerFunc func(http.ResponseWriter, *http.Request, TaifnContext)
