package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"sync"
)

func GetConfig() *viper.Viper {
	var v *viper.Viper
	var once sync.Once
	once.Do(func() {
		v = viper.New()
		v.SetConfigType("yaml")
		v.SetConfigName("config")
		abs, err2 := filepath.Abs("")
		if err2 != nil {
			fmt.Println(err2)
		}
		v.AddConfigPath(abs + "/config")
		err := v.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
	})
	return v
}
