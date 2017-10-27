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
	e.POST("/crunch", func(c echo.Context) error {
		return cruncher(c)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
