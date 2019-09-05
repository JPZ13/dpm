package cmd

import (
	"github.com/JPZ13/dpm/cmd/activate"
	"github.com/JPZ13/dpm/cmd/deactivate"
	"github.com/JPZ13/dpm/cmd/run"
	"github.com/urfave/cli"
)

// CLI is the base cli
var CLI = cli.NewApp()

func init() {
	CLI.Commands = []cli.Command{
		{
			Name:  "activate",
			Usage: "Activate the project in the current shell",
			Action: func(c *cli.Context) error {
				activate.Command(c.Args())
				return nil
			},
		},
		{
			Name:            "run",
			Usage:           "Run an alias that is defined in the dpm yaml file",
			SkipFlagParsing: true,
			Action: func(c *cli.Context) error {
				run.Command(c.Args())
				return nil
			},
		},
		{
			Name:  "deactivate",
			Usage: "Deactivate the project in the current shell",
			Action: func(c *cli.Context) error {
				deactivate.Command(c.Args())
				return nil
			},
		},
	}
}
