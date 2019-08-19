package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/JPZ13/dpm/internal/pathtable"
)

// Runner holds methods related to running
// a command
type Runner interface {
	Run(ctx context.Context, args []string) error
}

type runner struct {
	baseService
}

// Run executes a command in a container
func (r *runner) Run(ctx context.Context, args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectInfo, err := r.pathTable.Get(pwd)
	if err != nil {
		return err
	}

	if !projectInfo.IsActive {
		return callBinary(args, projectInfo)
	}

	return runDockerizedCommand(args, projectInfo)
}

// callBinary calls the stored binary of an alias
func callBinary(args []string, project *pathtable.ProjectInfo) error {
	command := args[0]
	remainder := args[1:]

	for _, aliasInfo := range project.Commands {
		for alias, binaryPath := range aliasInfo.Aliases {
			if alias != command {
				continue
			}

			// execute command of binary and remainder
			cmd := exec.Command(binaryPath, remainder...)
			return cmd.Run()
		}
	}

	return errors.New("Native binary not found")
}

func runDockerizedCommand(args []string, project *pathtable.ProjectInfo) error {
	command := args[0]

	for _, aliasInfo := range project.Commands {
		for alias := range aliasInfo.Aliases {
			if alias != command {
				continue
			}
			return runDocker(args, &aliasInfo)
		}
	}

	return errors.New("Project error: alias not found")
}

func runDocker(args []string, alias *pathtable.AliasInfo) error {
	// create named volume if it doesn't exist
	err := maybeCreateVolume(alias.VolumeName)
	if err != nil {
		return err
	}

	// run volume mounted container container
	remainder := args[1:]
	return runContainer(alias.Image, alias.VolumeName, remainder)
}

func maybeCreateVolume(volumeName string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dockerCmd := fmt.Sprintf(`docker volume create --driver local
		--opt type=none device=%s --opt o=bind %s`, pwd, volumeName)

	cmd := exec.Command(dockerCmd)
	return cmd.Run()
}

func runContainer(imageName, volumeName string, remainder []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dockerCmd := fmt.Sprintf(`docker run --rm -it
		-v %s:%s -w %s %s`,
		volumeName, pwd, pwd, imageName)

	cmd := exec.Command(dockerCmd, remainder...)
	return cmd.Run()
}
