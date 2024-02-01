//go:build windows

package system_paths

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	paths "rockerbacon/ice-cream-machine-core/pkg/system_paths"
	testing "testing"
)

func TestGetConfigDirPathReturnsAppDataEnvironmentVariable(t *testing.T) {
	t.Setenv("AppData", "C:\\Users\\icmtests\\AppData\\Roaming")
	assert.Equals(
		t,
		paths.GetConfigDirPath(),
		"C:\\Users\\icmtests\\AppData\\Roaming",
	)
}
