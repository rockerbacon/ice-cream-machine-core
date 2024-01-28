package common_controllers

import (
	api_errors "rockerbacon/ice-cream-machine-core/internal/rest_api/errors"
	http "net/http"
)

func raiseError () (any, error) {
	return nil, api_errors.NewMethodNotAllowedError()
}

type methodNotAllowed struct {}

func (methodNotAllowed) Connect (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Delete (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Get (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Head (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Options (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Patch (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Post (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Put (*http.Request) (any, error) {
	return raiseError()
}

func (methodNotAllowed) Trace (*http.Request) (any, error) {
	return raiseError()
}

var MethodNotAllowed methodNotAllowed

