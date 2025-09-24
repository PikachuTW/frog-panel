package instance

type InstanceType int
type InstanceStatus int

const (
	Vanilla InstanceType = iota
	Fabric
	Forge
	Paper
)

const (
	Stop InstanceStatus = iota
	Start
	Error
)

type Instance struct {
	Type InstanceType
}

func New(config *Config) (*Instance, error) {
	panic("not implemented")
}

func Status() InstanceStatus {
	panic("not implemented")
}
