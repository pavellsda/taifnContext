package taifn

import (
	"io"
	"net/http"
)

type HttpHelper interface {
	ParseBody(body io.ReadCloser, target interface{}) error
	NewError(w http.ResponseWriter, status int, err error)
}
