package main

import (
"fmt"
"net/http"

"github.com/labstack/echo"
"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		req := c.Request()
		format := `
			<code>
				Protocol: %s<br>
				Host: %s<br>
				Remote Address: %s<br>
				Method: %s<br>
				Path: %s<br>
			</code>
		`
		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	})
	e.Logger.Fatal(e.StartTLS(":8080", "./proyect/ssl/server.pem", "./proyect/ssl/server.key"))
}