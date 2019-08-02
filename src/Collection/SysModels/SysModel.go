package SysModels

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserOrder1 struct {
	ID         int    `db:"ID"`
	OrderTitle string `db:"OrderTitle"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default2", "mysql", "root:123456@tcp(10.10.15.202:3306)/lionjihua?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(UserOrder1))

	orm.RunSyncdb("default2", false, true)
}
