package entrypoints

import (
	http "net/http"
	version "rockerbacon/ice-cream-machine-core/internal/version"
)

type VersionController struct {}

func (VersionController) Get (r *http.Request) (any, error) {
	return version.Get(), nil
}

