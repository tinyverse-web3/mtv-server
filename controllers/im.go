package controllers

import (
	"encoding/json"
	"mtv/models"
	"mtv/utils"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
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

// @Title SearchUser
// @Description 根据用户名或公钥查询用户，并返回用户信息(需验签)
// @Param param query string true "用户名或公钥"
// @Success 200 {object} controllers.RespJson
// @router /searchuser [get]
func (c *ImController) SearchUser() {
	param := c.GetString("param")
	logs.Info("param = ", param)

	o := orm.NewOrm()

	var user models.User
	user.Name = param
	err := o.Read(&user, "name")
	if err == orm.ErrNoRows {
		user.PublicKey = param
		err = o.Read(&user, "public_key")
		if err == orm.ErrNoRows {
			c.ErrorJson("400000", "用户不存在")
			return
		}
	}

	var data models.User
	data.NostrPublicKey = user.NostrPublicKey

	ipfsGateWay, _ := config.String("ipfs_gate_way")
	data.ImgCid = ipfsGateWay + "/" + user.ImgCid

	data.PublicKey = user.PublicKey

	c.SuccessJson("", data)
}

// @Title Friends
// @Description 获取当前用户的好友列表(需验签)
// @Success 200 {object} controllers.RespJson
// @router /friends [get]
func (c *ImController) Friends() {
	curUser := c.CurUser()
	logs.Info("cur user = ", curUser)

	ipfsGateWay, _ := config.String("ipfs_gate_way")
	o := orm.NewOrm()
	var users = []models.User{}
	_, err := o.Raw("select name, public_key, nostr_public_key, case when img_cid = '' then '' when img_cid <> '' then CONCAT(?, img_cid) end as img_cid from user where public_key in (select to_public_key from im_friend where from_public_key = ?)", ipfsGateWay, curUser.PublicKey).QueryRows(&users)
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "获取好友列表失败")
		return
	}

	c.SuccessJson("", users)
}

// @Title AddFriend
// @Description 添加好友(需验签)
// @Param toPublicKey body string true "public key"
// @Success 200 {object} controllers.RespJson
// @router /addfriend [post]
func (c *ImController) AddFriend() {
	curUser := c.CurUser()
	logs.Info("cur user = ", curUser)

	var info ImPublicKey
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &info)
	logs.Info("info = ", info)

	var friend models.ImFriend
	o := orm.NewOrm()
	friend.FromPublicKey = curUser.PublicKey
	friend.ToPublicKey = info.ToPublicKey
	err := o.Read(&friend, "from_public_key", "to_public_key")
	if err == orm.ErrNoRows {
		friend = models.ImFriend{FromPublicKey: friend.FromPublicKey, ToPublicKey: friend.ToPublicKey}
		_, err := o.Insert(&friend)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "添加好友失败")
			return
		} else {
			friend = models.ImFriend{FromPublicKey: friend.ToPublicKey, ToPublicKey: friend.FromPublicKey} // 互加好友
			_, err := o.Insert(&friend)
			if err != nil {
				logs.Error(err)
				c.ErrorJson("400000", "添加好友失败")
				return
			} else {
				// 添加到redis中，用于通过websocket通知被添加者
				key := "friend_" + friend.FromPublicKey
				logs.Info("key = ", key)
				tmp, _ := utils.GetStr(key)
				var names []string
				if tmp != "" {
					names = strings.Split(tmp, ",")
				}
				var user models.User
				user.PublicKey = friend.ToPublicKey

				err = o.Read(&user, "public_key")
				if err == nil {
					names = append(names, user.Name)
					logs.Info("names = ", names)
					utils.SetStr(key, strings.Join(names, ","), 24*time.Hour)

					c.SuccessJson("", friend.Status)
				} else {
					logs.Error(err)
					c.ErrorJson("400000", "添加好友失败")
					return
				}
			}
		}
	} else {
		c.SuccessJson("", friend.Status)
	}
}
