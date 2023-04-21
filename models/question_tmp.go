package models

import "time"

// 问题模版
type QuestionTmp struct {
	Id         int       `orm:"auto; pk"`
	Title      string    `orm:"size(128);" json:"title"`
	Content    string    `orm:"size(512);" json:"content"` //问题内容(json格式)
	Type       int       `orm:"size(11);" json:"type"`     // 1：默认问题；2：自定义问题
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime time.Time `orm:"type(datetime); auto_now; null"`
}

func (u *QuestionTmp) TableName() string {
	return "question_tmp"
}

func (u *QuestionTmp) TableEngine() string {
	return "INNODB"
}

func (u *QuestionTmp) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
