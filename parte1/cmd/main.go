package main

import (
	shopping "app/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(shopping.ServerHeader)

	e.GET("/resumen/:dayId", shopping.GetSummary)

	e.Logger.Fatal(e.Start(":8080"))
}
