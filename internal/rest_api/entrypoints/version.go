package entrypoints

import (
	common "rockerbacon/ice-cream-machine-core/internal/rest_api/common_handlers"
	http "net/http"
)

func getPath() string {
	return "/version/"
}

func get(*http.Request) (any, error) {
	return "0.0.1", nil
}

func Version() *Entrypoint {
	entrypoint := Entrypoint {
		Connect: common.MethodNotAllowed,
		Delete: common.MethodNotAllowed,
		Get: get,
		GetPath: getPath,
		Head: common.MethodNotAllowed,
		Options: common.MethodNotAllowed,
		Patch: common.MethodNotAllowed,
		Post: common.MethodNotAllowed,
		Put: common.MethodNotAllowed,
		Trace: common.MethodNotAllowed,
	}

	return &entrypoint
}

