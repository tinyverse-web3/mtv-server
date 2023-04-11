package models

import (
	"time"
)

// 守护者
type Guardian struct {
	Id          int       `orm:"auto; pk"`
	UserId      int       `orm:"size(11)" json:"userId"`
	Type        string    `orm:"size(32);" json:"type"`
	Account     string    `orm:"size(255);" json:"account"`
	AccountMask string    `orm:"size(255);" json:"accountMask"`
	CreateTime  time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime  time.Time `orm:"type(datetime); auto_now; null"`
}

func (u *Guardian) TableName() string {
	return "guardian"
}

func (u *Guardian) TableEngine() string {
	return "INNODB"
}

func (u *Guardian) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}

func init() {
	//注册模型
	// orm.RegisterModel(new(User))
}
