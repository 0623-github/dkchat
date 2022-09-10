package sql

import (
	"dkchat/model"
	"github.com/jinzhu/gorm"
)

func createUserTable(){
	pool := NewPool()
	pool.Exec(func(db *gorm.DB) interface{} {
		db.AutoMigrate(&model.User{})
		u1 := model.User{
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