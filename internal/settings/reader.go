package settings

import (
	errors "errors"
	io "io"
	json "encoding/json"
	os "os"
	paths "rockerbacon/ice-cream-machine-core/internal/paths"
)

func NewDefaultSettings() Settings {
	return Settings {
		Host: "localhost",
		Port: 6533,
	}
}

func Read(reader io.Reader) (Settings, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	settings := NewDefaultSettings()
	err := decoder.Decode(&settings)
	return settings, err
}

func ReadFromPath(path string) (Settings, error) {
	file, err := os.Open(path)
	// consider other errors to be undefined behaviour for now
	if errors.Is(err, os.ErrNotExist) {
		return NewDefaultSettings(), nil
	}

	settings, err := Read(file)

	file.Close()

	return settings, err
}

func ReadFromDefaultPath() (Settings, error) {
	return ReadFromPath(paths.GetConfigFilePath())
}
