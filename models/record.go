package models

import "time"

type Record struct {
	Id         int       `orm:"auto; pk;" json:"id"`
	GroupName  string    `orm:"size(32);" json:"groupName"`
	Data       string    `orm:"size(1024);" json:"Data"`
	Type       string    `orm:"size(32);" json:"type"`
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null" json:"createTime"`
}

func (u *Record) TableName() string {
	return "record"
}

func (u *Record) TableEngine() string {
	return "INNODB"
}

func (u *Record) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
