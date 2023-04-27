package controllers

import (
	"encoding/json"
	"hash/fnv"

	"github.com/beego/beego/v2/core/logs"
)

type StorageController struct {
	BaseController
}

type Storage struct {
	Aaa string `json:"aaa"`
}

// @Title Test
// @Description 用于测试
// @Success 200 {object} controllers.RespJson
// @router /test [post]
func (c *StorageController) Test() {
	var storage Storage
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &storage)
	logs.Info("aaa = ", storage.Aaa)

	arg := c.GetString("arg")
	logs.Info("arg = ", arg)
	h := fnv.New64a()
	h.Write([]byte("18098922101@189.cn"))
	seed := h.Sum64()
	logs.Info(seed)

	var publicKey string
	tmp := c.Ctx.Request.Header["Public_key"]
	if tmp == nil {
		logs.Info("public key 为空")
		c.Abort("401")
	} else {
		publicKey = tmp[0]
	}
	logs.Info("public key = ", publicKey)

	c.SuccessJson("", seed)
	// c.Abort("403")
}

func (c *StorageController) Index() {
	c.TplName = "index.tpl"
}
