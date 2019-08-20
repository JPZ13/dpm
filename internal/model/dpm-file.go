package model

// Command contains YAML definition for a command
type Command struct {
	Name        string   `yaml:"-"`
	Image       string   `yaml:"image"`
	Entrypoints []string `yaml:"entrypoints,omitempty"`
	VolumeName  string   `yaml:"volume-name,omitempty"`
}

// DPMFile is the type for the entire YAML file
type DPMFile struct {
	Commands map[string]Command `yaml:"commands"`
}
