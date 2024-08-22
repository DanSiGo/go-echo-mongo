package main

import (
	"echo-mongo/configs"
	"echo-mongo/routes"
	// "fmt"
	// "net"
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


	// port := 8080
	var addr string
	// for{
	// 	addr := fmt.Sprintf(":%d", port)
	// 	if ln, err := net.Listen("tcp", addr); err == nil {
	// 		ln.Close()
	// 		break
	// 	}
	// 	fmt.Printf("Port %d is in use. Trying next available port...\n", port)
	// 	port++
	// }
	// fmt.Printf("Listening on port %d\n", port)
	
	e.Logger.Fatal(e.Start(addr))
}