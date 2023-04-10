package controllers

import (
	"encoding/json"
	"hash/fnv"
	"math/rand"
	"mtv/models"

	"github.com/beego/beego/v2/client/orm"
)

type QuestionController struct {
	BaseController
}

// @Title TmpList
// @Description 获取问题模版列表
// @Success 200 {object} controllers.RespJson
// @router /tmplist [get]
func (c *QuestionController) TmpList() {
	CurUser := c.CurUser()
	var data []models.QuestionTmp
	var tmp []models.QuestionTmp
	question := new(models.QuestionTmp)
	qt := orm.NewOrm().QueryTable(question)
	qt.OrderBy("id").All(&tmp, "Id", "Content")

	h := fnv.New64a()
	h.Write([]byte(CurUser.Email))
	seed := int64(h.Sum64())
	rand.Seed(seed)

	for i := 0; i <= 2; i++ {
		index := rand.Intn(len(tmp))
		item := tmp[index]
		data = append(data, item)
	}
	c.SuccessJson("", data)
}

// @Title List
// @Description 获取用户问题列表
// @Success 200 {object} controllers.RespJson
// @router /list [get]
func (c *QuestionController) List() {
	CurUser := c.CurUser()
	var data []models.Question
	question := new(models.Question)
	qt := orm.NewOrm().QueryTable(question)
	qt.Filter("user_id", CurUser.Id).All(&data, "Id", "Content")
	c.SuccessJson("", data)
}

// @Title Add
// @Description 设置用户问题
// @Success 200 {object} controllers.RespJson
// @router /add [post]
func (c *QuestionController) Add() {
	CurUser := c.CurUser()
	question := new(models.Question)
	question.UserId = CurUser.Id
	o := orm.NewOrm()
	o.Delete(question, "user_id")

	var contents []string
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &contents)
	for _, content := range contents {
		question = &models.Question{UserId: CurUser.Id, Content: content}
		o.Insert(question)
	}

	c.SuccessJson("", "")
}
