package version

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	testing "testing"
	version "rockerbacon/ice-cream-machine-core/internal/version"
)

func TestReturnsVersionString(t *testing.T) {
	assert.Matches(
		t,
		version.Get(),
		"([[:digit:]]+\\.){2}[[:digit:]]+",
	)
}
