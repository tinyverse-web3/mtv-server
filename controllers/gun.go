package controllers

import (
	"encoding/json"
	"mtv/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// GUN相关API
type GunController struct {
	BaseController
}

// @Title Register
// @Description 注册(需验签)
// @Param public_key header string true "public key"
// @Param sign header string true "加签数据"
// @Param address header string true "钱包地址"
// @Param name body string true "名称"
// @Success 200 {object} controllers.RespJson
// @router /register [post]
func (c *GunController) Register() {
	curUser := c.CurUser()

	var gun models.Gun
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &gun)

	gun.UserId = curUser.Id

	o := orm.NewOrm()
	err := o.Read(&gun, "name")
	if err == nil {
		logs.Error("GUN已存在")
		c.ErrorJson("400000", "GUN已存在")
		return
	}

	gun.UserId = curUser.Id
	err = o.Read(&gun, "user_id", "name")
	if err == nil {
		logs.Error("用户已设置GUN")
		c.ErrorJson("400000", "用户已设置GUN")
		return
	}

	_, err = o.Insert(&gun)
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "Add GUN faild!")
		return
	}

	c.SuccessJson("", "")
}

// @Title List
// @Description 获取GUN列表(需验签)
// @Param public_key header string true "public key"
// @Param sign header string true "加签数据"
// @Param address header string true "钱包地址"
// @Success 200 {object} controllers.RespJson
// @router /list [get]
func (c *GunController) List() {
	var data []models.Gun
	gun := new(models.Gun)
	qt := orm.NewOrm().QueryTable(gun)
	qt.All(&data, "user_id", "name")
	c.SuccessJson("", data)
}
