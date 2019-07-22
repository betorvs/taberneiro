package main

import (
	"context"

	"github.com/betorvs/taberneiro/config"
	"github.com/betorvs/taberneiro/controller"
	"github.com/betorvs/taberneiro/gateway/slackclient"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := echo.New()
	g := e.Group("/taberneiro/v1")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/health", controller.CheckHealth)
	g.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	g.POST("/events", controller.ReceiveEvents)
	g.POST("/messages", controller.ReceiveMessages)

	ctx := context.Background()
	s, err := slackclient.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := s.Run(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":" + config.Port))

}
