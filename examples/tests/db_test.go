package tests

import (
	"os"

	"github.com/linrongjian/cavy/core/logger"
	"github.com/linrongjian/cavy/core/store/redis"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: redis.Flags,
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(redis.Opts.RedisUrl)
}
