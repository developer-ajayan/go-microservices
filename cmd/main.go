package main

import (
	"github.com/gin-gonic/gin"
	"auth-service/pkg"
	"net/http/pprof"
	"net/http"
)

func main() {
	r := gin.Default()
	// profiling
	// Enable pprof routes
	r.GET("/debug/pprof/", gin.WrapF(http.HandlerFunc(pprof.Index)))
	r.GET("/debug/pprof/cmdline", gin.WrapF(http.HandlerFunc(pprof.Cmdline)))
	r.GET("/debug/pprof/profile", gin.WrapF(http.HandlerFunc(pprof.Profile)))
	r.GET("/debug/pprof/symbol", gin.WrapF(http.HandlerFunc(pprof.Symbol)))
	r.GET("/debug/pprof/trace", gin.WrapF(http.HandlerFunc(pprof.Trace)))


	r.POST("/register", pkg.RegisterHandler)
	r.POST("/login", pkg.LoginHandler)
	
	r.Use(pkg.AuthenticationMiddlware)

	r.GET("/protected", pkg.ProtectedHandler)
	r.Run(":8002")

}