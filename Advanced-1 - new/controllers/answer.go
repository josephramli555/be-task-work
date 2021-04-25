package controllers

import (
	"net/http"
	"surveyapp/config"
	"surveyapp/lib/db"
	"surveyapp/models"
	"github.com/labstack/echo/v4"
)

func DeleteAnswerController(c echo.Context) error {
	deletedId  := c.Param("id")
	db.DeleteAnswer(deletedId)
	return c.JSON(http.StatusOK,map[string]interface{}{
		"code" : 200,
		"status" : "Success",
		"message" : "Delete data successful",
	})
}

func GetAnswerController(c echo.Context)error{
	answers := db.GetAllAnswers()
	message := ""
	if len(answers) == 0{
		message = "There are no answers, please fill the answer through survey"
	}
	return c.Render(http.StatusOK, "report.html", echo.Map{
		"answerList": answers,
		"message": message,
	})
}

func CreateAnswerController(c echo.Context) error {
	name := c.FormValue("name")
	userAnswers := []string{c.FormValue("answer1"), c.FormValue("answer2"), c.FormValue("answer3")}
	answers := []models.Answer{}
	for i := 0; i < len(config.QuestionList); i++ {
		answers = append(answers, models.Answer{
			User:     name,
			Question: config.QuestionList[i],
			Answer: userAnswers[i],
		})
	}
	db.InsertAnswers(answers)
	return c.Redirect(http.StatusSeeOther,"/")
}

func UpdateAnswerController(c echo.Context)error{
	updateID := c.Param("id")
	newAnswer := c.FormValue("answer")
	db.UpdateAnswer(updateID,newAnswer)
	return c.JSON(http.StatusOK,map[string]interface{}{
		"code" : 200,
		"status" : "Success",
		"message" : "Update data successful",
	})
}