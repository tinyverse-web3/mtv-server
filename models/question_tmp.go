package models

import "time"

// table：问题模版
type QuestionTmp struct {
	Id         int       `orm:"auto; pk"`
	Content    string    `orm:"size(255);" json:"content"` //问题内容
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime time.Time `orm:"type(datetime); auto_now; null"`
}

// 定义表名
func (u *QuestionTmp) TableName() string {
	return "question_tmp"
}

// 定义引擎
func (u *QuestionTmp) TableEngine() string {
	return "INNODB"
}

// 定义普通索引
func (u *QuestionTmp) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
