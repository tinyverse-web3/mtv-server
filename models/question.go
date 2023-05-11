package models

import "time"

// 用户问题
type Question struct {
	Id         int       `orm:"auto; pk;" json:"id"`
	UserId     int       `orm:"size(11);" json:"userId"`
	Title      string    `orm:"size(128);" json:"title"`
	Content    string    `orm:"size(1024);" json:"content"` //问题内容(json格式)
	Type       int       `orm:"size(11);" json:"type"`      // 1：默认问题；2：自定义问题
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null" json:"createTime"`
	UpdateTime time.Time `orm:"type(datetime); auto_now; null" json:"updateTime"`
}

func (u *Question) TableName() string {
	return "question"
}

func (u *Question) TableEngine() string {
	return "INNODB"
}

func (u *Question) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
