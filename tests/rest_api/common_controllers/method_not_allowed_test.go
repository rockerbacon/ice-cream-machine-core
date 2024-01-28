package common_controllers

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	api_errors "rockerbacon/ice-cream-machine-core/internal/rest_api/errors"
	controllers "rockerbacon/ice-cream-machine-core/internal/rest_api/common_controllers"
	testing "testing"
	tuple "rockerbacon/ice-cream-machine-core/pkg/tuple"
)

func TestShouldReturnAMethodNotAllowedErrorOnConnect(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Connect(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnDelete(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Delete(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnGet(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Get(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnHead(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Head(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnOptions(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Options(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnPatch(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Patch(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnPost(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Post(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnPut(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Put(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}
func TestShouldReturnAMethodNotAllowedErrorOnTrace(t *testing.T) {
	assert.ErrorEquals(
		t,
		tuple.Second(
			controllers.MethodNotAllowed.Trace(nil),
		),
		api_errors.NewMethodNotAllowedError(),
	)
}

