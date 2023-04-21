package controllers

import (
	"encoding/json"
	"hash/fnv"
	"math/rand"
	"mtv/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type QuestionController struct {
	BaseController
}

// @Title TmpList
// @Description 获取问题模版列表。
// @Param type query type true "如果type为1，则表示返回默认问题列表；如果type为2，则表示返回自定义问题模版列表，且随机返回3个问题模版。"
// @Success 200 {object} controllers.RespJson
// @router /tmplist [get]
func (c *QuestionController) TmpList() {
	CurUser := c.CurUser()

	qType, _ := c.GetInt("type")
	logs.Info("type = ", qType)

	var data []models.QuestionTmp
	var tmp []models.QuestionTmp

	question := new(models.QuestionTmp)
	qt := orm.NewOrm().QueryTable(question)
	qt.Filter("type", qType).OrderBy("id").All(&tmp, "title", "content")

	if qType == 2 {
		h := fnv.New64a()
		h.Write([]byte(CurUser.Email))
		seed := int64(h.Sum64())
		rand.Seed(seed)

		for i := 0; i <= 2; i++ {
			index := rand.Intn(len(tmp))
			logs.Info("index = ", index)
			item := tmp[index]
			data = append(data, item)
		}
	} else {
		data = tmp
	}

	c.SuccessJson("", data)
}

// @Title List
// @Description 获取用户问题列表(type为1表示默认问题；type为2表示自定义问题)
// @Success 200 {object} controllers.RespJson
// @router /list [get]
func (c *QuestionController) List() {
	CurUser := c.CurUser()
	var data []models.Question
	question := new(models.Question)
	qt := orm.NewOrm().QueryTable(question)
	qt.Filter("user_id", CurUser.Id).All(&data, "type", "content")
	c.SuccessJson("", data)
}

type QuestionInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"` //问题内容(json格式)
	Type    int    `json:"type"`    // 1：默认问题；2：自定义问题
}

// @Title Add
// @Description 设置用户问题(type为1表示默认问题；type为2表示自定义问题)
// @Param questions body controllers.QuestionInfo true "问题列表"
// @Success 200 {object} controllers.RespJson
// @router /add [post]
func (c *QuestionController) Add() {
	CurUser := c.CurUser()
	question := new(models.Question)
	question.UserId = CurUser.Id
	o := orm.NewOrm()
	o.Delete(question, "user_id")

	var questions []QuestionInfo
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &questions)
	for _, q := range questions {
		question = &models.Question{UserId: CurUser.Id, Content: q.Content, Type: q.Type}
		o.Insert(question)
	}

	c.SuccessJson("", "")
}
