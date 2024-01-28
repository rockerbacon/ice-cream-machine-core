package version

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	testing "testing"
	version "rockerbacon/ice-cream-machine-core/internal/version"
)

func TestReturnsVersionString(t *testing.T) {
	manager := version.StaticVersionManager{}

	assert.Matches(
		t,
		manager.Get(),
		"([[:digit:]]+\\.){2}[[:digit:]]+",
	)
}
