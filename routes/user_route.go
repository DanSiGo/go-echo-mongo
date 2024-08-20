package routes

import (
	"echo-mongo/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo)  {
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetAUser)
	e.PUT("user/:userId", controllers.UpdateAUser)
	e.DELETE("user/:userId", controllers.DeleteAUser)
}