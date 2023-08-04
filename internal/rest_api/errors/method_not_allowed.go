package api_errors

import (
	http "net/http"
)

type MethodNotAllowedError struct {}

func (MethodNotAllowedError)StatusCode() int {
	return http.StatusMethodNotAllowed
}

func (self MethodNotAllowedError)Error() string {
	return http.StatusText(self.StatusCode())
}

func NewMethodNotAllowedError() MethodNotAllowedError {
	return MethodNotAllowedError{}
}

