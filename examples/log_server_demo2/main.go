package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/linrongjian/cavy"
)

var (
	n int32 = 0
)

func main() {

	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
	runtime.GOMAXPROCS(cpuNum - 1)

	p, err := cavy.NewLogProducer()
	if err != nil {
		log.Panicf("logconsumer err: %s", err)
	}

	// 声明只读通道
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			// log.Println("start Publish 10w")
			go func() {
				for i := 0; i < 300; i++ {
					// t2 := time.Now().UnixMilli()
					// if t2-t1 > 1000 {
					// 	t1 = t2
					// 	n = 0
					// }
					atomic.AddInt32(&n, 1)
					p.Options().RmqProducer.Publish([]byte(strconv.Itoa(int(n))))
				}
			}()
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Println(n)
			n = 0
		}
	}()

	select {}
}
