package main

import (
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"os"

	// echo
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// local packages
)

// Handler
func helloHandler(c echo.Context) error {
	c.Echo().Logger.Debug("debug test")
	return c.String(http.StatusOK, "Hello, World!")
}

func uploadHandler(c echo.Context) error {
	name := c.FormValue("name")

	// File Upload
	c.Echo().Logger.Debug("read file")
	file, err := c.FormFile("file_name")
	if err != nil {
		c.Echo().Logger.Warn(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		c.Echo().Logger.Warn(err)
		return err
	}
	defer src.Close()

	// File Destination
	dst, err := os.Create("./img/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		c.Echo().Logger.Warn(err)
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("File %s Uploaded, name=%s", file.Filename, name))
}

func octetFileHandler(c echo.Context) error {
	// check content-type
	contentType := c.Request().Header.Get("Content-Type")
	if contentType != "image/jpeg" {
		c.Echo().Logger.Warn("upload file is not jpeg")
		return c.HTML(http.StatusBadRequest, "Bad Request")
	}
	// get body data
	body := c.Request().Body
	img, err := jpeg.Decode(body)
	if err != nil {
		c.Echo().Logger.Warn(err)
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}
	dst, err := os.Create("./img/test_file.jpg")
	defer dst.Close()
	if err != nil {
		c.Echo().Logger.Warn(err)
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := jpeg.Encode(dst, img, &jpeg.Options{100}); err != nil {
		c.Echo().Logger.Warn(err)
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.String(http.StatusOK, "128,128")
}

// Main
func main() {
	var addr = ":8080"

	e := echo.New()

	// Access Log
	accessLogFile, _ := os.Create("logs/access.log")
	defer accessLogFile.Close()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: accessLogFile,
	}))
	e.Use(middleware.Recover())

	// Debug Log
	//	debugLogFile, _ := os.Create("logs/debug.log")
	//	defer debugLogFile.Close()
	//	e.Logger.SetOutput(debugLogFile)
	e.Debug = true

	// Route
	e.File("/quu6uriyex4eedahxieroox4rue8zu2e", "./public/index.html")
	e.GET("/wee9reiw9ieth3air2shahthuu0haibu", helloHandler)
	e.POST("/no5eepeiyeil9eevaij4eed5ohghahva", uploadHandler)
	e.POST("/id1aefup8oozahlo6etai4gei2aew5ee", octetFileHandler)

	e.Logger.Fatal(e.Start(addr))
}
