package models

import (
	"time"
)

type ImFriend struct {
	Id            int       `orm:"auto; pk"`
	From          int       `orm:"size(11);" json:"from"` // 聊天发起者id
	FromPublicKey string    `orm:"size(128); null" json:"fromPublicKey"`
	To            int       `orm:"size(11);" json:"to"` // 聊天参与者id
	ToPublicKey   string    `orm:"size(128); null" json:"toPublicKey"`
	CreateTime    time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime    time.Time `orm:"type(datetime); auto_now; null"`
	Status        int       `orm:"default(0)"`
}

func (u *ImFriend) TableName() string {
	return "im_friend"
}

func (u *ImFriend) TableEngine() string {
	return "INNODB"
}

func (u *ImFriend) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
