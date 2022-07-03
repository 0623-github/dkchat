package UnitTest

import (
	_ "dkchat/config"
	"dkchat/model"
	"dkchat/sql"
	"github.com/jinzhu/gorm"
	"testing"
)

func init() {

}

func TestSql(t *testing.T) {
	pool := sql.NewPool(10)
	ans := pool.Exec(func(db *gorm.DB) interface{}{
		return db.First(&model.User{})
	})
	println(ans.(*gorm.DB).RowsAffected)
	sql.DeletePool(pool)
}
