package main

import (
	_ "mtv/routers"
	"mtv/utils"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

func main() {

	//数据库初始化
	utils.InitMySQL()
	// utils.InitRedis()

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
