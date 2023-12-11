package httpsvr

import (
	"github.com/linrongjian/cavy/common/cache"
	"github.com/linrongjian/cavy/common/mlog"
)

//"battlePlatform/wjrgit.qianz.com/common/cache"
//"battlePlatform/wjrgit.qianz.com/common/mlog"

var (
	// maxreqpersec 连续1s内请求限制次数
	maxreqpersec = int32(100)
	// blacktime 黑名单时长
	blacktime = int64(15 * 60)
	// reqinterval
	reqinerval = int64(1)
)

// BlackList 黑名单策略
func BlackList(playerid int64) (bool, error) {
	log := mlog.NewLogger(nil)

	exist := cache.IsInBlackList(playerid)
	if exist {
		return false, nil
	}

	reqnum := cache.GetRequestNum(playerid)
	if reqnum > maxreqpersec {
		log.Warn("add black list !!!")
		err := cache.AddBlackList(playerid, blacktime)
		if err != nil {
			return true, err
		}

		return false, nil
	}

	reqnum++
	err := cache.AddRequestNum(playerid, reqinerval, reqnum)
	if err != nil {
		return true, err
	}

	return true, nil
}
