package controllers

import (
	"encoding/json"
	"mtv/models"
	"mtv/utils"
	"mtv/utils/crypto"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// 守护者相关API
type GuardianController struct {
	BaseController
}

type GuardianInfo struct {
	Type       string `json:"type"`
	Account    string `json:"account"`
	VerifyCode string `json:"verifyCode"`
}

// @Title Add
// @Description 添加守护者
// @Param public_key header string true "public key"
// @Param sign header string true "加签数据"
// @Param address header string true "钱包地址"
// @Param type body string true "守护者账号类型"
// @Param account body string true "守护者账号"
// @Param verifyCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /add [post]
func (c *GuardianController) Add() {
	curUser := c.CurUser()

	var tmp GuardianInfo
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &tmp)

	hashAccount := crypto.Md5(tmp.Account)

	var guardian models.Guardian
	guardian.Type = tmp.Type
	guardian.Account = hashAccount

	verifyCode := tmp.VerifyCode
	if guardian.Type == "mail" {
		tmpVerifyCode := utils.GetStr(guardian.Account)
		if tmpVerifyCode == "" {
			logs.Error("验证码超时")
			c.ErrorJson("400000", "验证码超时，请重新获取")
			return
		}
		if verifyCode != tmpVerifyCode {
			logs.Error("验证码错误")
			c.ErrorJson("400000", "验证码错误")
			return
		}

	}

	o := orm.NewOrm()
	err := o.Read(&guardian, "type", "account")
	if err == nil {
		logs.Error("守护者已存在")
		c.ErrorJson("400000", "守护者已存在")
		return
	}

	guardian.UserId = curUser.Id
	guardian.AccountMask = utils.Mask(tmp.Account)
	_, err = o.Insert(&guardian)
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "Add guardian faild!")
		return
	}

	c.SuccessJson("", "")
}

// @Title List
// @Description 获取守护者列表
// @Param public_key header string true "public key"
// @Param sign header string true "加签数据"
// @Param address header string true "钱包地址"
// @Success 200 {object} controllers.RespJson
// @router /list [get]
func (c *GuardianController) List() {
	curUser := c.CurUser()

	var data []models.Guardian

	guardian := new(models.Guardian)
	qt := orm.NewOrm().QueryTable(guardian)
	qt.Filter("user_id", curUser.Id).Exclude("account__in", curUser.Email).All(&data, "type", "account", "accountMask")
	c.SuccessJson("", data)
}
