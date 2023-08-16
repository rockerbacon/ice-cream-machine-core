package version

type Service struct {
}

func (self *Service) Get() string {
	return "0.0.1"
}
