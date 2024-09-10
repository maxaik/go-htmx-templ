package main

import (
	"context"
	"go-htmx/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	component := templates.Index()

	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8080"))
}
