package paths

import (
	filepath "path/filepath"
	syspaths "rockerbacon/ice-cream-machine-core/pkg/system_paths"
)

func GetConfigDirPath() string {
	return filepath.Join(syspaths.GetConfigDirPath(), "icm")
}

func GetConfigFilePath() string {
	return filepath.Join(GetConfigDirPath(), "settings.json")
}
