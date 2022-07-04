package main

import (
	"fmt"
	"github.com/spf13/viper"
)



func main() {
	get := viper.Get("user")
	fmt.Println(get)
}
