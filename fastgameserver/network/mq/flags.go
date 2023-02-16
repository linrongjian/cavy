package mq

import "github.com/urfave/cli/v2"

var (
	Opts = &struct {
		Host     string
		Account  string
		Password string
	}{}

	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "MQHOST",
			Value:       "127.0.0.1",
			Usage:       "-MQHOST 127.0.0.1",
			EnvVars:     []string{"MQHOST"},
			Destination: &Opts.Host,
		},
		&cli.StringFlag{
			Name:        "MQACCOUNT",
			Value:       "guest",
			Usage:       "-MQACCOUNT guest",
			EnvVars:     []string{"MQACCOUNT"},
			Destination: &Opts.Account,
		},
		&cli.StringFlag{
			Name:        "MQPASSWORD",
			Value:       "guest",
			Usage:       "-MQPASSWORD guest",
			EnvVars:     []string{"MQPASSWORD"},
			Destination: &Opts.Password,
		},
	}
)
