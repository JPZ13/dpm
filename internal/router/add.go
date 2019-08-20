package router

import (
	"fmt"

	"github.com/JPZ13/dpm/internal/utils"
)

func (r *router) Add(alias string) error {
	hasAlias, err := r.Has(alias)
	if err != nil {
		return err
	}

	if !hasAlias {
		aliasPath := r.getAliasPath(alias)
		return writeAliasBashScript(aliasPath, alias)
	}

	return nil
}

func writeAliasBashScript(location string, alias string) error {
	script := makeAliasBashScript(alias)
	return utils.WriteBashScript(location, script)
}

func makeAliasBashScript(alias string) string {
	// TODO: modify to include the path of the caller
	script := `dpm run %s "$@"`
	return fmt.Sprintf(script, alias)
}
