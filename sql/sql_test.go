package sql

import (
	_ "dkchat/config"
	"dkchat/model"
	"github.com/jinzhu/gorm"
	"testing"
)


func TestSql(t *testing.T) {
	pool := NewPool()
	ans := pool.Exec(func(db *gorm.DB) interface{}{
		return db.First(&model.User{})
	})
	println(ans.(*gorm.DB).RowsAffected)
}
