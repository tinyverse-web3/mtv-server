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
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &chat)

	fromPublicKey := chat.FromPublicKey

	o := orm.NewOrm()
	err := o.Read(&chat, "from_public_key")
	if err == orm.ErrNoRows {
		chat = models.Chat{FromPublicKey: fromPublicKey}
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

	// createTime := chat.CreateTime
	// curTime := time.Now()
	// if curTime.Sub(createTime).Seconds() < 600 {
	// 	c.ErrorJson("400000", "聊天室已超时")
	// 	return
	// }

	c.SuccessJson("", "")
}

func (c *ImController) ExchangeImPkey() {
	//websock server推送给创建者
}
