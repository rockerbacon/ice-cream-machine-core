//go:build darwin

package system_paths

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	paths "rockerbacon/ice-cream-machine-core/pkg/system_paths"
	testing "testing"
)

func TestGetConfigDirPathFollowsAppleGuidelines(t *testing.T) {
	t.Setenv("HOME", "/Users/icmtests")
	assert.Equals(
		t,
		paths.GetConfigDirPath(),
		"/Users/icmtests/Library/Application Support",
	)
}
