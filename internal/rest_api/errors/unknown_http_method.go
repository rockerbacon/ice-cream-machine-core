package api_errors

import (
	http "net/http"
)

type UnknownHttpMethodError struct {}

func (UnknownHttpMethodError) StatusCode() int {
	return http.StatusTeapot
}

func (UnknownHttpMethodError) Error() string {
	return "Unknown HTTP method"
}
