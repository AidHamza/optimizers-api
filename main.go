package main

import (
	"github.com/AidHamza/optimizers-api/config"
	"github.com/AidHamza/optimizers-api/log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := config.App.Init()
	if err != nil {
		log.Logger.Error("Cannot load app configuration", "VIPER_ERROR", err.Error())
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	  AllowOrigins: []string{config.App.Service.CrossOrigin},
	  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.POST("/crunch", func(c echo.Context) error {
		return cruncher(c)
	})

	e.Logger.Fatal(e.Start(":" + config.App.Service.Port))
}
