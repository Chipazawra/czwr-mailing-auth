package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/Chipazawra/czwrMailing-auth/doc"
	auth "github.com/Chipazawra/czwrMailing-auth/internal"
	"github.com/Chipazawra/czwrMailing-auth/pkg/pprofwrapper"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	defaultHost = "127.0.0.1"
	defaultPort = "1488"
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
	service.Register(httpEngine)

	//profile
	pprofrp := pprofwrapper.New()
	pprofrp.Register(httpEngine)

	//doc
	httpEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := httpEngine.Run(fmt.Sprintf("%v:%v", host, port))

	if err != nil {
		panic(err)
	}

}
