package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/Chipazawra/czwr-mailing-auth/doc"
	auth "github.com/Chipazawra/czwr-mailing-auth/internal/auth"
	"github.com/Chipazawra/czwr-mailing-auth/pkg/pprofwrapper"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	defaultHost = "0.0.0.0"
	defaultPort = "8885"
)

// @title czwrMailing - auth service
// @version 1.0
// @description This is a sample mailing servivce.
func main() {

	var host, port string

	flag.StringVar(&host, "host", "", "Host on which to start listening")
	flag.StringVar(&port, "port", "", "Port on which to start listening")
	flag.Parse()

	if host == "" {
		host = os.Getenv("AUTH_HOST")
		if host == "" {
			host = defaultHost
		}
	}

	if port == "" {
		port = os.Getenv("AUTH_PORT")
		if port == "" {
			port = defaultPort
		}
	}

	httpEngine := gin.New()
	httpEngine.Use(gin.Recovery())

	service := auth.New(auth.DefaultConfig)
	group := service.Register(httpEngine)
	//profiler
	pprofrp := pprofwrapper.New()
	pprofrp.Register(group)
	//doc
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := httpEngine.Run(fmt.Sprintf("%v:%v", host, port))

	if err != nil {
		panic(err)
	}

}
