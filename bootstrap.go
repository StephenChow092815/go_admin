package irisweb

import (
	"fmt"
	"irisweb/route"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

type Bootstrap struct {
	Application *iris.Application
	Port        int
	LoggerLevel string
}

func New(port int, loggerLevel string) *Bootstrap {
	var bootstrap Bootstrap
	bootstrap.Application = iris.New()
	bootstrap.Port = port
	bootstrap.LoggerLevel = loggerLevel
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // 允许所有来源
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	bootstrap.Application.UseRouter(corsMiddleware)
	return &bootstrap
}

func (bootstrap *Bootstrap) Serve() {
	bootstrap.Application.Logger().SetLevel(bootstrap.LoggerLevel)
	route.Register(bootstrap.Application)
	bootstrap.Application.Run(
		iris.Addr(fmt.Sprintf("127.0.0.1:%d", bootstrap.Port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutBodyConsumptionOnUnmarshal,
	)
}
