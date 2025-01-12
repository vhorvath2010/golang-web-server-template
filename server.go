package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vhorvath2010/what-do-they-pay/templates"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return templates.HomePage().Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start("localhost:8080"))
}
