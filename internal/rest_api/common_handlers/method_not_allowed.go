package common_handlers

import (
	http "net/http"
	api_errors "rockerbacon/ice-cream-machine-core/internal/rest_api/errors"
)

func MethodNotAllowed(*http.Request) (any, error) {
	return nil, api_errors.NewMethodNotAllowedError()
}
