package pathtable

// AliasInfo contains info on CLI commands
// that will be run in containers
type AliasInfo struct {
	Aliases        map[string]string `json:"aliases"`
	Image          string            `json:"image"`
	VolumeName     string            `json:"volumeName"`
	BinaryLocation string            `json:"binaryLocation"`
}

// ProjectInfo holds configuration information
// about a project
type ProjectInfo struct {
	IsActive bool        `json:"isActive"`
	Commands []AliasInfo `json:"commands"`
}
