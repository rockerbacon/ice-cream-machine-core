package version

type VersionManager interface {
	Get() string
}

type StaticVersionManager struct {}
func (StaticVersionManager) Get() string {
	return "0.0.1"
}
