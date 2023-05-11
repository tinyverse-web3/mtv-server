package controllers

import (
	"encoding/json"
	"mtv/models"

	"github.com/beego/beego/v2/client/orm"
)

type RecordController struct {
	BaseController
}

// @Title Report
// @Description 上报异常
// @Param groupName body string true "组织名称"
// @Param data body string true "内容"
// @Param type body string true "类型"
// @Success 200 {object} controllers.RespJson
// @router /report [post]
func (c *RecordController) Report() {
	var record models.Record
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &record)

	o := orm.NewOrm()
	_, err := o.Insert(&record)
	if err != nil {
		c.ErrorJson("400000", "上报异常失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title List
// @Description 获取异常列表
// @Success 200 {object} controllers.RespJson
// @router /list [get]
func (c *RecordController) List() {
	var data []models.Record
	record := new(models.Record)
	qt := orm.NewOrm().QueryTable(record)
	qt.All(&data)
	c.SuccessJson("", data)
}
