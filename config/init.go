package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

func GetConfig() *viper.Viper {
	var v *viper.Viper
	var once sync.Once
	once.Do(func() {
		v = viper.New()
		v.SetConfigType("yaml")
		v.SetConfigName("config")
		v.AddConfigPath("config")
		err := v.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
	})
	return v
}
