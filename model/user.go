package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	AccountNumber string `json:"account_number"`
	// 密码
	Pwd string `json:"pwd"`
	// 姓名
	RealName string `json:"real_name"`
	// 昵称
	NickName string `json:"nick_name"`
	// 年龄
	Age int `json:"age"`
	// 等级
	Grade int `json:"grade"`
	// 性别
	Gender int `json:"gender"`
	// 电话号码
	Phone string `json:"phone"`
	// 是否存在
	IsExist sql.NullInt16 `json:"is_exist" gorm:"default:1"`
}

func (User)TableName() string {
	return "user"
}
