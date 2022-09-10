package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	// 姓名
	RealName string
	// 昵称
	NickName string
	// 年龄
	Age uint8
	// 密码
	Pwd string
	// 等级
	Grade string
}

func (User)TableName() string {
	return "user"
}
