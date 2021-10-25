package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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
	//log template
	httpEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: logfmt,
		Output:    os.Stdout,
	}))

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

func logfmt(params gin.LogFormatterParams) string {

	var statusColor, methodColor, resetColor string
	if params.IsOutputColor() {
		statusColor = params.StatusCodeColor()
		methodColor = params.MethodColor()
		resetColor = params.ResetColor()
	}

	if params.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		params.Latency = params.Latency - params.Latency%time.Second
	}

	return fmt.Sprintf("[AUTH-LOG] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, params.StatusCode, resetColor,
		params.Latency,
		params.ClientIP,
		methodColor, params.Method, resetColor,
		params.Path,
		params.ErrorMessage,
	)
}
