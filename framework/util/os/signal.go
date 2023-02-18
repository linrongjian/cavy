package os

import (
	"fastgameserver/fastgameserver/logger"
	"os"
	"runtime"
	"runtime/pprof"
)

func dumpGoRoutinesInfo() {
	logger.Info("current goroutine count:", runtime.NumGoroutine())
	// use DEBUG=2, to dump stack like golang dying due to an unrecovered panic.
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
}

func reLoadConfig() {

}
