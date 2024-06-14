package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
  e.Renderer = 

	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
