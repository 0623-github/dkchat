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

type ExecFunc func(db *gorm.DB) interface{}

func NewPool(len int) *Pool {
	p := &Pool{
		db: make(chan *gorm.DB, len),
		size: len,
	}
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True", viper.Get("sqluser"), viper.Get("sqlpwd"), viper.Get("sqlhost"), viper.Get("dbname"))
	for i := 0; i < len; i ++ {
		db, err := gorm.Open("mysql", conn)
		if err != nil {
			logger := log.Default()
			logger.Println("errorerrorerrorerrorerrorerror")
			logger.Println(conn)
			logger.Println(err)
			return nil
		}
		p.db <- db
	}
	return p
}

func (p *Pool)Exec(execFunc ExecFunc) interface{} {
	db := <- p.db
	defer func() {
		p.db <- db
	}()
	return execFunc(db)
}

func DeletePool(pool *Pool) {
	for i := 0; i < pool.size; i ++ {
		db := <-pool.db
		err := db.Close()
		if err != nil {
			return
		}
	}
}