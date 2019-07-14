package router

import "path"

func (r *router) getAliasPath(alias string) string {
	return path.Join(r.baseDirectory, alias)
}
