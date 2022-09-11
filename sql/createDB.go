package sql

import (
	"dkchat/model"
	"github.com/jinzhu/gorm"
)

func CreateUserTable(){
	pool := NewPool()
	pool.Exec(func(db *gorm.DB) interface{} {
		//db.AutoMigrate(&model.User{})
		db.CreateTable(&model.User{})
		return db
	})
}