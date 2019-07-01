package project

import (
	"os"
	"path"

	"github.com/JPZ13/dpm/internal/utils"
)

const (
	configName = ".dpm-config.json"
	// DotDPMFolder houses commands
	DotDPMFolder = ".dpm"
)

// ProjectPath TODO: update
var ProjectPath string

// ProjectCmdPath TODO: update
var ProjectCmdPath string

// ProjectFilePath TODO: update
var ProjectFilePath string

// ProjectName TODO: update
var ProjectName string

// ProjectActive TODO: update
var ProjectActive bool

// TODO: fix this pattern to use a constructor
func init() {
	wd, err := os.Getwd()
	utils.HandleFatalError(err)
	ProjectPath = wd
	ProjectCmdPath = path.Join(wd, ".dpm")
	ProjectFilePath = path.Join(wd, "dpm.yml")
	ProjectName = path.Base(ProjectPath)
	ProjectActive = false
}

// IsProjectInitialized checks the environment variable for whether the
// ProjectFilePath is set
func IsProjectInitialized() bool {
	_, err := os.Stat(ProjectFilePath)
	return err == nil
}

// IsProjectInstalled checks the environment variable for whether the
// ProjectCmdPath is set
func IsProjectInstalled() bool {
	_, err := os.Stat(ProjectCmdPath)
	return err == nil
}
