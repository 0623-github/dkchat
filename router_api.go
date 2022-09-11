package main

import (
	"dkchat/handler"
	"github.com/gin-gonic/gin"
)

func registerApi(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	r.POST("/signUp", handler.SignUp)
	r.GET("/signIn", handler.SignIn)
}
