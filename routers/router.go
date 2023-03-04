package routers

import (
	"mtv/controllers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	// var LogInFilter = func(ctx *context.Context) {

	// 	logs.Info("LogInFilter start")
	// 	uri := ctx.Request.RequestURI
	// 	token := ctx.Input.Session("token")

	// 	logs.Info(uri)
	// 	if (uri != "/v0/user/login" || uri != "/v0/user/sendmail") {
	// 		if (token == nil) {
	// 			ctx.Redirect(400, "/login")
	// 		} else {

	// 		}
	// 	}

	// 	logs.Info("LogInFilter end")
	// }

	// beego.InsertFilter("/v0/*", beego.BeforeRouter, LogInFilter)

	logs.Info("router start")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v0",
		beego.NSAutoRouter(&controllers.QuestionController{}),
		beego.NSAutoRouter(&controllers.StorageController{}),
		beego.NSAutoRouter(&controllers.UserController{}),
	)
	beego.AddNamespace(ns)
	// beego.Router("/", &controllers.MainController{})
	// beego.AutoRouter(&controllers.StorageController{})
	logs.Info("router end")

}
