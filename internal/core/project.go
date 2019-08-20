package core

import "context"

// Project holds methods related to configuring
// a project
type Project interface {
	InstallProject(ctx context.Context, dpmFileLocation string) error
}

type project struct {
	baseService
}

// InstallProject installs a project from a dpm file but does not activate it
func (p *project) InstallProject(ctx context.Context, dpmFileLocation string) error {
	return nil
}
