package entrypoints

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	entrypoints "rockerbacon/ice-cream-machine-core/internal/rest_api/entrypoints"
	testing "testing"
)

type MockVersionManager struct {}
func (MockVersionManager) Get() string {
	return "4.2.0"
}

func TestRespondsVersionString(t *testing.T) {
	controller := entrypoints.VersionController{
		Manager: MockVersionManager{},
	}

	version, e := controller.Get(nil)

	assert.Equals(t, e, nil)
	assert.Equals(t, version, "4.2.0")
}

func TestBuildsWithSensibleDefaults(t *testing.T) {
	controller := entrypoints.NewVersionController()

	version, e := controller.Get(nil)

	assert.Equals(t, e, nil)
	assert.NotEquals(t, version, nil)
}
