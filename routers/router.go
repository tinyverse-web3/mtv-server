// @APIVersion V0
// @Title MTV API
// @Description MTV API
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"mtv/controllers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	logs.Info("router start")

	ns := beego.NewNamespace("/v0",
		// beego.NSAutoRouter(&controllers.AuthController{}),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/guardian",
			beego.NSInclude(
				&controllers.GuardianController{},
			),
		),
		beego.NSNamespace("/gun",
			beego.NSInclude(
				&controllers.GunController{},
			),
		),
		// beego.NSAutoRouter(&controllers.ImController{}),
		beego.NSNamespace("/im",
			beego.NSInclude(
				&controllers.ImController{},
			),
		),
		// beego.NSAutoRouter(&controllers.QuestionController{}),
		beego.NSNamespace("/question",
			beego.NSInclude(
				&controllers.QuestionController{},
			),
		),
		// beego.NSAutoRouter(&controllers.StorageController{}),
		beego.NSNamespace("/storage",
			beego.NSInclude(
				&controllers.StorageController{},
			),
		),
		// beego.NSAutoRouter(&controllers.UserController{}),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	logs.Info("router end")

}
