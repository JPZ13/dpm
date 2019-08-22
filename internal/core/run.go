package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/JPZ13/dpm/internal/pathtable"
	"github.com/JPZ13/dpm/internal/utils"
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
		if err == pathtable.ErrNotFound {
			return runBinaryFromPATH(args)
		}

		return err
	}

	if !projectInfo.IsActive {
		return callBinary(args, projectInfo)
	}

	return runDockerizedCommand(args, projectInfo)
}

func runBinaryFromPATH(args []string) error {
	userPath := os.Getenv("PATH")

	subPaths := strings.Split(userPath, ":")

	for _, subPath := range subPaths {
		if strings.Contains(subPath, ".dpm") {
			continue
		}

		cmd := args[0]
		binaryPath := path.Join(subPath, cmd)
		doesExist, _ := utils.DoesFileExist(binaryPath)
		if doesExist {
			remainder := args[1:]
			return runShellCommand(binaryPath, remainder...)
		}
	}

	return errors.New("command not found")
}

// runShellCommand is a helper function for running commands
// on the terminal
func runShellCommand(cmdName string, args ...string) error {
	command := exec.Command(cmdName, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
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
			return runShellCommand(binaryPath, remainder...)
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

	return attachToContainer(dockerClient, container, args)
}

func attachToContainer(dockerClient *docker.Client, container *container.ContainerCreateCreatedBody, args []string) error {
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
		Cmd:          args,
	})
	if err != nil {
		return err
	}

	resp, err := dockerClient.ContainerExecAttach(ctx, exec.ID, types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	})
	if err != nil {
		return err
	}
	defer resp.Close()

	_, err = io.Copy(os.Stdout, resp.Reader)

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

	err = pullImageIfNotInDockerHost(dockerClient, imageName)
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

func pullImageIfNotInDockerHost(dockerClient *docker.Client, imageName string) error {
	ctx := context.Background()
	images, err := dockerClient.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}

	// don't pull if image already in host
	for _, image := range images {
		for _, repoTag := range image.RepoTags {
			if repoTag == imageName {
				return nil
			}
		}
	}

	reader, err := dockerClient.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err := reader.Close(); err != nil {
			log.Println("Error: ", err)
		}
	}()

	_, err = io.Copy(os.Stdout, reader)

	return err
}
