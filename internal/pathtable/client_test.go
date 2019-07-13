package pathtable

import "testing"

func TestPathTableClient(t *testing.T) {
	client := makeTestClient()

	project := makeProjectInfo()
	goAlias := makeGoAliasInfo()

	project.Commands = append(project.Commands, goAlias)

	testSet(t, client, project)

	testGet(t, client, project)

	removeTestLocationFile(t)
}
