package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"surveyapp/config"
	"surveyapp/controllers"
	"surveyapp/models"
)

type response struct{
	Code  int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
}

func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func insertDummyData()[]models.Answer{
	answers :=[]models.Answer{
		models.NewAnswer("Naya","Warna favorit?","Hijau"),
		models.NewAnswer("Naya","Makanan favorit?","Jengkol"),
		models.NewAnswer("Naya","Wisata favorit?","Eiffel"),
		models.NewAnswer("Aldi","Warna favorit?","Merah"),
		models.NewAnswer("Aldi","Makanan favorit?","Gorengan"),
		models.NewAnswer("Aldi","Wisata favorit?","Candi"),
	}
	res := config.DB.Create(answers)
	if res.Error != nil{
		panic(res.Error)
	}
	return answers
}

func TestDeleteAnswerController(t *testing.T){
	e:= InitEcho()
	answers := insertDummyData()
	currentLen := len(answers)
	for idx,answer := range answers{
		req := httptest.NewRequest(http.MethodDelete,"/",nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req,rec)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(answer.Answer_id))

		if assert.NoError(t,controllers.DeleteAnswerController(c)){
			body := rec.Body.String()
			var response response
			err := json.Unmarshal([]byte(body),&response)
			if err!=nil{
				t.Error("Failed to unmarshall json when id is ",answer.Answer_id)
			}
			
			//search if data still inside db after deletion, it should return 0 because the data already deleted
			res := config.DB.First(&models.Answer{},answer.Answer_id)
			//check if length of array of answers, every time answers is deleted
			returnedAnswer := []models.Answer{}
			config.DB.Find(&returnedAnswer)
			
			assert.Equal(t,200,response.Code,"Error getting code %d",response.Code)
			assert.Equal(t,currentLen-idx-1,len(returnedAnswer),"Error array len expected:%d, returned %d",(currentLen-idx-1),len(returnedAnswer))
			assert.Equal(t,int64(0),res.RowsAffected,"Error when deleting id : %d, row affected %d",answer.Answer_id,res.RowsAffected)
		}
	}
}

func TestPostAnswerController(t *testing.T){
	testCasePostAnswer := []struct{
		name string
		answers []string
	}{
		{
			name: "naya",
			answers: []string{"Nasi Goreng","Korea","Tom Cruise"},
		},
		{
			name: "rudi",
			answers: []string{"Ayama Bakar","Jepang","Rooney"},
		},
	}

	e:=InitEcho()
	for idx,answer := range testCasePostAnswer{
		f := make(url.Values)
		f.Set("answer1",answer.answers[0])
		f.Set("answer2",answer.answers[1])
		f.Set("answer3",answer.answers[2])
		f.Set("name",answer.name)

		req := httptest.NewRequest(http.MethodPost,"/",strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

		rec := httptest.NewRecorder()
		c := e.NewContext(req,rec)
		if assert.NoError(t,controllers.CreateAnswerController(c)){
			returnedAnswer := []models.Answer{}
			config.DB.Find(&returnedAnswer)
			//check if len of answer equal to number of inserted answer
			totalAnswer := len(returnedAnswer)
			assert.Equal(t,(idx+1)*3,totalAnswer,"Expected length : %d, returned length %d",(idx+1)*3,totalAnswer)
			
			assert.Equal(t,answer.answers[2],returnedAnswer[totalAnswer-1].Answer,"Last answer didnt match expected:%s  actual:%s",answer.answers[2],returnedAnswer[totalAnswer-1].Answer)
			assert.Equal(t,answer.name,returnedAnswer[totalAnswer-1].User,"Usename didnt match expected:%s actual:%s",answer.name,returnedAnswer[totalAnswer-1].User)
		}
	}
}

func TestUpdateAnswerController(t *testing.T){
	//set 6 new updated answer since insert dummy data insert 6 answer
	testCaseUpdatedAnswer := []string{"Merah","Daging","Paris","Biru","Jambu","Paris"}
	e:=InitEcho()
	answers := insertDummyData()

	for idx,newAnswer := range testCaseUpdatedAnswer{
		f := make(url.Values)
		f.Set("answer",newAnswer)

		req := httptest.NewRequest(http.MethodPut,"/",strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

		rec := httptest.NewRecorder()
		c := e.NewContext(req,rec)
		//set updated id parameter
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(answers[idx].Answer_id))

		if assert.NoError(t,controllers.UpdateAnswerController(c)){
			body := rec.Body.String()
			var response response
			err := json.Unmarshal([]byte(body),&response)
			if err!=nil{
				t.Error("Failed to unmarshall json when id is ",answers[idx].Answer_id)
			}

			//take the latest data that has been updated
			var latestAnswer models.Answer
			searchId :=answers[idx].Answer_id
			config.DB.First(&latestAnswer,searchId)
			assert.Equal(t,200,response.Code,"Error getting code %d",response.Code)
			//check if answer is updated in db
			assert.Equal(t,newAnswer,latestAnswer.Answer,"Expected Answer:%s Latest Answer:%s",newAnswer,latestAnswer.Answer)
		}
	}
}