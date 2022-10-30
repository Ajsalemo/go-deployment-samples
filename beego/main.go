package main

import "github.com/beego/beego/v2/server/web"
import "github.com/azureossd/go-deployment-samples/beego/controller"

func main() {
	web.Router("/", &controller.MainController{})

	web.Run(":8080")
}