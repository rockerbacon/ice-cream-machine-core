//go:build linux

package system_paths

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	filepath "path/filepath"
	os "os"
	paths "rockerbacon/ice-cream-machine-core/pkg/system_paths"
	testing "testing"
)

func TestGetConfigDirPathUsesXDGEnvironmentVariable(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "/home/icmtests/Documents/config")
	assert.Equals(
		t,
		paths.GetConfigDirPath(),
		"/home/icmtests/Documents/config",
	)
}

func TestGetConfigDirPathUsesXDGDefault(t *testing.T) {
	assert.Equals(
		t,
		paths.GetConfigDirPath(),
		filepath.Join(os.Getenv("HOME"), ".config"),
	)
}
