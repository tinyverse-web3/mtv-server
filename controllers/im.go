package controllers

import (
	"encoding/json"
	"mtv/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type ImController struct {
	BaseController
}

// @Title Relays
// @Description 获取中继列表
// @Success 200 {object} controllers.RespJson
// @router /relays [get]
func (c *ImController) Relays() {
	var data []models.NostrRelay

	relay := new(models.NostrRelay)
	qt := orm.NewOrm().QueryTable(relay)
	qt.Filter("status", 1).Limit(20, 1).All(&data, "WsServer")
	c.SuccessJson("", data)
}

// @Title CreateShareIm
// @Description 分享聊天
// @Success 200 {object} controllers.RespJson
// @router /createshareim [post]
func (c *ImController) CreateShareIm() {
	var chat models.Chat
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &chat)

	o := orm.NewOrm()
	err := o.Read(&chat, "from_public_key")
	if err == orm.ErrNoRows {
		chat = models.Chat{FromPublicKey: chat.FromPublicKey, Status: 1}
		_, err := o.Insert(&chat)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "Share im faild!")
			return
		}
	} else {
		chat.UpdateTime = time.Now()
		chat.Status = 1
		_, err := o.Update(&chat)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "Share im faild!")
			return
		}
	}

	c.SuccessJson("", "")
}

type ImPublicKey struct {
	FromPublicKey string `json:"fromPublicKey"`
	ToPublicKey   string `json:"toPublicKey"`
}

// @Title ExchangeImPkey
// @Description 交换聊天公钥
// @Success 200 {object} controllers.RespJson
// @router /exchangeimpkey [post]
func (c *ImController) ExchangeImPkey() {
	var info ImPublicKey
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &info)
	logs.Info("info = ", info)

	var chat models.Chat
	o := orm.NewOrm()
	chat.FromPublicKey = info.FromPublicKey
	err := o.Read(&chat, "from_public_key")
	if err == orm.ErrNoRows {
		logs.Error(err)
		c.ErrorJson("400000", "聊天室不存在")
		return
	}
	logs.Info("111111111")
	if chat.Status == 0 {
		c.ErrorJson("400000", "聊天室已超时")
		return
	}
	logs.Info("2222222")
	curTime := time.Now()
	if curTime.Sub(chat.UpdateTime).Minutes() > 10 {
		chat.UpdateTime = time.Now()
		chat.Status = 0
		_, err := o.Update(&chat)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "聊天室已超时")
			return
		}
	}
	logs.Info("333333333")

	var chatNotify models.ChatNotify
	chatNotify.FromPublicKey = info.FromPublicKey
	chatNotify.ToPublicKey = info.ToPublicKey
	err = o.Read(&chatNotify, "from_public_key", "to_public_key")
	if err == orm.ErrNoRows {
		chatNotify = models.ChatNotify{FromPublicKey: chatNotify.FromPublicKey, ToPublicKey: chatNotify.ToPublicKey, Status: 1}
		_, err := o.Insert(&chatNotify)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "系统错误")
			return
		}
	} else {
		chatNotify.UpdateTime = time.Now()
		chatNotify.Status = 1
		_, err := o.Update(&chatNotify)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "系统错误")
			return
		}
	}
	logs.Info("44444444444444")

	c.SuccessJson("", "")

	//websock server推送给创建者
}

// @Title Notify
// @Description 获取聊天请求列表
// @Success 200 {object} controllers.RespJson
// @router /notify [get]
func (c *ImController) Notify() {
	CurUser := c.CurUser()
	fromPublicKey := CurUser.NostrPublicKey
	if fromPublicKey == "" {
		c.ErrorJson("400000", "Public Key不能为空")
		return
	}

	var tmp []models.ChatNotify
	var chatNotify models.ChatNotify
	o := orm.NewOrm()
	qt := o.QueryTable(chatNotify)
	qt.Filter("from_public_key", fromPublicKey).Filter("status", 1).All(&tmp, "toPublicKey")

	var data []models.ChatNotify
	curTime := time.Now()
	for _, value := range tmp {
		if curTime.Sub(value.UpdateTime).Minutes() > 10 {
			chatNotify.UpdateTime = time.Now()
			chatNotify.Status = 0
			_, err := o.Update(&chatNotify)
			if err != nil {
				logs.Error(err)
				continue
			}
		}
		data = append(data, value)
	}

	c.SuccessJson("", data)
}
