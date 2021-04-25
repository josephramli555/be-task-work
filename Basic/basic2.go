package main

import (
	"basictest/utility"
	"net/http"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "fmt"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = echoview.Default()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", echo.Map{})
	})

	
	e.GET("/soal1",func(c echo.Context) error {
		return c.Render(http.StatusOK,"soal1.html",echo.Map{})
	})

	e.POST("/soal1",func(c echo.Context)error{
		output := utility.ReverseUpperLower(c.FormValue("input1"))
		return c.Render(http.StatusOK,"soal1.html",echo.Map{
			"output" : output,
		})
	})

	e.GET("/soal3",func(c echo.Context) error {
		return c.Render(http.StatusOK,"soal3.html",echo.Map{})
	})

	e.POST("/soal3",func(c echo.Context)error{
		output := utility.ProcessWord(c.FormValue("input3"))
		return c.Render(http.StatusOK,"soal3.html",echo.Map{
			"output" : output,	
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
