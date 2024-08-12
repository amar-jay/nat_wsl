package config

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type Forwarding struct {
	Protocol string `yaml:",omitempty"` // default:"tcp"
	Type     string `yaml:",omitempty"` // default:"v4tov4"
	Wsl      struct {
		Listenport int
		Listenhost string
	}
	Remote struct {
		Connectport int
		Connectip   string
	}
}

type Config map[string]Forwarding // random name for now

// set default values
func (f *Config) SetDefaults() {
	// set default values for each forwarding
	for _, forwarded := range *f {
		if forwarded.Protocol == "" {
			forwarded.Protocol = "tcp"
		}
		if forwarded.Type == "" {
			forwarded.Type = "v4tov4"
		}
	}
}
