package entrypoints

import (
	http "net/http"
	version "rockerbacon/ice-cream-machine-core/internal/version"
)

type VersionEntrypoint struct {
	*BaseEntrypoint
	service *version.Service
}

func (self *VersionEntrypoint) Get(*http.Request) (any, error) {
	return self.service.Get(), nil
}

func Version() Entrypoint {
	return &VersionEntrypoint{
		service: &version.Service{},
		BaseEntrypoint: &BaseEntrypoint{
			path: "/version/",
		},
	}
}
