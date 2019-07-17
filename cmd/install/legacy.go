package install

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/JPZ13/dpm/internal/parser"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/utils"
)

// YAMLPackages writes bash scripts for
// each of the commands listed in the dpm.yml file
func YAMLPackages() error {
	commands := parser.GetCommands(project.ProjectFilePath)
	data, err := ioutil.ReadFile(project.ProjectFilePath)
	utils.HandleFatalError(err)

	// TODO: figure out what this is used for
	err = ioutil.WriteFile(path.Join(project.ProjectCmdPath, "dpm.yml"), data, 0644)
	utils.HandleFatalError(err)

	fmt.Printf("Installing %d commands...\n", len(commands))

	commandNames := []string{}

	for name, command := range commands {
		commandNames = append(commandNames, name)
		cliCommands := commandToDockerCLIs(command)
		err = writeDockerBashCommands(cliCommands)
		utils.HandleFatalError(err)
	}

	fmt.Printf("Installed: %s\n", strings.Join(commandNames, ", "))

	fmt.Print("Now you can run `dpm activate` to start using your new commands\n")

	return nil
}

// writeDockerBashCommands :
func writeDockerBashCommands(cliCommands map[string]string) error {
	for entrypoint, bashCommand := range cliCommands {
		targetPath := path.Join(project.ProjectCmdPath, entrypoint)
		contents := fmt.Sprintf("#!/bin/sh\nexec %s", bashCommand)

		err := ioutil.WriteFile(targetPath, []byte(contents), 0755)
		utils.HandleFatalError(err)
	}

	return nil
}

// commandToDockerCLIs takes a command and translates it into
// a docker cli command. It returns a map of the entrypoint to the
// docker command
func commandToDockerCLIs(command parser.Command) map[string]string {
	volumes := ""
	for _, volume := range command.Volumes {
		volumes = fmt.Sprintf("%s -v %s", volumes, volume)
	}

	cliCommands := make(map[string]string)
	for _, entrypoint := range command.Entrypoints {
		cliCommands[entrypoint] = fmt.Sprintf("docker run -it --rm -v $(pwd):%s %s -w %s --entrypoint %s %s \"$@\"",
			command.Context, volumes, command.Context, entrypoint, command.Image)
	}

	return cliCommands
}
