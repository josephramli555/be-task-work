package main

import (
	"fmt"
	"surveyapp/config"
	"surveyapp/controllers"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"surveyapp/modules"
)

func main() {

	new(modules.Configuration).Init()
	configJson := new(modules.Configuration).GetConfig()
	port := fmt.Sprintf(":%v",configJson.GetString("APP_PORT"))
	config.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")

	e.Renderer = echoview.Default()

	e.GET("/", controllers.HomePageController)
	e.GET("/survey",controllers.SurveyPageController)
	e.GET("/formupdate/:id",controllers.FormUpdatePageController)

	e.GET("/report",controllers.GetAnswerController)
	e.POST("/answers", controllers.CreateAnswerController)
	e.PUT("/answers/:id",controllers.UpdateAnswerController)
	e.DELETE("/answers/:id",controllers.DeleteAnswerController)

	e.Logger.Fatal(e.Start(port))
}
