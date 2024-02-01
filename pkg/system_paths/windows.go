//go:build windows

package system_paths

import (
	os "os"
)

func GetConfigDirPath() string {
	return os.Getenv("AppData")
}
