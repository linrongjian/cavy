package mqwrap

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
			Value:       "127.0.0.1:5672/xw-debug",
			Usage:       "-MQHOST 127.0.0.1",
			EnvVars:     []string{"MQHOST"},
			Destination: &Opts.Host,
		},
		&cli.StringFlag{
			Name:        "MQACCOUNT",
			Value:       "xw",
			Usage:       "-MQACCOUNT guest",
			EnvVars:     []string{"MQACCOUNT"},
			Destination: &Opts.Account,
		},
		&cli.StringFlag{
			Name:        "MQPASSWORD",
			Value:       "123456",
			Usage:       "-MQPASSWORD guest",
			EnvVars:     []string{"MQPASSWORD"},
			Destination: &Opts.Password,
		},
	}
)
