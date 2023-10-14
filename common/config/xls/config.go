package xls

import (
	"github.com/linrongjian/cavy/common/mlog"
)

// Config 配置数据
var Config struct {
	Third []ThirdConfig //渠道配置
}

var xlsfiles = map[string]cfg.XlsxFile{
	"Third.xlsx": &ThirdConfig{},
}

var path = "../../common/config/xls/"

// Init 加载xls配置
func Init() {
	log := mlog.NewLogger(nil)

	for key := range cfg.CfgMap {
		delete(cfg.CfgMap, key)
	}

	for file, val := range xlsfiles {
		cfg.CfgMap[file] = val
	}
	cfg.ReadConfig()

	log.Info("load xls config success")

	go cfg.WatchConfig(path)
}
