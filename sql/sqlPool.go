package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

type Pool struct {
	db chan *gorm.DB
	size int
}

func NewPool(len int) *Pool {
	p := &Pool{
		db: make(chan *gorm.DB, len),
		size: len,
	}
	conn := fmt.Sprintf("%s:%s@(%s)/%s", viper.Get("sqluser"), viper.Get("sqlpwd"), viper.Get("sqlhost"), viper.Get("dbname"))
	for i := 0; i < len; i ++ {
		db, err := gorm.Open("mysql", conn)
		if err != nil {
			logger := log.Default()
			logger.Println(err)
			return nil
		}
		p.db <- db
	}
	return p
}