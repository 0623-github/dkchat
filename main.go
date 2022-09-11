package main

import (
	"dkchat/config"
	"fmt"
)



func main() {
	getConfig := config.GetConfig()
	get := getConfig.Get("user")
	fmt.Println(get)
}
