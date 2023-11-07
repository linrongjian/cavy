package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type XormProxy struct {
	Xorm *xorm.Engine
}

func (my *XormProxy) Connect(addr string, name string, account string, password string) error {
	// sqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", account, password, addr, name)
	// var err error
	// my.Xorm, err := xorm.NewEngine("mysql", sqlInfo)
	// if err != nil
	// return err

	// my.Xorm.ShowSQL()

	// err = my.Xorm.Ping()
	// if err != nil
	// return err

	return nil
}
