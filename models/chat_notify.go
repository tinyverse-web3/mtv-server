package models

import (
	"time"
)

type ChatNotify struct {
	Id            int       `orm:"auto; pk"`
	From          int       `orm:"size(11);" json:"from"` // 聊天发起者id
	FromPublicKey string    `orm:"size(128); null" json:"fromPublicKey"`
	To            int       `orm:"size(11);" json:"to"` // 聊天参与者id
	ToPublicKey   string    `orm:"size(128); null" json:"toPublicKey"`
	CreateTime    time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime    time.Time `orm:"type(datetime); auto_now; null"`
	Status        int       `orm:"default(0)"`
}

// 定义表名
func (u *ChatNotify) TableName() string {
	return "chat_notify"
}

// 定义引擎
func (u *ChatNotify) TableEngine() string {
	return "INNODB"
}

// 定义普通索引
func (u *ChatNotify) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
