package controllers

import (
	"net/http"
	"surveyapp/config"
	"surveyapp/lib/db"
	"github.com/labstack/echo/v4"
)

func HomePageController(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", echo.Map{})
}

func SurveyPageController(c echo.Context) error {
	return c.Render(http.StatusOK, "survey.html", echo.Map{
		"Q1":config.QuestionList[0],
		"Q2":config.QuestionList[1],
		"Q3":config.QuestionList[2],
	})
}

func FormUpdatePageController(c echo.Context)error{
	updateId  := c.Param("id")
	answer := db.GetAnswer(updateId)
	return c.Render(http.StatusOK, "formupdate.html", echo.Map{
		"prevAnswer" : answer,
	})
}