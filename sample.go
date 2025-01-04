package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(echoprometheus.NewMiddleware("myapp"))   // adds middleware to gather metrics
	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	type message struct {
		Message string `json:"message"`
	}

	e.GET("/", func(c echo.Context) error {
		m := &message{
			Message: "Bonjour, tous le monde !",
		}
		return c.JSON(http.StatusOK, m)
	})

	e.GET("/ping", func(c echo.Context) error {
		m := &message{
			Message: "pong",
		}
		return c.JSON(http.StatusOK, m)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
