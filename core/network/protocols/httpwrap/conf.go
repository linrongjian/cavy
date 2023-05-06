package httpwrap

import "github.com/urfave/cli/v2"

type Config struct {
	Port int  `csv:"noHolidaysTime"`
	Test bool `csv:"holidaysTime"`
}

var (
	Conf = Config{
		Port: 15001,
		Test: true,
	}

	Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "HTTP_PORT",
			Value:       5001,
			Usage:       "-HTTP_PORT 7001",
			EnvVars:     []string{"HTTP_PORT"},
			Destination: &Conf.Port,
		},
		&cli.BoolFlag{
			Name:        "HTTP_TEST",
			Value:       true,
			Usage:       "-HTTP_TEST true",
			EnvVars:     []string{"HTTP_TEST"},
			Destination: &Conf.Test,
		},
	}
)
