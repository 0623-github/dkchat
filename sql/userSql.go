package sql

import (
	"dkchat/model"
	"github.com/jinzhu/gorm"
)

func NewUser(u *model.User) {
	pool := NewPool()
	pool.Exec(func(db *gorm.DB) interface{} {
		db.Create(u)
		return nil
	})
	SelectUserByAccount(u)
}

func SelectMaxAccount() string {
	pool := NewPool()
	number := pool.Exec(func(db *gorm.DB) interface{} {
		u := &model.User{}
		db.Order("account_number desc").First(u)
		return u.AccountNumber
	})
	return number.(string)
}

func SelectUserByAccount(user *model.User) {
	pool := NewPool()
	pool.Exec(func(db *gorm.DB) interface{} {
		db.Where("account_number = ?", user.AccountNumber).Take(user)
		return user
	})
}