package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	glob := new(globals)
	glob.Name = "main-api"
	glob.Log = log15.New("config", glob.Name)

	initLogger()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/crunch", func(c echo.Context) error {
		return cruncher(c, glob)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
