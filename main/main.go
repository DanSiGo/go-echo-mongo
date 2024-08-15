package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()
	
	e.GET("/", func (c echo.Context) error{
		return c.String(http.StatusOK, "yallo from the web side")
	})
	e.Logger.Fatal(e.Start(":8080"))
}