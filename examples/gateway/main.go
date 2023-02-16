/*
 * @Author: zheng
 * @Date: 2019-08-09 18:50:07
 * @Description:
 */
package main

import (
	"fastgameserver/fastgameserver/logger"
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"runtime/pprof"
	"strconv"
	"time"

	"os"
	"runtime"

	_ "MiniGamePushServer/handles/connect"
)

var (
	cfgFilepath    = ""
	redisServerURL = ""
	serverUUID     = ""
	genInPath      = ""
	genOutPath     = ""
)

var (
	// 显示版本
	showVer = false
	// BuildVersion 编译版本
	BuildVersion string
	//BuildTime 编译时间
	BuildTime string
	//CommitID 提供ID
	CommitID string
)

func init() {
	flag.StringVar(&cfgFilepath, "c", "servercfg/x.json", "specify the config file path name")
	flag.StringVar(&serverUUID, "u", "", "specify the server UUID")
	flag.StringVar(&redisServerURL, "r", "", "redis server address")
	flag.StringVar(&genInPath, "gi", "", "input path")
	flag.StringVar(&genOutPath, "go", "", "output path")
}

func main() {

	runtime.GOMAXPROCS(1)

	flag.BoolVar(&showVer, "v", false, "show version")

	flag.Parse()

	if showVer {
		fmt.Println("Build Version:", BuildVersion)
		fmt.Println("Build Time:", BuildTime)
		fmt.Println("CommitID:", CommitID)
		os.Exit(0)
	}

	if genInPath != "" && genOutPath != "" {
		gamecfg.Gen(genInPath, genOutPath)
		os.Exit(0)
	}

	if redisServerURL != "" {
		servercfg.RedisServer = redisServerURL
	}

	if serverUUID != "" {
		serverInt, err := strconv.Atoi(serverUUID)
		if err != nil {
			log.Fatal("serverUUID must be integer")
		}
		servercfg.ServerID = serverInt
	}

	if cfgFilepath == "" {
		// 如果没有配置json文件，则必须提供uuid以及redis地址
		if serverUUID == "" || redisServerURL == "" {
			log.Fatal("must provide redis and uuid when json config file is omit")
		}
	}

	if cfgFilepath != "" {
		r := servercfg.ParseConfigFile(cfgFilepath)
		if r != true {
			log.Fatal("can't parse configure file:", cfgFilepath)
		}
	} else {
		log.Fatal("please specify a valid config file path")
	}

	// server startTime
	servercfg.StartTime = int(time.Now().Unix())

	log.Println("try to start  stupid server...")

	// start http server
	server.CreateHTTPServer()

	// clear old online
	server.ClearOnline()

	log.Println("start stupid server ok!")

	if servercfg.Daemon == "yes" {
		server.WaitForSignal()
	} else {
		waitInput()
	}
	return
}

func waitInput() {
	var cmd string
	for {
		_, err := fmt.Scanf("%s\n", &cmd)
		if err != nil {
			continue
		}

		switch cmd {
		case "exit", "quit":
			log.Println("exit by user")
			return
		case "gr":
			log.Println("current goroutine count:", runtime.NumGoroutine())
			break
		case "gd":
			pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
			break
		default:
			break
		}
	}
}

func ClearOnline() {
	conn := pool.Get()
	defer conn.Close()

	serveridstr := fmt.Sprintf("%d", servercfg.ServerID)

	conn.Send("MULTI")
	conn.Send("DEL", rconst.SetStatisticsOnlineNewPlayerPrefix+serveridstr)
	conn.Send("DEL", rconst.SetStatisticsOnlineOldPlayerPrefix+serveridstr)
	conn.Send("DEL", rconst.SetOnlinePrefix+serveridstr)
	_, err := conn.Do("EXEC")
	if err != nil {
		logger.Error("exec err, err:", err.Error())
		return
	}
}
