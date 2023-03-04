package main

import (
	_ "mtv/routers"
	"mtv/utils"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

func main() {
	// env, _ := config.String("runmode")
	// logs.Info("adapterName:", env)
	// configFilePath, err := filepath.Abs("conf/" + env + ".conf")
	// if err != nil {
	// 	logs.Error("configFilePathError", err)
	// }
	// err = beego.LoadAppConfig("ini", configFilePath)
	// if err != nil {
	// 	logs.Error("loadConfigFileError", err)
	// }

	//数据库初始化
	utils.InitMySQL()
	// utils.InitRedis()

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
