package activate

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/JPZ13/dpm/cmd/tools"
)

// Command handles the activate command
func Command(args []string) {
	core, err := tools.MakeCoreInHomeDirectory()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	ctx := context.Background()
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("DPM error: %s", err)
	}

	dpmLocation := path.Join(pwd, "dpm.yml")

	err = core.InstallProject(ctx, dpmLocation)
	if err != nil {
		log.Fatalf("DPM error: %s", err)
	}
}
