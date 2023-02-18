package client

import "github.com/urfave/cli/v2"

var (
	Opts = new(struct {
		Dev         bool // 是否开发模式
		Port        int
		ServerID    int
		Channel     string
		Daemon      bool
		CfgFilepath string
	})

	defaultFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "dev",
			Usage:       "设置是否为开发模式",
			Value:       true,
			EnvVars:     []string{"GAME_DEV"},
			Destination: &Opts.Dev,
		},
		&cli.IntFlag{
			Name:        "port",
			Usage:       "set port",
			Value:       3009,
			EnvVars:     []string{"GAME_PORT"},
			Destination: &Opts.Port,
		},
		&cli.IntFlag{
			Name:        "serverId",
			Usage:       "set port",
			Value:       1,
			EnvVars:     []string{"GAME_SERVER_ID"},
			Destination: &Opts.ServerID,
		},
		&cli.StringFlag{
			Name:        "channel",
			Usage:       "set port",
			Value:       "",
			EnvVars:     []string{"CHANNEL"},
			Destination: &Opts.Channel,
		},
		&cli.BoolFlag{
			Name:        "Daemon",
			Usage:       "后台运行",
			Value:       true,
			EnvVars:     []string{"GAME_DEV"},
			Destination: &Opts.Daemon,
		},
		&cli.StringFlag{
			Name:        "CfgFilepath",
			Usage:       "set CfgFilepath",
			Value:       "../game_conf.json",
			EnvVars:     []string{"CFG_FILE_PATH"},
			Destination: &Opts.CfgFilepath,
		},
	}
)
