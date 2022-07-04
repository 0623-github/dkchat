package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}