package core

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/JPZ13/dpm/internal/model"
	docker "github.com/fsouza/go-dockerclient"
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
func callBinary(args []string, project *model.ProjectInfo) error {
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

func runDockerizedCommand(args []string, project *model.ProjectInfo) error {
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

func runDocker(args []string, alias *model.AliasInfo) error {
	dockerClient, err := docker.NewClientFromEnv()
	if err != nil {
		return err
	}

	// create named volume if it doesn't exist
	volume, err := maybeCreateVolume(dockerClient, alias.VolumeName)
	if err != nil {
		return err
	}

	// run volume mounted container container
	remainder := args[1:]
	return runContainer(dockerClient, alias.Image, volume, remainder)
}

func maybeCreateVolume(dockerClient *docker.Client, volumeName string) (*docker.Volume, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	volume, err := dockerClient.CreateVolume(docker.CreateVolumeOptions{
		Driver: "local",
		DriverOpts: map[string]string{
			"type":   "none",
			"device": pwd,
			"o":      "bind",
		},
		Name: volumeName,
	})
	if err != nil {
		return nil, err
	}

	return volume, nil
}

func runContainer(dockerClient *docker.Client, imageName string, volume *docker.Volume, remainder []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	container, err := dockerClient.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:        imageName,
			Cmd:          remainder,
			Tty:          true,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			WorkingDir:   pwd,
			// Volumes: map[string]struct{}{
			// 	volume.Name: *volume,
			// },
		},
	})
	if err != nil {
		return err
	}

	return dockerClient.StartContainer(container.ID, nil)
}
