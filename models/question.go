package models

import "time"

// table：用户问题
type Question struct {
	Id         int       `orm:"auto; pk"`
	UserId     int       `orm:"size(11)" json:"userId"`
	Content    string    `orm:"size(255);" json:"content"` //问题内容
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime time.Time `orm:"type(datetime); auto_now; null"`
}

// 定义表名
func (u *Question) TableName() string {
	return "question"
}

// 定义引擎
func (u *Question) TableEngine() string {
	return "INNODB"
}

// 定义普通索引
func (u *Question) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
