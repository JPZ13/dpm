package router

import (
	"path"

	"github.com/JPZ13/dpm/internal/utils"
)

func (r *router) getAliasPath(alias string) string {
	return path.Join(r.baseDirectory, alias)
}

func (r *router) ensureBaseDirectory() error {
	return utils.EnsureDirectory(r.baseDirectory)
}
