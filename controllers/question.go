package controllers

import (
	"encoding/json"
	"mtv/models"

	"github.com/beego/beego/v2/client/orm"
)

type QuestionController struct {
	BaseController
}

func (c *QuestionController) TmpList() {
	var data []models.QuestionTmp
	question := new(models.QuestionTmp)
	qt := orm.NewOrm().QueryTable(question)
	qt.All(&data, "Id", "Content")
	c.SuccessJson("", data)
}

func (c *QuestionController) List() {
	var data []models.Question
	question := new(models.Question)
	qt := orm.NewOrm().QueryTable(question)
	qt.Filter("user_id", CurUser.Id).All(&data, "Id", "Content")
	c.SuccessJson("", data)
}

func (c *QuestionController) Add() {
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
