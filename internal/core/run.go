package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/volume"
	docker "github.com/docker/docker/client"
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
	dockerClient, err := docker.NewEnvClient()
	if err != nil {
		return err
	}

	// create named volume if it doesn't exist
	volume, err := maybeCreateVolume(dockerClient, alias.VolumeName)
	if err != nil {
		return err
	}

	// run volume mounted container container
	container, err := runContainer(dockerClient, alias.Image, volume)
	if err != nil {
		return err
	}

	remainder := args[1:]
	return attachToContainer(dockerClient, container, remainder)
}

func attachToContainer(dockerClient *docker.Client, container *container.ContainerCreateCreatedBody, remainder []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	ctx := context.Background()
	exec, err := dockerClient.ContainerExecCreate(ctx, container.ID, types.ExecConfig{
		Tty:          true,
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		WorkingDir:   pwd,
		Cmd:          remainder,
	})
	if err != nil {
		return err
	}

	_, err = dockerClient.ContainerExecAttach(ctx, exec.ID, types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	})

	return err
}

func maybeCreateVolume(dockerClient *docker.Client, volumeName string) (*types.Volume, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	volume, err := dockerClient.VolumeCreate(ctx, volume.VolumeCreateBody{
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

	return &volume, nil
}

func runContainer(dockerClient *docker.Client, imageName string, volume *types.Volume) (*container.ContainerCreateCreatedBody, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	container, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image:        imageName,
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir:   pwd,
	}, &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:%s", volume.Name, pwd),
		},
	}, nil, "")
	if err != nil {
		return nil, err
	}

	err = dockerClient.ContainerStart(ctx, container.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}

	return &container, nil
}
