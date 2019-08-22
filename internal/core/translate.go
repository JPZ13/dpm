package core

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/JPZ13/dpm/internal/model"
)

func translateDPMFileToProjectInfo(dpmFile *model.DPMFile) (*model.ProjectInfo, error) {
	aliases := []model.AliasInfo{}

	for _, cmd := range dpmFile.Commands {
		alias, err := translateCommandToAliasInfo(&cmd)
		if err != nil {
			return nil, err
		}

		aliases = append(aliases, *alias)
	}

	return &model.ProjectInfo{
		Commands: aliases,
	}, nil
}

func translateCommandToAliasInfo(cmd *model.Command) (*model.AliasInfo, error) {
	aliasTable, err := translateEntrypointsToAliasTable(&cmd.Entrypoints)
	if err != nil {
		return nil, err
	}

	return &model.AliasInfo{
		Aliases:    *aliasTable,
		Image:      cmd.Image,
		VolumeName: cmd.VolumeName,
	}, nil
}

func translateEntrypointsToAliasTable(entrypoints *[]string) (*map[string]string, error) {
	aliasTable := make(map[string]string)

	for _, entrypoint := range *entrypoints {
		binaryLocation, err := getBinaryLocation(entrypoint)
		if err != nil {
			return nil, err
		}

		aliasTable[entrypoint] = binaryLocation
	}

	return &aliasTable, nil
}

func getBinaryLocation(entrypoint string) (string, error) {
	cmd := exec.Command("which", entrypoint)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil && isCommandNotFoundError(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	location := out.String()
	location = strings.TrimSuffix(location, "\n")

	return location, nil
}

func isCommandNotFoundError(err error) bool {
	if err.Error() == "exit status 1" {
		return true
	}

	return false
}
