package models

import (
	"time"
)

type Gun struct {
	Id         int       `orm:"auto; pk"`
	UserId     int       `orm:"size(11)" json:"userId"`
	Name       string    `orm:"size(128); null" json:"name"`
	CreateTime time.Time `orm:"type(datetime); auto_now_add; null"`
	UpdateTime time.Time `orm:"type(datetime); auto_now; null"`
}

func (u *Gun) TableName() string {
	return "gun"
}

func (u *Gun) TableEngine() string {
	return "INNODB"
}

func (u *Gun) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}
