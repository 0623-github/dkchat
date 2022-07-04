package model

import (
	"dkchat/sql"
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

func init() {
	pool := sql.NewPool(1)
	pool.Exec(func(db *gorm.DB) interface{} {
		db.AutoMigrate(&User{})
		u1 := User{
			RealName: "ldk",
			NickName: "sqrt3",
			Age: 13,
			Pwd: "123456",
			Grade: "10",
		}
		db.Create(&u1)
		return nil
	})
}