package cas

// AliasInfo contains info on CLI commands
// that will be run in containers
type AliasInfo struct {
	Aliases        map[string]bool `json:"aliases"`
	BinaryLocation string          `json:"binaryLocation"`
	VolumeName     string          `json:"volumeName"`
}
