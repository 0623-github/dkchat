package main

import (
	"dkchat/sql"
	"github.com/gin-gonic/gin"
)

func initDB() {
	sql.CreateUserTable()
}

func runServer() {
	// 127.0.0.1:6789
	r := gin.Default()
	r.Use(Cors())
	registerApi(r)
	err := r.Run(":6789")
	if err != nil {
		return
	}
}

func main() {
	runServer()
	//initDB()
}
