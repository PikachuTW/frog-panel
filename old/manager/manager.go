package manager

type Manager struct {
	config *Config
}

func New(config *Config) *Manager {
	return &Manager{
		config: config,
	}
}
