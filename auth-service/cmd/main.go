package main

import (
	"github.com/gin-gonic/gin"
	"auth-service/pkg"
)

func main() {
	r := gin.Default()

	r.POST("/login", pkg.LoginHandler)

	r.Run()

}