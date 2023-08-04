package common_handlers

import (
	api_errors "rockerbacon/ice-cream-machine-core/internal/rest_api/errors"
	http "net/http"
)

func MethodNotAllowed (*http.Request) (any, error) {
	return nil, api_errors.NewMethodNotAllowedError()
}

