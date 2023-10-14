package xlsx

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/linrongjian/cavy/common/mlog"
)

type XlsxFile interface {
	Head() int
	Obj() interface{}
	Data() interface{}
}

// CfgMap 配置文件
var CfgMap = map[string]XlsxFile{}

// ReadConfig 读取xls配置
func ReadConfig() {
	log := mlog.NewLogger(nil)

	for k, v := range CfgMap {
		log.Debugf("Read Config %v", k)
		file := fmt.Sprintf("../../common/config/xls/%s", k) // golang相对路径是相对于执行命令时的目录
		buf, err := ReadFile(file)
		if err != nil {
			msg := fmt.Sprintf("ReadFile %s err:%v", k, err)
			panic(msg)
		}
		if err := NewSessionBinary(buf, "Sheet1", v.Head(), v.Obj()).Get(v.Data()); err != nil {
			msg := fmt.Sprintf("Read %s Config err:%v", k, err)
			panic(msg)
		}
	}
}

// ReloadConfig 重载配置
func ReloadConfig(k string) error {
	log := mlog.NewLogger(nil)

	if v, exist := CfgMap[k]; !exist || v == nil {
		log.Warnf("reload xls isn't in cfgMap:(%v)", k)
		return nil
	}
	file := fmt.Sprintf("config/xls/%s", k) // golang相对路径是相对于执行命令时的目录
	buf, err := ReadFile(file)
	if err != nil {
		msg := fmt.Sprintf("ReadFile %s err:%v", k, err)
		panic(msg)
	}
	v := CfgMap[k]
	if err := NewSessionBinary(buf, "Sheet1", v.Head(), v.Obj()).Get(v.Data()); err != nil {
		msg := fmt.Sprintf("Read %s Config err:%v", k, err)
		panic(msg)
	}
	log.Debugf("reload xls config:(%v)", v.Data())
	return nil
}

// WatchConfig 监听配置变化
func WatchConfig(path string) {
	log := mlog.NewLogger(nil)
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	watcher.Add(path)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok { // 'Events' channel is closed
				return
			}

			// we only care about the config file with the following cases:
			// if the xlsx file was modified or created
			const writeOrCreateMask = fsnotify.Write | fsnotify.Create
			if strings.HasSuffix(event.Name, ".xlsx") &&
				event.Op&writeOrCreateMask != 0 {
				_, fileName := filepath.Split(event.Name)
				log.Infof("write or create file:%v, name:%v", event.Name, fileName)
				ReloadConfig(fileName)
			}
		case err, ok := <-watcher.Errors:
			if ok { // 'Errors' channel is not closed
				log.Infof("watcher error: %v\n", err)
			}
			return
		}
	}
}

func ReadFile(fileName string) ([]byte, error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
