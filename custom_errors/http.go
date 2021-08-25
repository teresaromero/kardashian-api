package custom_errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Err        error
}

func (r *HttpError) Error() string {
	return r.Err.Error()
}

func (r *HttpError) Status() int {
	return r.StatusCode
}

func BadRequest(err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusBadRequest,
		Err:        err,
	}
}

func InternalServerError(err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}

func InvalidCollection(c string) *HttpError {
	return BadRequest(fmt.Errorf("invalid collection %v", c))
}
