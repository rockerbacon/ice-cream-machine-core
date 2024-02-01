//go:build linux

package system_paths

import (
	filepath "path/filepath"
	os "os"
)

func GetConfigDirPath() string {
	xdgPath, isXdgPathSet := os.LookupEnv("XDG_CONFIG_HOME")

	if isXdgPathSet {
		return xdgPath
	}

	return filepath.Join(os.Getenv("HOME"), ".config")
}
