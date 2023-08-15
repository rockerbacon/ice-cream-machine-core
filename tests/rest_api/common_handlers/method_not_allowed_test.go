package common_handlers

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	handlers "rockerbacon/ice-cream-machine-core/internal/rest_api/common_handlers"
	api_errors "rockerbacon/ice-cream-machine-core/internal/rest_api/errors"
	testing "testing"
	tuple "rockerbacon/ice-cream-machine-core/pkg/tuple"
)

func TestShouldReturnAMethodNotAllowedError(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(handlers.MethodNotAllowed(nil)),
		api_errors.NewMethodNotAllowedError(),
	)
}

