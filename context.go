package taifn

import (
	"io"
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

type KafkaProducer interface {
	ProduceFull(topic string, key string, value string, Headers map[string]string, brokers []string)
	ProduceWithHeaders(topic string, key string, value string, Headers map[string]string)
	Produce(topic string, key string, value string)
}

type HttpHelper interface {
	ParseBody(body io.ReadCloser, target interface{})
	NewError(w http.ResponseWriter, status int, err error)
}

type TaifnContext interface {
	Logger() TaifnLogger
	KafkaProducer() KafkaProducer
	HttpHelper() HttpHelper
}

type TaifnHandlerFunc func(http.ResponseWriter, *http.Request, TaifnContext)
