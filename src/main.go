package main

import (
	"net/http"
	"os"

	// echo
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// local packages
)

func main() {
	var addr = ":8080"

	e := echo.New()

	// Log
	logFile, _ := os.Create("logs/echo.log")
	defer logFile.Close()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))
	e.Use(middleware.Recover())

	// Route
	e.Static("/", "./public")
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(addr))
}
