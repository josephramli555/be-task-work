package main

import(
	"fmt"
	"sales-api/modules"
	"sales-api/config"
	"sales-api/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	config.InitDB()
	e:= echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	new(modules.Configuration).Init()
	configJson := new(modules.Configuration).GetConfig()
	port := fmt.Sprintf(":%v",configJson.GetString("APP_PORT"))


	e.GET("/order/:id",controllers.GetOrderController)

	e.Logger.Fatal(e.Start(port))
}	