package version

import (
	version "rockerbacon/ice-cream-machine-core/internal/version"
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	testing "testing"
)

func TestReturnsVersionString(t *testing.T) {
	var versionService = version.Service{}
	assert.Matches(
		t,
		versionService.Get(),
		"([[:digit:]]+\\.){2}[[:digit:]]+",
	)
}
