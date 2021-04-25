package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"surveyapp/config"
	"surveyapp/controllers"
	"surveyapp/models"
	"testing"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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
		models.NewAnswer("Aldi","Warna favorit?","Hijau"),
		models.NewAnswer("Aldi","Makanan favorit?","Gorengan"),
		models.NewAnswer("Aldi","Wisata favorit?","Eiffel"),
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

	for _,answer := range answers{
		req := httptest.NewRequest(http.MethodDelete,"/",nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req,rec)
		c.SetPath("/answers")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(answer.Answer_id))
	
		if assert.NoError(t,controllers.DeleteAnswerController(c)){
			body := rec.Body.String()
			var response response
			err := json.Unmarshal([]byte(body),&response)
			if err!=nil{
				t.Error("Failed to unmarshall json when id is ",answer.Answer_id)
			}
			
			//search if data still inside db after deletion
			res := config.DB.First(&models.Answer{},answer.Answer_id)

			assert.Equal(t,200,response.Code,"Error getting code %d",response.Code)
			assert.Equal(t,int64(0),res.RowsAffected,"Error when deleting id : %d, row affected %d",answer.Answer_id,res.RowsAffected)
		}
	}

}

