package settings

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	filepath "path/filepath"
	io "io"
	os "os"
	settings "rockerbacon/ice-cream-machine-core/internal/settings"
	testing "testing"
)

func NewMockFile(t *testing.T, contents string) *os.File {
	tmpdir := t.TempDir()
	file, err := os.OpenFile(
		filepath.Join(tmpdir, "settings.json"),
		os.O_RDWR | os.O_CREATE,
		0600,
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = file.WriteString(contents)

	if err != nil {
		t.Fatal(err)
	}

	file.Seek(0, io.SeekStart)

	return file
}

func TestReadsCompleteSettings(t *testing.T) {
	file := NewMockFile(t, "{\"host\":\"tests.com\",\"port\":8080}")
	parsedSettings, err := settings.Read(file)
	file.Close()
	assert.Equals(t, err, nil)
	assert.Equals(t, parsedSettings.Host, "tests.com")
	assert.Equals(t, parsedSettings.Port, 8080)
}

func TestUsesCorrectDefaultHost(t *testing.T) {
	file := NewMockFile(t, "{\"port\":8080}")
	parsedSettings, err := settings.Read(file)
	file.Close()
	assert.Equals(t, err, nil)
	assert.Equals(t, parsedSettings.Host, "localhost")
	assert.Equals(t, parsedSettings.Port, 8080)
}

func TestUsesCorrectDefaultPort(t *testing.T) {
	file := NewMockFile(t, "{\"host\":\"tests.com\"}")
	parsedSettings, err := settings.Read(file)
	file.Close()
	assert.Equals(t, err, nil)
	assert.Equals(t, parsedSettings.Host, "tests.com")
	assert.Equals(t, parsedSettings.Port, 6533)
}

func TestRaisesErrorWhenUnknownFieldIsPresent(t *testing.T) {
	file := NewMockFile(t, "{\"test_field\":\"value\"}")
	_, err := settings.Read(file)
	file.Close()
	assert.NotEquals(t, err, nil)
	assert.Equals(t, err.Error(), "json: unknown field \"test_field\"")
}

func TestReadsDirectlyFromFilePath(t *testing.T) {
	file := NewMockFile(t, "{\"host\":\"tests.com\",\"port\":8080}")
	file.Close()
	parsedSettings, err := settings.ReadFromPath(file.Name())
	assert.Equals(t, err, nil)
	assert.Equals(t, parsedSettings.Host, "tests.com")
	assert.Equals(t, parsedSettings.Port, 8080)
}

func TestReturnsDefaultSettingsWhenFilePathDoesNotExist(t *testing.T) {
	tmpdir := t.TempDir()
	path := filepath.Join(tmpdir, "settings.json")
	parsedSettings, err := settings.ReadFromPath(path)
	assert.Equals(t, err, nil)
	assert.Equals(t, parsedSettings.Host, "localhost")
	assert.Equals(t, parsedSettings.Port, 6533)
}
