package controllers

import (
	"CreatePlan/Models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Email"] = "astaxie@gmail.com"

	var u Models.UserCreatePlan

	Models.Engine.Get(&u)
	logs.Info(u.Name)

	c.Data["Website"] = "beego.me" + u.Name
	c.TplName = "index.tpl"
}
