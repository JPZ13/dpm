package pathtable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	client := makeTestClient()

	project := makeProjectInfo()
	goAlias := makeGoAliasInfo()

	project.Commands = append(project.Commands, goAlias)

	err := client.Set(testLocation, project)
	require.NoError(t, err)

	testGet(t, client, project)

	removeTestLocationFile(t)
}
