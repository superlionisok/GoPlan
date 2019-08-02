package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(10.10.15.202:3306)/superplan?charset=utf8")
	//orm.RegisterDataBase("default2", "mysql", "root:123456@tcp(10.10.15.202:3306)/lionjihua?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
	//orm.RunSyncdb("default2",false,true)
}

type Manager struct {
	ID        int    // id
	LoginName string // 登陆名称
	LoginPwd  string
	Age       int
	IsLock    int
	Email     string
}

type User struct {
	ID         int    // id
	LoginName  string // 登陆名称
	LoginPwd   string
	CreateTime time.Time
}
type SysUser struct {
	ID        int    // id
	Name      string // 登陆名称
	Age       int
	LoginName string
}

type UserOrder struct {
	ID         int    `db:"ID"`
	OrderTitle string `db:"OrderTitle"`
}
