package models

import (
	"time"
)

// 定义表结构
type NostrRelay struct {
	Id         int       `orm:"auto; pk"`
	WsServer   string    `orm:"size(128);" json:"wsServer"`
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime time.Time `orm:"type(datetime); auto_now; null"`
	Remark     string    `orm:"size(512);" json:"remark"`
	Status     int       `orm:"default(0)"` // 1:不可用 2:可用 3:检查
}

// 定义表名
func (u *NostrRelay) TableName() string {
	return "nostr_relay"
}

// 定义引擎
func (u *NostrRelay) TableEngine() string {
	return "INNODB"
}

// 定义普通索引
func (u *NostrRelay) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
