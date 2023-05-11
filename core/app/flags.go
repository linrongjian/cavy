package app

import "github.com/urfave/cli/v2"

var (
	DefaultFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "env",
			EnvVars: []string{"env"},
			Usage:   "set up the environment",
		},
	}
)
