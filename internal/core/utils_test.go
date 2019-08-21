package core

import "github.com/JPZ13/dpm/internal/model"

func makeTestGlideCommand() *model.Command {
	return &model.Command{
		Image:       "dockerepo/glide",
		Entrypoints: []string{"glide"},
	}
}

func makeTestGoCommand() *model.Command {
	return &model.Command{
		Image:       "golang:1.7.5",
		Entrypoints: []string{"go"},
	}
}

func makeTestPythonCommand() *model.Command {
	return &model.Command{
		Image:       "python:latest",
		Entrypoints: []string{"python", "pip"},
	}
}

func makeTestDPMFile() *model.DPMFile {
	glideCmd := makeTestGlideCommand()
	goCmd := makeTestGoCommand()
	pythonCmd := makeTestPythonCommand()

	cmdMap := map[string]model.Command{
		"glide":  *glideCmd,
		"go":     *goCmd,
		"python": *pythonCmd,
	}

	return &model.DPMFile{
		Commands: cmdMap,
	}
}
