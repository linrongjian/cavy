package gamehttp

import "github.com/urfave/cli/v2"

var (
	Opts = &struct {
		Port int
		Test bool
	}{}

	Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "HTTP_PORT",
			Value:       7001,
			Usage:       "-HTTP_PORT 7001",
			EnvVars:     []string{"HTTP_PORT"},
			Destination: &Opts.Port,
		},
		&cli.BoolFlag{
			Name:        "HTTP_TEST",
			Value:       true,
			Usage:       "-HTTP_TEST true",
			EnvVars:     []string{"HTTP_TEST"},
			Destination: &Opts.Test,
		},
	}
)
