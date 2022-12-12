package taifn

import (
	"io"
	"net/http"
)

type HttpHelper interface {
	ParseBody(body io.ReadCloser, target interface{})
	NewError(w http.ResponseWriter, status int, err error)
}
