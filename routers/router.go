package routers

import (
	"mtv/controllers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	logs.Info("router start")

	ns := beego.NewNamespace("/v0",
		beego.NSAutoRouter(&controllers.QuestionController{}),
		beego.NSAutoRouter(&controllers.StorageController{}),
		beego.NSAutoRouter(&controllers.UserController{}),
	)
	beego.AddNamespace(ns)
	logs.Info("router end")

}
