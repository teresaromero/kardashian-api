package http_errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Err        error
	Message    string
}

func (r *HttpError) Error() string {
	return r.Message
}

func (r *HttpError) Status() int {
	return r.StatusCode
}

func BadRequest(err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusBadRequest,
		Err:        err,
		Message:    "⚠️ Oops... Seems you requested something wrong!",
	}
}

func InternalServerError(err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Err:        err,
		Message:    "🔥 Oops... Seems something is wrong on our servers!",
	}
}

func InvalidCollection(c string) *HttpError {
	return BadRequest(fmt.Errorf("invalid collection %v", c))
}
