package main

import (
	"mtv/controllers"
	_ "mtv/routers"
	"mtv/task"
	"mtv/utils"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

func main() {

	//数据库初始化
	utils.InitMySQL()
	utils.InitRedis()
	go controllers.InitWebSocket()

	beego.AddFuncMap("i18n", i18n.Tr)
	task.StartTask()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
