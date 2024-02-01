//go:build darwin

package system_paths

import (
	filepath "path/filepath"
	os "os"
)

func GetConfigDirPath() string {
	return filepath.Join(os.Getenv("HOME"), "Library", "Application Support")
}
