package controller

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	r := make(map[string]string)
	r["message"] = "go-deployment-samples-beego"

	this.Data["json"] = r
	this.ServeJSON()
}