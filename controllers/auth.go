package controllers

import (
	"github.com/beego/beego/v2/core/logs"
)

// Auth API
type AuthController struct {
	BaseController
}

// @Title CheckSign
// @Description 验签
// @Param public_key header string true "public key"
// @Param sign header string true "加签数据"
// @Param address header string true "钱包地址"
// @Param data header string true "未加签数据"
// @Success 200 {object} controllers.RespJson
// @Failure 401 Unauthorized
// @router /checksign [get]
func (c *AuthController) CheckSign() {

	var publicKey string
	tmp := c.Ctx.Request.Header["Public_key"]
	if tmp == nil {
		logs.Error("public_key不能为空")
		c.Abort("401")
		return
	} else {
		publicKey = tmp[0]
	}
	logs.Info("public key = ", publicKey)

	var signature string
	tmp = c.Ctx.Request.Header["Sign"]
	if tmp == nil {
		logs.Error("sign不能为空")
		c.Abort("401")
		return
	} else {
		signature = tmp[0]
	}
	logs.Info("sign = ", signature)

	var address string
	tmp = c.Ctx.Request.Header["Address"]
	if tmp == nil {
		logs.Error("Address不能为空")
		c.Abort("401")
		return
	} else {
		address = tmp[0]
	}
	logs.Info("address = ", address)

	var data string
	tmp = c.Ctx.Request.Header["Data"]
	if tmp == nil {
		logs.Error("data不能为空")
		c.Abort("401")
		return
	} else {
		data = tmp[0]
	}
	logs.Info("data = ", data)

	match := sign(address, data, signature, publicKey)
	if !match {
		logs.Error("验签失败")
		c.Abort("401")
		return
	}
	c.SuccessJson("", "")
}
