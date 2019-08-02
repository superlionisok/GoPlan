package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var Engine *xorm.Engine

func init() {

	var err error
	Engine, err = xorm.NewEngine("mysql", "root:123456@tcp(10.10.15.202:3306)/lionjihua?charset=utf8")
	if err != nil {
		fmt.Println("mysql connect err. err=", err.Error())
		return
	}
	Engine.ShowSQL(true)
	//	engine.ShowWarn=true
	err2 := Engine.Sync2(new(UserCreatePlan))
	if err2 != nil {
		fmt.Println("mysql sync2 err. err=", err2.Error())
		return
	}

}

type UserCreatePlan struct {
	ID   int `xorm:"pk autoincr"`
	Name string
	//CreateTime time.Time `xorm:"created"`
	//UpdateTime time.Time `xorm:"updated"`
	CreateTime time.Time
	UpdateTime time.Time
}
