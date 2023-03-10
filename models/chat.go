package models

import (
	"time"
)

type Chat struct {
	Id            int       `orm:"auto; pk"`
	From          int       `orm:"size(11);" json:"from"` // 聊天发起者id
	FromPublicKey string    `orm:"size(128); null" json:"nostrPublicKey"`
	CreateTime    time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime    time.Time `orm:"type(datetime); auto_now; null"`
	Status        int       `orm:"default(0)"`
}

// 定义表名
func (u *Chat) TableName() string {
	return "chat"
}

// 定义引擎
func (u *Chat) TableEngine() string {
	return "INNODB"
}

// 定义普通索引
func (u *Chat) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
