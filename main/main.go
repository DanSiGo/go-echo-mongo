package main

import (
	"echo-mongo/configs"
	"echo-mongo/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "yallo from the web side")
	})

	e.GET("/products", controller.GetProducts)

	e.Logger.Fatal(e.Start(":8000"))
}
