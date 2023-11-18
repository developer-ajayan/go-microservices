package main

import (
	"github.com/gin-gonic/gin"
	"auth-service/pkg"
)

func main() {
	r := gin.Default()

	r.POST("/register", pkg.RegisterHandler)
	r.POST("/login", pkg.LoginHandler)
	
	r.Use(pkg.AuthenticationMiddlware)

	r.GET("/protected", pkg.ProtectedHandler)
	r.Run()

}