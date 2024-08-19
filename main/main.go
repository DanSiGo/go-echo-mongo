package main

import (
	"echo-mongo/configs"
	"echo-mongo/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()

	configs.ConnectDB()

	routes.UserRoute(e)
	
	e.GET("/", func (c echo.Context) error{
		return c.String(http.StatusOK, "yallo from the web side")
	})

	e.GET("/home2", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello via json"})
	})

	e.Logger.Fatal(e.Start(":8000"))
}