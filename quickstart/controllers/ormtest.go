package controllers

import (
	"encoding/json"
	"quickstart/models"

	"github.com/astaxie/beego"
)

type OrmController struct {
	beego.Controller
}

func (c *OrmController) Get() {
	userInfo := models.AddUser()
	json, _ := json.Marshal(*userInfo)
	c.Ctx.WriteString("已经添加了用户" + string(json))
}
