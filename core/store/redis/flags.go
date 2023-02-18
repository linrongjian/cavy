package redis

import "github.com/urfave/cli/v2"

var (
	Opts = &struct {
		RedisUrl      string // URL
		RedisPassword string // 密码
		RedisDb       int    // Redis Db
		RedisPrefix   string // 前缀
		RedisIdeConns int    // Redis 最小空闲连接数
		RedisMaxPool  int    // Redis 最大连接数
	}{}

	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "redis",
			Value:       "redis://10.10.60.16:16379",
			Usage:       "设置redis连接地址. 格式: [:password@]hostname:port/[db]",
			EnvVars:     []string{"REDIS_URL"},
			Destination: &Opts.RedisUrl,
		},
		&cli.IntFlag{
			Name:        "redis_db",
			Value:       0,
			Usage:       "设置redis数据库",
			EnvVars:     []string{"REDIS_DB"},
			Destination: &Opts.RedisDb,
		},
		&cli.IntFlag{
			Name:        "redis_ide_conns",
			Value:       40,
			Usage:       "设置redis最大空闲连接数",
			EnvVars:     []string{"REDIS_IDE_CONNS"},
			Destination: &Opts.RedisIdeConns,
		},
		&cli.IntFlag{
			Name:        "redis_max_pool",
			Value:       80,
			Usage:       "设置redis最大连接数",
			EnvVars:     []string{"REDIS_MAX_POOL"},
			Destination: &Opts.RedisMaxPool,
		},
		&cli.StringFlag{
			Name:        "redis_password",
			Value:       "",
			Usage:       "设置redis密码",
			EnvVars:     []string{"REDIS_PASSWORD"},
			Destination: &Opts.RedisPassword,
		},
		&cli.StringFlag{
			Name:        "redis_prefix",
			Value:       "train",
			Usage:       "设置redis前缀",
			EnvVars:     []string{"REDIS_PREFIX"},
			Destination: &Opts.RedisPrefix,
		},
	}
)
