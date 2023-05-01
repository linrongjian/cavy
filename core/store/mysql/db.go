package mysql

import (
	"eventgo/core/logger"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	showSQL   = false
	DbConnect *xorm.Engine
)

func init() {
	flag.BoolVar(&showSQL, "showsql", false, "显示调用数据库sql")
}

func Startup() {
	if Opts.DbName == "" || Opts.Addr == "" || Opts.User == "" {
		logger.Error("Must specify the DbServer info in config json")
		return
	}
	logger.Infof("connect db addr:%s name:%s account:%s password:%s", Opts.Addr, Opts.DbName, Opts.User, Opts.Password)

	var err error
	DbConnect, err = connect(Opts.Addr, Opts.DbName, Opts.User, Opts.Password)
	if err != nil || DbConnect == nil {
		logger.Errorf("connect log db addr:%s name:%s account:%s password:%s err:%v", Opts.Addr, Opts.DbName, Opts.User, Opts.Password, err)
	}

	if err = createLogTables(DbConnect); err != nil {
		logger.Errorf("createLogTables err:%v", err)
	}
}

// 连接数据库
func connect(addr string, name string, account string, password string) (*xorm.Engine, error) {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", account, password, addr, name)
	engine, err := xorm.NewEngine("mysql", sqlInfo)

	if err != nil {
		return nil, err
	}

	if showSQL {
		engine.ShowSQL()
	}

	err = engine.Ping()
	if err != nil {
		return nil, err
	}
	return engine, nil
}

func createLogTables(engine *xorm.Engine) error {

	// TODO
	// loginLog := &tables.LoginLog{}
	// logoutLog := &tables.LogoutLog{}

	// // 创建表
	// if err := engine.CreateTables(
	// 	loginLog,
	// 	logoutLog); err != nil {
	// 	logger.Errorf("CreateTable Player err:%v", err)
	// 	return err
	// }

	// // 同步表结构
	// if err := engine.Sync2(
	// 	loginLog,
	// 	logoutLog); err != nil {
	// 	logger.Errorf("Syn2 Tables err:%v", err)
	// 	return err
	// }
	return nil
}
