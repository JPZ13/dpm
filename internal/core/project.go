package core

import (
	"context"
	"io/ioutil"

	"github.com/JPZ13/dpm/internal/model"
	"gopkg.in/yaml.v2"
)

// Project holds methods related to configuring
// a project
type Project interface {
	InstallProject(ctx context.Context, dpmFileLocation string) error
}

type project struct {
	baseService
}

// InstallProject installs a project from a dpm file and activates it
func (p *project) InstallProject(ctx context.Context, dpmFileLocation string) error {
	// parse dpm file
	dpmFile := &model.DPMFile{}
	err := parseYAMLFile(dpmFileLocation, dpmFile)
	if err != nil {
		return err
	}

	// translate to project info
	projectInfo, err := translateDPMFileToProjectInfo(dpmFile)
	if err != nil {
		return err
	}

	projectInfo.IsActive = true

	// set project info
	err = p.pathTable.Set(dpmFileLocation, projectInfo)
	if err != nil {
		return err
	}

	// add all aliases
	return p.addAliasesToRouter(projectInfo)
}

func parseYAMLFile(location string, obj interface{}) error {
	fileBytes, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileBytes, obj)
	if err != nil {
		return err
	}

	return nil
}

func (p *project) addAliasesToRouter(projectInfo *model.ProjectInfo) error {
	for _, aliasInfo := range projectInfo.Commands {
		for alias := range aliasInfo.Aliases {
			err := p.router.Add(alias)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
