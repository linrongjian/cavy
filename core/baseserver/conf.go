package baseserver

import "github.com/urfave/cli/v2"

type Config struct {
	ServerID string `csv:"cfgId"`
	Port     int    `csv:"noHolidaysTime"`
	Dev      bool   `csv:"holidaysTime"`
}

var (
	Conf = Config{
		ServerID: "g1",
		Port:     7001,
		Dev:      true,
	}

	//TODO
	defaultFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "dev",
			Usage:       "设置是否为开发模式",
			Value:       true,
			EnvVars:     []string{"GAME_DEV"},
			Destination: &Conf.Dev,
		},
	}
)
