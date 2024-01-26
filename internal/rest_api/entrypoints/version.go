package entrypoints

import (
	http "net/http"
	version "rockerbacon/ice-cream-machine-core/internal/version"
)

type VersionController struct {
	Manager version.VersionManager
}

func (self VersionController) Get (r *http.Request) (any, error) {
	return self.Manager.Get(), nil
}

func NewVersionController() VersionController {
	return VersionController {
		Manager: version.StaticVersionManager{},
	}
}
