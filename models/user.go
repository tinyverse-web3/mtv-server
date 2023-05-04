package models

import (
	"time"
)

// 用户
type User struct {
	Id              int    `orm:"auto; pk"`
	Name            string `orm:"size(32);" json:"name"`
	Password        string `orm:"size(256);" json:"password"`
	Email           string `orm:"size(255);" json:"email"`
	Mobile          string `orm:"size(255);" json:"mobile"`
	NostrPublicKey  string `orm:"size(128); null" json:"nostrPublicKey"`
	PublicKey       string `orm:"size(128); null"`
	GuardianSssData string `orm:"size(512); null" json:"guardianSssData"`
	QuestionSssData string `orm:"size(512); null" json:"questionSssData"`
	Address         string `orm:"size(128); null"`
	Sign            string `orm:"size(255); null"`
	Ipns            string `orm:"size(255); null"`
	DbAddress       string `orm:"size(255); null"`
	// ConfirmCode           string    `orm:"size(16); null"`
	// ConfirmCodeUpdateTime time.Time `orm:"type(datetime); null"`
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime time.Time `orm:"type(datetime); auto_now; null"`
	Status     int       `orm:"default(0)"`
	SafeLevel  int       `orm:"default(0)" json:"safeLevel"` // 安全等级
	ImgCid     string    `orm:"size(64)" json:"imgCid"`      // 头像存储在IPFS上的cid
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) TableEngine() string {
	return "INNODB"
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Email"},
	}
}

func init() {
	//注册模型
	// orm.RegisterModel(new(User))
}
