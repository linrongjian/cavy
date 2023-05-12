package mongodb

import (
	"fmt"
	"testing"
	"time"
)

type Town struct {
	TownId            string    `bson:"townId"`            // TownId
	Name              string    `bson:"name"`              // 角色名
	TownCode          string    `bson:"townCode"`          // 邀请码
	Diamond           int       `bson:"diamond"`           // 钻石
	Coin              int       `bson:"coin"`              // 金币
	Shell             int       `bson:"shell"`             // 贝壳
	Lv                int       `bson:"lvl"`               // 等级
	Exp               int       `bson:"exp"`               // 经验
	CurPopulation     int       `bson:"curPopulation"`     // 当前人口
	MaxPopulation     int       `bson:"maxPopulation"`     // 人口上限
	WarehouseLevel    int       `bson:"warehouseLevel"`    // 仓库等级
	WarehouseCapacity int       `bson:"warehouseCapacity"` // 仓库容量
	LastOnline        float64   `bson:"lastOnline"`        // 最后上线时间
	Channel           string    `bson:"channel"`           // 渠道
	CreateTime        time.Time `bson:"createTime"`        // 创建时间
}

func TestStore_Connect(t *testing.T) {
	db := NewStore(
		WithDbName("xw-debug"),
		WithMinPoolSize(20),
		WithMaxPoolSize(100),
		WithUrl("mongodb://user-debug:123456@127.0.0.1:27017/xw-debug"),
	)

	if err := db.Connect(); err != nil {
		panic(err)
	}
	cl := db.C("town")

	var townlist []Town
	FindAll(cl, nil, &townlist, nil)

	defer db.Disconnect()

	fmt.Printf("MongoDB连接成功\n")
}
