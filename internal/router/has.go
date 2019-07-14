package router

import (
	"github.com/JPZ13/dpm/internal/utils"
)

func (r *router) Has(alias string) (bool, error) {
	err := r.ensureBaseDirectory()
	if err != nil {
		return false, err
	}

	aliasPath := r.getAliasPath(alias)
	return utils.DoesFileExist(aliasPath)
}
