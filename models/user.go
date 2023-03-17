package models

import (
	"time"
)

// 定义表结构
type User struct {
	Id                    int       `orm:"auto; pk"`
	Name                  string    `orm:"size(32); unique;" json:"name"`
	Email                 string    `orm:"size(255); unique;" json:"email"`
	Mobile                string    `orm:"size(255); unique;" json:"mobile"`
	NostrPublicKey        string    `orm:"size(128); null" json:"nostrPublicKey"`
	PublicKey             string    `orm:"size(128); null"`
	SssData               string    `orm:"size(512); null"`
	Address               string    `orm:"size(128); null"`
	Sign                  string    `orm:"size(255); null"`
	Ipns                  string    `orm:"size(255); null"`
	DbAddress             string    `orm:"size(255); null"`
	ConfirmCode           string    `orm:"size(16); null"`
	ConfirmCodeUpdateTime time.Time `orm:"type(datetime); null"`
	CreateTime            time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime            time.Time `orm:"type(datetime); auto_now; null"`
	Status                int       `orm:"default(0)"`
}

// 定义表名
func (u *User) TableName() string {
	return "user"
}

// 定义引擎
func (u *User) TableEngine() string {
	return "INNODB"
}

// 定义普通索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Email"},
	}
}

func init() {
	//注册模型
	// orm.RegisterModel(new(User))
}
