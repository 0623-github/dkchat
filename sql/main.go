package main

import (
	"dkchat/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)


func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initSqlTable() {
	pool := NewPool(1)
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



func main() {
	initSqlTable()
	pool := NewPool(10)
	ans := pool.Exec(func(db *gorm.DB) interface{}{
		return db.First(&model.User{})
	})
	print(ans.(*gorm.DB).RowsAffected)
	DeletePool(pool)
}
