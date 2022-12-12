package taifn

import (
	"net/http"
)

type TaifnContext interface {
	Logger() TaifnLogger
	KafkaProducer() KafkaProducer
	HttpHelper() HttpHelper
}

type TaifnHandlerFunc func(http.ResponseWriter, *http.Request, TaifnContext)
