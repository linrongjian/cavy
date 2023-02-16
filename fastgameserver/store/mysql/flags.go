package mysql

import "github.com/urfave/cli/v2"

var (
	Opts = &struct {
		DbName   string
		Addr     string
		User     string
		Password string
	}{}

	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "dbname",
			Value:       "gamecenter",
			Usage:       "设置数据库名字",
			EnvVars:     []string{"MYSQL_DBNAME"},
			Destination: &Opts.DbName,
		},
		&cli.StringFlag{
			Name:        "addr",
			Value:       "10.10.60.16:3306",
			Usage:       "设置数据库地址",
			EnvVars:     []string{"MYSQL_ADDR"},
			Destination: &Opts.Addr,
		},
		&cli.StringFlag{
			Name:        "user",
			Value:       "root",
			Usage:       "设置数据库用户",
			EnvVars:     []string{"MYSQL_USER"},
			Destination: &Opts.User,
		},
		&cli.StringFlag{
			Name:        "password",
			Value:       "123456",
			Usage:       "设置数据库密码",
			EnvVars:     []string{"MYSQL_PASSWORD"},
			Destination: &Opts.Password,
		},
	}
)
