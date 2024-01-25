package entrypoints

import (
	http "net/http"
	version "rockerbacon/ice-cream-machine-core/internal/version"
)

func getPath() string {
	return "/version/"
}

func get(*http.Request) (any, error) {
	return version.Get(), nil
}

func Version() *Entrypoint {
	entrypoint := Entrypoint {
		Get: get,
		GetPath: getPath,
	}

	return &entrypoint
}

