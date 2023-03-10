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

func (c *ImController) Relays() {
	var data []models.NostrRelay

	relay := new(models.NostrRelay)
	qt := orm.NewOrm().QueryTable(relay)
	qt.Filter("status", 1).Limit(20, 1).All(&data, "WsServer")
	c.SuccessJson("", data)
}

func (c *ImController) CreateShareIm() {
	var chat models.Chat

	fromPublicKey := CurUser.NostrPublicKey
	chat.FromPublicKey = fromPublicKey

	o := orm.NewOrm()
	err := o.Read(&chat, "from_public_key")
	if err == orm.ErrNoRows {
		chat = models.Chat{FromPublicKey: fromPublicKey, Status: 1}
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
	PublicKey string `json:"publicKey"`
}

func (c *ImController) ExchangeImPkey() {
	var info ImPublicKey
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &info)

	var chat models.Chat
	o := orm.NewOrm()
	chat.FromPublicKey = CurUser.NostrPublicKey
	err := o.Read(&chat, "from_public_key")
	if err == orm.ErrNoRows {
		logs.Error(err)
		c.ErrorJson("400000", "聊天室不存在")
		return
	}

	if chat.Status == 0 {
		c.ErrorJson("400000", "聊天室已超时")
		return
	}

	curTime := time.Now()
	if curTime.Sub(chat.UpdateTime).Minutes() > 10 {
		chat.UpdateTime = time.Now()
		chat.Status = 0
		_, err := o.Update(&chat)
		if err != nil {
			logs.Error(err)
		}
		c.ErrorJson("400000", "聊天室已超时")
		return
	}

	var chatNotify models.ChatNotify
	chatNotify.FromPublicKey = CurUser.NostrPublicKey
	chatNotify.ToPublicKey = info.PublicKey
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
	c.SuccessJson("", "")

	//websock server推送给创建者
}

func (c *ImController) Notify() {
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
