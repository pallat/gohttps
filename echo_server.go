package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func users(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte(`{"username":"pallat", "name":"pallat","lastname":"anchaleechamaikorn"}`))
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users", users)

	// Start server
	e.Logger.Fatal(e.StartTLS(":443", "local.crt", "server.key"))
}
