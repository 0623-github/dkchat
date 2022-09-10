package sql

import (
	"dkchat/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type Pool struct {
	db chan *gorm.DB
	size int
}

type ExecFunc func(db *gorm.DB) interface{}

// NewPool Singleton mode
func NewPool() *Pool {
	var p *Pool
	var once sync.Once
	poolSize := 10
	once.Do(func() {
		p = &Pool{
			db:   make(chan *gorm.DB, poolSize),
			size: poolSize,
		}
		v := config.GetConfig()
		conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True", v.Get("sqluser"), v.Get("sqlpwd"), v.Get("sqlhost"), v.Get("dbname"))
		for i := 0; i < poolSize; i ++ {
			db, err := gorm.Open("mysql", conn)
			if err != nil {
				logger := log.Default()
				logger.Println("errorerrorerrorerrorerrorerror")
				logger.Println(conn)
				logger.Println(err)
			}
			p.db <- db
		}
	})
	return p
}

func (p *Pool)Exec(execFunc ExecFunc) interface{} {
	db := <- p.db
	defer func() {
		p.db <- db
	}()
	return execFunc(db)
}
