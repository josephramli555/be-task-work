package db_test

import (
	"testing"
	"surveyapp/config"
	"surveyapp/models"
	"surveyapp/lib/db"
	"github.com/stretchr/testify/assert"
)

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

func TestGetAllAnswer(t *testing.T){
	var testCase = []struct{
		insertData bool
		expectedLen int
		answers []models.Answer
	}{
		{
			false,
			0,
			[]models.Answer{},
		},
		{
			true,
			6,
			[]models.Answer{
				models.NewAnswer("Naya","Warna favorit?","Hijau"),
				models.NewAnswer("Naya","Makanan favorit?","Jengkol"),
				models.NewAnswer("Naya","Wisata favorit?","Eiffel"),
				models.NewAnswer("Aldi","Warna favorit?","Hijau"),
				models.NewAnswer("Aldi","Makanan favorit?","Gorengan"),
				models.NewAnswer("Aldi","Wisata favorit?","Eiffel"),
			},
		},
	}
	config.InitDBTest()
	for _,test := range testCase{
		if !test.insertData{
			answerList := db.GetAllAnswers()
			assert.Equal(t,test.expectedLen,len(answerList),"Length of data not equal")
		}else{
			insertDummyData()
			answerList := db.GetAllAnswers()
			assert.Equal(t,test.expectedLen,len(answerList),"Length of data not equal")
			for idx,actualAnswer := range answerList{
				assert.Equal(t,test.answers[idx].Answer,actualAnswer.Answer,"Answer not equal expected "+ test.answers[idx].Answer)
				assert.Equal(t,test.answers[idx].Question,actualAnswer.Question,"Question not equal expected "+test.answers[idx].Answer)
				assert.Equal(t,test.answers[idx].User,actualAnswer.User,"User not equal expected "+test.answers[idx].Answer)
			}
		} 
	}
}



