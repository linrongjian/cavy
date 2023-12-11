package servercfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"cavy/common/api"
)

// Cfg server config
var Cfg *api.Conf

// ParseConfigFile 解析配置
func ParseConfigFile(configFile string) *api.StConfig {
	if configFile != "" {
		buf, err := ioutil.ReadFile(configFile)
		if err != nil {
			msg := fmt.Sprintf("read config file %s err:%v", configFile, err)
			panic(msg)
		}
		config := &api.StConfig{}
		err = json.Unmarshal(buf, config)
		if err != nil {
			msg := fmt.Sprintf("parse config file %s err:%v", configFile, err)
			panic(msg)
		}
		return config
	}
	return nil
}
